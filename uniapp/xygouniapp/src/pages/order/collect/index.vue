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
      <wd-datetime-picker
        v-model="dateValue"
        type="date"
        :min-date="MIN_DATE"
        :default-value="todayTs"
        @confirm="onDateConfirm"
      >
        <view class="date-cell">
          <text :class="form.visitDate ? 'date-text' : 'date-placeholder'">{{ form.visitDate || '请选择上门日期' }}</text>
        </view>
      </wd-datetime-picker>

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
      <wd-input v-model="form.isNew" placeholder="请输入是否纯新增" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

      <!-- 附件 -->
      <text class="f-label">附件</text>
      <view class="attachment-grid">
        <template v-for="(item, i) in attachments" :key="i">
          <view v-if="isImage(item)" class="att-item">
            <image :src="item.url" class="att-img" mode="aspectFill" @tap="onTapAtt(i)" />
            <view class="att-del" @tap="removeAtt(i)">×</view>
          </view>
          <view v-else class="att-item file-item" @tap="onTapAtt(i)">
            <text class="file-ext">{{ fileExt(item.url).toUpperCase() || '文件' }}</text>
            <text class="file-name">{{ item.name || fileName(item.url) }}</text>
            <view class="att-del" @tap.stop="removeAtt(i)">×</view>
          </view>
        </template>
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
import { ref, reactive } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import { getCurrentOrder } from '@/utils/orderBus'
import config from '@/utils/config'
import { getOrderView, orderCollect, uploadFile } from '@/api/staff'
import { isImage, fileExt, fileName, openAttachment } from '@/utils/attachment'

const order = ref({})
const dateValue = ref('')
const attachments = ref([]) // [{ url, attachmentId }]

// 上门日期最小年份 2026，选择器打开默认今天
const MIN_DATE = new Date(2026, 0, 1).getTime()
const todayTs = Date.now()

function todayStr() {
  const d = new Date()
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}

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

function fmtDate(v) {
  if (!v) return ''
  return String(v).slice(0, 10)
}

// 相对 URL 补全为绝对地址（本地存储返回 /attachment/...，用资产源拼成完整后端地址）
function resolveUrl(u) {
  if (!u) return u
  if (/^https?:\/\//.test(u)) return u
  if (!u.startsWith('/')) u = '/' + u
  return config.ASSET_URL + u
}

function fillFormFromOrder() {
  const o = order.value || {}
  form.visitDate = fmtDate(o.visitDate) || todayStr()
  form.dealtBusinessType = o.dealtBusinessType || ''
  form.portingStatus = o.portingStatus || ''
  form.customerRealName = o.customerRealName || ''
  form.customerIdNumber = o.customerIdNumber || ''
  form.paidAmount = o.paidAmount != null ? String(o.paidAmount) : ''
  form.isRuralOrder = o.isRuralOrder === 1 ? '是' : o.isRuralOrder === 0 ? '否' : ''
  form.newPhoneNo = o.newPhoneNo || ''
  form.deviceSerial = o.deviceSerial || ''
  form.subsidyAmount = o.subsidyAmount != null ? String(o.subsidyAmount) : ''
  form.isNew = o.isNew || ''
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
  // 暂存后再次进入：回填已保存字段
  fillFormFromOrder()
  // 回填已保存附件（详情接口联查返回的 attachments 对象，attachmentId 为数字 ID）
  attachments.value = (order.value.attachments || []).map((a) => ({
    url: resolveUrl(a.url),
    attachmentId: a.id,
    mimetype: a.mimetype,
    name: a.name || fileName(a.url)
  }))
})

function onDateConfirm({ value }) {
  const d = new Date(value)
  form.visitDate = `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}
function callPhone() {
  if (!order.value.contactPhone) return
  uni.makePhoneCall({ phoneNumber: String(order.value.contactPhone), fail: () => {} })
}

function chooseImage() {
  const remain = 5 - attachments.value.length
  if (remain <= 0) return
  uni.chooseImage({
    count: remain,
    success: async (res) => {
      uni.showLoading({ title: '上传中...' })
      try {
        for (const path of res.tempFilePaths) {
          const data = await uploadFile(path)
          attachments.value.push({
            url: resolveUrl(data.url) || path,
            attachmentId: data.attachmentId,
            mimetype: data.mime,
            ext: data.ext,
            name: data.name || fileName(data.url)
          })
        }
      } catch (e) {
        uni.showToast({ title: '上传失败', icon: 'none' })
      } finally {
        uni.hideLoading()
      }
    }
  })
}
function removeAtt(i) {
  attachments.value.splice(i, 1)
}
function onTapAtt(i) {
  const item = attachments.value[i]
  if (!item) return
  if (isImage(item)) {
    const imageUrls = attachments.value.filter((a) => isImage(a)).map((a) => a.url)
    uni.previewImage({ urls: imageUrls, current: item.url })
    return
  }
  openAttachment(item)
}

async function submit(finish) {
  try {
    await orderCollect({
      id: order.value.id,
      finish,
      visitDate: form.visitDate,
      dealtBusinessType: form.dealtBusinessType.trim(),
      portingStatus: form.portingStatus,
      customerRealName: form.customerRealName,
      customerIdNumber: form.customerIdNumber,
      paidAmount: form.paidAmount,
      isRuralOrder: form.isRuralOrder === '是' ? 1 : 0,
      newPhoneNo: form.newPhoneNo,
      deviceSerial: form.deviceSerial,
      subsidyAmount: form.subsidyAmount,
      isNew: form.isNew,
      attachmentId: attachments.value
        .map((a) => a.attachmentId)
        .filter(Boolean)
        .join(',')
    })
    uni.showToast({ title: finish ? '已完成' : '已暂存', icon: 'success' })
    setTimeout(() => uni.navigateBack(), 500)
  } catch (e) {
    // 错误已由 request 提示
  }
}
</script>

<style scoped lang="scss">
.form-page { min-height: 100vh; background: #f4f6fa; padding: 24rpx 32rpx; box-sizing: border-box; }
.info-card, .section-card, .edit-card {
  background: #ffffff; border-radius: 28rpx; padding: 32rpx; margin-bottom: 32rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06); box-shadow: 0 6rpx 24rpx rgba(17, 24, 39, 0.06);
}
.card-head { display: flex; align-items: center; gap: 16rpx; margin-bottom: 32rpx; }
.business { font-size: 32rpx; font-weight: 600; color: #111827; }
.field { display: flex; align-items: baseline; font-size: 28rpx; line-height: 1.4; margin-bottom: 24rpx; }
.label { color: #9ca3af; flex-shrink: 0; }
.val { color: #4b5563; word-break: break-all; }
.val.link { color: #2563eb; }
.sec-title { display: flex; align-items: center; gap: 14rpx; font-size: 30rpx; font-weight: 700; color: #111827; margin-bottom: 24rpx; }
.sec-title::before { content: ''; width: 8rpx; height: 28rpx; border-radius: 4rpx; background: #2563eb; }
.sec-content { font-size: 28rpx; color: #4b5563; line-height: 1.5; }

.edit-title { display: flex; align-items: center; gap: 14rpx; font-size: 30rpx; font-weight: 700; color: #111827; margin-bottom: 24rpx; }
.edit-title::before { content: ''; width: 8rpx; height: 28rpx; border-radius: 4rpx; background: #2563eb; }
.f-label { display: block; font-size: 28rpx; color: #111827; font-weight: 500; margin: 32rpx 0 16rpx; }
:deep(.edit-input) { background: #f4f6fa; border-radius: 16rpx; padding: 0 24rpx; }

.date-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 24rpx;
  background: #f4f6fa;
  border-radius: 16rpx;
}
.date-text { font-size: 28rpx; color: #111827; }
.date-placeholder { font-size: 28rpx; color: #b9bec6; }

.segment { display: flex; gap: 24rpx; }
.seg-item {
  width: 120rpx; height: 72rpx; border-radius: 16rpx; background: #f4f6fa;
  display: flex; align-items: center; justify-content: center; font-size: 28rpx; color: #111827;
}
.seg-item.on { background: #2563eb; color: #ffffff; }

.attachment-grid { display: flex; flex-wrap: wrap; gap: 16rpx; }
.att-item {
  width: 176rpx; height: 176rpx; border-radius: 16rpx; border: 2rpx dashed #d0d5dd;
  position: relative; overflow: hidden; background: #ffffff;
}
.att-img { width: 100%; height: 100%; }
.att-del {
  position: absolute; top: 0; right: 0; width: 40rpx; height: 40rpx;
  background: rgba(17, 24, 39, 0.55); color: #fff; text-align: center; line-height: 40rpx; border-radius: 0 0 0 16rpx;
}
.att-item.add { display: flex; flex-direction: column; align-items: center; justify-content: center; }
.add-plus { font-size: 48rpx; color: #9ca3af; line-height: 1; }
.add-text { font-size: 24rpx; color: #9ca3af; margin-top: 8rpx; }

.att-item.file-item {
  width: 100%;
  height: auto;
  min-height: 96rpx;
  border: none;
  background: #f4f6fa;
  display: flex;
  align-items: center;
  gap: 16rpx;
  padding: 20rpx 24rpx;
  box-sizing: border-box;
  overflow: visible;
}
.file-ext {
  flex-shrink: 0;
  font-size: 24rpx;
  font-weight: 700;
  color: #2563eb;
  background: #e8effd;
  border-radius: 8rpx;
  padding: 6rpx 12rpx;
}
.file-name {
  flex: 1;
  font-size: 28rpx;
  color: #111827;
  word-break: break-all;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.att-item.file-item .att-del {
  position: relative;
  top: auto;
  right: auto;
  border-radius: 50%;
  width: 44rpx;
  height: 44rpx;
  line-height: 44rpx;
  flex-shrink: 0;
}

.bottom { padding-top: 16rpx; }
</style>
