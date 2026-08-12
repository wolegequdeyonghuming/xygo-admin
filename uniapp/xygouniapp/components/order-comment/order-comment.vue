<template>
  <view class="order-comment">
    <view class="head" @tap="expanded = !expanded">
      <text class="title">详细情况</text>
      <text class="action">{{ expanded ? '收起' : `全部详细情况 (${list.length})` }}</text>
    </view>

    <view v-if="expanded" class="body">
      <view v-if="!list.length" class="empty">暂无详细情况</view>
      <view v-for="item in list" :key="item.id" class="item">
        <view class="item-head">
          <text class="user">{{ item.userName || '未知用户' }}</text>
          <text v-if="canDelete(item)" class="del" @tap.stop="handleDelete(item)">删除</text>
        </view>
        <text class="time">{{ item.createdAt || '' }}</text>
        <text class="content">{{ item.content }}</text>
      </view>
    </view>

    <view class="add" @tap="showAdd = true">
      <text class="add-icon">＋</text>
      <text>添加详情</text>
    </view>

    <wd-popup v-model="showAdd" position="bottom" custom-style="padding:32rpx">
      <view class="add-panel">
        <text class="panel-title">添加详情</text>
        <textarea
          v-model="input"
          class="panel-input"
          maxlength="500"
          placeholder="请输入详细情况"
          placeholder-style="color:#C0C4CC"
        />
        <view class="panel-count">{{ input.length }}/500</view>
        <view class="panel-btns">
          <wd-button size="small" plain @click="showAdd = false">取消</wd-button>
          <wd-button size="small" type="primary" :loading="submitting" @click="handleAdd">保存</wd-button>
        </view>
      </view>
    </wd-popup>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getDetailList, addDetail, deleteDetail } from '@/api/staff'
import { useStaffStore } from '@/store/staff'

const props = defineProps({
  orderId: { type: [Number, String], required: true }
})

const store = useStaffStore()
const list = ref([])
const expanded = ref(false)
const showAdd = ref(false)
const input = ref('')
const submitting = ref(false)

async function load() {
  try {
    list.value = (await getDetailList(props.orderId)) || []
  } catch (e) {
    list.value = []
  }
}

function canDelete(item) {
  return store.userInfo && item.userId && String(item.userId) === String(store.userInfo.id)
}

async function handleAdd() {
  if (!input.value.trim()) {
    uni.showToast({ title: '请输入内容', icon: 'none' })
    return
  }
  submitting.value = true
  try {
    await addDetail(props.orderId, input.value.trim())
    input.value = ''
    showAdd.value = false
    uni.showToast({ title: '已添加', icon: 'success' })
    await load()
  } catch (e) {
    uni.showToast({ title: '添加失败', icon: 'none' })
  } finally {
    submitting.value = false
  }
}

async function handleDelete(item) {
  try {
    await deleteDetail(item.id)
    uni.showToast({ title: '已删除', icon: 'success' })
    await load()
  } catch (e) {
    uni.showToast({ title: '删除失败', icon: 'none' })
  }
}

onMounted(load)
</script>

<style scoped lang="scss">
.order-comment {
  background: #f8f9fa;
  border-radius: 24rpx;
  padding: 24rpx;
}
.head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16rpx;
}
.title {
  font-size: 32rpx;
  font-weight: 700;
  color: #333333;
}
.action {
  font-size: 28rpx;
  color: #336fff;
  font-weight: 600;
}
.empty {
  color: #8b8c8f;
  font-size: 28rpx;
  padding: 24rpx 0;
  text-align: center;
}
.item {
  padding: 20rpx 0;
  border-bottom: 1rpx solid #eef0f2;
}
.item-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.user {
  font-size: 28rpx;
  font-weight: 700;
  color: #333333;
}
.del {
  font-size: 26rpx;
  color: #d92400;
}
.time {
  display: block;
  font-size: 24rpx;
  color: #8b8c8f;
  margin-top: 8rpx;
}
.content {
  display: block;
  font-size: 28rpx;
  color: #333333;
  margin-top: 16rpx;
  line-height: 1.5;
}
.add {
  display: inline-flex;
  align-items: center;
  gap: 8rpx;
  margin-top: 24rpx;
  font-size: 28rpx;
  color: #336fff;
  background: #e6f0ff;
  border-radius: 16rpx;
  padding: 12rpx 24rpx;
}
.add-icon {
  font-size: 32rpx;
  line-height: 1;
}
.add-panel {
  .panel-title {
    font-size: 32rpx;
    font-weight: 700;
    color: #333333;
  }
  .panel-input {
    width: 100%;
    height: 200rpx;
    background: #f8f9fa;
    border-radius: 16rpx;
    padding: 20rpx;
    margin-top: 24rpx;
    font-size: 28rpx;
    box-sizing: border-box;
  }
  .panel-count {
    text-align: right;
    color: #8b8c8f;
    font-size: 24rpx;
    margin-top: 8rpx;
  }
  .panel-btns {
    display: flex;
    justify-content: flex-end;
    gap: 24rpx;
    margin-top: 24rpx;
  }
}
</style>
