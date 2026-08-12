<template>
  <view class="order-card" @tap="$emit('tap', order)">
    <view class="card-head">
      <order-tag :status="order.orderStatus" />
      <text class="business">{{ order.businessType || '-' }}</text>
    </view>

    <view class="fields">
      <view class="field"><text class="label">客户姓名：</text><text class="val">{{ order.customerName || '-' }}</text></view>
      <view class="field">
        <text class="label">客户电话：</text>
        <text class="val link" @tap.stop="callPhone(order.contactPhone)">{{ order.contactPhone || '-' }}</text>
      </view>
      <view class="field"><text class="label">预约时间：</text><text class="val">{{ order.visitDate || order.scheduleDate || '-' }}</text></view>
      <view class="field">
        <text class="label">办理地址：</text>
        <text class="val link" @tap.stop="copyAddress(order.installAddress)">{{ order.installAddress || '-' }}</text>
      </view>
    </view>
  </view>
</template>

<script setup>
const props = defineProps({
  order: { type: Object, default: () => ({}) }
})

const emit = defineEmits(['tap'])

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
  background: #ffffff;
  border-radius: 20rpx;
  padding: 20rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
}
.card-head {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 24rpx;
}
.business {
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
}
.fields {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}
.field {
  display: flex;
  font-size: 28rpx;
  line-height: 1.4;
  align-items: baseline;
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
