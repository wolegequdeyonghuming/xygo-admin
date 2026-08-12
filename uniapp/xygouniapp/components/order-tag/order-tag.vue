<template>
  <view class="order-tag" :style="{ color: cfg.color, background: cfg.bg }">{{ cfg.text }}</view>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  status: { type: [Number, String], default: 0 }
})

// 订单状态 → 标签（收单员视角）
// 待预约=红(#D92400/#FFDED6) 待收单=黄(#FE8B00/#FDE9B1) 已处理=蓝(#1F61FF/#D0E3FF)
const MAP = {
  0: { text: '未成交', color: '#8B8C8F', bg: '#F0F0F0' },
  1: { text: '已录单', color: '#8B8C8F', bg: '#F0F0F0' },
  2: { text: '待预约', color: '#D92400', bg: '#FFDED6' },
  3: { text: '待收单', color: '#FE8B00', bg: '#FDE9B1' },
  4: { text: '已上门', color: '#1F61FF', bg: '#D0E3FF' },
  5: { text: '已回访', color: '#1F61FF', bg: '#D0E3FF' },
  6: { text: '已完工', color: '#1F61FF', bg: '#D0E3FF' }
}

const cfg = computed(() => MAP[String(props.status)] || { text: '-', color: '#8B8C8F', bg: '#F0F0F0' })
</script>

<style scoped lang="scss">
.order-tag {
  display: inline-flex;
  align-items: center;
  height: 40rpx;
  padding: 0 24rpx;
  border-radius: 16rpx;
  font-size: 28rpx;
  font-weight: 700;
  line-height: 1;
}
</style>
