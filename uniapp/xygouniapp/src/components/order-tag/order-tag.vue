<template>
  <view class="order-tag" :style="{ color: cfg.color, background: cfg.bg }">{{ cfg.text }}</view>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  status: { type: [Number, String], default: 0 }
})

// 订单状态 → 标签（收单员视角）
// 未收单（已录单/待预约）=橙(#FE8B00/#FDE9B1) 待收单=红(#D92400/#FFDED6) 已收单/已回访/已完工=蓝(#1F61FF/#D0E3FF)
const MAP = {
  0: { text: '未成交', color: '#8B8C8F', bg: '#F0F0F0' },
  1: { text: '已录单', color: '#FE8B00', bg: '#FDE9B1' },
  2: { text: '待预约', color: '#FE8B00', bg: '#FDE9B1' },
  3: { text: '待收单', color: '#D92400', bg: '#FFDED6' },
  4: { text: '已收单', color: '#1F61FF', bg: '#D0E3FF' },
  5: { text: '已回访', color: '#1F61FF', bg: '#D0E3FF' },
  6: { text: '已完工', color: '#1F61FF', bg: '#D0E3FF' }
}

const cfg = computed(() => MAP[String(props.status)] || { text: '-', color: '#8B8C8F', bg: '#F0F0F0' })
</script>

<style scoped lang="scss">
.order-tag {
  display: inline-flex;
  align-items: center;
  height: 60rpx;
  padding: 0 24rpx;
  border-radius: 16rpx;
  font-size: 28rpx;
  font-weight: 700;
  line-height: 1;
}
</style>
