<template>
  <view class="home-page">
    <!-- 浅蓝舱头 -->
    <view class="deck-head" :style="{ paddingTop: headPad + 'rpx' }">
      <view class="head-grid"></view>
      <view class="head-glow"></view>
      <view class="user-row" @tap="goUser">
        <view class="avatar-wrap">
          <image v-if="avatarUrl" :src="avatarUrl" class="avatar-img" mode="aspectFill" />
          <view v-else class="avatar">{{ avatarChar }}</view>
        </view>
        <view class="user-meta">
          <text class="nickname">{{ displayName }}</text>
          <text class="role">{{ roleName }}</text>
        </view>
        <wd-icon name="arrow-right" size="14" color="#9CA3AF" />
      </view>
    </view>

    <!-- 统计舱 -->
    <view class="stat-wrap">
      <stat-cards :items="statItems" :active-index="activeTab" @change="switchTab" />
    </view>

    <view class="content">
      <!-- 列表标题（已处理 tab 右侧显示月份选择器） -->
      <view class="list-title-row">
        <text class="list-title">{{ currentTitle }}</text>
        <wd-datetime-picker
          v-if="showMonthPicker"
          v-model="monthTs"
          type="year-month"
          @confirm="onMonthConfirm"
        >
          <view class="month-picker"><text class="month-text">{{ month }}</text></view>
        </wd-datetime-picker>
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
        <text class="empty-sub">下拉刷新试试</text>
      </view>
    </view>

    <tab-bar />
  </view>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { onLoad, onShow, onPullDownRefresh, onReachBottom } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import { tabActive } from '@/utils/tab'
import { setCurrentOrder } from '@/utils/orderBus'
import { getOrderList, getOrderStat } from '@/api/staff'
import { getStatusBarHeightRpx } from '@/utils/system'

const store = useStaffStore()
const activeTab = ref(0)
const page = ref(1)
const pageSize = 20
const list = ref([])
const total = ref(0)
const loading = ref(false)
const month = ref(defaultMonth())
const monthTs = ref(Date.now())

const headPad = ref(getStatusBarHeightRpx() + 24)

const stat = ref({ todayRecorded: 0, recorded: 0, pendingSchedule: 0, todo: 0, today: 0, done: 0 })

const RED = '#E5484D'
const GREEN = '#30A46C'
const BLUE = '#2563EB'
const ORANGE = '#E8930C'

// 按角色生成首页统计 tab（label/color/statKey/statuses/mode/title）
const tabs = computed(() => {
  const r = store.userInfo.value?.roleKey || ''
  if (r === 'telemarketer') {
    return [
      { label: '今日录单', color: GREEN, statKey: 'todayRecorded', statuses: [], mode: 'createdToday', title: '今日录单订单' },
      { label: '已录单', color: BLUE, statKey: 'recorded', statuses: [1], title: '已录单订单' }
    ]
  }
  if (r === 'agent') {
    return [
      { label: '待收单', color: RED, statKey: 'todo', statuses: [2, 3], title: '待收单订单' },
      { label: '今日收单', color: GREEN, statKey: 'today', statuses: [4, 5, 6], mode: 'today', title: '今日收单订单' },
      { label: '已收单', color: BLUE, statKey: 'done', statuses: [4, 5, 6], mode: 'month', title: '已收单订单' }
    ]
  }
  // 收单员管理员 / 文员 / 管理员 / 超管
  return [
    { label: '待排单', color: ORANGE, statKey: 'pendingSchedule', statuses: [1], mode: 'pending', title: '待排单订单' },
    { label: '待收单', color: RED, statKey: 'todo', statuses: [2, 3], title: '待收单订单' },
    { label: '已收单', color: BLUE, statKey: 'done', statuses: [4, 5, 6], mode: 'month', title: '已收单订单' }
  ]
})

watch(tabs, () => {
  if (activeTab.value >= tabs.value.length) {
    activeTab.value = 0
  }
  refresh()
})

const nickname = computed(() => store.userInfo.value?.nickname || store.userInfo.value?.realName || '')
const displayName = computed(() => store.userInfo.value?.realName || store.userInfo.value?.nickname || '')
const avatarUrl = computed(() => store.userInfo.value?.avatar || '')
const avatarChar = computed(() => (displayName.value || '员').slice(0, 1))
const roleName = computed(() => store.userInfo.value?.roleName || store.userInfo.value?.postName || '')
const currentTitle = computed(() => (tabs.value[activeTab.value] || {}).title || '')
const statItems = computed(() => tabs.value.map((t) => ({ label: t.label, color: t.color, value: stat.value[t.statKey] ?? 0 })))
const showMonthPicker = computed(() => (tabs.value[activeTab.value] || {}).mode === 'month')

function defaultMonth() {
  const d = new Date()
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}`
}

function todayStr() {
  const d = new Date()
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}

function monthRange(ym) {
  const [y, m] = ym.split('-').map(Number)
  const lastDay = new Date(y, m, 0).getDate()
  return {
    start: `${y}-${String(m).padStart(2, '0')}-01`,
    end: `${y}-${String(m).padStart(2, '0')}-${String(lastDay).padStart(2, '0')}`
  }
}

function fmtDate(v) {
  if (!v) return ''
  return String(v).slice(0, 10)
}

async function loadStat() {
  try {
    stat.value = (await getOrderStat()) || { todayRecorded: 0, recorded: 0, pendingSchedule: 0, todo: 0, today: 0, done: 0 }
  } catch (e) {
    // 忽略
  }
}

async function refresh() {
  page.value = 1
  await fetchList(true)
}

async function fetchList(reset) {
  if (loading.value) return
  loading.value = true
  try {
    const tab = tabs.value[activeTab.value] || {}
    const params = { page: page.value, pageSize }
    if (tab.statuses && tab.statuses.length) {
      params.statuses = tab.statuses.join(',')
    }
    if (tab.mode === 'createdToday') {
      const today = todayStr()
      params.createdDateStart = today
      params.createdDateEnd = today
    } else if (tab.mode === 'pending') {
      params.pendingOnly = true
    } else if (tab.mode === 'today') {
      const today = todayStr()
      params.visitDateStart = today
      params.visitDateEnd = today
    } else if (tab.mode === 'month') {
      const range = monthRange(month.value)
      params.visitDateStart = range.start
      params.visitDateEnd = range.end
    }
    const data = await getOrderList(params)
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

onLoad(() => {
  // 不允许游客登录
  if (!store.isLoggedIn.value) {
    uni.reLaunch({ url: '/pages/login/index' })
    return
  }
  if (!store.userInfo.value) {
    store.fetchProfile()
  }
})

onShow(() => {
  tabActive.value = 0
  // 确保用户资料（姓名/头像）已加载
  if (!store.userInfo.value) {
    store.fetchProfile()
  }
  // 每次显示（含详情/表单返回）刷新统计与当前列表
  loadStat()
  refresh()
})

function switchTab(idx) {
  if (activeTab.value === idx) return
  activeTab.value = idx
  refresh()
}

function onMonthConfirm({ value }) {
  const d = new Date(value)
  month.value = `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}`
  refresh()
}

function goDetail(order) {
  setCurrentOrder(order)
  uni.navigateTo({ url: `/pages/order/detail/index?id=${order.id}` })
}

function goUser() {
  uni.reLaunch({ url: '/pages/user/index' })
}

onPullDownRefresh(() => {
  loadStat()
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
.home-page {
  min-height: 100vh;
  background: #f4f6fa;
}

/* ===== 浅蓝舱头 ===== */
.deck-head {
  position: relative;
  overflow: hidden;
  padding-left: 32rpx;
  padding-right: 32rpx;
  padding-bottom: 48rpx;
  background: linear-gradient(160deg, #f2f6fe 0%, #e7eefb 55%, #e2ebfa 100%);
  border-bottom-left-radius: 40rpx;
  border-bottom-right-radius: 40rpx;
}
.head-grid {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background-image: linear-gradient(rgba(37, 99, 235, 0.05) 1rpx, transparent 1rpx),
    linear-gradient(90deg, rgba(37, 99, 235, 0.05) 1rpx, transparent 1rpx);
  background-size: 40rpx 40rpx;
  pointer-events: none;
}
.head-glow {
  position: absolute;
  right: -80rpx;
  top: -100rpx;
  width: 280rpx;
  height: 280rpx;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(37, 99, 235, 0.14) 0%, rgba(37, 99, 235, 0) 70%);
  pointer-events: none;
}
.user-row {
  position: relative;
  display: flex;
  align-items: center;
  gap: 20rpx;
}
.avatar,
.avatar-img {
  width: 76rpx;
  height: 76rpx;
  border-radius: 50%;
  flex-shrink: 0;
}
.avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #2563eb;
  color: #ffffff;
  font-size: 28rpx;
  font-weight: 800;
}
.avatar-img {
  border: 2rpx solid #ffffff;
  box-shadow: 0 4rpx 12rpx rgba(37, 99, 235, 0.2);
}
.user-meta {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
}
.nickname {
  font-size: 30rpx;
  font-weight: 700;
  color: #111827;
}
.role {
  margin-top: 4rpx;
  font-size: 22rpx;
  color: #4b5563;
}

/* ===== 统计舱 ===== */
.stat-wrap {
  position: relative;
  z-index: 2;
  margin-top: -36rpx;
  padding: 0 32rpx;
}

/* ===== 内容区 ===== */
.content {
  padding: 32rpx 32rpx 220rpx;
}
.list-title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24rpx;
}
.list-title {
  font-size: 32rpx;
  font-weight: 700;
  color: #111827;
}
.month-picker {
  background: #e8effd;
  border-radius: 999rpx;
  padding: 14rpx 26rpx;
}
.month-text {
  font-size: 26rpx;
  font-weight: 600;
  color: #2563eb;
}
.order-list {
  margin-top: 8rpx;
}

/* ===== 空态 ===== */
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
</style>
