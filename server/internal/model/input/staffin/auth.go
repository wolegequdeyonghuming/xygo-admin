// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package staffin

import (
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/model/input/form"
)

// LoginInp 收单员登录入参
type LoginInp struct {
	Username string `p:"username" v:"required#用户名不能为空" json:"username" dc:"用户名"`
	Password string `p:"password" v:"required#密码不能为空" json:"password" dc:"密码"`
}

// LoginModel 收单员登录出参（复用后台 admin token）
type LoginModel struct {
	Id               uint64 `json:"id" dc:"用户ID"`
	Username         string `json:"username" dc:"用户名"`
	Nickname         string `json:"nickname" dc:"昵称"`
	AccessToken      string `json:"accessToken" dc:"访问令牌"`
	ExpiresIn        int64  `json:"expiresIn" dc:"访问令牌过期时间（秒）"`
	RefreshToken     string `json:"refreshToken" dc:"刷新令牌"`
	RefreshExpiresIn int64  `json:"refreshExpiresIn" dc:"刷新令牌过期时间（秒）"`
}

// RefreshInp 刷新 accessToken 入参
type RefreshInp struct {
	RefreshToken string `p:"refreshToken" v:"required#刷新令牌不能为空" json:"refreshToken" dc:"刷新令牌"`
}

// RefreshModel 刷新 accessToken 出参
type RefreshModel struct {
	AccessToken string `json:"accessToken" dc:"新的访问令牌"`
	ExpiresIn   int64  `json:"expiresIn" dc:"访问令牌过期时间（秒）"`
}

// ProfileModel 收单员信息出参
type ProfileModel struct {
	Id       uint64 `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	RealName string `json:"realName" dc:"真实姓名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Mobile   string `json:"mobile" dc:"手机号"`
	DeptName string `json:"deptName" dc:"部门名称"`
	PostName string `json:"postName" dc:"岗位名称"`
	RoleKey  string `json:"roleKey" dc:"角色标识"`
	RoleName string `json:"roleName" dc:"角色名称"`
}

// StaffOrderListInp 收单员订单列表入参
type StaffOrderListInp struct {
	form.PageReq
	Statuses       string `json:"statuses" dc:"订单状态列表（逗号分隔）"`
	VisitDateStart string `json:"visitDateStart" dc:"上门日期开始（YYYY-MM-DD）"`
	VisitDateEnd   string `json:"visitDateEnd" dc:"上门日期结束（YYYY-MM-DD）"`
	Keyword        string `json:"keyword" dc:"客户姓名/电话模糊搜索"`
}

// StaffOrderStatModel 收单员订单统计出参
type StaffOrderStatModel struct {
	Todo  int `json:"todo" dc:"待处理（待预约+待收单）"`
	Today int `json:"today" dc:"今日已处理"`
	Done  int `json:"done" dc:"已处理（累计）"`
}

// StaffOrderAppointInp 预约入参（步骤3）
type StaffOrderAppointInp struct {
	Id              int64       `json:"id" v:"required#订单ID不能为空" dc:"订单ID"`
	AppointmentDesc string      `json:"appointmentDesc" dc:"收单预约情况"`
	VisitDate       *gtime.Time `json:"visitDate" dc:"上门日期"`
}

// StaffOrderCollectInp 收单入参（步骤4）
type StaffOrderCollectInp struct {
	Id                int64       `json:"id" v:"required#订单ID不能为空" dc:"订单ID"`
	Finish            bool        `json:"finish" dc:"是否完成（true=保存并推进到已上门；false=暂存）"`
	VisitDate         *gtime.Time `json:"visitDate" dc:"上门日期"`
	DealtBusinessType string      `json:"dealtBusinessType" dc:"成交业务"`
	PortingStatus     string      `json:"portingStatus" dc:"携转情况"`
	CustomerRealName  string      `json:"customerRealName" dc:"客户实际姓名"`
	CustomerIdNumber  string      `json:"customerIdNumber" dc:"客户身份证号"`
	PaidAmount        float64     `json:"paidAmount" dc:"实缴额度（元）"`
	IsRuralOrder      int         `json:"isRuralOrder" dc:"是否乡下单"`
	NewPhoneNo        string      `json:"newPhoneNo" dc:"新开号码"`
	DeviceSerial      string      `json:"deviceSerial" dc:"终端串码"`
	SubsidyAmount     float64     `json:"subsidyAmount" dc:"话补"`
	IsNew             string      `json:"isNew" dc:"是否纯新增"`
	AttachmentId      string      `json:"attachmentId" dc:"附件"`
}

// StaffAttachmentItem 附件信息项
type StaffAttachmentItem struct {
	Id  int64  `json:"id" dc:"附件ID"`
	Url string `json:"url" dc:"附件URL"`
}

// StaffAttachmentListModel 附件列表出参
type StaffAttachmentListModel struct {
	List []StaffAttachmentItem `json:"list"`
}
