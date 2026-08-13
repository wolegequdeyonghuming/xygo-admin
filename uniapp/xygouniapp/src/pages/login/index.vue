<template>
  <view class="login-page">
    <view class="logo-area">
      <image class="logo" src="/static/logo.png" mode="aspectFit"></image>
      <text class="app-name">{{ siteName || '移动业务平台' }}</text>
      <text class="app-desc">账号密码登录</text>
    </view>

    <view class="form-area">
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

    <view class="agreement">
      <text class="agree-text">登录即表示同意</text>
      <text class="agree-link">《用户协议》</text>
      <text class="agree-text">和</text>
      <text class="agree-link">《隐私政策》</text>
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
  background: linear-gradient(180deg, #e8f0fe 0%, #f5f6f8 55%);
  padding-top: 20px;
  box-sizing: border-box;
}

.logo-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 120rpx;
}
.logo {
  width: 160rpx;
  height: 160rpx;
  border-radius: 32rpx;
}
.app-name {
  font-size: 44rpx;
  font-weight: 700;
  color: #1a1a1a;
  margin-top: 32rpx;
}
.app-desc {
  font-size: 26rpx;
  color: #8b8c8f;
  margin-top: 12rpx;
}

.form-area {
  width: 100%;
  padding: 0 80rpx;
  margin-top: 100rpx;
  box-sizing: border-box;
}
.field {
  margin-bottom: 40rpx;
}
.field-label {
  display: block;
  font-size: 28rpx;
  color: #333333;
  font-weight: 600;
  margin-bottom: 16rpx;
}
:deep(.form-input) {
  background: #ffffff;
  border-radius: 16rpx;
  padding: 0 24rpx;
}

.login-btn {
  width: 100%;
  height: 96rpx;
  line-height: 96rpx;
  border-radius: 48rpx;
  font-size: 32rpx;
  font-weight: 700;
  letter-spacing: 2rpx;
  border: none;
  background: #336fff;
  color: #ffffff;
  margin-top: 60rpx;
}
.login-btn::after {
  border: none;
}
.login-btn:active {
  background: #2a5bd6;
}

.agreement {
  position: fixed;
  bottom: 60rpx;
  display: flex;
  align-items: center;
}
.agree-text {
  font-size: 22rpx;
  color: #999;
}
.agree-link {
  font-size: 22rpx;
  color: #336fff;
}
</style>
