// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BizOrder is the golang structure of table xy_biz_order for DAO operations like Where/Data.
type BizOrder struct {
	g.Meta            `orm:"table:xy_biz_order, do:true"`
	Id                any         // 主键
	ScheduleDate      *gtime.Time // 排单日期
	VisitDate         *gtime.Time // 上门日期
	CustomerName      any         // 客户姓名
	AvailableTimeDesc any         // 可联系时间
	Area              any         // 区县（字典）
	InstallAddress    any         // 安装地址
	ContactPhone      any         // 联系电话
	BusinessType      any         // 预约业务
	ExpiryDate        any         // 到期时间
	PrimaryTelNo      any         // 主卡号码
	Details           any         // 详细情况（关联）
	TelemarketerId    any         // 话务员id
	AgentId           any         // 收单员id
	AppointmentDesc   any         // 收单预约情况
	FollowUpDesc      any         // 话务二次回访情况
	DealtBusinessType any         // 成交业务
	PortingStatus     any         // 携转情况
	CustomerRealName  any         // 客户实际姓名
	CustomerIdNumber  any         // 客户身份证号
	PaidAmount        any         // 实缴额度（元）
	IsRuralOrder      any         // 是否乡下单
	NewPhoneNo        any         // 新开号码
	DeviceSerial      any         // 终端串码
	BroadbandAccount  any         // 宽带账号
	IsCompleted       any         // 是否完工
	SubsidyAmount     any         // 话补
	AgencyNo          any         // 工号
	IsNew             any         // 是否纯新增
	CreatedAt         any         // 创建时间
	UpdatedAt         any         // 更新时间
	IsDeleted         any         // 是否删除
	Remark            any         // 备注
	OrderStatus       any         // 订单状态
	AttachmentId      any         // 附件id
	CreatedBy         any         // 创建人id
}
