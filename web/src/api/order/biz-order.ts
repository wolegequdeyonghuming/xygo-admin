/**
 * 订单主表管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchBizOrderList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/biz-order/list',
    params
  })
}

/** 详情 */
export function fetchBizOrderView(id: number) {
  return adminRequest.get<any>({
    url: '/biz-order/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchBizOrderEdit(params: any) {
  return adminRequest.post<any>({
    url: '/biz-order/edit',
    params
  })
}

/** 删除 */
export function fetchBizOrderDelete(id: number) {
  return adminRequest.post<any>({
    url: '/biz-order/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchBizOrderExport(params?: any) {
  return adminRequest.get<any>({
    url: '/biz-order/export',
    params,
    responseType: 'blob'
  })
}
