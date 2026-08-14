<template>
  <view
    class="order-card"
    :style="{ animationDelay: `${Math.min(index || 0, 6) * 90}ms` }"
    @tap="$emit('select', order)"
  >
    <view class="rail" :style="{ background: cfg.color }"></view>

    <view class="card-head">
      <order-tag :status="order.orderStatus" />
      <text class="business">{{ order.businessType || '-' }}</text>
      <text v-if="order.orderNo" class="order-no">#{{ order.orderNo }}</text>
    </view>

    <text class="customer">{{ order.customerName || '-' }}</text>

    <view class="fields">
      <view class="f-cell">
        <text class="f-label">客户电话</text>
        <text class="f-val link" @tap.stop="callPhone(order.contactPhone)">{{ order.contactPhone || '-' }}</text>
      </view>
      <view class="f-cell">
        <text class="f-label">预约时间</text>
        <text class="f-val">{{ order.visitDate || order.scheduleDate || '-' }}</text>
      </view>
      <view class="f-cell full">
        <text class="f-label">办理地址</text>
        <text class="f-val link" @tap.stop="copyAddress(order.installAddress)">{{ order.installAddress || '-' }}</text>
      </view>
      <view class="f-cell">
        <text class="f-label">收单员</text>
        <text class="f-val">{{ order.agent_real_name || '-' }}</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { computed } from 'vue'
import { orderStatus } from '@/utils/status'

const props = defineProps({
  order: { type: Object, default: () => ({}) },
  index: { type: Number, default: 0 }
})

const emit = defineEmits(['select'])

const cfg = computed(() => orderStatus(props.order.orderStatus))

function callPhone(phone) {
  if (!phone) return
  uni.makePhoneCall({ phoneNumber: String(phone), fail: () => {} })
}

function copyAddress(address) {
  if (!address) return
  uni.setClipboardData({
    data: address,
    success: () => uni.showToast({ title: '地址已复制', icon: 'none' })
  })
}
</script>

<style scoped lang="scss">
.order-card {
  position: relative;
  padding: 24rpx 24rpx 18rpx 36rpx;
  margin-bottom: 28rpx;
  background: #ffffff;
  border-radius: 24rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 6rpx 24rpx rgba(17, 24, 39, 0.06);
  overflow: hidden;
  animation: cardUp 0.45s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.15s ease;
}
.order-card:active {
  transform: scale(0.985);
}
.rail {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 8rpx;
}
.card-head {
  display: flex;
  align-items: center;
  gap: 12rpx;
}
.business {
  flex: 1;
  min-width: 0;
  font-size: 30rpx;
  font-weight: 600;
  color: #1a2332;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.order-no {
  flex-shrink: 0;
  font-size: 22rpx;
  color: #8a8f98;
}
.customer {
  display: block;
  margin: 6rpx 0 16rpx;
  font-size: 36rpx;
  font-weight: 700;
  letter-spacing: 1rpx;
  color: #1a2332;
}
.fields {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16rpx 20rpx;
}
.f-cell {
  min-width: 0;
}
.f-cell.full {
  grid-column: 1 / -1;
}
.f-label {
  display: block;
  font-size: 22rpx;
  color: #8a8f98;
  margin-bottom: 4rpx;
}
.f-val {
  font-size: 28rpx;
  font-weight: 400;
  color: #1a2332;
  word-break: break-all;
  line-height: 1.4;
}
.f-val.link {
  color: #2563eb;
}
@keyframes cardUp {
  from {
    opacity: 0;
    transform: translateY(24rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
