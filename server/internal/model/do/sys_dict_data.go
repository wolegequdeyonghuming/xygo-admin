// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictData is the golang structure of table xy_sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta     `orm:"table:xy_sys_dict_data, do:true"`
	Id         any // 主键
	DictTypeId any // 关联字典类型ID
	Label      any // 字典标签
	Value      any // 字典值
	CssClass   any // 样式类名
	ListClass  any // 表格样式类名（ElTag type）
	IsDefault  any // 是否默认
	Status     any // 状态：0=禁用 1=启用
	Sort       any // 排序
	Remark     any // 备注
	CreatedBy  any // 创建人ID
	UpdatedBy  any // 更新人ID
	CreateTime any // 创建时间（Unix秒）
	UpdateTime any // 更新时间（Unix秒）
}
