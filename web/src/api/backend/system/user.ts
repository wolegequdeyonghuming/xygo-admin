/**
 * 用户管理 API
 * @module api/backend/system/user
 */
import { adminRequest } from '@/utils/http'

/**
 * 获取用户列表
 */
export function fetchGetUserList(params: any) {
  return adminRequest.get<{ 
    list: Api.SystemManage.UserListItem[]
    page: number
    pageSize: number
    total: number 
  }>({
    url: '/user/list',
    params
  })
}

/**
 * 获取用户详情（编辑用，未脱敏）
 */
export function fetchGetUserDetail(id: number) {
  return adminRequest.get<any>({
    url: '/user/detail',
    params: { id }
  })
}

/**
 * 保存用户（新增/编辑）
 */
export function fetchSaveUser(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/user/save',
    params
  })
}

/**
 * 删除用户
 */
export function fetchDeleteUser(id: number) {
  return adminRequest.post({
    url: '/user/delete',
    params: { id }
  })
}

/**
 * 强制用户下线
 */
export function fetchKickUser(id: number) {
  return adminRequest.post({
    url: '/user/kick',
    params: { id }
  })
}

/**
 * 人员选择器（部门树，叶子为用户）
 */
export interface PersonSelectorNode {
  value: number
  label: string
  disabled?: boolean
  children?: PersonSelectorNode[]
}

export interface PersonSelectorParams {
  deptIds?: number[]
  roleIds?: number[]
  postIds?: number[]
  keyword?: string
}

export function fetchPersonSelector(params: PersonSelectorParams) {
  return adminRequest.get<{ list: PersonSelectorNode[] }>({
    url: '/user/selector',
    params
  })
}
