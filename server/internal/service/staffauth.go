package service

import (
	"context"

	"xygo/internal/model/input/staffin"
)

// IStaffAuth 收单小程序认证服务
type IStaffAuth interface {
	// Login 收单员登录（账号密码 → 校验角色 → 签发后台 token）
	Login(ctx context.Context, in *staffin.LoginInp) (*staffin.LoginModel, error)
	// Refresh 用 refreshToken 刷新 accessToken
	Refresh(ctx context.Context, refreshToken string) (*staffin.RefreshModel, error)
	// Profile 当前收单员信息
	Profile(ctx context.Context) (*staffin.ProfileModel, error)
}

var localStaffAuth IStaffAuth

func StaffAuth() IStaffAuth {
	if localStaffAuth == nil {
		panic("implement not found for interface IStaffAuth, forgot register?")
	}
	return localStaffAuth
}

func RegisterStaffAuth(i IStaffAuth) {
	localStaffAuth = i
}
