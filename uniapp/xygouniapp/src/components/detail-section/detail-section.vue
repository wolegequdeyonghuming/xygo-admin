<template>
  <view class="detail-section">
    <text class="title">{{ title }}</text>
    <view class="rows">
      <view v-for="(row, idx) in rows" :key="idx" class="row">
        <text class="label">{{ row.label }}：</text>
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
  border-radius: 24rpx;
  padding: 32rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
}
.title {
  display: block;
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
  margin-bottom: 24rpx;
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
  color: #8b8c8f;
  flex-shrink: 0;
}
.val {
  color: #8b8c8f;
  word-break: break-all;
}
.val.link {
  color: #336fff;
}
</style>
