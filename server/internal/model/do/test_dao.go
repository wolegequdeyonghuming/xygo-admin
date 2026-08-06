// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TestDao is the golang structure of table xy_test_dao for DAO operations like Where/Data.
type TestDao struct {
	g.Meta    `orm:"table:xy_test_dao, do:true"`
	Id        any // 主键
	Title     any // 标题
	Content   any // 内容
	Sort      any // 排序
	Status    any // 状态
	CreatedAt any // 创建时间
	UpdatedAt any // 更新时间
}
