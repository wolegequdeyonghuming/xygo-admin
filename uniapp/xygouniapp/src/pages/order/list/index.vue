<template>
  <view class="home-page">
    <!-- 顶部栏 -->
    <view class="top-bar">
      <text class="sys-name">{{ siteName || config.SYS_NAME }}</text>
      <view class="user-info" @tap="goUser">
        <image v-if="avatarUrl" :src="avatarUrl" class="avatar-img" mode="aspectFill" />
        <view v-else class="avatar">{{ avatarChar }}</view>
        <text class="nickname">{{ nickname }}</text>
      </view>
    </view>

    <view class="content">
      <!-- 统计卡：点击切换列表状态 -->
      <stat-cards :items="statItems" :active-index="activeTab" @change="switchTab" />

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
        <order-card v-for="order in list" :key="order.id" :order="order" @select="goDetail" />
      </view>
      <view v-else class="empty">
        <text class="empty-text">暂无订单</text>
      </view>
    </view>

    <tab-bar />
  </view>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { onLoad, onShow, onPullDownRefresh, onReachBottom } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import config from '@/utils/config'
import { tabActive } from '@/utils/tab'
import { setCurrentOrder } from '@/utils/orderBus'
import { siteName, loadSiteName } from '@/utils/site'
import { getOrderList, getOrderStat } from '@/api/staff'

const store = useStaffStore()
const activeTab = ref(0)
const page = ref(1)
const pageSize = 20
const list = ref([])
const total = ref(0)
const loading = ref(false)
const month = ref(defaultMonth())
const monthTs = ref(Date.now())

const stat = ref({ todayRecorded: 0, recorded: 0, pendingSchedule: 0, todo: 0, today: 0, done: 0 })

const RED = '#D92400'
const GREEN = '#299A0C'
const BLUE = '#1F61FF'
const ORANGE = '#FE8B00'

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
  loadSiteName()
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
  background: #f5f6f8;
}
.top-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 20px;
  padding-left: 32rpx;
  padding-right: 32rpx;
  padding-bottom: 24rpx;
}
.sys-name {
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 16rpx;
}
.nickname {
  font-size: 24rpx;
  font-weight: 700;
  color: #333333;
}
.avatar {
  width: 72rpx;
  height: 72rpx;
  border-radius: 50%;
  background: #e0e0e0;
  color: #333333;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28rpx;
  font-weight: 700;
}
.avatar-img {
  width: 72rpx;
  height: 72rpx;
  border-radius: 50%;
}
.content {
  padding: 0 32rpx 130rpx;
}
.list-title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 32rpx 0 24rpx;
}
.list-title {
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
}
.month-picker {
  background: #f8f9fa;
  border-radius: 24rpx;
  padding: 16rpx 24rpx;
}
.month-text {
  font-size: 28rpx;
  font-weight: 700;
  color: #333333;
}
.order-list {
  margin-top: 8rpx;
}
.empty {
  padding: 120rpx 0;
  text-align: center;
}
.empty-text {
  color: #8b8c8f;
  font-size: 28rpx;
}
</style>
