// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysQueueConfigDao is the data access object for the table xy_sys_queue_config.
type SysQueueConfigDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysQueueConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysQueueConfigColumns defines and stores column names for the table xy_sys_queue_config.
type SysQueueConfigColumns struct {
	Id            string // ID
	Topic         string // Topic 标识（唯一）
	Title         string // 显示名称
	Workers       string // 并行 Worker 数
	MaxRetry      string // 最大重试次数
	RetryDelaySec string // 重试间隔（秒，0=立即重试）
	Status        string // 状态:0禁用,1启用
	Remark        string // 备注
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// sysQueueConfigColumns holds the columns for the table xy_sys_queue_config.
var sysQueueConfigColumns = SysQueueConfigColumns{
	Id:            "id",
	Topic:         "topic",
	Title:         "title",
	Workers:       "workers",
	MaxRetry:      "max_retry",
	RetryDelaySec: "retry_delay_sec",
	Status:        "status",
	Remark:        "remark",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysQueueConfigDao creates and returns a new DAO object for table data access.
func NewSysQueueConfigDao(handlers ...gdb.ModelHandler) *SysQueueConfigDao {
	return &SysQueueConfigDao{
		group:    "default",
		table:    "xy_sys_queue_config",
		columns:  sysQueueConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysQueueConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysQueueConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysQueueConfigDao) Columns() SysQueueConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysQueueConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysQueueConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysQueueConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
