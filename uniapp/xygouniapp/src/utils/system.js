// 系统信息（状态栏高度等），供自定义导航避让使用
let cached = null

export function initSystem() {
  try {
    cached = uni.getSystemInfoSync()
  } catch (e) {
    cached = {}
  }
  return cached
}

export function getSystemInfo() {
  if (!cached) initSystem()
  return cached || {}
}

// 状态栏高度（rpx）：小程序自定义导航需在顶部避让
export function getStatusBarHeightRpx() {
  const info = getSystemInfo()
  const sw = info.screenWidth || 375
  const sh = info.statusBarHeight || 0
  return Math.round((sh / sw) * 750)
}
