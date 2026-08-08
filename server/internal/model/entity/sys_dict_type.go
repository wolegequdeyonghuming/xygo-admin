// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	Id         uint64 `json:"id"         orm:"id"          description:"主键"`           // 主键
	Name       string `json:"name"       orm:"name"        description:"字典名称"`         // 字典名称
	Type       string `json:"type"       orm:"type"        description:"字典标识"`         // 字典标识
	Remark     string `json:"remark"     orm:"remark"      description:"备注"`           // 备注
	Status     int    `json:"status"     orm:"status"      description:"状态：0=禁用 1=启用"` // 状态：0=禁用 1=启用
	Sort       int    `json:"sort"       orm:"sort"        description:"排序"`           // 排序
	CreatedBy  uint64 `json:"createdBy"  orm:"created_by"  description:"创建人ID"`        // 创建人ID
	UpdatedBy  uint64 `json:"updatedBy"  orm:"updated_by"  description:"更新人ID"`        // 更新人ID
	CreateTime uint64 `json:"createTime" orm:"create_time" description:"创建时间（Unix秒）"`  // 创建时间（Unix秒）
	UpdateTime uint64 `json:"updateTime" orm:"update_time" description:"更新时间（Unix秒）"`  // 更新时间（Unix秒）
}
