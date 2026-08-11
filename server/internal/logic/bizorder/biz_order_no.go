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
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// genBizOrderNo 生成当日订单编号：yyyyMMdd + 4 位当日序号（每天从 0001 递增、次日重置）。
// 使用 Redis 原子自增（INCR），键名含日期，天然按天重置。
func genBizOrderNo(ctx context.Context) (string, error) {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("biz_order_no:%s", date)
	seq, err := g.Redis().Incr(ctx, key)
	if err != nil {
		return "", err
	}
	if seq > 9999 {
		return "", gerror.New("当日订单编号序号已用尽")
	}
	return fmt.Sprintf("%s%04d", date, seq), nil
}

// isDuplicateEntryErr 判断是否为唯一索引冲突错误（MySQL Duplicate entry）。
// 用于订单编号生成时遇到重复编号（如缓存重置导致的撞号）自动重试。
func isDuplicateEntryErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "Duplicate entry")
}
