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

// BizOrderStepEditReq 订单主表分步保存请求
type BizOrderStepEditReq struct {
	g.Meta `path:"/admin/biz-order/stepEdit" method:"post" tags:"BizOrder" summary:"订单分步保存"`
	adminin.BizOrderStepEditInp
}

type BizOrderStepEditRes struct {
	Id uint64 `json:"id" dc:"订单ID（录单新增时返回）"`
}

// BizOrderStepNextReq 订单推进下一步请求
type BizOrderStepNextReq struct {
	g.Meta `path:"/admin/biz-order/stepNext" method:"post" tags:"BizOrder" summary:"订单推进下一步"`
	Id     uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type BizOrderStepNextRes struct{}

// BizOrderDetailListReq 订单详细情况列表请求
type BizOrderDetailListReq struct {
	g.Meta `path:"/admin/biz-order/detailList" method:"get" tags:"BizOrder" summary:"订单详细情况列表"`
	adminin.BizOrderDetailListInp
}

type BizOrderDetailListRes struct {
	*adminin.BizOrderDetailListModel
}

// BizOrderDetailAddReq 新增订单详细情况请求
type BizOrderDetailAddReq struct {
	g.Meta `path:"/admin/biz-order/detailAdd" method:"post" tags:"BizOrder" summary:"新增订单详细情况"`
	adminin.BizOrderDetailAddInp
}

type BizOrderDetailAddRes struct{}

// BizOrderDetailDeleteReq 删除订单详细情况请求
type BizOrderDetailDeleteReq struct {
	g.Meta `path:"/admin/biz-order/detailDelete" method:"post" tags:"BizOrder" summary:"删除订单详细情况"`
	adminin.BizOrderDetailDeleteInp
}

type BizOrderDetailDeleteRes struct{}

// BizOrderExportReq 订单导出请求（复用列表筛选条件）
type BizOrderExportReq struct {
	g.Meta `path:"/admin/biz-order/export" method:"get" tags:"BizOrder" summary:"订单导出"`
	adminin.BizOrderListInp
}

type BizOrderExportRes struct{}
