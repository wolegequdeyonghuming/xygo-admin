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
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/orderin"
	"xygo/internal/model/input/staffin"
)

// OrderListReq 收单员订单列表请求
type OrderListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"Staff" summary:"收单员订单列表"`
	staffin.StaffOrderListInp
}

type OrderListRes struct {
	*orderin.BizOrderListModel
}

// OrderStatReq 收单员订单统计请求
type OrderStatReq struct {
	g.Meta `path:"/order/stat" method:"get" tags:"Staff" summary:"收单员订单统计"`
}

type OrderStatRes struct {
	*staffin.StaffOrderStatModel
}

// OrderViewReq 订单详情请求
type OrderViewReq struct {
	g.Meta `path:"/order/view" method:"get" tags:"Staff" summary:"订单详情"`
	Id     uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type OrderViewRes struct {
	*orderin.BizOrderViewModel
}

// OrderAppointReq 预约请求
type OrderAppointReq struct {
	g.Meta `path:"/order/appoint" method:"post" tags:"Staff" summary:"填写预约情况（步骤3）"`
	staffin.StaffOrderAppointInp
}

type OrderAppointRes struct{}

// OrderCollectReq 收单请求
type OrderCollectReq struct {
	g.Meta `path:"/order/collect" method:"post" tags:"Staff" summary:"填写收单情况（步骤4）"`
	staffin.StaffOrderCollectInp
}

type OrderCollectRes struct{}

// OrderScheduleReq 排单（派单）请求
type OrderScheduleReq struct {
	g.Meta `path:"/order/schedule" method:"post" tags:"Staff" summary:"排单派单（步骤2）"`
	staffin.StaffOrderScheduleInp
}

type OrderScheduleRes struct{}

// OrderCompleteReq 完工请求
type OrderCompleteReq struct {
	g.Meta `path:"/order/complete" method:"post" tags:"Staff" summary:"完工（步骤6）"`
	staffin.StaffOrderCompleteInp
}

type OrderCompleteRes struct{}

// UserAgentListReq 收单员列表请求（排单选人）
type UserAgentListReq struct {
	g.Meta `path:"/user/agents" method:"get" tags:"Staff" summary:"收单员列表"`
}

type UserAgentListRes struct {
	*staffin.StaffUserAgentListModel
}

// OrderDetailListReq 订单详细情况列表请求
type OrderDetailListReq struct {
	g.Meta `path:"/order/detailList" method:"get" tags:"Staff" summary:"订单详细情况列表"`
	orderin.BizOrderDetailListInp
}

type OrderDetailListRes struct {
	*orderin.BizOrderDetailListModel
}

// OrderDetailAddReq 新增订单详细情况请求
type OrderDetailAddReq struct {
	g.Meta `path:"/order/detailAdd" method:"post" tags:"Staff" summary:"新增订单详细情况"`
	orderin.BizOrderDetailAddInp
}

type OrderDetailAddRes struct{}

// OrderDetailDeleteReq 删除订单详细情况请求
type OrderDetailDeleteReq struct {
	g.Meta `path:"/order/detailDelete" method:"post" tags:"Staff" summary:"删除订单详细情况"`
	orderin.BizOrderDetailDeleteInp
}

type OrderDetailDeleteRes struct{}

// AttachmentListReq 附件信息列表请求
type AttachmentListReq struct {
	g.Meta `path:"/attachment/list" method:"get" tags:"Staff" summary:"按附件ID列表获取URL"`
	Ids    string `json:"ids" dc:"附件ID列表（逗号分隔）"`
}

type AttachmentListRes struct {
	*staffin.StaffAttachmentListModel
}
