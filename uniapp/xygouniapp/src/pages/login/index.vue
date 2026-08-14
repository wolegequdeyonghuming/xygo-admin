<template>
  <view class="login-page">
    <view class="bg-glow g1"></view>
    <view class="bg-glow g2"></view>

    <view class="logo-area">
      <view class="logo-box">
        <image class="logo" src="/static/logo.png" mode="aspectFit"></image>
      </view>
      <text class="app-name">{{ siteName || '移动业务平台' }}</text>
      <text class="app-desc">账号密码登录 · 外勤收单工作台</text>
    </view>

    <view class="form-card">
      <view class="field">
        <text class="field-label">账号</text>
        <wd-input v-model="username" placeholder="请输入收单员账号" clearable custom-class="form-input" />
      </view>
      <view class="field">
        <text class="field-label">密码</text>
        <wd-input v-model="password" type="password" placeholder="请输入密码" show-password custom-class="form-input" />
      </view>

      <button class="login-btn" :loading="loading" @tap="handleLogin">登 录</button>
    </view>
  </view>
</template>

<script setup>
import { ref } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import { staffLogin } from '@/api/staff'
import { useStaffStore } from '@/store/staff'
import config from '@/utils/config'
import { siteName, loadSiteName } from '@/utils/site'

const username = ref('')
const password = ref('')
const loading = ref(false)

const store = useStaffStore()

onLoad(() => {
  loadSiteName()
  // 已登录则直接进入首页（不允许游客登录）
  if (store.isLoggedIn.value) {
    uni.reLaunch({ url: '/pages/order/list/index' })
  }
})

async function handleLogin() {
  if (loading.value) return
  if (!username.value.trim() || !password.value) {
    uni.showToast({ title: '请输入账号和密码', icon: 'none' })
    return
  }
  loading.value = true
  try {
    const data = await staffLogin(username.value.trim(), password.value)
    store.setToken(data.accessToken)
    store.setRefreshToken(data.refreshToken)
    await store.fetchProfile()
    uni.showToast({ title: '登录成功', icon: 'success' })
    setTimeout(() => uni.reLaunch({ url: '/pages/order/list/index' }), 400)
  } catch (e) {
    // mock 预览：后端不可达时直接以 mock 身份进入，便于浏览器预览
    if (config.MOCK_PREVIEW) {
      store.setToken('mock-token')
      await store.fetchProfile()
      uni.showToast({ title: '已进入预览模式', icon: 'success' })
      setTimeout(() => uni.reLaunch({ url: '/pages/order/list/index' }), 400)
    }
    // 否则错误已由 request 提示
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.login-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(180deg, #e7eefb 0%, #f4f6fa 60%);
  padding-top: 20px;
  box-sizing: border-box;
  position: relative;
  overflow: hidden;
}
.bg-glow {
  position: absolute;
  border-radius: 50%;
  pointer-events: none;
}
.bg-glow.g1 {
  width: 480rpx;
  height: 480rpx;
  right: -160rpx;
  top: -120rpx;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.18) 0%, rgba(37, 99, 235, 0) 70%);
}
.bg-glow.g2 {
  width: 400rpx;
  height: 400rpx;
  left: -160rpx;
  bottom: 120rpx;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.1) 0%, rgba(37, 99, 235, 0) 70%);
}

.logo-area {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 120rpx;
}
.logo-box {
  width: 176rpx;
  height: 176rpx;
  border-radius: 40rpx;
  background: #ffffff;
  box-shadow: 0 16rpx 40rpx rgba(37, 99, 235, 0.18);
  display: flex;
  align-items: center;
  justify-content: center;
}
.logo {
  width: 128rpx;
  height: 128rpx;
}
.app-name {
  font-size: 44rpx;
  font-weight: 800;
  color: #111827;
  margin-top: 28rpx;
}
.app-desc {
  font-size: 24rpx;
  color: #4b5563;
  margin-top: 12rpx;
}

.form-card {
  position: relative;
  width: calc(100% - 96rpx);
  background: #ffffff;
  border-radius: 32rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 16rpx 48rpx rgba(37, 99, 235, 0.14);
  padding: 40rpx 36rpx;
  margin-top: 64rpx;
  box-sizing: border-box;
}
.field {
  margin-bottom: 32rpx;
}
.field-label {
  display: block;
  font-size: 26rpx;
  color: #4b5563;
  font-weight: 500;
  margin-bottom: 14rpx;
}
:deep(.form-input) {
  background: #f4f6fa;
  border-radius: 16rpx;
  padding: 0 24rpx;
}

.login-btn {
  width: 100%;
  height: 96rpx;
  line-height: 96rpx;
  border-radius: 24rpx;
  font-size: 32rpx;
  font-weight: 700;
  letter-spacing: 2rpx;
  border: none;
  background: #2563eb;
  color: #ffffff;
  margin-top: 40rpx;
  box-shadow: 0 12rpx 28rpx rgba(37, 99, 235, 0.3);
}
.login-btn::after {
  border: none;
}
.login-btn:active {
  background: #1d4ed8;
}

</style>
