// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BizOrderDetails is the golang structure for table biz_order_details.
type BizOrderDetails struct {
	Id        int64  `json:"id"        orm:"id"         description:"主键"`   // 主键
	OrderId   int64  `json:"orderId"   orm:"order_id"   description:"订单id"` // 订单id
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"修改时间"` // 修改时间
	UserId    int64  `json:"userId"    orm:"user_id"    description:"操作人"`  // 操作人
	Content   string `json:"content"   orm:"content"    description:"内容"`   // 内容
	IsDeleted int    `json:"isDeleted" orm:"is_deleted" description:"是否删除"` // 是否删除
}
