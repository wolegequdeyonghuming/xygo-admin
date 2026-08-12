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
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/staffin"
)

// LoginReq 收单员登录请求
type LoginReq struct {
	g.Meta `path:"/auth/login" method:"post" tags:"Staff" summary:"收单员登录"`
	staffin.LoginInp
}

type LoginRes struct {
	*staffin.LoginModel
}

// RefreshReq 刷新 accessToken 请求
type RefreshReq struct {
	g.Meta `path:"/auth/refresh" method:"post" tags:"Staff" summary:"刷新访问令牌"`
	staffin.RefreshInp
}

type RefreshRes struct {
	*staffin.RefreshModel
}

// ProfileReq 收单员信息请求
type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"Staff" summary:"收单员信息"`
}

type ProfileRes struct {
	*staffin.ProfileModel
}
