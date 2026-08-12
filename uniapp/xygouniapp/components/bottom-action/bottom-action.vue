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
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
  font-weight: 700;
}
.primary {
  background: #336fff;
  color: #ffffff;
}
.ghost {
  background: #f5f7fa;
  color: #333333;
}
</style>
