// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDataDao is the data access object for the table xy_sys_dict_data.
type SysDictDataDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysDictDataColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysDictDataColumns defines and stores column names for the table xy_sys_dict_data.
type SysDictDataColumns struct {
	Id         string // 主键
	DictTypeId string // 关联字典类型ID
	Label      string // 字典标签
	Value      string // 字典值
	CssClass   string // 样式类名
	ListClass  string // 表格样式类名（ElTag type）
	IsDefault  string // 是否默认
	Status     string // 状态：0=禁用 1=启用
	Sort       string // 排序
	Remark     string // 备注
	CreatedBy  string // 创建人ID
	UpdatedBy  string // 更新人ID
	CreateTime string // 创建时间（Unix秒）
	UpdateTime string // 更新时间（Unix秒）
}

// sysDictDataColumns holds the columns for the table xy_sys_dict_data.
var sysDictDataColumns = SysDictDataColumns{
	Id:         "id",
	DictTypeId: "dict_type_id",
	Label:      "label",
	Value:      "value",
	CssClass:   "css_class",
	ListClass:  "list_class",
	IsDefault:  "is_default",
	Status:     "status",
	Sort:       "sort",
	Remark:     "remark",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysDictDataDao creates and returns a new DAO object for table data access.
func NewSysDictDataDao(handlers ...gdb.ModelHandler) *SysDictDataDao {
	return &SysDictDataDao{
		group:    "default",
		table:    "xy_sys_dict_data",
		columns:  sysDictDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysDictDataDao) Columns() SysDictDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysDictDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
