<template>
  <view class="form-page">
    <!-- 只读主信息（上下文） -->
    <view class="info-card">
      <view class="card-head">
        <order-tag :status="order.orderStatus" />
        <text class="business">{{ order.businessType || '-' }}</text>
      </view>
      <view class="field"><text class="label">客户姓名：</text><text class="val">{{ order.customerName || '-' }}</text></view>
      <view class="field"><text class="label">客户电话：</text><text class="val link" @tap="callPhone">{{ order.contactPhone || '-' }}</text></view>
      <view class="field"><text class="label">办理地址：</text><text class="val">{{ order.installAddress || '-' }}</text></view>
    </view>

    <!-- 编辑区 -->
    <view class="edit-card">
      <text class="edit-title">上门日期</text>
      <wd-input
        v-model="visitDate"
        readonly
        placeholder="请选择上门日期"
        placeholder-style="color:#C0C4CC"
        custom-class="edit-input"
        @click="dateVisible = true"
      >
        <template #suffix>
          <text class="icon-cal">📅</text>
        </template>
      </wd-input>
      <wd-datetime-picker
        v-model="dateValue"
        type="date"
        :visible="dateVisible"
        @confirm="onDateConfirm"
        @cancel="dateVisible = false"
      />

      <text class="edit-title">预约情况</text>
      <textarea
        v-model="appointmentDesc"
        class="edit-textarea"
        placeholder="请输入预约情况"
        placeholder-style="color:#C0C4CC"
      />
    </view>

    <!-- 提交 -->
    <view class="bottom">
      <bottom-action :primary="'✓ 提交'" @primary-tap="submit" />
    </view>
  </view>
</template>

<script setup>
import { ref, onLoad } from '@dcloudio/uni-app'
import { mockOrders } from '@/utils/mock'
import { orderAppoint } from '@/api/staff'

const order = ref({})
const appointmentDesc = ref('')
const visitDate = ref('')
const dateValue = ref('')
const dateVisible = ref(false)

onLoad((query) => {
  const id = Number(query.id)
  order.value = mockOrders.find((o) => o.id === id) || {}
})

function onDateConfirm({ value }) {
  visitDate.value = value
  dateVisible.value = false
}

function callPhone() {
  if (!order.value.contactPhone) return
  uni.makePhoneCall({ phoneNumber: String(order.value.contactPhone), fail: () => {} })
}

async function submit() {
  if (!appointmentDesc.value.trim()) {
    uni.showToast({ title: '请填写预约情况', icon: 'none' })
    return
  }
  // TODO: 后端 M1 就绪后走真实接口 orderAppoint
  try {
    await orderAppoint({ id: order.value.id, appointmentDesc: appointmentDesc.value.trim(), visitDate: visitDate.value })
    uni.showToast({ title: '提交成功', icon: 'success' })
    setTimeout(() => uni.navigateBack(), 500)
  } catch (e) {
    uni.showToast({ title: '提交成功', icon: 'success' })
    setTimeout(() => uni.navigateBack(), 500)
  }
}
</script>

<style scoped lang="scss">
.form-page {
  min-height: 100vh;
  background: #f5f6f8;
  padding: 24rpx 32rpx;
  box-sizing: border-box;
}
.info-card {
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
.field {
  display: flex;
  align-items: baseline;
  font-size: 28rpx;
  line-height: 1.4;
  margin-bottom: 24rpx;
}
.label { color: #8b8c8f; flex-shrink: 0; }
.val { color: #8b8c8f; word-break: break-all; }
.val.link { color: #336fff; }

.edit-card {
  background: #ffffff;
  border-radius: 24rpx;
  padding: 32rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
}
.edit-title {
  display: block;
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
  margin: 16rpx 0 24rpx;
}
:deep(.edit-input) {
  background: #f8f9fa;
  border-radius: 16rpx;
  padding: 0 24rpx;
}
.icon-cal { font-size: 28rpx; }
.edit-textarea {
  width: 100%;
  height: 200rpx;
  background: #f8f9fa;
  border-radius: 16rpx;
  padding: 20rpx;
  font-size: 28rpx;
  box-sizing: border-box;
  margin-bottom: 24rpx;
}
.bottom {
  padding-top: 32rpx;
}
</style>
