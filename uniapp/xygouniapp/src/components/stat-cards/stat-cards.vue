<template>
  <view class="stat-deck">
    <view class="deck-glow"></view>
    <view
      v-for="(item, idx) in items"
      :key="idx"
      class="stat"
      :class="{ active: idx === activeIndex }"
      @tap="$emit('change', idx)"
    >
      <view class="stat-top">
        <text class="stat-num" :style="{ color: item.color }">{{ nums[idx] ?? 0 }}</text>
      </view>
      <text class="stat-label">{{ item.label }}</text>
      <view
        class="stat-rail"
        :class="{ on: idx === activeIndex }"
        :style="{
          background: item.color,
          boxShadow: idx === activeIndex ? `0 0 16rpx ${item.color}` : 'none'
        }"
      ></view>
    </view>
  </view>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  items: { type: Array, default: () => [] },
  activeIndex: { type: Number, default: 0 }
})
defineEmits(['change'])

const nums = ref([])
let prev = []
let timer = null

function raf(cb) {
  if (typeof requestAnimationFrame === 'function') {
    requestAnimationFrame(cb)
  } else {
    setTimeout(() => cb(Date.now()), 16)
  }
}

function animateTo(next) {
  if (timer) {
    clearTimeout(timer)
    timer = null
  }
  const from = prev.length === next.length ? prev : next.map(() => 0)
  const start = Date.now()
  const dur = 700
  const tick = () => {
    const t = Math.min(1, (Date.now() - start) / dur)
    const e = 1 - Math.pow(1 - t, 3)
    nums.value = next.map((v, i) => Math.round(from[i] + (v - from[i]) * e))
    if (t < 1) {
      raf(tick)
    } else {
      prev = next.slice()
    }
  }
  raf(tick)
}

watch(
  () => props.items,
  (val) => {
    const next = val.map((i) => Number(i.value) || 0)
    animateTo(next)
  },
  { immediate: true }
)
</script>

<style scoped lang="scss">
.stat-deck {
  position: relative;
  overflow: hidden;
  display: flex;
  gap: 16rpx;
  padding: 36rpx 28rpx;
  background: linear-gradient(135deg, #ffffff 0%, #f2f6fe 100%);
  border-radius: 32rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 16rpx 40rpx rgba(37, 99, 235, 0.1);
}
.deck-glow {
  position: absolute;
  right: -80rpx;
  top: -120rpx;
  width: 280rpx;
  height: 280rpx;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.12) 0%, rgba(37, 99, 235, 0) 70%);
  pointer-events: none;
}
.stat {
  flex: 1;
  min-width: 0;
  padding: 12rpx 10rpx;
  border-radius: 24rpx;
  transition: opacity 0.2s, transform 0.2s;
}
.stat:not(.active) {
  opacity: 0.5;
}
.stat:active {
  transform: scale(0.96);
}
.stat-top {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8rpx;
}
.stat-num {
  font-size: 56rpx;
  font-weight: 800;
  line-height: 1;
  letter-spacing: -1rpx;
  font-feature-settings: 'tnum';
}
.stat-label {
  display: block;
  margin-top: 14rpx;
  font-size: 24rpx;
  color: #4b5563;
  text-align: center;
}
.stat-rail {
  height: 6rpx;
  border-radius: 3rpx;
  margin-top: 16rpx;
  opacity: 0.25;
  transition: opacity 0.2s;
}
.stat-rail.on {
  opacity: 1;
}
</style>
