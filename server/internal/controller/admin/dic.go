// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

// DictTypeList 字典类型列表
func (c *ControllerV1) DictTypeList(ctx context.Context, req *api.DictTypeListReq) (res *api.DictTypeListRes, err error) {
	list, total, err := service.Dict().TypeList(ctx, &req.DictTypeListInp)
	if err != nil {
		return nil, err
	}

	res = &api.DictTypeListRes{
		DictTypeListModel: adminin.DictTypeListModel{
			List: list,
			PageRes: form.PageRes{
				Page:     req.Page,
				PageSize: req.PageSize,
				Total:    total,
			},
		},
	}
	return
}

// DictTypeDetail 字典类型详情
func (c *ControllerV1) DictTypeDetail(ctx context.Context, req *api.DictTypeDetailReq) (res *api.DictTypeDetailRes, err error) {
	detail, err := service.Dict().TypeDetail(ctx, uint64(req.Id))
	if err != nil {
		return nil, err
	}

	res = &api.DictTypeDetailRes{DictTypeListItem: *detail}
	return
}

// DictTypeSave 字典类型保存
func (c *ControllerV1) DictTypeSave(ctx context.Context, req *api.DictTypeSaveReq) (res *api.DictTypeSaveRes, err error) {
	id, err := service.Dict().TypeSave(ctx, &req.DictTypeSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.DictTypeSaveRes{Id: id}
	return
}

// DictTypeDelete 字典类型删除
func (c *ControllerV1) DictTypeDelete(ctx context.Context, req *api.DictTypeDeleteReq) (res *api.DictTypeDeleteRes, err error) {
	err = service.Dict().TypeDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &api.DictTypeDeleteRes{}
	return
}

// DictTypeOptions 字典类型选项（下拉）
func (c *ControllerV1) DictTypeOptions(ctx context.Context, req *api.DictTypeOptionsReq) (res *api.DictTypeOptionsRes, err error) {
	list, err := service.Dict().TypeOptions(ctx)
	if err != nil {
		return nil, err
	}

	res = &api.DictTypeOptionsRes{List: list}
	return
}

// DictDataList 字典数据列表
func (c *ControllerV1) DictDataList(ctx context.Context, req *api.DictDataListReq) (res *api.DictDataListRes, err error) {
	list, total, err := service.Dict().DataList(ctx, &req.DictDataListInp)
	if err != nil {
		return nil, err
	}

	res = &api.DictDataListRes{
		DictDataListModel: adminin.DictDataListModel{
			List: list,
			PageRes: form.PageRes{
				Page:     req.Page,
				PageSize: req.PageSize,
				Total:    total,
			},
		},
	}
	return
}

// DictDataSave 字典数据保存
func (c *ControllerV1) DictDataSave(ctx context.Context, req *api.DictDataSaveReq) (res *api.DictDataSaveRes, err error) {
	id, err := service.Dict().DataSave(ctx, &req.DictDataSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.DictDataSaveRes{Id: id}
	return
}

// DictDataDelete 字典数据删除
func (c *ControllerV1) DictDataDelete(ctx context.Context, req *api.DictDataDeleteReq) (res *api.DictDataDeleteRes, err error) {
	err = service.Dict().DataDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &api.DictDataDeleteRes{}
	return
}
