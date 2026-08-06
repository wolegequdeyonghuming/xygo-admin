// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BizOrderDetails is the golang structure of table xy_biz_order_details for DAO operations like Where/Data.
type BizOrderDetails struct {
	g.Meta    `orm:"table:xy_biz_order_details, do:true"`
	Id        any // 主键
	OrderId   any // 订单id
	CreatedAt any // 创建时间
	UpdatedAt any // 修改时间
	UserId    any // 操作人
	Content   any // 内容
}
