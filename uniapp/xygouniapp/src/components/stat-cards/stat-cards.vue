<template>
  <view class="stat-cards">
    <view
      v-for="(item, idx) in items"
      :key="idx"
      class="stat-card"
      :class="{ active: idx === activeIndex }"
      :style="idx === activeIndex ? `border-bottom-color: ${item.color}` : ''"
      @tap="$emit('change', idx)"
    >
      <text class="num" :style="{ color: item.color }">{{ item.value }}</text>
      <text class="label">{{ item.label }}</text>
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
  height: 176rpx; /* 高度约 88px（比正方形矮 10px） */
  display: flex;
  flex-direction: column;
  align-items: center; /* 内容垂直居中 */
  justify-content: center;
  background: #ffffff;
  border-radius: 20rpx;
  border-bottom: 20rpx solid transparent; /* 选中时用 10px 底部边框 */
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
  box-sizing: border-box;
}
.num {
  display: block;
  font-size: 48rpx;
  font-weight: 700;
  line-height: 1.2;
  text-align: center; /* 数字水平居中 */
}
.label {
  display: block;
  margin-top: 24rpx;
  font-size: 28rpx;
  color: #333333;
  text-align: center;
}
</style>
