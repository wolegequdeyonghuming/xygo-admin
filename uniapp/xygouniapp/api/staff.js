import { get, post } from '@/utils/request'

/** 收单员登录（账号密码） */
export const staffLogin = (username, password) => post('/staff/auth/login', { username, password })

/** 当前收单员信息 */
export const getProfile = () => get('/staff/user/profile')

/** 订单列表：status 订单状态，month 月份筛选(YYYY-MM)，keyword 客户姓名/电话 */
export const getOrderList = (params) => get('/staff/order/list', params)

/** 订单详情 */
export const getOrderView = (id) => get('/staff/order/view', { id })

/** 预约（步骤3 保存并推进） */
export const orderAppoint = (params) => post('/staff/order/appoint', params)

/** 收单（步骤4 保存并推进 / 已上门重编辑） */
export const orderCollect = (params) => post('/staff/order/collect', params)

/** 详细情况列表 */
export const getDetailList = (orderId) => get('/staff/order/detailList', { orderId })

/** 新增详细情况 */
export const addDetail = (orderId, content) => post('/staff/order/detailAdd', { orderId, content })

/** 删除详细情况 */
export const deleteDetail = (id) => post('/staff/order/detailDelete', { id })
