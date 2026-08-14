<template>
  <view class="order-tag" :style="{ color: cfg.color, background: cfg.bg }">{{ cfg.text }}</view>
</template>

<script setup>
import { computed } from 'vue'
import { orderStatus } from '@/utils/status'

const props = defineProps({
  status: { type: [Number, String], default: 0 }
})

const cfg = computed(() => orderStatus(props.status))
</script>

<style scoped lang="scss">
.order-tag {
  display: inline-flex;
  align-items: center;
  height: 52rpx;
  padding: 0 20rpx;
  border-radius: 16rpx;
  clip-path: polygon(0 0, calc(100% - 12rpx) 0, 100% 12rpx, 100% 100%, 0 100%);
  font-size: 24rpx;
  font-weight: 700;
  letter-spacing: 1rpx;
  line-height: 1;
}
.order-tag::before {
  content: '';
  width: 10rpx;
  height: 10rpx;
  border-radius: 50%;
  background: currentColor;
  margin-right: 10rpx;
}
</style>
