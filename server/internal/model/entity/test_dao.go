// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TestDao is the golang structure for table test_dao.
type TestDao struct {
	Id        uint64 `json:"id"        orm:"id"         description:"主键"`   // 主键
	Title     string `json:"title"     orm:"title"      description:"标题"`   // 标题
	Content   string `json:"content"   orm:"content"    description:"内容"`   // 内容
	Sort      int    `json:"sort"      orm:"sort"       description:"排序"`   // 排序
	Status    int    `json:"status"    orm:"status"     description:"状态"`   // 状态
	CreatedAt uint64 `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt uint64 `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
