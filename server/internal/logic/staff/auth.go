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

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/library/contexts"
	"xygo/internal/library/token"
	"xygo/internal/model"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/staffin"
)

// Login 收单员登录：账号密码校验 → 角色守卫（agent/agent_manager） → 签发后台 token
func (s *sStaffAuth) Login(ctx context.Context, in *staffin.LoginInp) (*staffin.LoginModel, error) {
	var user *entity.AdminUser
	if err := dao.AdminUser.Ctx(ctx).Where("username", in.Username).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil || user.Status != 1 {
		return nil, gerror.New("账号或密码错误")
	}
	// 密码校验：md5(password + salt)
	hashed := gmd5.MustEncryptString(in.Password + user.Salt)
	if hashed != user.Password {
		return nil, gerror.New("账号或密码错误")
	}

	// 查询用户第一个启用的角色
	var role *entity.AdminRole
	_ = dao.AdminRole.Ctx(ctx).
		LeftJoin(dao.AdminUserRole.Table()+" aur", "aur.role_id = "+dao.AdminRole.Table()+".id").
		Where("aur.user_id", user.Id).
		Where(dao.AdminRole.Table()+".status", 1).
		OrderAsc(dao.AdminRole.Table() + ".id").
		Limit(1).
		Scan(&role)

	var roleId uint64
	var roleKey string
	if role != nil {
		roleId = uint64(role.Id)
		roleKey = role.Key
	}

	// 角色守卫：仅收单员/收单员管理员可登录小程序
	if roleKey != consts.RoleAgent && roleKey != consts.RoleAgentManager {
		return nil, gerror.New("非收单员账号，无法登录收单小程序")
	}

	authUser := model.AuthUser{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Pid:      user.Pid,
		DeptId:   user.DeptId,
		RoleId:   roleId,
		RoleKey:  roleKey,
	}

	accessToken, refreshToken, expiresIn, refreshExpiresIn, err := token.Generate(ctx, authUser)
	if err != nil {
		return nil, err
	}

	return &staffin.LoginModel{
		Id:               user.Id,
		Username:         user.Username,
		Nickname:         user.Nickname,
		AccessToken:      accessToken,
		ExpiresIn:        expiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
	}, nil
}

// Refresh 用 refreshToken 刷新 accessToken（复用后台 admin 会话）
func (s *sStaffAuth) Refresh(ctx context.Context, refreshToken string) (*staffin.RefreshModel, error) {
	userId, err := token.ValidateRefreshToken(ctx, token.AppAdmin, refreshToken)
	if err != nil {
		return nil, gerror.New("刷新令牌无效或已过期，请重新登录")
	}

	var user *entity.AdminUser
	if err = dao.AdminUser.Ctx(ctx).Where("id", userId).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil || user.Status != 1 {
		return nil, gerror.New("用户不存在或已禁用")
	}

	// 重新查询角色，保持上下文角色正确
	var role *entity.AdminRole
	_ = dao.AdminRole.Ctx(ctx).
		LeftJoin(dao.AdminUserRole.Table()+" aur", "aur.role_id = "+dao.AdminRole.Table()+".id").
		Where("aur.user_id", user.Id).
		Where(dao.AdminRole.Table()+".status", 1).
		OrderAsc(dao.AdminRole.Table() + ".id").
		Limit(1).
		Scan(&role)

	var roleId uint64
	var roleKey string
	if role != nil {
		roleId = uint64(role.Id)
		roleKey = role.Key
	}
	if roleKey != consts.RoleAgent && roleKey != consts.RoleAgentManager {
		return nil, gerror.New("非收单员账号，无法使用收单小程序")
	}

	authUser := model.AuthUser{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Pid:      user.Pid,
		DeptId:   user.DeptId,
		RoleId:   roleId,
		RoleKey:  roleKey,
	}

	accessToken, expiresIn, err := token.RefreshAccessAdmin(ctx, refreshToken, authUser)
	if err != nil {
		return nil, gerror.New("刷新令牌无效或已过期，请重新登录")
	}
	return &staffin.RefreshModel{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
	}, nil
}

// Profile 当前收单员信息（姓名/部门/手机号/岗位）
func (s *sStaffAuth) Profile(ctx context.Context) (*staffin.ProfileModel, error) {
	userId := contexts.GetUserId(ctx)
	var user *entity.AdminUser
	if err := dao.AdminUser.Ctx(ctx).Where("id", userId).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 部门名
	var deptName string
	if user.DeptId > 0 {
		var dept *entity.AdminDept
		_ = dao.AdminDept.Ctx(ctx).Where("id", user.DeptId).Scan(&dept)
		if dept != nil {
			deptName = dept.Name
		}
	}

	// 岗位名（第一个启用的岗位）
	var postName string
	var postRows []struct {
		Name string `json:"name"`
	}
	_ = dao.AdminPost.Ctx(ctx).
		LeftJoin(dao.AdminUserPost.Table()+" aup", "aup.post_id = "+dao.AdminPost.Table()+".id").
		Where("aup.user_id", user.Id).
		Where(dao.AdminPost.Table()+".status", 1).
		OrderAsc(dao.AdminPost.Table() + ".id").
		Limit(1).
		Fields(dao.AdminPost.Columns().Name).
		Scan(&postRows)
	if len(postRows) > 0 {
		postName = postRows[0].Name
	}

	// 角色名
	roleKey := contexts.GetRoleKey(ctx)
	roleName := ""
	if roleKey != "" {
		var role *entity.AdminRole
		_ = dao.AdminRole.Ctx(ctx).Where("key", roleKey).Scan(&role)
		if role != nil {
			roleName = role.Name
		}
	}

	return &staffin.ProfileModel{
		Id:       user.Id,
		Username: user.Username,
		RealName: user.RealName,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		DeptName: deptName,
		PostName: postName,
		RoleKey:  roleKey,
		RoleName: roleName,
	}, nil
}
