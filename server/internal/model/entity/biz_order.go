// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BizOrder is the golang structure for table biz_order.
type BizOrder struct {
	Id                int64       `json:"id"                orm:"id"                  description:"主键"`       // 主键
	ScheduleDate      *gtime.Time `json:"scheduleDate"      orm:"schedule_date"       description:"排单日期"`     // 排单日期
	VisitDate         *gtime.Time `json:"visitDate"         orm:"visit_date"          description:"上门日期"`     // 上门日期
	CustomerName      string      `json:"customerName"      orm:"customer_name"       description:"客户姓名"`     // 客户姓名
	AvailableTimeDesc string      `json:"availableTimeDesc" orm:"available_time_desc" description:"可联系时间"`    // 可联系时间
	Area              int         `json:"area"              orm:"area"                description:"区县（字典）"`   // 区县（字典）
	InstallAddress    string      `json:"installAddress"    orm:"install_address"     description:"安装地址"`     // 安装地址
	ContactPhone      string      `json:"contactPhone"      orm:"contact_phone"       description:"联系电话"`     // 联系电话
	BusinessType      string      `json:"businessType"      orm:"business_type"       description:"预约业务"`     // 预约业务
	ExpiryDate        string      `json:"expiryDate"        orm:"expiry_date"         description:"到期时间"`     // 到期时间
	PrimaryTelNo      string      `json:"primaryTelNo"      orm:"primary_tel_no"      description:"主卡号码"`     // 主卡号码
	Details           string      `json:"details"           orm:"details"             description:"详细情况（关联）"` // 详细情况（关联）
	TelemarketerId    int64       `json:"telemarketerId"    orm:"telemarketer_id"     description:"话务员id"`    // 话务员id
	AgentId           int64       `json:"agentId"           orm:"agent_id"            description:"收单员id"`    // 收单员id
	AppointmentDesc   string      `json:"appointmentDesc"   orm:"appointment_desc"    description:"收单预约情况"`   // 收单预约情况
	FollowUpDesc      string      `json:"followUpDesc"      orm:"follow_up_desc"      description:"话务二次回放情况"` // 话务二次回放情况
	DealtBusinessType string      `json:"dealtBusinessType" orm:"dealt_business_type" description:"成交业务"`     // 成交业务
	PortingStatus     string      `json:"portingStatus"     orm:"porting_status"      description:"携转情况"`     // 携转情况
	CustomerRealName  string      `json:"customerRealName"  orm:"customer_real_name"  description:"客户实际姓名"`   // 客户实际姓名
	CustomerIdNumber  int64       `json:"customerIdNumber"  orm:"customer_id_number"  description:"客户身份证号"`   // 客户身份证号
	PaidAmount        float64     `json:"paidAmount"        orm:"paid_amount"         description:"实缴额度（元）"`  // 实缴额度（元）
	IsRuralOrder      int         `json:"isRuralOrder"      orm:"is_rural_order"      description:"是否乡下单"`    // 是否乡下单
	NewPhoneNo        string      `json:"newPhoneNo"        orm:"new_phone_no"        description:"新开号码"`     // 新开号码
	DeviceSerial      string      `json:"deviceSerial"      orm:"device_serial"       description:"终端串码"`     // 终端串码
	BroadbandAccount  string      `json:"broadbandAccount"  orm:"broadband_account"   description:"宽带账号"`     // 宽带账号
	IsCompleted       int         `json:"isCompleted"       orm:"is_completed"        description:"是否完工"`     // 是否完工
	SubsidyAmount     float64     `json:"subsidyAmount"     orm:"subsidy_amount"      description:"话补"`       // 话补
	AgencyNo          string      `json:"agencyNo"          orm:"agency_no"           description:"工号"`       // 工号
	IsNew             string      `json:"isNew"             orm:"is_new"              description:"是否纯新增"`    // 是否纯新增
	CreatedAt         int64       `json:"createdAt"         orm:"created_at"          description:"创建时间"`     // 创建时间
	UpdatedAt         int64       `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`     // 更新时间
	IsDeleted         int         `json:"isDeleted"         orm:"is_deleted"          description:"是否删除"`     // 是否删除
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`       // 备注
}
