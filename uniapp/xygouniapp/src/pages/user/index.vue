<template>
  <view class="profile-page">
    <!-- 浅蓝舱头 -->
    <view class="deck-head" :style="{ paddingTop: headPad + 'rpx' }">
      <view class="head-glow"></view>
      <view class="profile-head">
        <view class="avatar">{{ avatarChar }}</view>
        <view class="info">
          <text class="username">{{ name }}</text>
          <text class="role">{{ postName }}</text>
        </view>
      </view>
    </view>

    <view class="content">
      <!-- 信息卡 -->
      <view class="fields-card">
        <view class="field-row">
          <text class="label">手机号</text>
          <text class="val">{{ userInfo?.mobile || '-' }}</text>
        </view>
        <view class="field-row">
          <text class="label">岗位</text>
          <text class="val">{{ userInfo?.postName || '-' }}</text>
        </view>
        <view class="field-row">
          <text class="label">部门</text>
          <text class="val">{{ userInfo?.deptName || '-' }}</text>
        </view>
      </view>

      <view class="logout">
        <bottom-action :primary="'退出登录'" @primary-tap="logout" />
      </view>
    </view>

    <tab-bar />
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import config from '@/utils/config'
import { tabActive } from '@/utils/tab'
import { getStatusBarHeightRpx } from '@/utils/system'

const store = useStaffStore()
const userInfo = ref(null)
const headPad = ref(getStatusBarHeightRpx() + 24)

const name = computed(() => userInfo.value?.realName || userInfo.value?.nickname || '')
const nickname = computed(() => userInfo.value?.nickname || '')
const postName = computed(() => userInfo.value?.postName || userInfo.value?.roleName || '')
const avatarChar = computed(() => (name.value || '员').slice(0, 1))

onShow(async () => {
  tabActive.value = 2
  if (config.MOCK_PREVIEW && !store.isLoggedIn.value) {
    store.setToken('mock-token')
  }
  if (!store.isLoggedIn.value) {
    uni.reLaunch({ url: '/pages/login/index' })
    return
  }
  if (!store.userInfo.value) {
    await store.fetchProfile()
  }
  userInfo.value = store.userInfo.value
})

function logout() {
  uni.showModal({
    title: '提示',
    content: '确定退出登录吗？',
    success: (res) => {
      if (res.confirm) store.logout()
    }
  })
}
</script>

<style scoped lang="scss">
.profile-page {
  min-height: 100vh;
  background: #f4f6fa;
}

/* ===== 浅蓝舱头 ===== */
.deck-head {
  position: relative;
  overflow: hidden;
  padding-left: 32rpx;
  padding-right: 32rpx;
  padding-bottom: 48rpx;
  background: linear-gradient(160deg, #f2f6fe 0%, #e7eefb 100%);
  border-bottom-left-radius: 40rpx;
  border-bottom-right-radius: 40rpx;
}
.head-glow {
  position: absolute;
  right: -80rpx;
  top: -100rpx;
  width: 280rpx;
  height: 280rpx;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.14) 0%, rgba(37, 99, 235, 0) 70%);
  pointer-events: none;
}
.profile-head {
  position: relative;
  display: flex;
  align-items: center;
  gap: 24rpx;
}
.avatar {
  width: 112rpx;
  height: 112rpx;
  border-radius: 50%;
  background: #2563eb;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 44rpx;
  font-weight: 800;
  flex-shrink: 0;
  box-shadow: 0 8rpx 24rpx rgba(37, 99, 235, 0.3);
}
.info {
  display: flex;
  flex-direction: column;
}
.username {
  font-size: 36rpx;
  font-weight: 700;
  color: #111827;
}
.role {
  margin-top: 6rpx;
  font-size: 24rpx;
  color: #4b5563;
}

/* ===== 内容 ===== */
.content {
  padding: 32rpx 32rpx 220rpx;
  box-sizing: border-box;
}
.fields-card {
  background: #ffffff;
  border-radius: 28rpx;
  padding: 8rpx 32rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 6rpx 24rpx rgba(17, 24, 39, 0.06);
}
.field-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 28rpx 0;
  font-size: 28rpx;
  border-bottom: 1rpx solid rgba(17, 24, 39, 0.05);
}
.field-row:last-child {
  border-bottom: none;
}
.label {
  color: #9ca3af;
  flex-shrink: 0;
}
.val {
  color: #111827;
  font-weight: 400;
  text-align: right;
  word-break: break-all;
}
.logout {
  padding-top: 32rpx;
}
</style>
