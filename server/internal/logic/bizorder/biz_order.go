package bizorder

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/dao"
	"xygo/internal/model/input/form"
	adminin "xygo/internal/model/input/orderin"
	"xygo/internal/service"
)

type sBizOrder struct{}

func init() {
	service.RegisterBizOrder(New())
}

func New() *sBizOrder {
	return &sBizOrder{}
}

// List 订单主表列表
func (s *sBizOrder) List(ctx context.Context, in *adminin.BizOrderListInp) (*adminin.BizOrderListModel, error) {
	model := dao.BizOrder.Ctx(ctx).As("t")
	// 关联表 LeftJoin
	model = model.LeftJoin("xy_admin_user telemarketer", "telemarketer.id = t.telemarketer_id")
	model = model.LeftJoin("xy_admin_user agent", "agent.id = t.agent_id")
	if in.OrderStatus != nil {
		model = model.Where("t.order_status", *in.OrderStatus)
	}
	if in.ScheduleDateStart != "" && in.ScheduleDateEnd != "" {
		model = model.WhereBetween("t.schedule_date", in.ScheduleDateStart, in.ScheduleDateEnd)
	}
	if in.VisitDateStart != "" && in.VisitDateEnd != "" {
		model = model.WhereBetween("t.visit_date", in.VisitDateStart, in.VisitDateEnd)
	}
	if in.CustomerName != "" {
		model = model.WhereLike("t.customer_name", "%"+in.CustomerName+"%")
	}
	if in.Area != nil {
		model = model.Where("t.area", *in.Area)
	}
	if in.InstallAddress != "" {
		model = model.WhereLike("t.install_address", "%"+in.InstallAddress+"%")
	}
	if in.ContactPhone != "" {
		model = model.WhereLike("t.contact_phone", "%"+in.ContactPhone+"%")
	}
	if in.BusinessType != "" {
		model = model.WhereLike("t.business_type", "%"+in.BusinessType+"%")
	}
	// 关联表搜索条件
	// 先计数（不带 Fields，避免 COUNT + 字段别名冲突）
	count, err := model.Clone().Count()
	if err != nil {
		return nil, err
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}
	// 计数后添加 Fields
	// real_name 为空时回退到 nickname，保证话务员/收单员姓名可读
	model = model.Fields("t.*, " +
		"CASE WHEN telemarketer.real_name != '' THEN telemarketer.real_name ELSE telemarketer.nickname END as telemarketer_real_name, " +
		"CASE WHEN agent.real_name != '' THEN agent.real_name ELSE agent.nickname END as agent_real_name")
	var list []adminin.BizOrderListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("t.id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.BizOrderListItem{}
	}

	return &adminin.BizOrderListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 订单主表详情
func (s *sBizOrder) View(ctx context.Context, id uint64) (*adminin.BizOrderViewModel, error) {
	var item adminin.BizOrderViewModel
	model := dao.BizOrder.Ctx(ctx).As("t")
	model = model.LeftJoin("xy_admin_user telemarketer", "telemarketer.id = t.telemarketer_id")
	model = model.LeftJoin("xy_admin_user agent", "agent.id = t.agent_id")
	// real_name 为空时回退到 nickname，保证话务员/收单员姓名可读
	model = model.Fields("t.*, " +
		"CASE WHEN telemarketer.real_name != '' THEN telemarketer.real_name ELSE telemarketer.nickname END as telemarketer_real_name, " +
		"CASE WHEN agent.real_name != '' THEN agent.real_name ELSE agent.nickname END as agent_real_name")
	err := model.Where("t.id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存订单主表
func (s *sBizOrder) Edit(ctx context.Context, in *adminin.BizOrderEditInp) error {
	data := g.Map{
		"order_status":        in.OrderStatus,
		"schedule_date":       in.ScheduleDate,
		"visit_date":          in.VisitDate,
		"customer_name":       in.CustomerName,
		"available_time_desc": in.AvailableTimeDesc,
		"area":                in.Area,
		"install_address":     in.InstallAddress,
		"contact_phone":       in.ContactPhone,
		"business_type":       in.BusinessType,
		"expiry_date":         in.ExpiryDate,
		"primary_tel_no":      in.PrimaryTelNo,
		"details":             in.Details,
		"telemarketer_id":     in.TelemarketerId,
		"agent_id":            in.AgentId,
		"appointment_desc":    in.AppointmentDesc,
		"follow_up_desc":      in.FollowUpDesc,
		"dealt_business_type": in.DealtBusinessType,
		"porting_status":      in.PortingStatus,
		"customer_real_name":  in.CustomerRealName,
		"customer_id_number":  in.CustomerIdNumber,
		"paid_amount":         in.PaidAmount,
		"is_rural_order":      in.IsRuralOrder,
		"new_phone_no":        in.NewPhoneNo,
		"device_serial":       in.DeviceSerial,
		"broadband_account":   in.BroadbandAccount,
		"is_completed":        in.IsCompleted,
		"subsidy_amount":      in.SubsidyAmount,
		"agency_no":           in.AgencyNo,
		"is_new":              in.IsNew,
		"remark":              in.Remark,
		"attachment_id":       in.AttachmentId,
	}

	// 非自增主键：为空时生成主键值，并加入 INSERT 数据
	if in.Id == 0 {
		in.Id = gtime.TimestampNano()
	}
	data["id"] = in.Id

	// 判断新增/更新：自增主键用 0 判断，非自增（UUID/时间戳）用空值判断
	isNew := false
	isNew = in.Id == 0
	if isNew {
		// 新增
		_, err := dao.BizOrder.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新（updated_at 由 GoFrame 自动维护）
	_, err := dao.BizOrder.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除订单主表
func (s *sBizOrder) Delete(ctx context.Context, id uint64) error {
	_, err := dao.BizOrder.Ctx(ctx).Where("id", id).Delete()
	return err
}
