package service

import (
	"context"

	"xygo/internal/model/input/orderin"
	"xygo/internal/model/input/staffin"
)

// IStaffOrder 收单小程序订单服务
type IStaffOrder interface {
	// List 收单员订单列表（收单员只看派给自己的单）
	List(ctx context.Context, in *staffin.StaffOrderListInp) (*orderin.BizOrderListModel, error)
	// Stat 收单员订单统计（待处理/今日已处理/已处理）
	Stat(ctx context.Context) (*staffin.StaffOrderStatModel, error)
	// Appoint 填写预约情况（步骤3 保存并推进）
	Appoint(ctx context.Context, in *staffin.StaffOrderAppointInp) error
	// Collect 填写收单情况（步骤4 保存；finish=true 时推进到已上门）
	Collect(ctx context.Context, in *staffin.StaffOrderCollectInp) error
	// AttachmentList 按附件ID列表获取附件 URL
	AttachmentList(ctx context.Context, ids string) (*staffin.StaffAttachmentListModel, error)
}

var localStaffOrder IStaffOrder

func StaffOrder() IStaffOrder {
	if localStaffOrder == nil {
		panic("implement not found for interface IStaffOrder, forgot register?")
	}
	return localStaffOrder
}

func RegisterStaffOrder(i IStaffOrder) {
	localStaffOrder = i
}
