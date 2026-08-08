// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xygo/internal/model/input/adminin"
)

type (
	IAdminUser interface {
		// List 获取用户列表
		List(ctx context.Context, in *adminin.UserListInp) (list []adminin.UserListItem, total int, err error)
		// Detail 获取用户详情
		Detail(ctx context.Context, id uint64) (*adminin.UserListItem, error)
		// DetailForEdit 获取用户详情（未脱敏，编辑用，含角色和岗位ID）
		DetailForEdit(ctx context.Context, id uint64) (*adminin.UserDetailModel, error)
		// Save 保存用户（新增/编辑）
		Save(ctx context.Context, in *adminin.UserSaveInp) (uint, error)
		// Delete 删除用户
		Delete(ctx context.Context, id uint64) error
		// Selector 人员选择器：按部门/角色/岗位过滤后组装部门树（叶子为用户）
		// 过滤规则：同类型数组内为并集(OR)，不同类型之间为交集(AND)
		Selector(ctx context.Context, in *adminin.UserSelectorInp) ([]*adminin.UserSelectorNode, error)
	}
)

var (
	localAdminUser IAdminUser
)

func AdminUser() IAdminUser {
	if localAdminUser == nil {
		panic("implement not found for interface IAdminUser, forgot register?")
	}
	return localAdminUser
}

func RegisterAdminUser(i IAdminUser) {
	localAdminUser = i
}
