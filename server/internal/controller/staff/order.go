// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package staff

import (
	"context"

	api "xygo/api/staff"
	"xygo/internal/service"
)

// OrderList 收单员订单列表
func (c *ControllerV1) OrderList(ctx context.Context, req *api.OrderListReq) (res *api.OrderListRes, err error) {
	result, err := service.StaffOrder().List(ctx, &req.StaffOrderListInp)
	if err != nil {
		return nil, err
	}
	return &api.OrderListRes{result}, nil
}

// OrderStat 收单员订单统计
func (c *ControllerV1) OrderStat(ctx context.Context, req *api.OrderStatReq) (res *api.OrderStatRes, err error) {
	result, err := service.StaffOrder().Stat(ctx)
	if err != nil {
		return nil, err
	}
	return &api.OrderStatRes{result}, nil
}

// OrderView 订单详情
func (c *ControllerV1) OrderView(ctx context.Context, req *api.OrderViewReq) (res *api.OrderViewRes, err error) {
	result, err := service.BizOrder().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.OrderViewRes{result}, nil
}

// OrderAppoint 填写预约情况
func (c *ControllerV1) OrderAppoint(ctx context.Context, req *api.OrderAppointReq) (res *api.OrderAppointRes, err error) {
	err = service.StaffOrder().Appoint(ctx, &req.StaffOrderAppointInp)
	return &api.OrderAppointRes{}, err
}

// OrderCollect 填写收单情况
func (c *ControllerV1) OrderCollect(ctx context.Context, req *api.OrderCollectReq) (res *api.OrderCollectRes, err error) {
	err = service.StaffOrder().Collect(ctx, &req.StaffOrderCollectInp)
	return &api.OrderCollectRes{}, err
}

// OrderSchedule 排单（步骤2）
func (c *ControllerV1) OrderSchedule(ctx context.Context, req *api.OrderScheduleReq) (res *api.OrderScheduleRes, err error) {
	err = service.StaffOrder().Schedule(ctx, &req.StaffOrderScheduleInp)
	return &api.OrderScheduleRes{}, err
}

// OrderComplete 完工（步骤6）
func (c *ControllerV1) OrderComplete(ctx context.Context, req *api.OrderCompleteReq) (res *api.OrderCompleteRes, err error) {
	err = service.StaffOrder().Complete(ctx, &req.StaffOrderCompleteInp)
	return &api.OrderCompleteRes{}, err
}

// UserAgentList 收单员列表
func (c *ControllerV1) UserAgentList(ctx context.Context, req *api.UserAgentListReq) (res *api.UserAgentListRes, err error) {
	result, err := service.StaffOrder().UserAgents(ctx)
	if err != nil {
		return nil, err
	}
	return &api.UserAgentListRes{result}, nil
}

// OrderDetailList 订单详细情况列表
func (c *ControllerV1) OrderDetailList(ctx context.Context, req *api.OrderDetailListReq) (res *api.OrderDetailListRes, err error) {
	result, err := service.BizOrder().DetailList(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &api.OrderDetailListRes{result}, nil
}

// OrderDetailAdd 新增订单详细情况
func (c *ControllerV1) OrderDetailAdd(ctx context.Context, req *api.OrderDetailAddReq) (res *api.OrderDetailAddRes, err error) {
	err = service.BizOrder().DetailAdd(ctx, &req.BizOrderDetailAddInp)
	return &api.OrderDetailAddRes{}, err
}

// OrderDetailDelete 删除订单详细情况
func (c *ControllerV1) OrderDetailDelete(ctx context.Context, req *api.OrderDetailDeleteReq) (res *api.OrderDetailDeleteRes, err error) {
	err = service.BizOrder().DetailDelete(ctx, req.Id)
	return &api.OrderDetailDeleteRes{}, err
}

// AttachmentList 附件信息列表
func (c *ControllerV1) AttachmentList(ctx context.Context, req *api.AttachmentListReq) (res *api.AttachmentListRes, err error) {
	result, err := service.StaffOrder().AttachmentList(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &api.AttachmentListRes{result}, nil
}
