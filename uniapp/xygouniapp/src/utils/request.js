import config from './config'

// 业务状态码（与后端 internal/consts/code.go 保持一致）
const CODE_SUCCESS = 0
const CODE_UNAUTHORIZED = 61 // gcode.CodeNotAuthorized：未登录/登录失效
const CODE_KICKED_OUT = 10010 // 被踢下线（SSO单点登录/管理员强制下线）

const request = (options = {}) => {
  return new Promise((resolve, reject) => {
    const token = uni.getStorageSync(config.TOKEN_KEY)

    const header = {
      'Content-Type': 'application/json',
      ...options.header,
    }
    if (token) {
      header['Authorization'] = `Bearer ${token}`
    }

    uni.request({
      url: config.BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data,
      header,
      success: (res) => {
        if (res.statusCode === 200) {
          const data = res.data
          if (data.code === CODE_SUCCESS) {
            resolve(data.data)
          } else if (data.code === CODE_UNAUTHORIZED) {
            handleUnauthorized(options, resolve, reject, data)
          } else if (data.code === CODE_KICKED_OUT) {
            logout(data.message)
            reject(data)
          } else {
            uni.showToast({ title: data.message || '请求失败', icon: 'none' })
            reject(data)
          }
        } else {
          uni.showToast({ title: `网络错误 ${res.statusCode}`, icon: 'none' })
          reject(res)
        }
      },
      fail: (err) => {
        uni.showToast({ title: '网络连接失败', icon: 'none' })
        reject(err)
      },
    })
  })
}

// 清除本地会话并跳转登录页
function logout(message = '请先登录') {
  uni.removeStorageSync(config.TOKEN_KEY)
  uni.removeStorageSync(config.REFRESH_TOKEN_KEY)
  uni.showToast({ title: message, icon: 'none' })
  uni.reLaunch({ url: '/pages/login/index' })
}

let refreshPromise = null

// 用 refreshToken 换新 accessToken（48h 会话内自动续期）
// 单飞模式：并发的多个未授权请求共享同一个刷新 Promise，避免重复刷新
function refreshAccessToken() {
  if (refreshPromise) return refreshPromise
  refreshPromise = new Promise((resolve, reject) => {
    const refreshToken = uni.getStorageSync(config.REFRESH_TOKEN_KEY)
    if (!refreshToken) {
      reject(new Error('no refresh token'))
      return
    }
    uni.request({
      url: config.BASE_URL + '/staff/auth/refresh',
      method: 'POST',
      data: { refreshToken },
      header: { 'Content-Type': 'application/json' },
      success: (res) => {
        const data = res.data
        if (data && data.code === CODE_SUCCESS && data.data && data.data.accessToken) {
          uni.setStorageSync(config.TOKEN_KEY, data.data.accessToken)
          resolve(data.data.accessToken)
        } else {
          reject(new Error('refresh failed'))
        }
      },
      fail: reject,
    })
  }).finally(() => {
    refreshPromise = null
  })
  return refreshPromise
}

// 未授权处理：先尝试刷新 token 并重放原请求，失败则清除会话回登录页
function handleUnauthorized(options, resolve, reject, originalData) {
  if (options._retried) {
    logout(originalData.message)
    reject(originalData)
    return
  }
  refreshAccessToken()
    .then(() => request({ ...options, _retried: true }))
    .then(resolve)
    .catch(() => {
      logout(originalData.message)
      reject(originalData)
    })
}

export const get = (url, data) => request({ url, method: 'GET', data })
export const post = (url, data) => request({ url, method: 'POST', data })

export default request
