// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package middleware

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"xygo/internal/consts"
	"xygo/internal/library/contexts"
	"xygo/internal/library/token"
)

// StaffAuth 收单小程序接口鉴权中间件
// - 放行 /staff/auth/login
// - 其它 /staff/** 必须携带有效后台 accessToken（Authorization: Bearer XXX）
// - 角色守卫：仅收单员(agent)/收单员管理员(agent_manager)可访问
// - 校验通过后注入用户上下文（复用后台 admin 用户体系，bizorder 权限逻辑自动生效）
func StaffAuth(r *ghttp.Request) {
	path := r.URL.Path

	// 初始化自定义上下文
	customCtx := &contexts.Context{
		Module: "staff",
	}
	contexts.Init(r, customCtx)

	// 登录与刷新接口放行
	if path == "/staff/auth/login" || path == "/staff/auth/refresh" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "未登录"))
		return
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "未登录"))
		return
	}

	authUser, err := token.Parse(r.Context(), tokenStr)
	if err != nil {
		if errors.Is(err, token.ErrTokenKicked) {
			r.SetError(gerror.NewCode(consts.CodeKickedOut, "您的账号已在其他设备登录，请重新登录"))
			return
		}
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "登录已失效，请重新登录"))
		return
	}

	// 角色守卫：仅收单员/收单员管理员
	if authUser.RoleKey != consts.RoleAgent && authUser.RoleKey != consts.RoleAgentManager {
		r.SetError(gerror.NewCode(consts.CodeNoPermission, "无权限访问收单小程序"))
		return
	}

	contexts.SetUser(r.Context(), authUser)
	r.Middleware.Next()
}
