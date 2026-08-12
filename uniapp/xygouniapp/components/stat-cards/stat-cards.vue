<template>
  <view class="stat-cards">
    <view
      v-for="(item, idx) in items"
      :key="idx"
      class="stat-card"
      :class="{ active: idx === activeIndex }"
      @tap="$emit('change', idx)"
    >
      <text class="num" :style="{ color: item.color }">{{ item.value }}</text>
      <text class="label">{{ item.label }}</text>
      <view v-if="idx === activeIndex" class="bar" :style="{ background: item.color }"></view>
    </view>
  </view>
</template>

<script setup>
defineProps({
  items: { type: Array, default: () => [] },
  activeIndex: { type: Number, default: 0 }
})
defineEmits(['change'])
</script>

<style scoped lang="scss">
.stat-cards {
  display: flex;
  gap: 48rpx;
}
.stat-card {
  flex: 1;
  position: relative;
  background: #ffffff;
  border-radius: 20rpx;
  padding: 20rpx;
  min-height: 184rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
  overflow: hidden;
}
.num {
  display: block;
  font-size: 48rpx;
  font-weight: 700;
  line-height: 1.2;
}
.label {
  display: block;
  margin-top: 24rpx;
  font-size: 28rpx;
  color: #333333;
}
.bar {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 20rpx;
  border-radius: 20rpx 20rpx 0 0;
}
</style>
