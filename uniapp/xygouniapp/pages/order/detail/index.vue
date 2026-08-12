<template>
  <view class="detail-page">
    <!-- 主信息卡 -->
    <view class="main-card">
      <view class="card-head">
        <order-tag :status="order.orderStatus" />
        <text class="business">{{ order.businessType || '-' }}</text>
      </view>
      <view class="fields">
        <view class="field"><text class="label">客户姓名：</text><text class="val">{{ order.customerName || '-' }}</text></view>
        <view class="field">
          <text class="label">客户电话：</text>
          <text class="val link" @tap="callPhone(order.contactPhone)">{{ order.contactPhone || '-' }}</text>
        </view>
        <view class="field"><text class="label">预约时间：</text><text class="val">{{ order.visitDate || order.scheduleDate || '-' }}</text></view>
        <view class="field">
          <text class="label">办理地址：</text>
          <text class="val link" @tap="copyAddress(order.installAddress)">{{ order.installAddress || '-' }}</text>
        </view>
        <view class="field"><text class="label">订单编号：</text><text class="val">{{ order.orderNo || '-' }}</text></view>
        <view class="field"><text class="label">可联系时间：</text><text class="val">{{ order.availableTimeDesc || '-' }}</text></view>
      </view>
    </view>

    <!-- 分阶段只读区 -->
    <detail-section v-if="order.appointmentDesc" title="预约情况" :rows="[{ label: '预约情况', value: order.appointmentDesc }]" />
    <detail-section
      v-if="Number(order.orderStatus) >= 4"
      title="收单情况"
      :rows="[
        { label: '成交业务', value: order.dealtBusinessType },
        { label: '上门日期', value: order.visitDate },
        { label: '实缴额度', value: order.paidAmount },
        { label: '新开号码', value: order.newPhoneNo }
      ]"
    />
    <detail-section
      v-if="Number(order.orderStatus) >= 5"
      title="回访情况"
      :rows="[{ label: '话务二次回访', value: order.followUpDesc }]"
    />
    <detail-section
      v-if="Number(order.orderStatus) >= 6"
      title="完工情况"
      :rows="[{ label: '是否完工', value: '已完工' }]"
    />

    <!-- 详细情况 -->
    <order-comment v-if="order.id" :order-id="order.id" />

    <!-- 底部操作 -->
    <view class="bottom">
      <bottom-action
        v-if="actionText"
        :primary="actionText"
        @primary-tap="goAction"
      />
    </view>
  </view>
</template>

<script setup>
import { ref, computed, onLoad } from '@dcloudio/uni-app'
import { mockOrders } from '@/utils/mock'

const order = ref({})

const actionText = computed(() => {
  const s = Number(order.value.orderStatus)
  if (s === 2) return '填写预约情况'
  if (s === 3) return '填写收单情况'
  if (s === 4) return '编辑收单'
  return ''
})

onLoad((query) => {
  const id = Number(query.id)
  order.value = mockOrders.find((o) => o.id === id) || {}
})

function callPhone(phone) {
  if (!phone) return
  uni.makePhoneCall({ phoneNumber: String(phone), fail: () => {} })
}
function copyAddress(address) {
  if (!address) return
  uni.setClipboardData({ data: address, success: () => uni.showToast({ title: '地址已复制', icon: 'none' }) })
}
function goAction() {
  const s = Number(order.value.orderStatus)
  if (s === 2) {
    uni.navigateTo({ url: `/pages/order/appoint/index?id=${order.value.id}` })
  } else {
    uni.navigateTo({ url: `/pages/order/collect/index?id=${order.value.id}` })
  }
}
</script>

<style scoped lang="scss">
.detail-page {
  min-height: 100vh;
  background: #f5f6f8;
  padding: 24rpx 32rpx;
  box-sizing: border-box;
}
.main-card {
  background: #ffffff;
  border-radius: 24rpx;
  padding: 32rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
}
.card-head {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 32rpx;
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
.bottom {
  padding-top: 24rpx;
}
</style>
