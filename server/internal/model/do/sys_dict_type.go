// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictType is the golang structure of table xy_sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta     `orm:"table:xy_sys_dict_type, do:true"`
	Id         any //
	Name       any // 字典名称
	Type       any // 字典标识
	Remark     any // 备注
	Status     any // 状态：0=禁用 1=启用
	Sort       any // 排序
	CreatedBy  any // 创建人ID
	UpdatedBy  any // 更新人ID
	CreateTime any // 创建时间（Unix秒）
	UpdateTime any // 更新时间（Unix秒）
}
