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
