// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package site

import (
	"context"

	api "xygo/api/site"
	"xygo/internal/service"
)

// DictData 按类型标识获取启用的字典数据（公开接口，无需登录）
func (c *ControllerV1) DictData(ctx context.Context, req *api.DictDataReq) (res *api.DictDataRes, err error) {
	list, err := service.Dict().GetByType(ctx, req.Type)
	if err != nil {
		return nil, err
	}
	return &api.DictDataRes{List: list}, nil
}
