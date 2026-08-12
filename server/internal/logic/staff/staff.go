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
	"xygo/internal/service"
)

type sStaffAuth struct{}
type sStaffOrder struct{}

func init() {
	service.RegisterStaffAuth(NewStaffAuth())
	service.RegisterStaffOrder(NewStaffOrder())
}

func NewStaffAuth() *sStaffAuth {
	return &sStaffAuth{}
}

func NewStaffOrder() *sStaffOrder {
	return &sStaffOrder{}
}
