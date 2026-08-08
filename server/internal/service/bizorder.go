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
		// Delete 删除订单主表
		Delete(ctx context.Context, id uint64) error
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
