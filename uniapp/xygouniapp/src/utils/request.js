import config from './config'

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
          if (data.code === 0) {
            resolve(data.data)
          } else if (data.code === 401) {
            handleUnauthorized(options, resolve, reject, data)
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

// 用 refreshToken 换新 accessToken（48h 会话内自动续期）
function refreshAccessToken() {
  return new Promise((resolve, reject) => {
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
        if (data && data.code === 0 && data.data && data.data.accessToken) {
          uni.setStorageSync(config.TOKEN_KEY, data.data.accessToken)
          resolve(data.data.accessToken)
        } else {
          reject(new Error('refresh failed'))
        }
      },
      fail: reject,
    })
  })
}

// 401 处理：先尝试刷新 token 并重放原请求，失败则清除会话回登录页
function handleUnauthorized(options, resolve, reject, originalData) {
  const logout = () => {
    uni.removeStorageSync(config.TOKEN_KEY)
    uni.removeStorageSync(config.REFRESH_TOKEN_KEY)
    uni.showToast({ title: '请先登录', icon: 'none' })
    uni.reLaunch({ url: '/pages/login/index' })
    reject(originalData)
  }

  if (options._retried) {
    logout()
    return
  }
  refreshAccessToken()
    .then(() => request({ ...options, _retried: true }))
    .then(resolve)
    .catch(() => {
      logout()
    })
}

export const get = (url, data) => request({ url, method: 'GET', data })
export const post = (url, data) => request({ url, method: 'POST', data })

export default request
