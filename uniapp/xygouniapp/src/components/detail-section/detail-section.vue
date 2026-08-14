<template>
  <view class="detail-section">
    <text class="title">{{ title }}</text>
    <view class="rows">
      <view v-for="(row, idx) in rows" :key="idx" class="row">
        <text class="label">{{ row.label }}</text>
        <text class="val" :class="{ link: row.link }" @tap="row.link ? $emit('link', row) : ''">{{ row.value || '-' }}</text>
      </view>
    </view>
  </view>
</template>

<script setup>
defineProps({
  title: { type: String, default: '' },
  rows: { type: Array, default: () => [] }
})
defineEmits(['link'])
</script>

<style scoped lang="scss">
.detail-section {
  background: #ffffff;
  border-radius: 28rpx;
  padding: 32rpx;
  margin-bottom: 32rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 6rpx 24rpx rgba(17, 24, 39, 0.06);
}
.title {
  display: flex;
  align-items: center;
  gap: 14rpx;
  font-size: 30rpx;
  font-weight: 700;
  color: #111827;
  margin-bottom: 28rpx;
}
.title::before {
  content: '';
  width: 8rpx;
  height: 28rpx;
  border-radius: 4rpx;
  background: #2563eb;
}
.rows {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}
.row {
  display: flex;
  align-items: baseline;
  font-size: 28rpx;
  line-height: 1.4;
}
.label {
  color: #9ca3af;
  flex-shrink: 0;
  margin-right: 16rpx;
}
.val {
  color: #4b5563;
  word-break: break-all;
}
.val.link {
  color: #2563eb;
}
</style>
