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
      <wd-datetime-picker
        v-model="dateValue"
        type="date"
        :min-date="MIN_DATE"
        :default-value="todayTs"
        @confirm="onDateConfirm"
      >
        <view class="date-cell">
          <text :class="visitDate ? 'date-text' : 'date-placeholder'">{{ visitDate || '请选择上门日期' }}</text>
        </view>
      </wd-datetime-picker>

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
import { ref } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import { getCurrentOrder } from '@/utils/orderBus'
import { getOrderView, orderAppoint } from '@/api/staff'

const order = ref({})
const appointmentDesc = ref('')
const visitDate = ref('')
const dateValue = ref('')

// 上门日期最小年份 2026，选择器打开默认今天
const MIN_DATE = new Date(2026, 0, 1).getTime()
const todayTs = Date.now()

function todayStr() {
  const d = new Date()
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}

onLoad(async (query) => {
  const id = query.id
  order.value = getCurrentOrder() || {}
  if (id) {
    try {
      const detail = await getOrderView(id)
      if (detail) order.value = detail
    } catch (e) {
      // request 已提示
    }
  }
  // 暂存后再次进入：回填上门日期（空则默认今天）与预约情况
  visitDate.value = order.value.visitDate ? String(order.value.visitDate).slice(0, 10) : todayStr()
  appointmentDesc.value = order.value.appointmentDesc || ''
})

function onDateConfirm({ value }) {
  const d = new Date(value)
  visitDate.value = `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
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
  try {
    await orderAppoint({
      id: order.value.id,
      appointmentDesc: appointmentDesc.value.trim(),
      visitDate: visitDate.value
    })
    uni.showToast({ title: '提交成功', icon: 'success' })
    setTimeout(() => uni.navigateBack(), 500)
  } catch (e) {
    // 错误已由 request 提示
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
.date-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 24rpx;
  background: #f8f9fa;
  border-radius: 16rpx;
  margin-bottom: 24rpx;
}
.date-text { font-size: 28rpx; color: #333333; }
.date-placeholder { font-size: 28rpx; color: #c0c4cc; }
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
