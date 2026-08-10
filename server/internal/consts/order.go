// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package consts

// ============================================
// 订单状态
// ============================================

// 订单状态（与字典 order_status 保持一致，业务代码直接使用常量判断）
const (
	OrderStatusNotDeal   = 0 // 未成交（终止态）
	OrderStatusRecorded  = 1 // 已录单
	OrderStatusAssigned  = 2 // 已接单
	OrderStatusAppointed = 3 // 已预约
	OrderStatusVisited   = 4 // 已上门
	OrderStatusFollowed  = 5 // 已回访
	OrderStatusCompleted = 6 // 已完工
)

// ============================================
// 订单相关角色 key（集中管理，便于调整）
// ============================================

const (
	RoleAdmin               = "admin"                // 管理员
	RoleTelemarketerManager = "telemarketer_manager" // 话务员管理员
	RoleAgentManager        = "agent_manager"        // 收单员管理员
	RoleDocumentary         = "documentary"          // 财务文员
	RoleTelemarketer        = "telemarketer"         // 话务员
	RoleAgent               = "agent"                // 收单员
)
