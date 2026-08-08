// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenCodesColumnDao is the data access object for the table xy_sys_gen_codes_column.
type SysGenCodesColumnDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  SysGenCodesColumnColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// SysGenCodesColumnColumns defines and stores column names for the table xy_sys_gen_codes_column.
type SysGenCodesColumnColumns struct {
	Id              string // 主键
	GenId           string // 关联gen_codes.id
	Name            string // 数据库字段名
	GoName          string // Go字段名
	TsName          string // TS字段名
	DbType          string // 数据库类型
	GoType          string // Go类型
	TsType          string // TS类型
	Comment         string // 字段注释
	IsPk            string // 是否主键:0=否,1=是
	IsRequired      string // 是否必填:0=否,1=是
	IsList          string // 表格列显示:0=否,1=是
	IsEdit          string // 表单编辑显示:0=否,1=是
	IsQuery         string // 搜索条件:0=否,1=是
	QueryType       string // 查询方式:eq/like/between/gt/in
	DesignType      string // 设计类型(designType)
	Extra           string // 扩展配置JSON(关联表配置/表格属性/表单属性等)
	FormType        string // 表单组件类型
	DictType        string // 关联字典
	IsAutoIncrement string // 是否自增
	IsNullable      string // 是否可空
	Sort            string // 排序
}

// sysGenCodesColumnColumns holds the columns for the table xy_sys_gen_codes_column.
var sysGenCodesColumnColumns = SysGenCodesColumnColumns{
	Id:              "id",
	GenId:           "gen_id",
	Name:            "name",
	GoName:          "go_name",
	TsName:          "ts_name",
	DbType:          "db_type",
	GoType:          "go_type",
	TsType:          "ts_type",
	Comment:         "comment",
	IsPk:            "is_pk",
	IsRequired:      "is_required",
	IsList:          "is_list",
	IsEdit:          "is_edit",
	IsQuery:         "is_query",
	QueryType:       "query_type",
	DesignType:      "design_type",
	Extra:           "extra",
	FormType:        "form_type",
	DictType:        "dict_type",
	IsAutoIncrement: "is_auto_increment",
	IsNullable:      "is_nullable",
	Sort:            "sort",
}

// NewSysGenCodesColumnDao creates and returns a new DAO object for table data access.
func NewSysGenCodesColumnDao(handlers ...gdb.ModelHandler) *SysGenCodesColumnDao {
	return &SysGenCodesColumnDao{
		group:    "default",
		table:    "xy_sys_gen_codes_column",
		columns:  sysGenCodesColumnColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysGenCodesColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysGenCodesColumnDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysGenCodesColumnDao) Columns() SysGenCodesColumnColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysGenCodesColumnDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysGenCodesColumnDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysGenCodesColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
