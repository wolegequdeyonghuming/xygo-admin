// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package bizorder

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/consts"
	"xygo/internal/library/contexts"
)

// currentRole 获取当前登录用户的角色 key（带日志记录便于排查）。
func currentRole(ctx context.Context) string {
	role := contexts.GetRoleKey(ctx)
	if role == "" {
		// 正常情况下经过 Auth 中间件必然有角色，这里仅兜底，避免空角色被当作无权限
		g.Log().Warning(ctx, "bizorder.currentRole: 当前用户角色为空")
	}
	return role
}

// currentUserId 获取当前登录用户 ID。
func currentUserId(ctx context.Context) int64 {
	return int64(contexts.GetUserId(ctx))
}

// canViewAll 是否可查看全部订单（收单员管理员/财务文员/管理员/超管）。
func canViewAll(role string) bool {
	switch role {
	case consts.RoleAgentManager, consts.RoleDocumentary, consts.RoleAdmin, consts.SuperRoleKey:
		return true
	}
	return false
}

// applyScope 按角色约束列表可见范围：
//   - 话务员/话务员管理员：只能看自己录的单（created_by = 当前用户）
//   - 收单员：只能看派给自己的单（agent_id = 当前用户）
//   - 其余（收单员管理员/财务文员/管理员）：全部
func applyScope(ctx context.Context, model *gdb.Model, role string) *gdb.Model {
	if canViewAll(role) {
		return model
	}
	userId := currentUserId(ctx)
	switch role {
	case consts.RoleTelemarketer, consts.RoleTelemarketerManager:
		return model.Where("t.created_by", userId)
	case consts.RoleAgent:
		return model.Where("t.agent_id", userId)
	}
	// 未匹配到已知角色：默认只能看自己录的单，保证数据不越权
	return model.Where("t.created_by", userId)
}

// canStep 判断角色是否允许执行某一步骤。
// step 对应订单流程步骤：1 录单、2 派单、3 预约、4 上门、5 回访、6 完成。
// 管理员/超管可执行任意步骤（含补录）。
func canStep(role string, step int) bool {
	if role == consts.RoleAdmin || consts.IsSuperRole(role) {
		return step >= 1 && step <= 6
	}
	switch step {
	case consts.OrderStatusRecorded: // 1 录单
		return role == consts.RoleTelemarketer || role == consts.RoleTelemarketerManager
	case consts.OrderStatusAssigned: // 2 派单
		return role == consts.RoleAgentManager
	case consts.OrderStatusAppointed: // 3 预约
		return role == consts.RoleAgent || role == consts.RoleAgentManager
	case consts.OrderStatusVisited: // 4 上门
		return role == consts.RoleAgent || role == consts.RoleAgentManager
	case consts.OrderStatusFollowed: // 5 回访
		return role == consts.RoleTelemarketer || role == consts.RoleTelemarketerManager
	case consts.OrderStatusCompleted: // 6 完成
		return role == consts.RoleDocumentary
	}
	return false
}

// canStepFromStatus 判断对处于 status 状态的订单执行 nextStep 步骤是否被当前角色允许。
// status 为订单当前状态（0-6）。允许两种情形：
//   - 原地保存：nextStep 与当前状态相同（如录单后修改录单信息）
//   - 正常推进：nextStep == status+1（1 已录单 → 2 已接单 → ... → 6 已完工）
func canStepFromStatus(role string, status int, nextStep int) bool {
	if nextStep == status {
		return canStep(role, nextStep)
	}
	expected := status + 1
	if nextStep != expected {
		return false
	}
	// 录单步骤：订单状态 0（未成交/初始）时方可执行录入
	if nextStep == consts.OrderStatusRecorded && status != consts.OrderStatusNotDeal {
		return false
	}
	return canStep(role, nextStep)
}

// canOperateOrder 判断当前角色是否有权操作该订单：
//   - 话务员/话务员管理员：只能操作自己录的单（created_by）
//   - 收单员：只能操作派给自己的单（agent_id）
//   - 其余（收单员管理员/财务文员/管理员/超管）：不限
func canOperateOrder(ctx context.Context, role string, createdBy int64, agentId int64) bool {
	if canViewAll(role) {
		return true
	}
	userId := currentUserId(ctx)
	switch role {
	case consts.RoleTelemarketer, consts.RoleTelemarketerManager:
		return createdBy == userId
	case consts.RoleAgent:
		return agentId == userId
	}
	// 未匹配到已知角色：不允许操作，保证数据不越权
	return false
}

// canAgentManagerAppointCollect 收单员管理员操作预约(3)/收单(4)时的收单员归属校验。
// 收单员管理员可查看全部订单，但预约/收单仅限收单员为自己的单；其余步骤（如排单）不受限。
func canAgentManagerAppointCollect(ctx context.Context, role string, agentId int64, step int) bool {
	if role != consts.RoleAgentManager {
		return true
	}
	switch step {
	case consts.OrderStatusAppointed, consts.OrderStatusVisited:
		return agentId == currentUserId(ctx)
	}
	return true
}

// canDelete 判断角色是否允许删除某状态的订单。
//   - 话务员/话务员管理员：仅可删除"已录单"状态的单
//   - 管理员/超管：任意状态
func canDelete(role string, status int) bool {
	if role == consts.RoleAdmin || consts.IsSuperRole(role) {
		return true
	}
	if (role == consts.RoleTelemarketer || role == consts.RoleTelemarketerManager) && status == consts.OrderStatusRecorded {
		return true
	}
	return false
}

// permissionDenied 统一的权限不足错误。
func permissionDenied(msg string) error {
	if msg == "" {
		msg = "您没有权限执行该操作"
	}
	return gerror.New(msg)
}
