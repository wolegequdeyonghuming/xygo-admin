// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/adminin"
)

// DictTypeListReq 字典类型列表请求
type DictTypeListReq struct {
	g.Meta `path:"/admin/dict/type/list" method:"get" tags:"AdminDict" summary:"字典类型列表"`
	adminin.DictTypeListInp
}

// DictTypeListRes 字典类型列表响应
type DictTypeListRes struct {
	adminin.DictTypeListModel
}

// DictTypeDetailReq 字典类型详情请求
type DictTypeDetailReq struct {
	g.Meta `path:"/admin/dict/type/detail" method:"get" tags:"AdminDict" summary:"字典类型详情"`
	adminin.DictTypeDetailInp
}

// DictTypeDetailRes 字典类型详情响应
type DictTypeDetailRes struct {
	adminin.DictTypeListItem
}

// DictTypeSaveReq 字典类型保存请求
type DictTypeSaveReq struct {
	g.Meta `path:"/admin/dict/type/save" method:"post" tags:"AdminDict" summary:"字典类型保存"`
	adminin.DictTypeSaveInp
}

// DictTypeSaveRes 字典类型保存响应
type DictTypeSaveRes struct {
	Id uint `json:"id" dc:"字典类型ID"`
}

// DictTypeDeleteReq 字典类型删除请求
type DictTypeDeleteReq struct {
	g.Meta `path:"/admin/dict/type/delete" method:"post" tags:"AdminDict" summary:"字典类型删除"`
	adminin.DictTypeDeleteInp
}

// DictTypeDeleteRes 字典类型删除响应
type DictTypeDeleteRes struct{}

// DictTypeOptionsReq 字典类型选项（供下拉选择）请求
type DictTypeOptionsReq struct {
	g.Meta `path:"/admin/dict/type/options" method:"get" tags:"AdminDict" summary:"字典类型选项"`
}

// DictTypeOptionsRes 字典类型选项响应
type DictTypeOptionsRes struct {
	List []adminin.DictTypeOptionsItem `json:"list" dc:"选项列表"`
}

// DictDataListReq 字典数据列表请求
type DictDataListReq struct {
	g.Meta `path:"/admin/dict/data/list" method:"get" tags:"AdminDict" summary:"字典数据列表"`
	adminin.DictDataListInp
}

// DictDataListRes 字典数据列表响应
type DictDataListRes struct {
	adminin.DictDataListModel
}

// DictDataSaveReq 字典数据保存请求
type DictDataSaveReq struct {
	g.Meta `path:"/admin/dict/data/save" method:"post" tags:"AdminDict" summary:"字典数据保存"`
	adminin.DictDataSaveInp
}

// DictDataSaveRes 字典数据保存响应
type DictDataSaveRes struct {
	Id uint `json:"id" dc:"字典数据ID"`
}

// DictDataDeleteReq 字典数据删除请求
type DictDataDeleteReq struct {
	g.Meta `path:"/admin/dict/data/delete" method:"post" tags:"AdminDict" summary:"字典数据删除"`
	adminin.DictDataDeleteInp
}

// DictDataDeleteRes 字典数据删除响应
type DictDataDeleteRes struct{}
