// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	Id         uint64 `json:"id"         orm:"id"           description:""`                   //
	DictTypeId uint64 `json:"dictTypeId" orm:"dict_type_id" description:"关联字典类型ID"`           // 关联字典类型ID
	Label      string `json:"label"      orm:"label"        description:"字典标签"`               // 字典标签
	Value      string `json:"value"      orm:"value"        description:"字典值"`                // 字典值
	CssClass   string `json:"cssClass"   orm:"css_class"    description:"样式类名"`               // 样式类名
	ListClass  string `json:"listClass"  orm:"list_class"   description:"表格样式类名（ElTag type）"` // 表格样式类名（ElTag type）
	IsDefault  int    `json:"isDefault"  orm:"is_default"   description:"是否默认"`               // 是否默认
	Status     int    `json:"status"     orm:"status"       description:"状态：0=禁用 1=启用"`       // 状态：0=禁用 1=启用
	Sort       int    `json:"sort"       orm:"sort"         description:"排序"`                 // 排序
	Remark     string `json:"remark"     orm:"remark"       description:"备注"`                 // 备注
	CreatedBy  uint64 `json:"createdBy"  orm:"created_by"   description:"创建人ID"`              // 创建人ID
	UpdatedBy  uint64 `json:"updatedBy"  orm:"updated_by"   description:"更新人ID"`              // 更新人ID
	CreateTime uint64 `json:"createTime" orm:"create_time"  description:"创建时间（Unix秒）"`        // 创建时间（Unix秒）
	UpdateTime uint64 `json:"updateTime" orm:"update_time"  description:"更新时间（Unix秒）"`        // 更新时间（Unix秒）
}
