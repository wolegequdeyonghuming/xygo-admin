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

	api "xygo/api/staff"
	"xygo/internal/service"
)

// Login 收单员登录
func (c *ControllerV1) Login(ctx context.Context, req *api.LoginReq) (res *api.LoginRes, err error) {
	result, err := service.StaffAuth().Login(ctx, &req.LoginInp)
	if err != nil {
		return nil, err
	}
	return &api.LoginRes{result}, nil
}

// Refresh 刷新 accessToken
func (c *ControllerV1) Refresh(ctx context.Context, req *api.RefreshReq) (res *api.RefreshRes, err error) {
	result, err := service.StaffAuth().Refresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &api.RefreshRes{result}, nil
}

// Profile 收单员信息
func (c *ControllerV1) Profile(ctx context.Context, req *api.ProfileReq) (res *api.ProfileRes, err error) {
	result, err := service.StaffAuth().Profile(ctx)
	if err != nil {
		return nil, err
	}
	return &api.ProfileRes{result}, nil
}
