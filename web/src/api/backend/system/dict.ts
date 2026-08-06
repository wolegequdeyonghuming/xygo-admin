// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

/**
 * 字典管理 API
 * @module api/backend/system/dict
 */
import { adminRequest } from '@/utils/http'

export interface DictTypeItem {
  id: number
  name: string
  type: string
  remark: string
  status: number
  sort: number
  createTime: number
  updateTime: number
}

export interface DictDataItem {
  id: number
  dictTypeId: number
  label: string
  value: string
  cssClass: string
  listClass: string
  isDefault: number
  status: number
  sort: number
  remark: string
  createTime: number
  updateTime: number
}

export interface DictTypeOption {
  label: string
  value: string
}

/**
 * 获取字典类型列表
 */
export function fetchDictTypeList(params: any) {
  return adminRequest.get<{
    list: DictTypeItem[]
    total: number
    page: number
    pageSize: number
  }>({
    url: '/dict/type/list',
    params
  })
}

/**
 * 获取字典类型详情
 */
export function fetchDictTypeDetail(id: number) {
  return adminRequest.get<DictTypeItem>({
    url: '/dict/type/detail',
    params: { id }
  })
}

/**
 * 保存字典类型
 */
export function fetchSaveDictType(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/dict/type/save',
    params
  })
}

/**
 * 删除字典类型
 */
export function fetchDeleteDictType(id: number) {
  return adminRequest.post({
    url: '/dict/type/delete',
    params: { id }
  })
}

/**
 * 获取字典类型选项（下拉）
 */
export function fetchDictTypeOptions() {
  return adminRequest.get<{ list: DictTypeOption[] }>({
    url: '/dict/type/options'
  })
}

/**
 * 获取字典数据列表
 */
export function fetchDictDataList(params: any) {
  return adminRequest.get<{
    list: DictDataItem[]
    total: number
    page: number
    pageSize: number
  }>({
    url: '/dict/data/list',
    params
  })
}

/**
 * 保存字典数据
 */
export function fetchSaveDictData(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/dict/data/save',
    params
  })
}

/**
 * 删除字典数据
 */
export function fetchDeleteDictData(id: number) {
  return adminRequest.post({
    url: '/dict/data/delete',
    params: { id }
  })
}
