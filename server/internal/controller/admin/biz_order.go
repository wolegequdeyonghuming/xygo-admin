package admin

import (
	"context"

	api "xygo/api/order"
	"xygo/internal/service"
)

// BizOrderList 订单主表列表
func (c *ControllerV1) BizOrderList(ctx context.Context, req *api.BizOrderListReq) (res *api.BizOrderListRes, err error) {
	result, err := service.BizOrder().List(ctx, &req.BizOrderListInp)
	if err != nil {
		return nil, err
	}
	return &api.BizOrderListRes{result}, nil
}

// BizOrderView 订单主表详情
func (c *ControllerV1) BizOrderView(ctx context.Context, req *api.BizOrderViewReq) (res *api.BizOrderViewRes, err error) {
	result, err := service.BizOrder().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.BizOrderViewRes{result}, nil
}

// BizOrderEdit 保存订单主表
func (c *ControllerV1) BizOrderEdit(ctx context.Context, req *api.BizOrderEditReq) (res *api.BizOrderEditRes, err error) {
	err = service.BizOrder().Edit(ctx, &req.BizOrderEditInp)
	return &api.BizOrderEditRes{}, err
}

// BizOrderDelete 删除订单主表
func (c *ControllerV1) BizOrderDelete(ctx context.Context, req *api.BizOrderDeleteReq) (res *api.BizOrderDeleteRes, err error) {
	err = service.BizOrder().Delete(ctx, req.Id)
	return &api.BizOrderDeleteRes{}, err
}

// BizOrderStepEdit 订单分步保存
func (c *ControllerV1) BizOrderStepEdit(ctx context.Context, req *api.BizOrderStepEditReq) (res *api.BizOrderStepEditRes, err error) {
	id, err := service.BizOrder().StepEdit(ctx, &req.BizOrderStepEditInp)
	if err != nil {
		return nil, err
	}
	return &api.BizOrderStepEditRes{Id: id}, nil
}

// BizOrderStepNext 订单推进下一步（仅改变状态）
func (c *ControllerV1) BizOrderStepNext(ctx context.Context, req *api.BizOrderStepNextReq) (res *api.BizOrderStepNextRes, err error) {
	err = service.BizOrder().StepNext(ctx, req.Id)
	return &api.BizOrderStepNextRes{}, err
}

// BizOrderDetailList 订单详细情况列表
func (c *ControllerV1) BizOrderDetailList(ctx context.Context, req *api.BizOrderDetailListReq) (res *api.BizOrderDetailListRes, err error) {
	result, err := service.BizOrder().DetailList(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &api.BizOrderDetailListRes{result}, nil
}

// BizOrderDetailAdd 新增订单详细情况
func (c *ControllerV1) BizOrderDetailAdd(ctx context.Context, req *api.BizOrderDetailAddReq) (res *api.BizOrderDetailAddRes, err error) {
	err = service.BizOrder().DetailAdd(ctx, &req.BizOrderDetailAddInp)
	return &api.BizOrderDetailAddRes{}, err
}

// BizOrderDetailDelete 删除订单详细情况
func (c *ControllerV1) BizOrderDetailDelete(ctx context.Context, req *api.BizOrderDetailDeleteReq) (res *api.BizOrderDetailDeleteRes, err error) {
	err = service.BizOrder().DetailDelete(ctx, req.Id)
	return &api.BizOrderDetailDeleteRes{}, err
}
