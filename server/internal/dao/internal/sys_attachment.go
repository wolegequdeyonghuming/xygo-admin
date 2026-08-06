// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentDao is the data access object for the table xy_sys_attachment.
type SysAttachmentDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysAttachmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysAttachmentColumns defines and stores column names for the table xy_sys_attachment.
type SysAttachmentColumns struct {
	Id         string // ID
	Topic      string // 分组/主题标识
	UserId     string // 上传用户ID（预留）
	Url        string // 物理路径（相对路径）
	Width      string // 宽度
	Height     string // 高度
	Name       string // 原始名称
	Size       string // 大小（字节）
	Mimetype   string // MIME类型
	Quote      string // 上传(引用)次数
	Storage    string // 存储方式：local=本地
	Sha1       string // sha1摘要
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	TenantId   string // 所属租户ID（0=平台）
}

// sysAttachmentColumns holds the columns for the table xy_sys_attachment.
var sysAttachmentColumns = SysAttachmentColumns{
	Id:         "id",
	Topic:      "topic",
	UserId:     "user_id",
	Url:        "url",
	Width:      "width",
	Height:     "height",
	Name:       "name",
	Size:       "size",
	Mimetype:   "mimetype",
	Quote:      "quote",
	Storage:    "storage",
	Sha1:       "sha1",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	TenantId:   "tenant_id",
}

// NewSysAttachmentDao creates and returns a new DAO object for table data access.
func NewSysAttachmentDao(handlers ...gdb.ModelHandler) *SysAttachmentDao {
	return &SysAttachmentDao{
		group:    "default",
		table:    "xy_sys_attachment",
		columns:  sysAttachmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAttachmentDao) Columns() SysAttachmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
