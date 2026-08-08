package order

import (
	"github.com/gogf/gf/v2/frame/g"
	adminin "xygo/internal/model/input/orderin"
)

// BizOrderListReq 订单主表列表请求
type BizOrderListReq struct {
	g.Meta `path:"/admin/biz-order/list" method:"get" tags:"BizOrder" summary:"订单主表列表"`
	adminin.BizOrderListInp
}

type BizOrderListRes struct {
	*adminin.BizOrderListModel
}

// BizOrderViewReq 订单主表详情请求
type BizOrderViewReq struct {
	g.Meta `path:"/admin/biz-order/view" method:"get" tags:"BizOrder" summary:"订单主表详情"`
	Id     uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type BizOrderViewRes struct {
	*adminin.BizOrderViewModel
}

// BizOrderEditReq 订单主表保存请求
type BizOrderEditReq struct {
	g.Meta `path:"/admin/biz-order/edit" method:"post" tags:"BizOrder" summary:"保存订单主表"`
	adminin.BizOrderEditInp
}

type BizOrderEditRes struct{}

// BizOrderDeleteReq 订单主表删除请求
type BizOrderDeleteReq struct {
	g.Meta `path:"/admin/biz-order/delete" method:"post" tags:"BizOrder" summary:"删除订单主表"`
	Id     uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type BizOrderDeleteRes struct{}
