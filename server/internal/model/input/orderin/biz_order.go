package orderin

import (
	"github.com/gogf/gf/v2/os/gtime"
	"xygo/internal/model/input/form"
)

// ==================== 订单主表 ====================

// BizOrderListInp 订单主表列表入参
type BizOrderListInp struct {
	form.PageReq
	OrderStatus       *int   `json:"orderStatus" dc:"订单状态"`
	ScheduleDateStart string `json:"scheduleDateStart" dc:"排单日期开始值"`
	ScheduleDateEnd   string `json:"scheduleDateEnd" dc:"排单日期结束值"`
	VisitDateStart    string `json:"visitDateStart" dc:"上门日期开始值"`
	VisitDateEnd      string `json:"visitDateEnd" dc:"上门日期结束值"`
	CustomerName      string `json:"customerName" dc:"客户姓名"`
	Area              *int   `json:"area" dc:"区县（字典）"`
	InstallAddress    string `json:"installAddress" dc:"安装地址"`
	ContactPhone      string `json:"contactPhone" dc:"联系电话"`
	BusinessType      string `json:"businessType" dc:"预约业务"`
	// 关联表搜索字段
}

// BizOrderListItem 订单主表列表项
type BizOrderListItem struct {
	Id                int64       `json:"id" dc:"主键"`
	OrderNo           string      `json:"orderNo" dc:"订单编号"`
	OrderStatus       *int        `json:"orderStatus" dc:"订单状态"`
	ScheduleDate      *gtime.Time `json:"scheduleDate" dc:"排单日期"`
	VisitDate         *gtime.Time `json:"visitDate" dc:"上门日期"`
	CustomerName      string      `json:"customerName" dc:"客户姓名"`
	AvailableTimeDesc string      `json:"availableTimeDesc" dc:"可联系时间"`
	Area              *int        `json:"area" dc:"区县（字典）"`
	InstallAddress    string      `json:"installAddress" dc:"安装地址"`
	ContactPhone      string      `json:"contactPhone" dc:"联系电话"`
	BusinessType      string      `json:"businessType" dc:"预约业务"`
	ExpiryDate        string      `json:"expiryDate" dc:"到期时间"`
	PrimaryTelNo      string      `json:"primaryTelNo" dc:"主卡号码"`
	Details           string      `json:"details" dc:"详细情况（关联）"`
	TelemarketerId    *int64      `json:"telemarketerId" dc:"话务员"`
	AgentId           *int64      `json:"agentId" dc:"收单员"`
	AppointmentDesc   string      `json:"appointmentDesc" dc:"收单预约情况"`
	FollowUpDesc      string      `json:"followUpDesc" dc:"话务二次回访情况"`
	DealtBusinessType string      `json:"dealtBusinessType" dc:"成交业务"`
	PortingStatus     string      `json:"portingStatus" dc:"携转情况"`
	CustomerRealName  string      `json:"customerRealName" dc:"客户实际姓名"`
	CustomerIdNumber  *string     `json:"customerIdNumber" dc:"客户身份证号"`
	PaidAmount        *float64    `json:"paidAmount" dc:"实缴额度（元）"`
	IsRuralOrder      *int        `json:"isRuralOrder" dc:"是否乡下单"`
	NewPhoneNo        string      `json:"newPhoneNo" dc:"新开号码"`
	DeviceSerial      string      `json:"deviceSerial" dc:"终端串码"`
	BroadbandAccount  string      `json:"broadbandAccount" dc:"宽带账号"`
	IsCompleted       *int        `json:"isCompleted" dc:"是否完工"`
	SubsidyAmount     *float64    `json:"subsidyAmount" dc:"话补"`
	AgencyNo          string      `json:"agencyNo" dc:"工号"`
	IsNew             string      `json:"isNew" dc:"是否纯新增"`
	Remark            string      `json:"remark" dc:"备注"`
	CreatedBy         int64       `json:"createdBy" dc:"录单人ID"`
	// 关联表字段（来自 LeftJoin）
	TelemarketerRealName string `json:"telemarketer_real_name" dc:"Telemarketerreal_name"`
	AgentRealName        string `json:"agent_real_name" dc:"Agentreal_name"`
}

// BizOrderListModel 订单主表列表出参
type BizOrderListModel struct {
	List []BizOrderListItem `json:"list"`
	form.PageRes
}

// BizOrderViewModel 订单主表详情出参
type BizOrderViewModel struct {
	Id                int64       `json:"id" dc:"主键"`
	OrderNo           string      `json:"orderNo" dc:"订单编号"`
	OrderStatus       *int        `json:"orderStatus" dc:"订单状态"`
	ScheduleDate      *gtime.Time `json:"scheduleDate" dc:"排单日期"`
	VisitDate         *gtime.Time `json:"visitDate" dc:"上门日期"`
	CustomerName      string      `json:"customerName" dc:"客户姓名"`
	AvailableTimeDesc string      `json:"availableTimeDesc" dc:"可联系时间"`
	Area              *int        `json:"area" dc:"区县（字典）"`
	InstallAddress    string      `json:"installAddress" dc:"安装地址"`
	ContactPhone      string      `json:"contactPhone" dc:"联系电话"`
	BusinessType      string      `json:"businessType" dc:"预约业务"`
	ExpiryDate        string      `json:"expiryDate" dc:"到期时间"`
	PrimaryTelNo      string      `json:"primaryTelNo" dc:"主卡号码"`
	Details           string      `json:"details" dc:"详细情况（关联）"`
	TelemarketerId    *int64      `json:"telemarketerId" dc:"话务员"`
	AgentId           *int64      `json:"agentId" dc:"收单员"`
	AppointmentDesc   string      `json:"appointmentDesc" dc:"收单预约情况"`
	FollowUpDesc      string      `json:"followUpDesc" dc:"话务二次回访情况"`
	DealtBusinessType string      `json:"dealtBusinessType" dc:"成交业务"`
	PortingStatus     string      `json:"portingStatus" dc:"携转情况"`
	CustomerRealName  string      `json:"customerRealName" dc:"客户实际姓名"`
	CustomerIdNumber  *string     `json:"customerIdNumber" dc:"客户身份证号"`
	PaidAmount        *float64    `json:"paidAmount" dc:"实缴额度（元）"`
	IsRuralOrder      *int        `json:"isRuralOrder" dc:"是否乡下单"`
	NewPhoneNo        string      `json:"newPhoneNo" dc:"新开号码"`
	DeviceSerial      string      `json:"deviceSerial" dc:"终端串码"`
	BroadbandAccount  string      `json:"broadbandAccount" dc:"宽带账号"`
	IsCompleted       *int        `json:"isCompleted" dc:"是否完工"`
	SubsidyAmount     *float64    `json:"subsidyAmount" dc:"话补"`
	AgencyNo          string      `json:"agencyNo" dc:"工号"`
	IsNew             string      `json:"isNew" dc:"是否纯新增"`
	CreatedAt         int64       `json:"createdAt" dc:"创建时间"`
	UpdatedAt         int64       `json:"updatedAt" dc:"更新时间"`
	IsDeleted         int         `json:"isDeleted" dc:"是否删除"`
	Remark            string      `json:"remark" dc:"备注"`
	AttachmentId      string      `json:"attachmentId" dc:"附件"`
	CreatedBy         int64       `json:"createdBy" dc:"录单人ID"`

	// 关联表字段（来自 LeftJoin）
	TelemarketerRealName string `json:"telemarketer_real_name" dc:"话务员姓名"`
	AgentRealName        string `json:"agent_real_name" dc:"收单员姓名"`
}

// BizOrderEditInp 订单主表编辑入参
type BizOrderEditInp struct {
	Id                int64       `json:"id" dc:"主键"`
	OrderStatus       int         `json:"orderStatus" dc:"订单状态"`
	ScheduleDate      *gtime.Time `json:"scheduleDate" dc:"排单日期"`
	VisitDate         *gtime.Time `json:"visitDate" dc:"上门日期"`
	CustomerName      string      `json:"customerName" v:"required#客户姓名不能为空" dc:"客户姓名"`
	AvailableTimeDesc string      `json:"availableTimeDesc" dc:"可联系时间"`
	Area              int         `json:"area" dc:"区县（字典）"`
	InstallAddress    string      `json:"installAddress" dc:"安装地址"`
	ContactPhone      string      `json:"contactPhone" v:"required#联系电话不能为空" dc:"联系电话"`
	BusinessType      string      `json:"businessType" dc:"预约业务"`
	ExpiryDate        string      `json:"expiryDate" dc:"到期时间"`
	PrimaryTelNo      string      `json:"primaryTelNo" dc:"主卡号码"`
	Details           string      `json:"details" dc:"详细情况（关联）"`
	TelemarketerId    int64       `json:"telemarketerId" dc:"话务员"`
	AgentId           int64       `json:"agentId" dc:"收单员"`
	AppointmentDesc   string      `json:"appointmentDesc" dc:"收单预约情况"`
	FollowUpDesc      string      `json:"followUpDesc" dc:"话务二次回访情况"`
	DealtBusinessType string      `json:"dealtBusinessType" dc:"成交业务"`
	PortingStatus     string      `json:"portingStatus" dc:"携转情况"`
	CustomerRealName  string      `json:"customerRealName" dc:"客户实际姓名"`
	CustomerIdNumber  string      `json:"customerIdNumber" dc:"客户身份证号"`
	PaidAmount        float64     `json:"paidAmount" dc:"实缴额度（元）"`
	IsRuralOrder      int         `json:"isRuralOrder" dc:"是否乡下单"`
	NewPhoneNo        string      `json:"newPhoneNo" dc:"新开号码"`
	DeviceSerial      string      `json:"deviceSerial" dc:"终端串码"`
	BroadbandAccount  string      `json:"broadbandAccount" dc:"宽带账号"`
	IsCompleted       int         `json:"isCompleted" dc:"是否完工"`
	SubsidyAmount     float64     `json:"subsidyAmount" dc:"话补"`
	AgencyNo          string      `json:"agencyNo" dc:"工号"`
	IsNew             string      `json:"isNew" dc:"是否纯新增"`
	Remark            string      `json:"remark" dc:"备注"`
	AttachmentId      string      `json:"attachmentId" dc:"附件"`
}

// BizOrderStepEditInp 订单主表分步保存入参
// Step 对应订单流程步骤：1 录单、2 派单、3 预约、4 上门、5 回访、6 完成
type BizOrderStepEditInp struct {
	Id                int64       `json:"id" v:"required#订单ID不能为空" dc:"订单ID"`
	Step              int         `json:"step" v:"required#步骤不能为空|in:1,2,3,4,5,6" dc:"步骤(1-6)"`
	ScheduleDate      *gtime.Time `json:"scheduleDate" dc:"排单日期"`
	CustomerName      string      `json:"customerName" v:"required#客户姓名不能为空" dc:"客户姓名"`
	AvailableTimeDesc string      `json:"availableTimeDesc" dc:"可联系时间"`
	Area              int         `json:"area" dc:"区县（字典）"`
	InstallAddress    string      `json:"installAddress" dc:"安装地址"`
	ContactPhone      string      `json:"contactPhone" v:"required#联系电话不能为空" dc:"联系电话"`
	BusinessType      string      `json:"businessType" dc:"预约业务"`
	ExpiryDate        string      `json:"expiryDate" dc:"到期时间"`
	AgentId           int64       `json:"agentId" dc:"收单员"`
	AppointmentDesc   string      `json:"appointmentDesc" dc:"收单预约情况"`
	VisitDate         *gtime.Time `json:"visitDate" dc:"上门日期"`
	DealtBusinessType string      `json:"dealtBusinessType" dc:"成交业务"`
	PortingStatus     string      `json:"portingStatus" dc:"携转情况"`
	CustomerRealName  string      `json:"customerRealName" dc:"客户实际姓名"`
	CustomerIdNumber  string      `json:"customerIdNumber" dc:"客户身份证号"`
	PaidAmount        float64     `json:"paidAmount" dc:"实缴额度（元）"`
	IsRuralOrder      int         `json:"isRuralOrder" dc:"是否乡下单"`
	NewPhoneNo        string      `json:"newPhoneNo" dc:"新开号码"`
	DeviceSerial      string      `json:"deviceSerial" dc:"终端串码"`
	AttachmentId      string      `json:"attachmentId" dc:"附件"`
	SubsidyAmount     float64     `json:"subsidyAmount" dc:"话补"`
	IsNew             string      `json:"isNew" dc:"是否纯新增"`
	FollowUpDesc      string      `json:"followUpDesc" dc:"话务二次回访情况"`
	BroadbandAccount  string      `json:"broadbandAccount" dc:"宽带账号"`
	IsCompleted       int         `json:"isCompleted" dc:"是否完工"`
	AgencyNo          string      `json:"agencyNo" dc:"工号"`
}

// BizOrderDetailListItem 订单详细情况列表项
type BizOrderDetailListItem struct {
	Id        int64  `json:"id" dc:"主键"`
	OrderId   int64  `json:"orderId" dc:"订单id"`
	UserId    int64  `json:"userId" dc:"操作人"`
	UserName  string `json:"userName" dc:"操作人姓名"`
	Content   string `json:"content" dc:"内容"`
	CreatedAt int64  `json:"createdAt" dc:"创建时间"`
}

// BizOrderDetailListInp 订单详细情况列表入参
type BizOrderDetailListInp struct {
	OrderId int64 `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}

// BizOrderDetailListModel 订单详细情况列表出参
type BizOrderDetailListModel struct {
	List []BizOrderDetailListItem `json:"list"`
}

// BizOrderDetailAddInp 新增订单详细情况入参
type BizOrderDetailAddInp struct {
	OrderId int64  `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
	Content string `json:"content" v:"required#详细情况不能为空" dc:"详细情况"`
}

// BizOrderDetailDeleteInp 删除订单详细情况入参
type BizOrderDetailDeleteInp struct {
	Id int64 `json:"id" v:"required#详细情况ID不能为空" dc:"详细情况ID"`
}
