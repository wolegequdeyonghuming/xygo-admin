// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BizOrderDetailsDao is the data access object for the table xy_biz_order_details.
type BizOrderDetailsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BizOrderDetailsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BizOrderDetailsColumns defines and stores column names for the table xy_biz_order_details.
type BizOrderDetailsColumns struct {
	Id        string // 主键
	OrderId   string // 订单id
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	UserId    string // 操作人
	Content   string // 内容
}

// bizOrderDetailsColumns holds the columns for the table xy_biz_order_details.
var bizOrderDetailsColumns = BizOrderDetailsColumns{
	Id:        "id",
	OrderId:   "order_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	UserId:    "user_id",
	Content:   "content",
}

// NewBizOrderDetailsDao creates and returns a new DAO object for table data access.
func NewBizOrderDetailsDao(handlers ...gdb.ModelHandler) *BizOrderDetailsDao {
	return &BizOrderDetailsDao{
		group:    "default",
		table:    "xy_biz_order_details",
		columns:  bizOrderDetailsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BizOrderDetailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BizOrderDetailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BizOrderDetailsDao) Columns() BizOrderDetailsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BizOrderDetailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BizOrderDetailsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BizOrderDetailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
