<template>
  <view class="bottom-action">
    <template v-if="ghost">
      <view class="btn ghost" @tap="$emit('ghost-tap')">{{ ghost }}</view>
      <view class="btn primary" @tap="$emit('primary-tap')">{{ primary }}</view>
    </template>
    <view v-else class="btn primary" @tap="$emit('primary-tap')">{{ primary }}</view>
  </view>
</template>

<script setup>
defineProps({
  primary: { type: String, default: '' },
  ghost: { type: String, default: '' }
})
defineEmits(['primary-tap', 'ghost-tap'])
</script>

<style scoped lang="scss">
.bottom-action {
  display: flex;
  gap: 32rpx;
  padding: 24rpx 0;
}
.btn {
  flex: 1;
  height: 96rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
  font-weight: 700;
  transition: transform 0.15s ease, opacity 0.15s ease;
}
.btn:active {
  transform: scale(0.97);
}
.primary {
  background: #2563eb;
  color: #ffffff;
  box-shadow: 0 8rpx 20rpx rgba(37, 99, 235, 0.25);
}
.primary:active {
  background: #1d4ed8;
}
.ghost {
  background: #eff2f7;
  color: #111827;
}
</style>
