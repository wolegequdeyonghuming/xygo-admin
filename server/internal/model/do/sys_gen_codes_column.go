// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCodesColumn is the golang structure of table xy_sys_gen_codes_column for DAO operations like Where/Data.
type SysGenCodesColumn struct {
	g.Meta          `orm:"table:xy_sys_gen_codes_column, do:true"`
	Id              any // 主键
	GenId           any // 关联gen_codes.id
	Name            any // 数据库字段名
	GoName          any // Go字段名
	TsName          any // TS字段名
	DbType          any // 数据库类型
	GoType          any // Go类型
	TsType          any // TS类型
	Comment         any // 字段注释
	IsPk            any // 是否主键:0=否,1=是
	IsRequired      any // 是否必填:0=否,1=是
	IsList          any // 表格列显示:0=否,1=是
	IsEdit          any // 表单编辑显示:0=否,1=是
	IsQuery         any // 搜索条件:0=否,1=是
	QueryType       any // 查询方式:eq/like/between/gt/in
	DesignType      any // 设计类型(designType)
	Extra           any // 扩展配置JSON(关联表配置/表格属性/表单属性等)
	FormType        any // 表单组件类型
	DictType        any // 关联字典
	IsAutoIncrement any // 是否自增
	IsNullable      any // 是否可空
	Sort            any // 排序
}
