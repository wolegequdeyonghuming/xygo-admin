<template>
  <view class="tab-bar">
    <view
      v-for="(item, i) in tabs"
      :key="i"
      class="tab-item"
      :class="{ active: selected === i }"
      @tap="onTap(i)"
    >
      <wd-icon :name="item.icon" :size="24" :color="selected === i ? '#336FFF' : '#999999'" />
      <text class="tab-title" :class="{ active: selected === i }">{{ item.title }}</text>
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
.tab-bar {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1000;
  display: flex;
  height: 100rpx; /* 50px */
  padding-bottom: env(safe-area-inset-bottom);
  background: #ffffff;
  box-shadow: 0 -2rpx 12rpx rgba(0, 0, 0, 0.05);
}
.tab-item {
  flex: 1;
  display: flex;
  flex-direction: row; /* 图标与文字横向排列 */
  align-items: center;
  justify-content: center;
  gap: 12rpx;
}
.tab-title {
  font-size: 32rpx; /* 16px */
  color: #999999;
  line-height: 1;
}
.tab-title.active {
  color: #336fff;
  font-weight: 700;
}
</style>
