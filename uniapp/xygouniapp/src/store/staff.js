import { ref, computed } from 'vue'
import { getProfile } from '@/api/staff'
import config from '@/utils/config'
import { mockProfile } from '@/utils/mock'

const token = ref(uni.getStorageSync(config.TOKEN_KEY) || '')
const refreshToken = ref(uni.getStorageSync(config.REFRESH_TOKEN_KEY) || '')
const userInfo = ref(null)

export function useStaffStore() {
  const isLoggedIn = computed(() => !!token.value)

  function setToken(val) {
    token.value = val
    if (val) {
      uni.setStorageSync(config.TOKEN_KEY, val)
    } else {
      uni.removeStorageSync(config.TOKEN_KEY)
    }
  }

  function setRefreshToken(val) {
    refreshToken.value = val
    if (val) {
      uni.setStorageSync(config.REFRESH_TOKEN_KEY, val)
    } else {
      uni.removeStorageSync(config.REFRESH_TOKEN_KEY)
    }
  }

  function setUserInfo(val) {
    userInfo.value = val
  }

  async function fetchProfile() {
    if (!token.value) return null
    // mock 预览：无后端时返回 mock 资料
    if (config.MOCK_PREVIEW) {
      userInfo.value = mockProfile
      return userInfo.value
    }
    try {
      const data = await getProfile()
      userInfo.value = data
      return data
    } catch (e) {
      console.error('获取用户信息失败', e)
      return null
    }
  }

  function logout() {
    setToken('')
    setRefreshToken('')
    userInfo.value = null
    uni.reLaunch({ url: '/pages/login/index' })
  }

  return { token, refreshToken, userInfo, isLoggedIn, setToken, setRefreshToken, setUserInfo, fetchProfile, logout }
}
