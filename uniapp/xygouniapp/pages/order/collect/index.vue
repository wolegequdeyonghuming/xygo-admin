<template>
  <view class="form-page">
    <!-- 只读主信息 -->
    <view class="info-card">
      <view class="card-head">
        <order-tag :status="order.orderStatus" />
        <text class="business">{{ order.businessType || '-' }}</text>
      </view>
      <view class="field"><text class="label">客户姓名：</text><text class="val">{{ order.customerName || '-' }}</text></view>
      <view class="field"><text class="label">客户电话：</text><text class="val link" @tap="callPhone">{{ order.contactPhone || '-' }}</text></view>
      <view class="field"><text class="label">办理地址：</text><text class="val">{{ order.installAddress || '-' }}</text></view>
    </view>

    <!-- 只读预约情况 -->
    <view v-if="order.appointmentDesc" class="section-card">
      <text class="sec-title">预约情况</text>
      <text class="sec-content">{{ order.appointmentDesc }}</text>
    </view>

    <!-- 收单情况编辑区（完整 12 项） -->
    <view class="edit-card">
      <text class="edit-title">收单情况</text>

      <text class="f-label">上门日期</text>
      <wd-input v-model="form.visitDate" readonly placeholder="请选择上门日期" placeholder-style="color:#C0C4CC" custom-class="edit-input" @click="dateVisible = true" />
      <wd-datetime-picker v-model="dateValue" type="date" :visible="dateVisible" @confirm="onDateConfirm" @cancel="dateVisible = false" />

      <text class="f-label">成交业务</text>
      <wd-input v-model="form.dealtBusinessType" placeholder="请输入成交业务" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">携转情况</text>
      <wd-input v-model="form.portingStatus" placeholder="请输入携转情况" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">客户实际姓名</text>
      <wd-input v-model="form.customerRealName" placeholder="请输入客户实际姓名" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">客户身份证号</text>
      <wd-input v-model="form.customerIdNumber" placeholder="请输入客户身份证号" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">实缴额度（元）</text>
      <wd-input v-model="form.paidAmount" type="number" placeholder="请输入实缴额度" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">是否乡下单</text>
      <view class="segment">
        <view class="seg-item" :class="{ on: form.isRuralOrder === '是' }" @tap="form.isRuralOrder = '是'">是</view>
        <view class="seg-item" :class="{ on: form.isRuralOrder === '否' }" @tap="form.isRuralOrder = '否'">否</view>
      </view>

      <text class="f-label">新开号码</text>
      <wd-input v-model="form.newPhoneNo" placeholder="请输入新开号码" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">终端串码</text>
      <wd-input v-model="form.deviceSerial" placeholder="请输入终端串码" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">话补（元）</text>
      <wd-input v-model="form.subsidyAmount" type="number" placeholder="请输入话补" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <text class="f-label">是否纯新增</text>
      <view class="segment">
        <view class="seg-item" :class="{ on: form.isNew === '是' }" @tap="form.isNew = '是'">是</view>
        <view class="seg-item" :class="{ on: form.isNew === '否' }" @tap="form.isNew = '否'">否</view>
      </view>

      <!-- 附件 -->
      <text class="f-label">附件</text>
      <view class="attachment-grid">
        <view v-for="(img, i) in attachments" :key="i" class="att-item">
          <image :src="img" class="att-img" mode="aspectFill" @tap="preview(i)" />
          <view class="att-del" @tap="removeAtt(i)">×</view>
        </view>
        <view v-if="attachments.length < 5" class="att-item add" @tap="chooseImage">
          <text class="add-plus">＋</text>
          <text class="add-text">添加图片</text>
        </view>
      </view>
    </view>

    <!-- 暂存 / 完成 -->
    <view class="bottom">
      <bottom-action :primary="'完成'" :ghost="'暂存'" @primary-tap="submit(true)" @ghost-tap="submit(false)" />
    </view>
  </view>
</template>

<script setup>
import { ref, reactive, onLoad } from '@dcloudio/uni-app'
import { mockOrders } from '@/utils/mock'
import { orderCollect } from '@/api/staff'

const order = ref({})
const dateValue = ref('')
const dateVisible = ref(false)
const attachments = ref([])

const form = reactive({
  visitDate: '',
  dealtBusinessType: '',
  portingStatus: '',
  customerRealName: '',
  customerIdNumber: '',
  paidAmount: '',
  isRuralOrder: '',
  newPhoneNo: '',
  deviceSerial: '',
  subsidyAmount: '',
  isNew: ''
})

onLoad((query) => {
  const id = Number(query.id)
  order.value = mockOrders.find((o) => o.id === id) || {}
})

function onDateConfirm({ value }) {
  form.visitDate = value
  dateVisible.value = false
}
function callPhone() {
  if (!order.value.contactPhone) return
  uni.makePhoneCall({ phoneNumber: String(order.value.contactPhone), fail: () => {} })
}

function chooseImage() {
  const remain = 5 - attachments.value.length
  uni.chooseImage({
    count: remain,
    success: (res) => {
      // TODO: 后端 M1 就绪后 uni.uploadFile 到 /admin/upload/file（带 Authorization），取返回 URL
      attachments.value = attachments.value.concat(res.tempFilePaths)
    }
  })
}
function removeAtt(i) {
  attachments.value.splice(i, 1)
}
function preview(i) {
  uni.previewImage({ urls: attachments.value, current: i })
}

function submit(finish) {
  if (!form.dealtBusinessType.trim()) {
    uni.showToast({ title: '请填写成交业务', icon: 'none' })
    return
  }
  // TODO: 后端 M1 就绪后走 orderCollect({ id, ...form, finish })
  uni.showToast({ title: finish ? '已完成' : '已暂存', icon: 'success' })
  setTimeout(() => uni.navigateBack(), 500)
}
</script>

<style scoped lang="scss">
.form-page { min-height: 100vh; background: #f5f6f8; padding: 24rpx 32rpx; box-sizing: border-box; }
.info-card, .section-card, .edit-card {
  background: #ffffff; border-radius: 24rpx; padding: 32rpx; margin-bottom: 40rpx; box-shadow: 0 2rpx 12rpx rgba(0,0,0,0.04);
}
.card-head { display: flex; align-items: center; gap: 16rpx; margin-bottom: 32rpx; }
.business { font-size: 32rpx; font-weight: 700; color: #333333; }
.field { display: flex; align-items: baseline; font-size: 28rpx; line-height: 1.4; margin-bottom: 24rpx; }
.label { color: #8b8c8f; flex-shrink: 0; }
.val { color: #8b8c8f; word-break: break-all; }
.val.link { color: #336fff; }
.sec-title { display: block; font-size: 32rpx; font-weight: 700; color: #333333; margin-bottom: 24rpx; }
.sec-content { font-size: 28rpx; color: #8b8c8f; line-height: 1.5; }

.edit-title { display: block; font-size: 32rpx; font-weight: 700; color: #333333; margin-bottom: 24rpx; }
.f-label { display: block; font-size: 28rpx; color: #333333; font-weight: 600; margin: 32rpx 0 16rpx; }
:deep(.edit-input) { background: #f8f9fa; border-radius: 16rpx; padding: 0 24rpx; }

.segment { display: flex; gap: 24rpx; }
.seg-item {
  width: 120rpx; height: 72rpx; border-radius: 16rpx; background: #f8f9fa;
  display: flex; align-items: center; justify-content: center; font-size: 28rpx; color: #333333;
}
.seg-item.on { background: #336fff; color: #ffffff; }

.attachment-grid { display: flex; flex-wrap: wrap; gap: 16rpx; }
.att-item {
  width: 176rpx; height: 176rpx; border-radius: 16rpx; border: 2rpx dashed #d9d9d9;
  position: relative; overflow: hidden;
}
.att-img { width: 100%; height: 100%; }
.att-del {
  position: absolute; top: 0; right: 0; width: 40rpx; height: 40rpx;
  background: rgba(0,0,0,0.5); color: #fff; text-align: center; line-height: 40rpx; border-radius: 0 0 0 16rpx;
}
.att-item.add { display: flex; flex-direction: column; align-items: center; justify-content: center; }
.add-plus { font-size: 48rpx; color: #8b8c8f; line-height: 1; }
.add-text { font-size: 24rpx; color: #8b8c8f; margin-top: 8rpx; }

.bottom { padding-top: 16rpx; }
</style>
