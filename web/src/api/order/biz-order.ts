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

/** 分步保存(录单/派单/预约/收单/回访/完成) */
export function fetchBizOrderStepEdit(params: any) {
  return adminRequest.post<any>({
    url: '/biz-order/stepEdit',
    params
  })
}

/** 推进订单到下一步（仅改变状态） */
export function fetchBizOrderStepNext(id: number) {
  return adminRequest.post<any>({
    url: '/biz-order/stepNext',
    params: { id }
  })
}

/** 详细情况列表 */
export function fetchBizOrderDetailList(orderId: number) {
  return adminRequest.get<any>({
    url: '/biz-order/detailList',
    params: { orderId }
  })
}

/** 新增详细情况 */
export function fetchBizOrderDetailAdd(params: { orderId: number; content: string }) {
  return adminRequest.post<any>({
    url: '/biz-order/detailAdd',
    params
  })
}

/** 删除详细情况 */
export function fetchBizOrderDetailDelete(id: number) {
  return adminRequest.post<any>({
    url: '/biz-order/detailDelete',
    params: { id }
  })
}

/** 删除 */
export function fetchBizOrderDelete(id: number) {
  return adminRequest.post<any>({
    url: '/biz-order/delete',
    params: { id }
  })
}

/** 导出（返回 Excel 文件 Blob） */
export function fetchBizOrderExport(params?: any): Promise<Blob> {
  return adminRequest.get<Blob>({
    url: '/biz-order/export',
    params,
    responseType: 'blob'
  })
}
