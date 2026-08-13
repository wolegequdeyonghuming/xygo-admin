// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package staff

import (
	"context"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/library/contexts"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/form"
	adminin "xygo/internal/model/input/orderin"
	"xygo/internal/model/input/staffin"
	"xygo/internal/service"
)

// List 收单员订单列表
func (s *sStaffOrder) List(ctx context.Context, in *staffin.StaffOrderListInp) (*adminin.BizOrderListModel, error) {
	model := dao.BizOrder.Ctx(ctx).As("t")
	model = model.LeftJoin("xy_admin_user telemarketer", "telemarketer.id = t.telemarketer_id")
	model = model.LeftJoin("xy_admin_user agent", "agent.id = t.agent_id")
	model = model.Where("t.is_deleted", 0)

	// 可见范围：收单员只看 agent_id=自己；话务员只看自己录的单；其余（收单员管理员/文员/管理员/超管）看全部
	switch contexts.GetRoleKey(ctx) {
	case consts.RoleAgent:
		model = model.Where("t.agent_id", contexts.GetUserId(ctx))
	case consts.RoleTelemarketer:
		model = model.Where("t.created_by", contexts.GetUserId(ctx))
	}

	if in.Statuses != "" {
		ids := make([]int, 0)
		for _, p := range strings.Split(in.Statuses, ",") {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			if v, err := strconv.Atoi(p); err == nil {
				ids = append(ids, v)
			}
		}
		if len(ids) > 0 {
			model = model.WhereIn("t.order_status", ids)
		}
	}
	if in.VisitDateStart != "" && in.VisitDateEnd != "" {
		model = model.WhereBetween("t.visit_date", in.VisitDateStart, in.VisitDateEnd)
	}
	if in.ScheduleDateStart != "" && in.ScheduleDateEnd != "" {
		model = model.WhereBetween("t.schedule_date", in.ScheduleDateStart, in.ScheduleDateEnd)
	}
	if in.CreatedDateStart != "" && in.CreatedDateEnd != "" {
		start := gtime.NewFromStr(in.CreatedDateStart).StartOfDay().Unix()
		end := gtime.NewFromStr(in.CreatedDateEnd).EndOfDay().Unix()
		model = model.WhereBetween("t.created_at", start, end)
	}
	if in.PendingOnly {
		model = model.Where("(t.agent_id IS NULL OR t.agent_id = 0)")
	}
	if in.Area > 0 {
		model = model.Where("t.area", in.Area)
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
	if in.Keyword != "" {
		kw := "%" + in.Keyword + "%"
		model = model.Where("(t.customer_name LIKE ? OR t.contact_phone LIKE ?)", kw, kw)
	}

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

// Stat 订单统计（按角色返回对应 tab 的统计）
// 话务员：今日录单 / 已录单；收单员：待收单 / 今日收单 / 已收单；其余：待排单 / 待收单 / 已收单
func (s *sStaffOrder) Stat(ctx context.Context) (*staffin.StaffOrderStatModel, error) {
	base := dao.BizOrder.Ctx(ctx).As("t").Where("t.is_deleted", 0)
	switch contexts.GetRoleKey(ctx) {
	case consts.RoleAgent:
		base = base.Where("t.agent_id", contexts.GetUserId(ctx))
	case consts.RoleTelemarketer:
		base = base.Where("t.created_by", contexts.GetUserId(ctx))
	}
	role := contexts.GetRoleKey(ctx)
	result := &staffin.StaffOrderStatModel{}

	switch role {
	case consts.RoleTelemarketer:
		todayStart := gtime.Now().StartOfDay().Unix()
		v, err := base.Clone().WhereGTE("t.created_at", todayStart).Count()
		if err != nil {
			return nil, err
		}
		result.TodayRecorded = v
		v, err = base.Clone().Where("t.order_status", consts.OrderStatusRecorded).Count()
		if err != nil {
			return nil, err
		}
		result.Recorded = v
	case consts.RoleAgent:
		v, err := base.Clone().WhereIn("t.order_status",
			[]int{consts.OrderStatusAssigned, consts.OrderStatusAppointed}).Count()
		if err != nil {
			return nil, err
		}
		result.Todo = v
		todayStart := gtime.Now().Format("Y-m-d")
		v, err = base.Clone().WhereGTE("t.order_status", consts.OrderStatusVisited).WhereLike("t.visit_date", todayStart+"%").Count()
		if err != nil {
			return nil, err
		}
		result.Today = v
		v, err = base.Clone().WhereGTE("t.order_status", consts.OrderStatusVisited).Count()
		if err != nil {
			return nil, err
		}
		result.Done = v
	default: // agent_manager / documentary / admin / super_admin
		v, err := base.Clone().Where("t.order_status", consts.OrderStatusRecorded).
			Where("(t.agent_id IS NULL OR t.agent_id = 0)").Count()
		if err != nil {
			return nil, err
		}
		result.PendingSchedule = v
		v, err = base.Clone().WhereIn("t.order_status",
			[]int{consts.OrderStatusAssigned, consts.OrderStatusAppointed}).Count()
		if err != nil {
			return nil, err
		}
		result.Todo = v
		v, err = base.Clone().WhereGTE("t.order_status", consts.OrderStatusVisited).Count()
		if err != nil {
			return nil, err
		}
		result.Done = v
	}
	return result, nil
}

// Appoint 填写预约情况（步骤3）：保存字段，状态为已接单时推进到已预约
func (s *sStaffOrder) Appoint(ctx context.Context, in *staffin.StaffOrderAppointInp) error {
	if forbidOperational(contexts.GetRoleKey(ctx)) {
		return gerror.New("当前角色无权限执行预约")
	}
	order, err := loadOrder(ctx, in.Id)
	if err != nil {
		return err
	}
	_, err = service.BizOrder().StepEdit(ctx, &adminin.BizOrderStepEditInp{
		Id:              in.Id,
		Step:            consts.OrderStatusAppointed,
		AppointmentDesc: in.AppointmentDesc,
		VisitDate:       in.VisitDate,
	})
	if err != nil {
		return err
	}
	if order.OrderStatus == consts.OrderStatusAssigned {
		return service.BizOrder().StepNext(ctx, uint64(in.Id))
	}
	return nil
}

// Collect 填写收单情况（步骤4）：保存字段；finish=true 且状态为已预约时推进到已上门
func (s *sStaffOrder) Collect(ctx context.Context, in *staffin.StaffOrderCollectInp) error {
	if forbidOperational(contexts.GetRoleKey(ctx)) {
		return gerror.New("当前角色无权限收单")
	}
	order, err := loadOrder(ctx, in.Id)
	if err != nil {
		return err
	}
	_, err = service.BizOrder().StepEdit(ctx, &adminin.BizOrderStepEditInp{
		Id:                in.Id,
		Step:              consts.OrderStatusVisited,
		VisitDate:         in.VisitDate,
		DealtBusinessType: in.DealtBusinessType,
		PortingStatus:     in.PortingStatus,
		CustomerRealName:  in.CustomerRealName,
		CustomerIdNumber:  in.CustomerIdNumber,
		PaidAmount:        in.PaidAmount,
		IsRuralOrder:      in.IsRuralOrder,
		NewPhoneNo:        in.NewPhoneNo,
		DeviceSerial:      in.DeviceSerial,
		SubsidyAmount:     in.SubsidyAmount,
		IsNew:             in.IsNew,
		AttachmentId:      in.AttachmentId,
	})
	if err != nil {
		return err
	}
	if in.Finish && order.OrderStatus == consts.OrderStatusAppointed {
		return service.BizOrder().StepNext(ctx, uint64(in.Id))
	}
	return nil
}

// loadOrder 加载订单（未删除），不存在则报错
func loadOrder(ctx context.Context, id int64) (*entity.BizOrder, error) {
	var order *entity.BizOrder
	if err := dao.BizOrder.Ctx(ctx).Where("id", id).Where("is_deleted", 0).Scan(&order); err != nil {
		return nil, err
	}
	if order == nil || order.Id == 0 {
		return nil, gerror.New("订单不存在")
	}
	return order, nil
}

// forbidOperational 管理员/文员是否禁止排单/预约/收单等运营操作（仅查看与完工，超管豁免）
func forbidOperational(role string) bool {
	if consts.IsSuperRole(role) {
		return false
	}
	return role == consts.RoleAdmin || role == consts.RoleDocumentary
}

// AttachmentList 按附件ID列表获取附件 URL（ids 为逗号分隔）
func (s *sStaffOrder) AttachmentList(ctx context.Context, ids string) (*staffin.StaffAttachmentListModel, error) {
	idList := make([]int64, 0)
	for _, p := range strings.Split(ids, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if v, err := strconv.ParseInt(p, 10, 64); err == nil {
			idList = append(idList, v)
		}
	}
	result := &staffin.StaffAttachmentListModel{List: []staffin.StaffAttachmentItem{}}
	if len(idList) == 0 {
		return result, nil
	}

	var rows []struct {
		Id  int64  `json:"id"`
		Url string `json:"url"`
	}
	if err := dao.SysAttachment.Ctx(ctx).
		WhereIn("id", idList).
		Fields("id, url").
		Scan(&rows); err != nil {
		return nil, err
	}
	for _, r := range rows {
		result.List = append(result.List, staffin.StaffAttachmentItem{
			Id:  r.Id,
			Url: r.Url,
		})
	}
	return result, nil
}

// UserAgents 收单员列表（排单选人）
func (s *sStaffOrder) UserAgents(ctx context.Context) (*staffin.StaffUserAgentListModel, error) {
	result := &staffin.StaffUserAgentListModel{List: []staffin.StaffUserAgentItem{}}
	var rows []struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	err := dao.AdminUser.Ctx(ctx).
		LeftJoin(dao.AdminUserRole.Table()+" aur", "aur.user_id = "+dao.AdminUser.Table()+".id").
		LeftJoin(dao.AdminRole.Table()+" r", "r.id = aur.role_id").
		Where("r.key", consts.RoleAgent).
		Where(dao.AdminUser.Table()+".status", 1).
		Where("r.status", 1).
		Fields(dao.AdminUser.Table()+"."+dao.AdminUser.Columns().Id,
			"CASE WHEN "+dao.AdminUser.Table()+".real_name != '' THEN "+dao.AdminUser.Table()+".real_name ELSE "+dao.AdminUser.Table()+".nickname END as name").
		OrderAsc(dao.AdminUser.Table() + "." + dao.AdminUser.Columns().Id).
		Scan(&rows)
	if err != nil {
		return nil, err
	}
	for _, r := range rows {
		result.List = append(result.List, staffin.StaffUserAgentItem{
			Id:   r.Id,
			Name: r.Name,
		})
	}
	return result, nil
}

// Schedule 排单（步骤2）：选择收单员，状态为已录单时推进到已接单
func (s *sStaffOrder) Schedule(ctx context.Context, in *staffin.StaffOrderScheduleInp) error {
	if forbidOperational(contexts.GetRoleKey(ctx)) {
		return gerror.New("当前角色无权限排单")
	}
	order, err := loadOrder(ctx, in.Id)
	if err != nil {
		return err
	}
	_, err = service.BizOrder().StepEdit(ctx, &adminin.BizOrderStepEditInp{
		Id:      in.Id,
		Step:    consts.OrderStatusAssigned, // 2 派单/排单
		AgentId: in.AgentId,
	})
	if err != nil {
		return err
	}
	if order.OrderStatus == consts.OrderStatusRecorded {
		return service.BizOrder().StepNext(ctx, uint64(in.Id))
	}
	return nil
}

// Complete 完工（步骤6）：
//   - 已回访(5)：StepEdit(6) + 推进到已完工
//   - 已完工(6)：StepEdit(6) 重编辑
//   - 其他状态：仅管理员/超管可直接补录完工
func (s *sStaffOrder) Complete(ctx context.Context, in *staffin.StaffOrderCompleteInp) error {
	order, err := loadOrder(ctx, in.Id)
	if err != nil {
		return err
	}
	role := contexts.GetRoleKey(ctx)
	inp := &adminin.BizOrderStepEditInp{
		Id:               in.Id,
		Step:             consts.OrderStatusCompleted,
		BroadbandAccount: in.BroadbandAccount,
		IsCompleted:      in.IsCompleted,
		AgencyNo:         in.AgencyNo,
	}

	switch order.OrderStatus {
	case consts.OrderStatusFollowed, consts.OrderStatusCompleted: // 5 / 6
		_, err = service.BizOrder().StepEdit(ctx, inp)
		if err != nil {
			return err
		}
		if order.OrderStatus == consts.OrderStatusFollowed {
			return service.BizOrder().StepNext(ctx, uint64(in.Id))
		}
		return nil
	default:
		// 其他状态：仅管理员/超管允许直接补录完工
		if role != consts.RoleAdmin && !consts.IsSuperRole(role) {
			return gerror.New("仅管理员可对当前状态的订单直接完工")
		}
		_, err = dao.BizOrder.Ctx(ctx).Where("id", in.Id).Data(g.Map{
			"broadband_account": in.BroadbandAccount,
			"is_completed":      in.IsCompleted,
			"agency_no":         in.AgencyNo,
			"order_status":      consts.OrderStatusCompleted,
		}).Update()
		return err
	}
}
