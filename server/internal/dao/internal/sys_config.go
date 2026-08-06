// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigDao is the data access object for the table xy_sys_config.
type SysConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysConfigColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysConfigColumns defines and stores column names for the table xy_sys_config.
type SysConfigColumns struct {
	Id         string // 主键
	Group      string // 分组标识，如 basics/mail/oss
	GroupName  string // 分组名称
	Name       string // 配置项显示名
	Key        string // 配置键（唯一）
	Value      string // 配置值（字符串/JSON）
	Type       string // 控件类型:text/textarea/number/switch/select/radio/checkbox/json/object/array/color/upload
	Options    string // 组件参数/选项 JSON
	Rules      string // 校验规则 JSON
	Sort       string // 排序(大靠后)
	Remark     string // 备注
	AllowDel   string // 允许删除:0=否,1=是
	CreatedBy  string // 创建人
	UpdatedBy  string // 更新人
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	TestField  string // 测试字段（v1.2.7）
}

// sysConfigColumns holds the columns for the table xy_sys_config.
var sysConfigColumns = SysConfigColumns{
	Id:         "id",
	Group:      "group",
	GroupName:  "group_name",
	Name:       "name",
	Key:        "key",
	Value:      "value",
	Type:       "type",
	Options:    "options",
	Rules:      "rules",
	Sort:       "sort",
	Remark:     "remark",
	AllowDel:   "allow_del",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	TestField:  "test_field",
}

// NewSysConfigDao creates and returns a new DAO object for table data access.
func NewSysConfigDao(handlers ...gdb.ModelHandler) *SysConfigDao {
	return &SysConfigDao{
		group:    "default",
		table:    "xy_sys_config",
		columns:  sysConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysConfigDao) Columns() SysConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
