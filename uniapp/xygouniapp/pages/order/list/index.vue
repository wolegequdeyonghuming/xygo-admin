<template>
  <view class="home-page">
    <!-- 顶部栏 -->
    <view class="top-bar" :style="{ paddingTop: statusBarHeight + 'px' }">
      <text class="sys-name">收单员工作台</text>
      <view class="user-info" @tap="goUser">
        <text class="nickname">{{ nickname }}</text>
        <view class="avatar">{{ avatarChar }}</view>
      </view>
    </view>

    <view class="content">
      <!-- 统计卡：点击切换列表状态 -->
      <stat-cards :items="statItems" :active-index="activeTab" @change="switchTab" />

      <!-- 搜索 -->
      <view class="search">
        <wd-search v-model="keyword" placeholder="搜索客户姓名 / 联系电话" custom-class="search-box" @search="onSearch" />
      </view>

      <!-- 列表标题 -->
      <view class="list-title">{{ currentTitle }}</view>

      <!-- 订单列表 -->
      <view v-if="list.length" class="order-list">
        <order-card v-for="order in list" :key="order.id" :order="order" @tap="goDetail" />
      </view>
      <view v-else class="empty">
        <text class="empty-text">暂无订单</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, computed, onLoad, onPullDownRefresh, onReachBottom } from '@dcloudio/uni-app'
import { useStaffStore } from '@/store/staff'
import { mockOrders, STAT_TABS, filterOrders } from '@/utils/mock'

const store = useStaffStore()
const statusBarHeight = ref(44)
const activeTab = ref(0)
const keyword = ref('')
const page = ref(1)
const pageSize = 20
const all = ref(mockOrders)
const list = ref([])

const nickname = computed(() => store.userInfo?.nickname || store.userInfo?.realName || '')
const avatarChar = computed(() => (nickname.value || '员').slice(0, 1))
const currentTitle = computed(() => STAT_TABS[activeTab.value].title)
const statItems = computed(() =>
  STAT_TABS.map((t, i) => ({
    label: t.label,
    color: t.color,
    value: filterOrders(i, all.value).length
  }))
)

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
  refresh()
})

function refresh() {
  page.value = 1
  list.value = filterOrders(activeTab.value, all.value, keyword.value).slice(0, pageSize)
}

function switchTab(idx) {
  activeTab.value = idx
  refresh()
}

function onSearch() {
  refresh()
}

function goDetail(order) {
  uni.navigateTo({ url: `/pages/order/detail/index?id=${order.id}` })
}

function goUser() {
  uni.switchTab({ url: '/pages/user/index' })
}

onPullDownRefresh(() => {
  refresh()
  setTimeout(() => uni.stopPullDownRefresh(), 300)
})

onReachBottom(() => {
  const total = filterOrders(activeTab.value, all.value, keyword.value).length
  if (list.value.length < total) {
    page.value++
    list.value = filterOrders(activeTab.value, all.value, keyword.value).slice(0, page.value * pageSize)
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
  padding: 0 32rpx;
}
.search {
  margin: 48rpx 0 16rpx;
}
:deep(.search-box) {
  background: #ffffff;
  border-radius: 20rpx;
}
.list-title {
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
  margin: 16rpx 0 24rpx;
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
