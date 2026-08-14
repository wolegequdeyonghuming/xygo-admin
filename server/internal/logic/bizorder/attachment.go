// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package bizorder

import (
	"context"
	"strconv"
	"strings"

	"xygo/internal/dao"
	"xygo/internal/model/entity"
	adminin "xygo/internal/model/input/orderin"
)

// parseOrderAttachmentIds 解析订单 attachment_id（逗号分隔的数字ID），忽略非数字 token。
func parseOrderAttachmentIds(ids string) []int64 {
	out := make([]int64, 0)
	for _, p := range strings.Split(ids, ",") {
		p = strings.TrimSpace(p)
		if v, err := strconv.ParseInt(p, 10, 64); err == nil {
			out = append(out, v)
		}
	}
	return out
}

// queryAttachmentMap 批量查询附件信息，返回 id → 附件对象映射。
func queryAttachmentMap(ctx context.Context, ids []int64) (map[int64]adminin.AttachmentInfo, error) {
	m := make(map[int64]adminin.AttachmentInfo)
	if len(ids) == 0 {
		return m, nil
	}
	var rows []entity.SysAttachment
	if err := dao.SysAttachment.Ctx(ctx).WhereIn("id", ids).Scan(&rows); err != nil {
		return nil, err
	}
	for _, r := range rows {
		m[int64(r.Id)] = adminin.AttachmentInfo{
			Id:       r.Id,
			Url:      r.Url,
			Name:     r.Name,
			Size:     r.Size,
			Mimetype: r.Mimetype,
		}
	}
	return m, nil
}

// resolveAttachments 按原顺序输出附件对象列表（数字ID在 map 中未命中的跳过）。
func resolveAttachments(ids string, m map[int64]adminin.AttachmentInfo) []adminin.AttachmentInfo {
	out := make([]adminin.AttachmentInfo, 0)
	for _, id := range parseOrderAttachmentIds(ids) {
		if a, ok := m[id]; ok {
			out = append(out, a)
		}
	}
	return out
}

// FillOrderListAttachments 批量联查订单列表的附件对象并填充到每一项。
func FillOrderListAttachments(ctx context.Context, list []adminin.BizOrderListItem) error {
	all := make([]int64, 0)
	for _, it := range list {
		all = append(all, parseOrderAttachmentIds(it.AttachmentId)...)
	}
	m, err := queryAttachmentMap(ctx, all)
	if err != nil {
		return err
	}
	for i := range list {
		list[i].Attachments = resolveAttachments(list[i].AttachmentId, m)
	}
	return nil
}

// FillOrderViewAttachments 联查订单详情的附件对象并填充。
func FillOrderViewAttachments(ctx context.Context, view *adminin.BizOrderViewModel) error {
	m, err := queryAttachmentMap(ctx, parseOrderAttachmentIds(view.AttachmentId))
	if err != nil {
		return err
	}
	view.Attachments = resolveAttachments(view.AttachmentId, m)
	return nil
}
