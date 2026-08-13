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
        <view class="field"><text class="label">话务员：</text><text class="val">{{ order.telemarketer_real_name || '-' }}</text></view>
        <view class="field"><text class="label">收单员：</text><text class="val">{{ order.agent_real_name || '-' }}</text></view>
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
      :rows="[
        { label: '是否完工', value: order.isCompleted ? '已完工' : '-' },
        { label: '宽带账号', value: order.broadbandAccount },
        { label: '工号', value: order.agencyNo }
      ]"
    />

    <!-- 附件 -->
    <view v-if="attachments.length" class="att-card">
      <text class="att-title">附件</text>
      <view class="att-grid">
        <image
          v-for="(img, i) in attachments"
          :key="i"
          :src="img"
          class="att-img"
          mode="aspectFill"
          @tap="preview(i)"
        />
      </view>
    </view>

    <!-- 详细情况 -->
    <order-comment v-if="order.id" :order-id="order.id" />

    <!-- 底部操作 -->
    <view class="bottom">
      <bottom-action
        v-if="actions.length"
        :primary="actions[0].text"
        :ghost="actions[1] ? actions[1].text : ''"
        @primary-tap="onPrimary"
        @ghost-tap="onGhost"
      />
    </view>

    <!-- 排单弹窗 -->
    <wd-popup v-model="showSchedule" position="bottom" :z-index="2000" custom-style="padding:32rpx">
      <view class="popup-panel">
        <text class="popup-title">排单</text>
        <text class="f-label">选择收单员</text>
        <wd-picker v-model="scheduleAgentId" :columns="agentOptions" use-default-slot>
          <view class="schedule-cell">
            <text :class="scheduleAgentId !== '' ? 'date-text' : 'date-placeholder'">{{ scheduleAgentLabel }}</text>
          </view>
        </wd-picker>
        <view class="filter-btns">
          <view class="btn ghost" @tap="showSchedule = false">取消</view>
          <view class="btn primary" @tap="doSchedule">确定</view>
        </view>
      </view>
    </wd-popup>

    <!-- 完工弹窗 -->
    <wd-popup v-model="showComplete" position="bottom" :z-index="2000" custom-style="padding:32rpx">
      <view class="popup-panel">
        <text class="popup-title">完工</text>
        <text class="f-label">宽带账号</text>
        <wd-input v-model="completeForm.broadbandAccount" placeholder="请输入宽带账号" placeholder-style="color:#C0C4CC" custom-class="edit-input" />
        <text class="f-label">是否完工</text>
        <wd-input v-model="completeForm.isCompleted" placeholder="0 或 1" placeholder-style="color:#C0C4CC" custom-class="edit-input" />
        <text class="f-label">工号</text>
        <wd-input v-model="completeForm.agencyNo" placeholder="请输入工号" placeholder-style="color:#C0C4CC" custom-class="edit-input" />
        <view class="filter-btns">
          <view class="btn ghost" @tap="showComplete = false">取消</view>
          <view class="btn primary" @tap="doComplete">确定</view>
        </view>
      </view>
    </wd-popup>
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'
import { onLoad, onShow } from '@dcloudio/uni-app'
import { getOrderView, getAttachmentList, getAgentList, orderSchedule, orderComplete } from '@/api/staff'
import { getCurrentOrder } from '@/utils/orderBus'
import { useStaffStore } from '@/store/staff'
import config from '@/utils/config'

const store = useStaffStore()
const order = ref({})
const attachments = ref([])
const orderId = ref('')

// 排单弹窗
const showSchedule = ref(false)
const agents = ref([])
const scheduleAgentId = ref('')

// 完工弹窗
const showComplete = ref(false)
const completeForm = ref({ broadbandAccount: '', isCompleted: '', agencyNo: '' })

const role = computed(() => store.userInfo.value?.roleKey || '')

// 预约/收单操作权限：收单员/收单员管理员/超管可操作；收单员管理员仅可操作收单员为自己的单；
// 管理员/文员不参与排单、预约、收单（仅查看与完工）
const canCollect = computed(() => {
  const r = role.value
  if (r === 'super_admin' || r === 'agent') return true
  if (r === 'agent_manager') {
    return Number(order.value.agentId) === Number(store.userInfo.value?.id)
  }
  return false
})

// 角色+状态 → 可执行操作（对齐后端 canStep）
const actions = computed(() => {
  const s = Number(order.value.orderStatus)
  const r = role.value
  const isSuper = r === 'super_admin'
  const list = []
  if (s === 1 && (r === 'agent_manager' || isSuper)) list.push({ text: '排单', type: 'schedule' })
  if (s === 2 && canCollect.value) list.push({ text: '填写预约情况', type: 'appoint' })
  if (s === 3 && canCollect.value) list.push({ text: '填写收单情况', type: 'collect' })
  if (s === 4 && canCollect.value) list.push({ text: '编辑收单', type: 'collect' })
  if ((r === 'documentary' && s === 5) || r === 'admin' || isSuper) list.push({ text: '完工', type: 'complete' })
  return list.slice(0, 2)
})

const agentOptions = computed(() => agents.value.map((a) => ({ label: a.name || a.id, value: String(a.id) })))
const scheduleAgentLabel = computed(() => {
  const found = agentOptions.value.find((o) => o.value === scheduleAgentId.value)
  return found ? found.label : '请选择收单员'
})

function fmtDate(v) {
  return v ? String(v).slice(0, 10) : ''
}

function resolveUrl(u) {
  if (!u) return u
  if (/^https?:\/\//.test(u)) return u
  if (!u.startsWith('/')) u = '/' + u
  return config.ASSET_URL + u
}

async function loadAttachments() {
  const attachmentIds = (order.value.attachmentId || '').trim()
  if (!attachmentIds) return
  try {
    const data = await getAttachmentList(attachmentIds)
    attachments.value = (data.list || []).map((a) => resolveUrl(a.url))
  } catch (e) {
    // 忽略
  }
}

async function loadOrder() {
  if (!orderId.value) return
  try {
    const detail = await getOrderView(orderId.value)
    if (detail) {
      order.value = {
        ...detail,
        visitDate: fmtDate(detail.visitDate),
        scheduleDate: fmtDate(detail.scheduleDate)
      }
    }
    await loadAttachments()
  } catch (e) {
    // request 已提示
  }
}

onLoad((query) => {
  orderId.value = query.id
  order.value = getCurrentOrder() || {}
  // 确保角色信息已加载（决定可执行操作）
  if (!store.userInfo.value) {
    store.fetchProfile()
  }
})

onShow(() => {
  // 每次显示（含表单返回）重新加载订单
  loadOrder()
})

function preview(i) {
  uni.previewImage({ urls: attachments.value, current: i })
}

function callPhone(phone) {
  if (!phone) return
  uni.makePhoneCall({ phoneNumber: String(phone), fail: () => {} })
}
function copyAddress(address) {
  if (!address) return
  uni.setClipboardData({ data: address, success: () => uni.showToast({ title: '地址已复制', icon: 'none' }) })
}

function onPrimary() {
  if (actions.value[0]) onAction(actions.value[0].type)
}
function onGhost() {
  if (actions.value[1]) onAction(actions.value[1].type)
}

function onAction(type) {
  if (type === 'schedule') {
    openSchedule()
  } else if (type === 'appoint') {
    uni.navigateTo({ url: `/pages/order/appoint/index?id=${order.value.id}` })
  } else if (type === 'collect') {
    uni.navigateTo({ url: `/pages/order/collect/index?id=${order.value.id}` })
  } else if (type === 'complete') {
    completeForm.value = { broadbandAccount: '', isCompleted: '', agencyNo: '' }
    showComplete.value = true
  }
}

async function openSchedule() {
  if (!agents.value.length) {
    try {
      const data = await getAgentList()
      agents.value = data.list || []
    } catch (e) {
      agents.value = []
    }
  }
  scheduleAgentId.value = String(order.value.agentId || '')
  showSchedule.value = true
}

async function doSchedule() {
  if (scheduleAgentId.value === '') {
    uni.showToast({ title: '请选择收单员', icon: 'none' })
    return
  }
  try {
    await orderSchedule({ id: Number(order.value.id), agentId: Number(scheduleAgentId.value) })
    uni.showToast({ title: '排单成功', icon: 'success' })
    showSchedule.value = false
    loadOrder()
  } catch (e) {
    // request 已提示
  }
}

async function doComplete() {
  try {
    await orderComplete({
      id: Number(order.value.id),
      broadbandAccount: completeForm.value.broadbandAccount,
      isCompleted: Number(completeForm.value.isCompleted) || 0,
      agencyNo: completeForm.value.agencyNo
    })
    uni.showToast({ title: '完工成功', icon: 'success' })
    showComplete.value = false
    loadOrder()
  } catch (e) {
    // request 已提示
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
.att-card {
  background: #ffffff;
  border-radius: 24rpx;
  padding: 32rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.04);
}
.att-title {
  display: block;
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
  margin-bottom: 24rpx;
}
.att-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 16rpx;
}
.att-img {
  width: 176rpx;
  height: 176rpx;
  border-radius: 16rpx;
}
.bottom {
  padding-top: 24rpx;
}

.popup-panel {
  padding-bottom: 40rpx;
}
.popup-title {
  display: block;
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
  margin-bottom: 24rpx;
}
.f-label {
  display: block;
  font-size: 28rpx;
  color: #333333;
  font-weight: 600;
  margin: 24rpx 0 16rpx;
}
.schedule-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 24rpx;
  background: #f8f9fa;
  border-radius: 16rpx;
}
.date-text { font-size: 28rpx; color: #333333; }
.date-placeholder { font-size: 28rpx; color: #c0c4cc; }
:deep(.edit-input) {
  background: #f8f9fa;
  border-radius: 16rpx;
  padding: 0 24rpx;
}
.filter-btns {
  display: flex;
  gap: 24rpx;
  margin-top: 40rpx;
}
.btn {
  flex: 1;
  height: 88rpx;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
  font-weight: 700;
}
.btn.primary { background: #336fff; color: #ffffff; }
.btn.ghost { background: #f5f7fa; color: #333333; }
</style>
