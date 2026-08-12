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

	// 可见范围：收单员只看 agent_id=自己
	if contexts.GetRoleKey(ctx) == consts.RoleAgent {
		model = model.Where("t.agent_id", contexts.GetUserId(ctx))
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

// Stat 收单员订单统计
func (s *sStaffOrder) Stat(ctx context.Context) (*staffin.StaffOrderStatModel, error) {
	base := dao.BizOrder.Ctx(ctx).As("t").Where("t.is_deleted", 0)
	if contexts.GetRoleKey(ctx) == consts.RoleAgent {
		base = base.Where("t.agent_id", contexts.GetUserId(ctx))
	}

	todo, err := base.Clone().WhereIn("t.order_status",
		[]int{consts.OrderStatusAssigned, consts.OrderStatusAppointed}).Count()
	if err != nil {
		return nil, err
	}
	todayStart := gtime.Now().Format("Y-m-d")
	today, err := base.Clone().
		WhereGTE("t.order_status", consts.OrderStatusVisited).
		WhereLike("t.visit_date", todayStart+"%").Count()
	if err != nil {
		return nil, err
	}
	done, err := base.Clone().WhereGTE("t.order_status", consts.OrderStatusVisited).Count()
	if err != nil {
		return nil, err
	}
	return &staffin.StaffOrderStatModel{
		Todo:  todo,
		Today: today,
		Done:  done,
	}, nil
}

// Appoint 填写预约情况（步骤3）：保存字段，状态为已接单时推进到已预约
func (s *sStaffOrder) Appoint(ctx context.Context, in *staffin.StaffOrderAppointInp) error {
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
