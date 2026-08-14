<template>
  <view class="query-page">
    <!-- 状态栏 spacer -->
    <view class="sb-spacer" :style="{ height: headPad + 'rpx' }"></view>

    <view class="content">
      <!-- 默认查询行：状态选择器 + 客户姓名输入框 + 筛选按钮 -->
      <view class="search-row">
        <wd-picker v-model="status" :columns="statusOptions" use-default-slot custom-class="status-picker" @confirm="refresh">
          <view class="status-cell">
            <text class="status-text">{{ statusLabel }}</text>
          </view>
        </wd-picker>
        <wd-input v-model="customerName" placeholder="客户姓名" clearable custom-class="name-input" @update:modelValue="onNameChange" />
        <view class="filter-btn" @tap="showFilter = true">
          <wd-icon name="filter" size="18" color="#FFFFFF" />
        </view>
      </view>

      <!-- 已设置更多条件提示 -->
      <view v-if="hasFilter" class="filter-tip" @tap="showFilter = true">
        <text class="filter-tip-text">已设置更多筛选条件，点击查看</text>
      </view>

      <!-- 订单列表 -->
      <view v-if="list.length" class="order-list">
        <order-card v-for="(order, i) in list" :key="order.id" :order="order" :index="i" @select="goDetail" />
      </view>
      <view v-else class="empty">
        <view class="radar">
          <view class="radar-ring r1"></view>
          <view class="radar-ring r2"></view>
          <view class="radar-ring r3"></view>
          <view class="radar-dot"></view>
        </view>
        <text class="empty-text">暂无订单</text>
        <text class="empty-sub">换个筛选条件试试</text>
      </view>
    </view>

    <!-- 筛选弹窗 -->
    <wd-popup v-model="showFilter" position="bottom" :z-index="2000" custom-style="padding:32rpx">
      <view class="filter-panel">
        <text class="filter-title">筛选</text>

        <text class="f-label">排单日期</text>
        <view class="date-pair">
          <view class="date-col">
            <wd-datetime-picker v-model="scheduleStartTs" type="date" :min-date="MIN_DATE" :default-value="todayTs" @confirm="(e) => onDateConfirm('scheduleStart', e)">
              <view class="date-cell"><text :class="filters.scheduleStart ? 'date-text' : 'date-placeholder'">{{ filters.scheduleStart || '开始日期' }}</text></view>
            </wd-datetime-picker>
          </view>
          <text class="date-sep">至</text>
          <view class="date-col">
            <wd-datetime-picker v-model="scheduleEndTs" type="date" :min-date="MIN_DATE" :default-value="todayTs" @confirm="(e) => onDateConfirm('scheduleEnd', e)">
              <view class="date-cell"><text :class="filters.scheduleEnd ? 'date-text' : 'date-placeholder'">{{ filters.scheduleEnd || '结束日期' }}</text></view>
            </wd-datetime-picker>
          </view>
        </view>

        <text class="f-label">上门日期</text>
        <view class="date-pair">
          <view class="date-col">
            <wd-datetime-picker v-model="visitStartTs" type="date" :min-date="MIN_DATE" :default-value="todayTs" @confirm="(e) => onDateConfirm('visitStart', e)">
              <view class="date-cell"><text :class="filters.visitStart ? 'date-text' : 'date-placeholder'">{{ filters.visitStart || '开始日期' }}</text></view>
            </wd-datetime-picker>
          </view>
          <text class="date-sep">至</text>
          <view class="date-col">
            <wd-datetime-picker v-model="visitEndTs" type="date" :min-date="MIN_DATE" :default-value="todayTs" @confirm="(e) => onDateConfirm('visitEnd', e)">
              <view class="date-cell"><text :class="filters.visitEnd ? 'date-text' : 'date-placeholder'">{{ filters.visitEnd || '结束日期' }}</text></view>
            </wd-datetime-picker>
          </view>
        </view>

        <text class="f-label">区县</text>
        <wd-picker v-model="filters.area" :columns="areaOptions" use-default-slot>
          <view class="area-cell">
            <text class="area-text">{{ areaLabel }}</text>
          </view>
        </wd-picker>

        <text class="f-label">安装地址</text>
        <wd-input v-model="filters.installAddress" placeholder="请输入安装地址" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

        <text class="f-label">联系电话</text>
        <wd-input v-model="filters.contactPhone" placeholder="请输入联系电话" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

        <text class="f-label">预约业务</text>
        <wd-input v-model="filters.businessType" placeholder="请输入预约业务" placeholder-style="color:#C0C4CC" custom-class="edit-input" />

        <view class="filter-btns">
          <view class="btn ghost" @tap="resetFilter">重置</view>
          <view class="btn primary" @tap="confirmFilter">确定</view>
        </view>
      </view>
    </wd-popup>

    <tab-bar />
  </view>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { onLoad, onShow, onPullDownRefresh, onReachBottom } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import { tabActive } from '@/utils/tab'
import { setCurrentOrder } from '@/utils/orderBus'
import { getOrderList, getDictData } from '@/api/staff'
import { getStatusBarHeightRpx } from '@/utils/system'

const store = useStaffStore()
const headPad = ref(getStatusBarHeightRpx() + 24)
const page = ref(1)
const pageSize = 20
const list = ref([])
const total = ref(0)
const loading = ref(false)

const status = ref('')
const customerName = ref('')
const showFilter = ref(false)

const filters = reactive({
  scheduleStart: '',
  scheduleEnd: '',
  visitStart: '',
  visitEnd: '',
  area: '',
  installAddress: '',
  contactPhone: '',
  businessType: ''
})

const scheduleStartTs = ref('')
const scheduleEndTs = ref('')
const visitStartTs = ref('')
const visitEndTs = ref('')

const MIN_DATE = new Date(2026, 0, 1).getTime()
const todayTs = Date.now()

const statusOptions = [
  { label: '全部状态', value: '' },
  { label: '未成交', value: '0' },
  { label: '已录单', value: '1' },
  { label: '已接单', value: '2' },
  { label: '已预约', value: '3' },
  { label: '已上门', value: '4' },
  { label: '已回访', value: '5' },
  { label: '已完工', value: '6' }
]

const areaOptions = ref([])

const statusLabel = computed(() => {
  const found = statusOptions.find((o) => o.value === status.value)
  return found ? found.label : '全部状态'
})
const areaLabel = computed(() => {
  const found = areaOptions.value.find((o) => o.value === filters.area)
  return found ? found.label : '全部区县'
})

const hasFilter = computed(
  () =>
    filters.scheduleStart ||
    filters.scheduleEnd ||
    filters.visitStart ||
    filters.visitEnd ||
    filters.area ||
    filters.installAddress ||
    filters.contactPhone ||
    filters.businessType
)

function fmtDate(v) {
  return v ? String(v).slice(0, 10) : ''
}

function buildParams(curPage) {
  const p = { page: curPage, pageSize }
  if (status.value !== '') p.statuses = status.value
  if (customerName.value.trim()) p.keyword = customerName.value.trim()
  if (filters.scheduleStart && filters.scheduleEnd) {
    p.scheduleDateStart = filters.scheduleStart
    p.scheduleDateEnd = filters.scheduleEnd
  }
  if (filters.visitStart && filters.visitEnd) {
    p.visitDateStart = filters.visitStart
    p.visitDateEnd = filters.visitEnd
  }
  if (filters.area !== '') p.area = Number(filters.area)
  if (filters.installAddress.trim()) p.installAddress = filters.installAddress.trim()
  if (filters.contactPhone.trim()) p.contactPhone = filters.contactPhone.trim()
  if (filters.businessType.trim()) p.businessType = filters.businessType.trim()
  return p
}

async function fetchList(reset) {
  if (loading.value) return
  loading.value = true
  try {
    const data = await getOrderList(buildParams(page.value))
    const items = (data.list || []).map((o) => ({
      ...o,
      visitDate: fmtDate(o.visitDate),
      scheduleDate: fmtDate(o.scheduleDate)
    }))
    total.value = data.total || 0
    list.value = reset ? items : list.value.concat(items)
  } catch (e) {
    // request 已提示
  } finally {
    loading.value = false
  }
}

async function refresh() {
  page.value = 1
  await fetchList(true)
}

let nameTimer = null
function onNameChange() {
  // 客户姓名即时查询（300ms 防抖）
  if (nameTimer) clearTimeout(nameTimer)
  nameTimer = setTimeout(() => refresh(), 300)
}

function onDateConfirm(target, { value }) {
  const d = new Date(value)
  filters[target] = `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}

function resetFilter() {
  filters.scheduleStart = ''
  filters.scheduleEnd = ''
  filters.visitStart = ''
  filters.visitEnd = ''
  filters.area = ''
  filters.installAddress = ''
  filters.contactPhone = ''
  filters.businessType = ''
}

function confirmFilter() {
  showFilter.value = false
  refresh()
}

function goDetail(order) {
  setCurrentOrder(order)
  uni.navigateTo({ url: `/pages/order/detail/index?id=${order.id}` })
}

onLoad(() => {
  if (!store.isLoggedIn.value) {
    uni.reLaunch({ url: '/pages/login/index' })
    return
  }
  if (!store.userInfo.value) {
    store.fetchProfile()
  }
})

onShow(() => {
  tabActive.value = 1
  refresh()
})

onMounted(async () => {
  try {
    const data = await getDictData('area')
    areaOptions.value = (data.list || []).map((i) => ({ label: i.label, value: i.value }))
  } catch (e) {
    areaOptions.value = []
  }
})

onPullDownRefresh(() => {
  refresh()
  setTimeout(() => uni.stopPullDownRefresh(), 300)
})

onReachBottom(() => {
  if (list.value.length < total.value) {
    page.value++
    fetchList(false)
  }
})
</script>

<style scoped lang="scss">
.query-page {
  min-height: 100vh;
  background: #f4f6fa;
}
.sb-spacer {
  width: 100%;
}

.content {
  padding-top: 24rpx;
  padding-left: 32rpx;
  padding-right: 32rpx;
  padding-bottom: 220rpx;
  box-sizing: border-box;
}
.search-row {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 24rpx;
}
:deep(.status-picker) {
  width: 220rpx;
  flex-shrink: 0;
}
.status-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 20rpx;
  background: #ffffff;
  border-radius: 16rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 4rpx 16rpx rgba(17, 24, 39, 0.05);
}
.status-text {
  font-size: 28rpx;
  color: #111827;
}
:deep(.name-input) {
  flex: 1;
  height: 80rpx;
  box-sizing: border-box;
  background: #ffffff;
  border-radius: 16rpx;
  border: 1rpx solid rgba(17, 24, 39, 0.06);
  box-shadow: 0 4rpx 16rpx rgba(17, 24, 39, 0.05);
  padding: 0 24rpx;
}
:deep(.name-input .wd-input__inner) {
  height: 80rpx;
  line-height: 80rpx;
  padding: 0;
}
.filter-btn {
  width: 96rpx;
  height: 80rpx;
  border-radius: 16rpx;
  background: #2563eb;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 8rpx 20rpx rgba(37, 99, 235, 0.25);
}
.filter-tip {
  background: #e8effd;
  border-radius: 12rpx;
  padding: 16rpx 24rpx;
  margin-bottom: 24rpx;
}
.filter-tip-text {
  font-size: 26rpx;
  color: #2563eb;
}
.order-list {
  margin-top: 8rpx;
}
.empty {
  padding: 100rpx 0;
  text-align: center;
}
.radar {
  position: relative;
  width: 140rpx;
  height: 140rpx;
  margin: 0 auto 28rpx;
}
.radar-ring {
  position: absolute;
  left: 50%;
  top: 50%;
  border-radius: 50%;
  border: 2rpx dashed rgba(37, 99, 235, 0.35);
  transform: translate(-50%, -50%);
}
.radar-ring.r1 {
  width: 140rpx;
  height: 140rpx;
}
.radar-ring.r2 {
  width: 96rpx;
  height: 96rpx;
  border-style: solid;
  border-color: rgba(37, 99, 235, 0.16);
}
.radar-ring.r3 {
  width: 52rpx;
  height: 52rpx;
  animation: ringPulse 1.6s ease-in-out infinite;
}
.radar-dot {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 16rpx;
  height: 16rpx;
  border-radius: 50%;
  background: #2563eb;
  box-shadow: 0 0 24rpx rgba(37, 99, 235, 0.55);
  transform: translate(-50%, -50%);
}
.empty-text {
  display: block;
  font-size: 30rpx;
  font-weight: 700;
  color: #4b5563;
}
.empty-sub {
  display: block;
  margin-top: 10rpx;
  font-size: 24rpx;
  color: #9ca3af;
}
@keyframes ringPulse {
  0%,
  100% {
    opacity: 1;
    transform: translate(-50%, -50%) scale(1);
  }
  50% {
    opacity: 0.4;
    transform: translate(-50%, -50%) scale(0.8);
  }
}

.filter-panel {
  max-height: 70vh;
  overflow-y: auto;
  padding-bottom: 60rpx;
}
.filter-title {
  display: block;
  font-size: 32rpx;
  font-weight: 700;
  color: #111827;
  margin-bottom: 24rpx;
}
.f-label {
  display: block;
  font-size: 28rpx;
  color: #111827;
  font-weight: 500;
  margin: 24rpx 0 16rpx;
}
.date-pair {
  display: flex;
  align-items: center;
  gap: 12rpx;
}
.date-col {
  flex: 1;
  min-width: 0;
}
.date-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 20rpx;
  background: #f4f6fa;
  border-radius: 16rpx;
}
.date-text { font-size: 28rpx; color: #111827; }
.date-placeholder { font-size: 28rpx; color: #b9bec6; }
.date-sep {
  font-size: 24rpx;
  color: #9ca3af;
  flex-shrink: 0;
}
.area-cell {
  height: 80rpx;
  display: flex;
  align-items: center;
  padding: 0 24rpx;
  background: #f4f6fa;
  border-radius: 16rpx;
}
.area-text { font-size: 28rpx; color: #111827; }
:deep(.edit-input) {
  background: #f4f6fa;
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
.btn.primary {
  background: #2563eb;
  color: #ffffff;
  box-shadow: 0 8rpx 20rpx rgba(37, 99, 235, 0.25);
}
.btn.ghost { background: #eff2f7; color: #111827; }
</style>
