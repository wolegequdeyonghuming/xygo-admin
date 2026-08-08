// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BizOrderDao is the data access object for the table xy_biz_order.
type BizOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BizOrderColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BizOrderColumns defines and stores column names for the table xy_biz_order.
type BizOrderColumns struct {
	Id                string // 主键
	ScheduleDate      string // 排单日期
	VisitDate         string // 上门日期
	CustomerName      string // 客户姓名
	AvailableTimeDesc string // 可联系时间
	Area              string // 区县（字典）
	InstallAddress    string // 安装地址
	ContactPhone      string // 联系电话
	BusinessType      string // 预约业务
	ExpiryDate        string // 到期时间
	PrimaryTelNo      string // 主卡号码
	Details           string // 详细情况（关联）
	TelemarketerId    string // 话务员id
	AgentId           string // 收单员id
	AppointmentDesc   string // 收单预约情况
	FollowUpDesc      string // 话务二次回放情况
	DealtBusinessType string // 成交业务
	PortingStatus     string // 携转情况
	CustomerRealName  string // 客户实际姓名
	CustomerIdNumber  string // 客户身份证号
	PaidAmount        string // 实缴额度（元）
	IsRuralOrder      string // 是否乡下单
	NewPhoneNo        string // 新开号码
	DeviceSerial      string // 终端串码
	BroadbandAccount  string // 宽带账号
	IsCompleted       string // 是否完工
	SubsidyAmount     string // 话补
	AgencyNo          string // 工号
	IsNew             string // 是否纯新增
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	IsDeleted         string // 是否删除
	Remark            string // 备注
	OrderStatus       string // 订单状态
	AttachmentId      string // 附件id
}

// bizOrderColumns holds the columns for the table xy_biz_order.
var bizOrderColumns = BizOrderColumns{
	Id:                "id",
	ScheduleDate:      "schedule_date",
	VisitDate:         "visit_date",
	CustomerName:      "customer_name",
	AvailableTimeDesc: "available_time_desc",
	Area:              "area",
	InstallAddress:    "install_address",
	ContactPhone:      "contact_phone",
	BusinessType:      "business_type",
	ExpiryDate:        "expiry_date",
	PrimaryTelNo:      "primary_tel_no",
	Details:           "details",
	TelemarketerId:    "telemarketer_id",
	AgentId:           "agent_id",
	AppointmentDesc:   "appointment_desc",
	FollowUpDesc:      "follow_up_desc",
	DealtBusinessType: "dealt_business_type",
	PortingStatus:     "porting_status",
	CustomerRealName:  "customer_real_name",
	CustomerIdNumber:  "customer_id_number",
	PaidAmount:        "paid_amount",
	IsRuralOrder:      "is_rural_order",
	NewPhoneNo:        "new_phone_no",
	DeviceSerial:      "device_serial",
	BroadbandAccount:  "broadband_account",
	IsCompleted:       "is_completed",
	SubsidyAmount:     "subsidy_amount",
	AgencyNo:          "agency_no",
	IsNew:             "is_new",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	IsDeleted:         "is_deleted",
	Remark:            "remark",
	OrderStatus:       "order_status",
	AttachmentId:      "attachment_id",
}

// NewBizOrderDao creates and returns a new DAO object for table data access.
func NewBizOrderDao(handlers ...gdb.ModelHandler) *BizOrderDao {
	return &BizOrderDao{
		group:    "default",
		table:    "xy_biz_order",
		columns:  bizOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BizOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BizOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BizOrderDao) Columns() BizOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BizOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BizOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BizOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
