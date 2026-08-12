<template>
  <view class="home-page">
    <!-- 顶部栏 -->
    <view class="top-bar" :style="{ paddingTop: statusBarHeight + 'px' }">
      <text class="sys-name">{{ config.SYS_NAME }}</text>
      <view class="user-info" @tap="goUser">
        <text class="nickname">{{ nickname }}</text>
        <view class="avatar">{{ avatarChar }}</view>
      </view>
    </view>

    <view class="content">
      <!-- 统计卡：点击切换列表状态 -->
      <stat-cards :items="statItems" :active-index="activeTab" @change="switchTab" />

      <!-- 列表标题（已处理 tab 右侧显示月份选择器） -->
      <view class="list-title-row">
        <text class="list-title">{{ currentTitle }}</text>
        <wd-datetime-picker
          v-if="activeTab === 2"
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
import { ref, computed } from 'vue'
import { onLoad, onShow, onPullDownRefresh, onReachBottom } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import config from '@/utils/config'
import { tabActive } from '@/utils/tab'
import { setCurrentOrder } from '@/utils/orderBus'
import { getOrderList, getOrderStat } from '@/api/staff'

const store = useStaffStore()
const statusBarHeight = ref(44)
const activeTab = ref(0)
const page = ref(1)
const pageSize = 20
const list = ref([])
const total = ref(0)
const loading = ref(false)
const month = ref(defaultMonth())
const monthTs = ref(Date.now())

const stat = ref({ todo: 0, today: 0, done: 0 })

const TABS = [
  { label: '待处理', color: '#D92400', statuses: [2, 3], title: '待处理订单' },
  { label: '今日已处理', color: '#299A0C', statuses: [4, 5, 6], title: '今日已处理订单' },
  { label: '已处理', color: '#1F61FF', statuses: [4, 5, 6], title: '已处理订单' }
]

const nickname = computed(() => store.userInfo?.nickname || store.userInfo?.realName || '')
const avatarChar = computed(() => (nickname.value || '员').slice(0, 1))
const currentTitle = computed(() => TABS[activeTab.value].title)
const statItems = computed(() => [
  { label: TABS[0].label, color: TABS[0].color, value: stat.value.todo },
  { label: TABS[1].label, color: TABS[1].color, value: stat.value.today },
  { label: TABS[2].label, color: TABS[2].color, value: stat.value.done }
])

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
    stat.value = (await getOrderStat()) || { todo: 0, today: 0, done: 0 }
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
    const tab = TABS[activeTab.value]
    const params = { page: page.value, pageSize, statuses: tab.statuses.join(',') }
    // 今日已处理：上门日期=今天；已处理：上门日期在选中月份内
    if (activeTab.value === 1) {
      const today = todayStr()
      params.visitDateStart = today
      params.visitDateEnd = today
    } else if (activeTab.value === 2) {
      const range = monthRange(month.value)
      params.visitDateStart = range.start
      params.visitDateEnd = range.end
    }
    const data = await getOrderList(params)
    // TODO: 调试完成后移除
    console.log('list sample', data.list && data.list[0])
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
  const sysInfo = uni.getSystemInfoSync()
  statusBarHeight.value = sysInfo.statusBarHeight || 44
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
