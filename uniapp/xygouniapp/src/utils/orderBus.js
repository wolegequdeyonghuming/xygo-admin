// 订单详情数据总线：列表页跳转前写入，详情/表单页优先读取，避免 query 解析异常导致无数据
import { ref } from 'vue'

export const currentOrder = ref(null)

const STORAGE_KEY = 'xygo_current_order'

// 跳转前写入：模块态 + 本地存储双保险
export function setCurrentOrder(order) {
  currentOrder.value = order
  try {
    uni.setStorageSync(STORAGE_KEY, order)
  } catch (e) {
    // 存储失败忽略
  }
}

// 读取：模块态优先，其次本地存储
export function getCurrentOrder() {
  if (currentOrder.value) return currentOrder.value
  try {
    const saved = uni.getStorageSync(STORAGE_KEY)
    if (saved) {
      currentOrder.value = saved
      return saved
    }
  } catch (e) {
    // 忽略
  }
  return null
}
