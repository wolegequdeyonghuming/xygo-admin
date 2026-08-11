package bizorder

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xuri/excelize/v2"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/model/entity"
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
	// 默认只查未删除的记录
	model = model.Where("t.is_deleted", 0)
	// 按角色约束可见范围（话务员看自己录的、收单员看派给自己的、其余看全部）
	model = applyScope(ctx, model, currentRole(ctx))
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
	// 与 List 一致的数据可见性 + 默认过滤已删除
	model = model.Where("t.is_deleted", 0)
	model = applyScope(ctx, model, currentRole(ctx))
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

// StepEdit 订单分步保存：按步骤只更新本步骤字段，并校验角色与当前状态
// 返回更新后的订单 ID（录单新增时返回新生成的主键）。
func (s *sBizOrder) StepEdit(ctx context.Context, in *adminin.BizOrderStepEditInp) (id uint64, err error) {
	role := currentRole(ctx)
	userId := currentUserId(ctx)

	// 录单新增：订单 ID 为空且当前为录单步骤时创建新订单
	if in.Id == 0 {
		if in.Step != consts.OrderStatusRecorded {
			return 0, gerror.New("订单不存在")
		}
		if !canStepFromStatus(role, consts.OrderStatusNotDeal, in.Step) {
			return 0, permissionDenied("当前角色不允许录单")
		}
		// 生成时间戳主键并插入订单（同时保存录单步骤字段，避免新增后数据丢失）
		in.Id = gtime.TimestampNano()
		recordData, err := s.buildStepData(in)
		if err != nil {
			return 0, err
		}
		data := g.Map{
			"id":              in.Id,
			"order_status":    consts.OrderStatusRecorded,
			"created_by":      userId,
			"telemarketer_id": userId,
			"is_deleted":      0,
		}
		for k, v := range recordData {
			data[k] = v
		}
		_, err = dao.BizOrder.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return 0, err
		}
		// 已创建订单并置为已录单，直接返回新 ID，后续步骤字段走正常更新
		return uint64(in.Id), nil
	}

	// 校验当前订单是否存在且未删除
	var cur entity.BizOrder
	err = dao.BizOrder.Ctx(ctx).Where("id", in.Id).Where("is_deleted", 0).Scan(&cur)
	if err != nil {
		return 0, err
	}
	if cur.Id == 0 {
		return 0, gerror.New("订单不存在")
	}

	// 校验步骤权限（角色 + 当前状态）
	if !canStepFromStatus(role, cur.OrderStatus, in.Step) {
		return 0, permissionDenied("当前状态不允许执行该步骤")
	}

	// 数据可见性：只能操作自己可见/所属的订单
	if !canOperateOrder(ctx, role, cur.CreatedBy, cur.AgentId) {
		return 0, permissionDenied("只能操作自己名下/自己录的订单")
	}

	// 录单步骤：首次录入时校验未重复录入
	if in.Step == consts.OrderStatusRecorded && cur.OrderStatus == consts.OrderStatusNotDeal && cur.CreatedBy != 0 {
		return 0, permissionDenied("该订单已有人录入，不可重复录单")
	}

	// 组装本步骤需更新的字段
	data, err := s.buildStepData(in)
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, gerror.New("本步骤无字段可保存")
	}
	// 首次录单（原状态未成交）时锁定录单人与话务员为当前用户
	if in.Step == consts.OrderStatusRecorded && cur.OrderStatus == consts.OrderStatusNotDeal {
		data["created_by"] = userId
		data["telemarketer_id"] = userId
	}
	// 保存为暂存：仅更新本步骤字段，不改变订单状态（状态推进由 stepNext 接口负责）
	_, err = dao.BizOrder.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return uint64(in.Id), err
}

// StepNext 推进订单到下一步状态（仅更新 order_status，不保存任何字段）。
// 校验当前角色对当前状态到下一状态的推进权限。
func (s *sBizOrder) StepNext(ctx context.Context, id uint64) error {
	role := currentRole(ctx)

	var cur entity.BizOrder
	err := dao.BizOrder.Ctx(ctx).Where("id", id).Where("is_deleted", 0).Scan(&cur)
	if err != nil {
		return err
	}
	if cur.Id == 0 {
		return gerror.New("订单不存在")
	}
	if cur.OrderStatus >= consts.OrderStatusCompleted {
		return gerror.New("订单已处于最终状态，无法继续推进")
	}

	nextStep := cur.OrderStatus + 1
	if !canStepFromStatus(role, cur.OrderStatus, nextStep) {
		return permissionDenied("当前角色不允许推进到下一步")
	}
	if !canOperateOrder(ctx, role, cur.CreatedBy, cur.AgentId) {
		return permissionDenied("只能操作自己名下/自己录的订单")
	}

	_, err = dao.BizOrder.Ctx(ctx).Where("id", id).Data(g.Map{"order_status": nextStep}).Update()
	return err
}

// buildStepData 按步骤组装待更新的字段
func (s *sBizOrder) buildStepData(in *adminin.BizOrderStepEditInp) (g.Map, error) {
	data := g.Map{}
	switch in.Step {
	case consts.OrderStatusRecorded: // 1 录单
		data["schedule_date"] = in.ScheduleDate
		data["customer_name"] = in.CustomerName
		data["available_time_desc"] = in.AvailableTimeDesc
		data["area"] = in.Area
		data["install_address"] = in.InstallAddress
		data["contact_phone"] = in.ContactPhone
		data["business_type"] = in.BusinessType
		data["expiry_date"] = in.ExpiryDate
	case consts.OrderStatusAssigned: // 2 派单
		if in.AgentId <= 0 {
			return nil, gerror.New("请选择收单员")
		}
		data["agent_id"] = in.AgentId
	case consts.OrderStatusAppointed: // 3 预约
		data["appointment_desc"] = in.AppointmentDesc
	case consts.OrderStatusVisited: // 4 上门
		data["visit_date"] = in.VisitDate
		data["dealt_business_type"] = in.DealtBusinessType
		data["porting_status"] = in.PortingStatus
		data["customer_real_name"] = in.CustomerRealName
		data["customer_id_number"] = in.CustomerIdNumber
		data["paid_amount"] = in.PaidAmount
		data["is_rural_order"] = in.IsRuralOrder
		data["new_phone_no"] = in.NewPhoneNo
		data["device_serial"] = in.DeviceSerial
		data["attachment_id"] = in.AttachmentId
		data["subsidy_amount"] = in.SubsidyAmount
		data["is_new"] = in.IsNew
	case consts.OrderStatusFollowed: // 5 回访
		data["follow_up_desc"] = in.FollowUpDesc
	case consts.OrderStatusCompleted: // 6 完成
		data["broadband_account"] = in.BroadbandAccount
		data["is_completed"] = in.IsCompleted
		data["agency_no"] = in.AgencyNo
	}
	return data, nil
}

// Delete 删除订单主表（逻辑删除，按角色+状态限制）
func (s *sBizOrder) Delete(ctx context.Context, id uint64) error {
	role := currentRole(ctx)

	var cur entity.BizOrder
	err := dao.BizOrder.Ctx(ctx).Where("id", id).Where("is_deleted", 0).Scan(&cur)
	if err != nil {
		return err
	}
	if cur.Id == 0 {
		return gerror.New("订单不存在")
	}
	if !canDelete(role, cur.OrderStatus) {
		return permissionDenied("当前角色不允许删除该状态下的订单")
	}
	_, err = dao.BizOrder.Ctx(ctx).Where("id", id).Data(g.Map{"is_deleted": 1}).Update()
	return err
}

// DetailList 订单详细情况列表：含操作人姓名，时间正序
func (s *sBizOrder) DetailList(ctx context.Context, orderId int64) (*adminin.BizOrderDetailListModel, error) {
	model := dao.BizOrderDetails.Ctx(ctx).As("d")
	model = model.LeftJoin("xy_admin_user u", "u.id = d.user_id")
	model = model.Fields("d.id, d.order_id, d.user_id, d.content, d.created_at, " +
		"CASE WHEN u.real_name != '' THEN u.real_name ELSE u.nickname END as user_name")
	var list []adminin.BizOrderDetailListItem
	err := model.Where("d.order_id", orderId).
		Where("d.is_deleted", 0).
		OrderAsc("d.created_at").
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.BizOrderDetailListItem{}
	}
	return &adminin.BizOrderDetailListModel{List: list}, nil
}

// DetailAdd 新增订单详细情况：user_id 取当前登录用户
func (s *sBizOrder) DetailAdd(ctx context.Context, in *adminin.BizOrderDetailAddInp) error {
	// 校验订单存在且对当前用户可见（与 List 数据可见性一致，不可操作看不到的订单）
	role := currentRole(ctx)
	model := dao.BizOrder.Ctx(ctx).Where("id", in.OrderId).Where("is_deleted", 0)
	if !canViewAll(role) {
		userId := currentUserId(ctx)
		switch role {
		case consts.RoleTelemarketer, consts.RoleTelemarketerManager:
			model = model.Where("created_by", userId)
		case consts.RoleAgent:
			model = model.Where("agent_id", userId)
		default:
			model = model.Where("created_by", userId)
		}
	}
	count, err := model.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("订单不存在或没有权限操作")
	}
	_, err = dao.BizOrderDetails.Ctx(ctx).Data(g.Map{
		"id":         gtime.TimestampNano(),
		"order_id":   in.OrderId,
		"user_id":    currentUserId(ctx),
		"content":    in.Content,
		"is_deleted": 0,
	}).Insert()
	return err
}

// DetailDelete 删除订单详细情况：只能删除自己录入的（逻辑删除）
func (s *sBizOrder) DetailDelete(ctx context.Context, id int64) error {
	var cur entity.BizOrderDetails
	err := dao.BizOrderDetails.Ctx(ctx).Where("id", id).Where("is_deleted", 0).Scan(&cur)
	if err != nil {
		return err
	}
	if cur.Id == 0 {
		return gerror.New("详细情况不存在")
	}
	if cur.UserId != currentUserId(ctx) {
		return permissionDenied("只能删除自己录入的详细情况")
	}
	_, err = dao.BizOrderDetails.Ctx(ctx).Where("id", id).Data(g.Map{"is_deleted": 1}).Update()
	return err
}

// Export 订单导出：使用模板文件填充并返回 Excel 文件流。
// 导出当前筛选条件下的全部数据（不限分页），"详细情况"列取该订单最新一条详细情况。
func (s *sBizOrder) Export(ctx context.Context, r *ghttp.Request, in *adminin.BizOrderListInp) error {
	// 查询全部符合条件的订单（复用列表筛选 + 数据可见性，不分页）
	var list []adminin.BizOrderListItem
	model := dao.BizOrder.Ctx(ctx).As("t")
	model = model.LeftJoin("xy_admin_user telemarketer", "telemarketer.id = t.telemarketer_id")
	model = model.LeftJoin("xy_admin_user agent", "agent.id = t.agent_id")
	model = model.Where("t.is_deleted", 0)
	model = applyScope(ctx, model, currentRole(ctx))
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
	model = model.Fields("t.*, " +
		"CASE WHEN telemarketer.real_name != '' THEN telemarketer.real_name ELSE telemarketer.nickname END as telemarketer_real_name, " +
		"CASE WHEN agent.real_name != '' THEN agent.real_name ELSE agent.nickname END as agent_real_name")
	err := model.OrderDesc("t.id").Scan(&list)
	if err != nil {
		return err
	}

	// 区域字典：value → label
	areaLabel := make(map[string]string)
	if dictItems, err := service.Dict().GetByType(ctx, "area"); err == nil {
		for _, item := range dictItems {
			areaLabel[item.Value] = item.Label
		}
	}

	// 打开模板
	f, err := excelize.OpenFile(filepath.Join("resource", "template", "download", "移动业务表格.xlsx"))
	if err != nil {
		return gerror.Wrap(err, "打开导出模板失败")
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	// 逐行填充，从第 2 行开始（第 1 行为表头）
	for i, order := range list {
		row := i + 2
		// 数据超出模板预置行数时追加新行
		if row > 279 {
			if err := f.DuplicateRow(sheet, 279); err != nil {
				return gerror.Wrap(err, "追加导出行失败")
			}
		}
		s.fillExportRow(ctx, f, sheet, row, &order, areaLabel)
	}

	// 写入响应
	fileName := fmt.Sprintf("移动业务表格-%s.xlsx", time.Now().Format("20060102150405"))
	buf, err := f.WriteToBuffer()
	if err != nil {
		return gerror.Wrap(err, "生成导出文件失败")
	}
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	r.Response.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	r.Response.Write(buf.Bytes())
	return nil
}

// fillExportRow 将单个订单填入模板指定行
func (s *sBizOrder) fillExportRow(ctx context.Context, f *excelize.File, sheet string, row int, order *adminin.BizOrderListItem, areaLabel map[string]string) {
	scheduleDate := ""
	if order.ScheduleDate != nil {
		scheduleDate = order.ScheduleDate.Format("2006-01-02")
	}
	area := ""
	if order.Area != nil {
		area = areaLabel[strconv.Itoa(*order.Area)]
	}
	rowStr := strconv.Itoa(row)
	values := map[string]interface{}{
		"A":  order.IsNew,                       // 是否纯新增
		"B":  scheduleDate,                      // 排单日期
		"C":  order.CustomerName,                // 姓名
		"D":  order.AvailableTimeDesc,           // 时间
		"E":  area,                              // 区域
		"F":  order.InstallAddress,              // 安装地址
		"G":  order.ContactPhone,                // 联系电话
		"H":  order.BusinessType,                // 预约业务
		"I":  order.ExpiryDate,                  // 到期
		"J":  order.PrimaryTelNo,                // 主卡号码
		"L":  order.TelemarketerRealName,        // 话务员
		"M":  order.AgentRealName,               // 收单员
		"N":  order.AppointmentDesc,             // 收单预约情况
		"O":  order.FollowUpDesc,                // 话务二次回访情况
		"P":  order.DealtBusinessType,           // 成交业务
		"Q":  order.PortingStatus,               // 携转情况
		"R":  order.CustomerRealName,            // 客户实际姓名
		"S":  stringVal(order.CustomerIdNumber), // 身份证号码
		"T":  float64Val(order.PaidAmount),      // 实缴额度
		"U":  intVal(order.IsRuralOrder),        // 是否乡下单
		"V":  order.NewPhoneNo,                  // 新开号码
		"W":  order.DeviceSerial,                // 终端串码
		"X":  order.BroadbandAccount,            // 宽带账号
		"Y":  intVal(order.IsCompleted),         // 是否完工
		"Z":  float64Val(order.SubsidyAmount),   // 话补
		"AA": order.AgencyNo,                    // 工号
		"AB": order.Remark,                      // 备注
	}
	// K 列「详细情况」取该订单最新一条
	values["K"] = s.latestDetail(ctx, order.Id)

	for col, val := range values {
		_ = f.SetCellValue(sheet, col+rowStr, val)
	}
}

// latestDetail 查询订单最新一条详细情况内容（created_at 降序取首条）
func (s *sBizOrder) latestDetail(ctx context.Context, orderId int64) string {
	var detail entity.BizOrderDetails
	err := dao.BizOrderDetails.Ctx(ctx).
		Where("order_id", orderId).
		Where("is_deleted", 0).
		OrderDesc("created_at").
		Limit(1).
		Scan(&detail)
	if err != nil || detail.Id == 0 {
		return ""
	}
	return detail.Content
}

// stringVal 指针字符串取值兜底
func stringVal(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// intVal 指针 int 取值兜底
func intVal(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// float64Val 指针 float64 取值兜底
func float64Val(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
