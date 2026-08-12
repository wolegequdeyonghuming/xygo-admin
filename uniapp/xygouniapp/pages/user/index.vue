<template>
  <view class="profile-page">
    <!-- 用户信息卡 -->
    <view class="profile-card">
      <view class="head">
        <view class="avatar">{{ avatarChar }}</view>
        <view class="info">
          <text class="username">{{ name }}</text>
          <text class="nickname">{{ nickname }}</text>
        </view>
      </view>
      <view class="fields">
        <view class="field"><text class="label">手机号：</text><text class="val">{{ userInfo?.mobile || '-' }}</text></view>
        <view class="field"><text class="label">岗位：</text><text class="val">{{ userInfo?.postName || '-' }}</text></view>
        <view class="field"><text class="label">部门：</text><text class="val">{{ userInfo?.deptName || '-' }}</text></view>
      </view>
    </view>

    <view class="logout">
      <bottom-action :primary="'退出登录'" @primary-tap="logout" />
    </view>
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'

const store = useStaffStore()
const userInfo = ref(null)

const name = computed(() => userInfo.value?.realName || userInfo.value?.nickname || '')
const nickname = computed(() => userInfo.value?.nickname || '')
const avatarChar = computed(() => (name.value || '员').slice(0, 1))

onShow(async () => {
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
.profile-page { min-height: 100vh; background: #f5f6f8; padding: 24rpx 32rpx; box-sizing: border-box; }
.profile-card {
  background: #ffffff; border-radius: 24rpx; padding: 40rpx; box-shadow: 0 2rpx 12rpx rgba(0,0,0,0.04);
}
.head { display: flex; align-items: center; gap: 32rpx; }
.avatar {
  width: 128rpx; height: 128rpx; border-radius: 50%; background: #e0e0e0; color: #333333;
  display: flex; align-items: center; justify-content: center; font-size: 44rpx; font-weight: 700; flex-shrink: 0;
}
.info { display: flex; flex-direction: column; }
.username { font-size: 32rpx; font-weight: 700; color: #333333; }
.nickname { font-size: 28rpx; color: #8b8c8f; margin-top: 8rpx; }

.fields { margin-top: 40rpx; display: flex; flex-direction: column; gap: 32rpx; }
.field { display: flex; align-items: baseline; font-size: 28rpx; }
.label { color: #8b8c8f; flex-shrink: 0; }
.val { color: #8b8c8f; word-break: break-all; }

.logout { padding-top: 40rpx; }
</style>
