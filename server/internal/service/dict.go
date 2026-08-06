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
	IDict interface {
		// TypeList 获取字典类型列表
		TypeList(ctx context.Context, in *adminin.DictTypeListInp) (list []adminin.DictTypeListItem, total int, err error)
		// TypeDetail 获取字典类型详情
		TypeDetail(ctx context.Context, id uint64) (*adminin.DictTypeListItem, error)
		// TypeSave 保存字典类型（新增/编辑）
		TypeSave(ctx context.Context, in *adminin.DictTypeSaveInp) (uint, error)
		// TypeDelete 删除字典类型
		TypeDelete(ctx context.Context, id uint64) error
		// TypeOptions 获取启用的字典类型选项（供下拉选择）
		TypeOptions(ctx context.Context) ([]adminin.DictTypeOptionsItem, error)
		// DataList 获取字典数据列表
		DataList(ctx context.Context, in *adminin.DictDataListInp) (list []adminin.DictDataListItem, total int, err error)
		// DataSave 保存字典数据（新增/编辑）
		DataSave(ctx context.Context, in *adminin.DictDataSaveInp) (uint, error)
		// DataDelete 删除字典数据
		DataDelete(ctx context.Context, id uint64) error
		// GetByType 按字典类型标识获取启用的字典数据（公开接口使用）
		GetByType(ctx context.Context, dictType string) ([]adminin.DictDataItem, error)
	}
)

var (
	localDict IDict
)

func Dict() IDict {
	if localDict == nil {
		panic("implement not found for interface IDict, forgot register?")
	}
	return localDict
}

func RegisterDict(i IDict) {
	localDict = i
}
