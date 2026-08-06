// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package dict

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/model/do"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/adminin"
	"xygo/internal/service"
)

type sDict struct{}

func init() {
	service.RegisterDict(New())
}

// New 构造字典服务
func New() *sDict {
	return &sDict{}
}

// TypeList 获取字典类型列表
func (s *sDict) TypeList(ctx context.Context, in *adminin.DictTypeListInp) (list []adminin.DictTypeListItem, total int, err error) {
	model := dao.SysDictType.Ctx(ctx)

	if in.Name != "" {
		model = model.WhereLike("name", "%"+in.Name+"%")
	}
	if in.Type != "" {
		model = model.WhereLike("type", "%"+in.Type+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		model = model.Where("status", in.Status)
	}

	count, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}

	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}

	err = model.
		Fields("id, name, type, remark, status, sort, create_time, update_time").
		OrderAsc("sort, id").
		Page(in.Page, in.PageSize).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

// TypeDetail 获取字典类型详情
func (s *sDict) TypeDetail(ctx context.Context, id uint64) (*adminin.DictTypeListItem, error) {
	var dictType *entity.SysDictType

	if err := dao.SysDictType.Ctx(ctx).Where("id", id).Scan(&dictType); err != nil {
		return nil, err
	}
	if dictType == nil {
		return nil, gerror.NewCode(consts.CodeDataNotFound, "字典类型不存在")
	}

	return &adminin.DictTypeListItem{
		Id:         dictType.Id,
		Name:       dictType.Name,
		Type:       dictType.Type,
		Remark:     dictType.Remark,
		Status:     dictType.Status,
		Sort:       dictType.Sort,
		CreateTime: int(dictType.CreateTime),
		UpdateTime: int(dictType.UpdateTime),
	}, nil
}

// TypeSave 保存字典类型（新增/编辑）
func (s *sDict) TypeSave(ctx context.Context, in *adminin.DictTypeSaveInp) (uint, error) {
	count, err := dao.SysDictType.Ctx(ctx).
		Where("type", in.Type).
		WhereNot("id", in.Id).
		Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(consts.CodeInvalidParam, "字典标识已存在")
	}

	data := do.SysDictType{
		Name:   in.Name,
		Type:   in.Type,
		Sort:   in.Sort,
		Status: in.Status,
		Remark: in.Remark,
	}

	if in.Id == 0 {
		r, err := dao.SysDictType.Ctx(ctx).Data(data).OmitNil().Insert()
		if err != nil {
			return 0, err
		}
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		return uint(lastId), nil
	}

	_, err = dao.SysDictType.Ctx(ctx).
		Data(data).
		OmitNil().
		Where("id", in.Id).
		Update()
	if err != nil {
		return 0, err
	}
	return uint(in.Id), nil
}

// TypeDelete 删除字典类型
func (s *sDict) TypeDelete(ctx context.Context, id uint64) error {
	dataCount, err := dao.SysDictData.Ctx(ctx).
		Where("dict_type_id", id).
		Count()
	if err != nil {
		return err
	}
	if dataCount > 0 {
		return gerror.NewCode(consts.CodeInvalidParam, "该字典类型下还有字典数据，无法删除")
	}

	_, err = dao.SysDictType.Ctx(ctx).Where("id", id).Delete()
	return err
}

// TypeOptions 获取启用的字典类型选项（供下拉选择）
func (s *sDict) TypeOptions(ctx context.Context) ([]adminin.DictTypeOptionsItem, error) {
	var items []adminin.DictTypeOptionsItem

	err := dao.SysDictType.Ctx(ctx).
		Where("status", 1).
		Fields("name as label, type as value").
		OrderAsc("sort, id").
		Scan(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// DataList 获取字典数据列表
func (s *sDict) DataList(ctx context.Context, in *adminin.DictDataListInp) (list []adminin.DictDataListItem, total int, err error) {
	model := dao.SysDictData.Ctx(ctx).Where("dict_type_id", in.DictTypeId)

	if in.Label != "" {
		model = model.WhereLike("label", "%"+in.Label+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		model = model.Where("status", in.Status)
	}

	count, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}

	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}

	err = model.
		Fields("id, dict_type_id, label, value, css_class, list_class, is_default, status, sort, remark, create_time, update_time").
		OrderAsc("sort, id").
		Page(in.Page, in.PageSize).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

// DataSave 保存字典数据（新增/编辑）
func (s *sDict) DataSave(ctx context.Context, in *adminin.DictDataSaveInp) (uint, error) {
	// 校验所属字典类型是否存在
	typeCount, err := dao.SysDictType.Ctx(ctx).Where("id", in.DictTypeId).Count()
	if err != nil {
		return 0, err
	}
	if typeCount == 0 {
		return 0, gerror.NewCode(consts.CodeInvalidParam, "字典类型不存在")
	}

	data := do.SysDictData{
		DictTypeId: in.DictTypeId,
		Label:      in.Label,
		Value:      in.Value,
		CssClass:   in.CssClass,
		ListClass:  in.ListClass,
		IsDefault:  in.IsDefault,
		Sort:       in.Sort,
		Status:     in.Status,
		Remark:     in.Remark,
	}

	if in.Id == 0 {
		r, err := dao.SysDictData.Ctx(ctx).Data(data).OmitNil().Insert()
		if err != nil {
			return 0, err
		}
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		return uint(lastId), nil
	}

	_, err = dao.SysDictData.Ctx(ctx).
		Data(data).
		OmitNil().
		Where("id", in.Id).
		Update()
	if err != nil {
		return 0, err
	}
	return uint(in.Id), nil
}

// DataDelete 删除字典数据
func (s *sDict) DataDelete(ctx context.Context, id uint64) error {
	_, err := dao.SysDictData.Ctx(ctx).Where("id", id).Delete()
	return err
}

// GetByType 按字典类型标识获取启用的字典数据（公开接口使用）
func (s *sDict) GetByType(ctx context.Context, dictType string) ([]adminin.DictDataItem, error) {
	var dictTypeEntity *entity.SysDictType

	if err := dao.SysDictType.Ctx(ctx).Where("type", dictType).Where("status", 1).Scan(&dictTypeEntity); err != nil {
		return nil, err
	}
	if dictTypeEntity == nil {
		return nil, gerror.NewCode(consts.CodeDataNotFound, "字典类型不存在")
	}

	var items []adminin.DictDataItem
	err := dao.SysDictData.Ctx(ctx).
		Where("dict_type_id", dictTypeEntity.Id).
		Where("status", 1).
		Fields("label, value, css_class, list_class, is_default").
		OrderAsc("sort, id").
		Scan(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
