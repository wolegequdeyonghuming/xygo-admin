<template>
  <view class="dock">
    <view class="pill" :style="{ transform: `translateX(${selected * 100}%)` }"></view>
    <view
      v-for="(item, i) in tabs"
      :key="i"
      class="dock-item"
      :class="{ active: selected === i }"
      @tap="onTap(i)"
    >
      <wd-icon :name="item.icon" :size="22" :color="selected === i ? '#111827' : '#9CA3AF'" />
      <text class="dock-title" :class="{ active: selected === i }">{{ item.title }}</text>
    </view>
  </view>
</template>

<script setup>
import { ref, watch } from 'vue'
import { tabActive } from '@/utils/tab'

const tabs = [
  { icon: 'home', title: '首页' },
  { icon: 'list', title: '订单' },
  { icon: 'user', title: '我的' }
]

const selected = ref(tabActive.value)

watch(tabActive, (v) => {
  selected.value = v
})

const PATHS = ['/pages/order/list/index', '/pages/order/query/index', '/pages/user/index']

function onTap(i) {
  if (selected.value === i) return
  selected.value = i
  uni.reLaunch({ url: PATHS[i] })
}
</script>

<style scoped lang="scss">
.dock {
  position: fixed;
  left: 32rpx;
  right: 32rpx;
  bottom: calc(24rpx + env(safe-area-inset-bottom));
  z-index: 1000;
  display: flex;
  padding: 16rpx 0 calc(16rpx + env(safe-area-inset-bottom));
  background: #ffffff;
  border-radius: 56rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.08);
  box-shadow: 0 12rpx 40rpx rgba(17, 24, 39, 0.12);
}
.pill {
  position: absolute;
  left: 0;
  top: 16rpx;
  z-index: 0;
  width: 33.3333%;
  height: 88rpx;
  border-radius: 44rpx;
  background: #e8effd;
  transition: transform 0.28s cubic-bezier(0.4, 0, 0.2, 1);
}
.dock-item {
  position: relative;
  z-index: 1;
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 6rpx;
  height: 88rpx;
}
.dock-title {
  font-size: 22rpx;
  color: #9ca3af;
  line-height: 1;
  transition: color 0.2s;
}
.dock-title.active {
  color: #111827;
  font-weight: 700;
}
</style>
