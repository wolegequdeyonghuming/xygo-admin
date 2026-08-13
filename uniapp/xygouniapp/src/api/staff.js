import { get, post } from '@/utils/request'
import config from '@/utils/config'

/** 收单员登录（账号密码） */
export const staffLogin = (username, password) => post('/staff/auth/login', { username, password })

/** 当前收单员信息 */
export const getProfile = () => get('/staff/user/profile')

/** 订单列表：statuses 状态数组，month 月份筛选(YYYY-MM)，keyword 客户姓名/电话 */
export const getOrderList = (params) => get('/staff/order/list', params)

/** 订单统计：待处理/今日已处理/已处理 */
export const getOrderStat = () => get('/staff/order/stat')

/** 订单详情 */
export const getOrderView = (id) => get('/staff/order/view', { id })

/** 预约（步骤3 保存并推进） */
export const orderAppoint = (params) => post('/staff/order/appoint', params)

/** 收单（步骤4 保存；finish=true 推进到已上门） */
export const orderCollect = (params) => post('/staff/order/collect', params)

/** 排单（步骤2，收单员管理员） */
export const orderSchedule = (params) => post('/staff/order/schedule', params)

/** 完工（步骤6，文员/管理员） */
export const orderComplete = (params) => post('/staff/order/complete', params)

/** 收单员列表（排单选人） */
export const getAgentList = () => get('/staff/user/agents')

/** 详细情况列表 */
export const getDetailList = (orderId) => get('/staff/order/detailList', { orderId })

/** 新增详细情况 */
export const addDetail = (orderId, content) => post('/staff/order/detailAdd', { orderId, content })

/** 删除详细情况 */
export const deleteDetail = (id) => post('/staff/order/detailDelete', { id })

/** 附件信息列表（按 id 逗号分隔） */
export const getAttachmentList = (ids) => get('/staff/attachment/list', { ids })

/** 字典数据（公开接口，如 type=area 区县） */
export const getDictData = (type) => get('/site/dict/data', { type })

/** 上传附件（复用后台 /admin/upload/file，multipart，返回 attachmentId） */
export function uploadFile(filePath) {
  return new Promise((resolve, reject) => {
    const token = uni.getStorageSync(config.TOKEN_KEY)
    uni.uploadFile({
      url: config.BASE_URL + '/admin/upload/file',
      filePath,
      name: 'file',
      header: token ? { Authorization: `Bearer ${token}` } : {},
      success: (res) => {
        try {
          const data = JSON.parse(res.data)
          if (data.code === 0) {
            resolve(data.data)
          } else {
            reject(data)
          }
        } catch (e) {
          reject(e)
        }
      },
      fail: reject
    })
  })
}
