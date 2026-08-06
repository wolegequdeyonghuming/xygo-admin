// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package adminin

import (
	"xygo/internal/model/input/form"
)

// DictTypeListInp 字典类型列表查询入参
type DictTypeListInp struct {
	form.PageReq
	Name   string `p:"name"   json:"name"   dc:"按名称模糊搜索"`
	Type   string `p:"type"   json:"type"   dc:"按标识模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// DictTypeListItem 字典类型列表项
type DictTypeListItem struct {
	Id         uint64 `json:"id"         dc:"字典类型ID"`
	Name       string `json:"name"       dc:"字典名称"`
	Type       string `json:"type"       dc:"字典标识"`
	Remark     string `json:"remark"     dc:"备注"`
	Status     int    `json:"status"     dc:"状态:0禁用,1启用"`
	Sort       int    `json:"sort"       dc:"排序"`
	CreateTime int    `json:"createTime" dc:"创建时间"`
	UpdateTime int    `json:"updateTime" dc:"更新时间"`
}

// DictTypeListModel 字典类型列表响应模型
type DictTypeListModel struct {
	List []DictTypeListItem `json:"list" dc:"数据列表"`
	form.PageRes
}

// DictTypeDetailInp 字典类型详情入参
type DictTypeDetailInp struct {
	Id uint `p:"id" v:"required#字典类型ID不能为空" json:"id" dc:"字典类型ID"`
}

// DictTypeSaveInp 字典类型新增/编辑入参
type DictTypeSaveInp struct {
	Id     uint64 `p:"id"     json:"id"     dc:"字典类型ID（为空表示新增）"`
	Name   string `p:"name"   v:"required#字典名称不能为空" json:"name"   dc:"字典名称"`
	Type   string `p:"type"   v:"required#字典标识不能为空" json:"type"   dc:"字典标识"`
	Sort   int    `p:"sort"   d:"0"         json:"sort"   dc:"排序"`
	Status int    `p:"status" d:"1"         json:"status" dc:"状态:0禁用,1启用"`
	Remark string `p:"remark"               json:"remark" dc:"备注"`
}

// DictTypeDeleteInp 字典类型删除入参
type DictTypeDeleteInp struct {
	Id uint64 `p:"id" v:"required#字典类型ID不能为空" json:"id" dc:"字典类型ID"`
}

// DictDataListInp 字典数据列表查询入参
type DictDataListInp struct {
	form.PageReq
	DictTypeId uint64 `p:"dictTypeId" v:"required#字典类型ID不能为空" json:"dictTypeId" dc:"字典类型ID"`
	Label      string `p:"label"      json:"label"      dc:"按标签模糊搜索"`
	Status     int    `p:"status"     d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// DictDataListItem 字典数据列表项
type DictDataListItem struct {
	Id         uint64 `json:"id"         dc:"字典数据ID"`
	DictTypeId uint64 `json:"dictTypeId" dc:"字典类型ID"`
	Label      string `json:"label"      dc:"字典标签"`
	Value      string `json:"value"      dc:"字典值"`
	CssClass   string `json:"cssClass"   dc:"样式类名"`
	ListClass  string `json:"listClass"  dc:"表格样式类名"`
	IsDefault  int    `json:"isDefault"  dc:"是否默认"`
	Status     int    `json:"status"     dc:"状态:0禁用,1启用"`
	Sort       int    `json:"sort"       dc:"排序"`
	Remark     string `json:"remark"     dc:"备注"`
	CreateTime int    `json:"createTime" dc:"创建时间"`
	UpdateTime int    `json:"updateTime" dc:"更新时间"`
}

// DictDataListModel 字典数据列表响应模型
type DictDataListModel struct {
	List []DictDataListItem `json:"list" dc:"数据列表"`
	form.PageRes
}

// DictDataSaveInp 字典数据新增/编辑入参
type DictDataSaveInp struct {
	Id         uint64 `p:"id"         json:"id"         dc:"字典数据ID（为空表示新增）"`
	DictTypeId uint64 `p:"dictTypeId" v:"required#字典类型ID不能为空" json:"dictTypeId" dc:"字典类型ID"`
	Label      string `p:"label"      v:"required#字典标签不能为空" json:"label"   dc:"字典标签"`
	Value      string `p:"value"      v:"required#字典值不能为空" json:"value"   dc:"字典值"`
	CssClass   string `p:"cssClass"   json:"cssClass"   dc:"样式类名"`
	ListClass  string `p:"listClass"  json:"listClass"  dc:"表格样式类名"`
	IsDefault  int    `p:"isDefault"  d:"0" json:"isDefault" dc:"是否默认"`
	Sort       int    `p:"sort"       d:"0" json:"sort"      dc:"排序"`
	Status     int    `p:"status"     d:"1" json:"status"    dc:"状态:0禁用,1启用"`
	Remark     string `p:"remark"     json:"remark"     dc:"备注"`
}

// DictDataDeleteInp 字典数据删除入参
type DictDataDeleteInp struct {
	Id uint64 `p:"id" v:"required#字典数据ID不能为空" json:"id" dc:"字典数据ID"`
}

// DictDataItem 对外公开的字典数据项（不含敏感字段）
type DictDataItem struct {
	Label     string `json:"label"     dc:"字典标签"`
	Value     string `json:"value"     dc:"字典值"`
	CssClass  string `json:"cssClass"  dc:"样式类名"`
	ListClass string `json:"listClass" dc:"表格样式类名"`
	IsDefault int    `json:"isDefault" dc:"是否默认"`
}

// DictTypeOptionsItem 字典类型选项（供下拉选择）
type DictTypeOptionsItem struct {
	Label string `json:"label" dc:"显示标签（字典名称）"`
	Value string `json:"value" dc:"值（字典标识type）"`
}
