// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	adminin "xygo/internal/model/input/orderin"
)

type (
	IBizOrder interface {
		// List 订单主表列表
		List(ctx context.Context, in *adminin.BizOrderListInp) (*adminin.BizOrderListModel, error)
		// View 订单主表详情
		View(ctx context.Context, id uint64) (*adminin.BizOrderViewModel, error)
		// Edit 保存订单主表
		Edit(ctx context.Context, in *adminin.BizOrderEditInp) error
		// StepEdit 订单分步保存：按步骤只更新本步骤字段，并校验角色与当前状态
		// 返回更新后的订单 ID（录单新增时返回新生成的主键）。
		StepEdit(ctx context.Context, in *adminin.BizOrderStepEditInp) (id uint64, err error)
		// StepNext 推进订单到下一步状态（仅更新 order_status，不保存任何字段）
		StepNext(ctx context.Context, id uint64) error
		// Delete 删除订单主表（逻辑删除，按角色+状态限制）
		Delete(ctx context.Context, id uint64) error
		// DetailList 订单详细情况列表：含操作人姓名，时间正序
		DetailList(ctx context.Context, orderId int64) (*adminin.BizOrderDetailListModel, error)
		// DetailAdd 新增订单详细情况：user_id 取当前登录用户
		DetailAdd(ctx context.Context, in *adminin.BizOrderDetailAddInp) error
		// DetailDelete 删除订单详细情况：只能删除自己录入的（逻辑删除）
		DetailDelete(ctx context.Context, id int64) error
	}
)

var (
	localBizOrder IBizOrder
)

func BizOrder() IBizOrder {
	if localBizOrder == nil {
		panic("implement not found for interface IBizOrder, forgot register?")
	}
	return localBizOrder
}

func RegisterBizOrder(i IBizOrder) {
	localBizOrder = i
}
