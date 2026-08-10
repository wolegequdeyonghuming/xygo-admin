<!-- 订单详细情况区组件（折叠模式 + 全量弹窗模式） -->
<template>
  <div class="order-comment">
    <!-- 折叠区：标题 + 最新一条 -->
    <div class="order-comment__head">
      <span class="order-comment__title">详细情况</span>
      <ElButton link type="primary" size="small" @click="openAll"
        >全部详细情况 ({{ list.length }})</ElButton
      >
    </div>

    <div v-if="latest" class="order-comment__latest" @click="openAll">
      <div class="order-comment__meta">
        <span class="order-comment__name">{{ latest.userName || '未知用户' }}</span>
        <span class="order-comment__time">{{ formatTimestamp(latest.createdAt) }}</span>
      </div>
      <div class="order-comment__content order-comment__content--ellipsis" :title="latest.content">
        {{ latest.content }}
      </div>
    </div>
    <div v-else class="order-comment__empty">暂无详细情况</div>

    <!-- 输入区：默认折叠，点击「增加详情」展开 -->
    <div v-if="!inputVisible" class="order-comment__trigger">
      <ElButton type="primary" plain size="small" :icon="EditPen" @click="expandInput"
        >增加详情</ElButton
      >
    </div>
    <div v-else class="order-comment__input">
      <ElInput
        ref="inputRef"
        v-model="draft"
        type="textarea"
        :rows="2"
        maxlength="500"
        show-word-limit
        resize="none"
        placeholder="输入详细情况……"
      />
      <div class="order-comment__input-actions">
        <ElButton size="small" @click="collapseInput">取消</ElButton>
        <ElButton
          type="primary"
          size="small"
          :loading="submitting"
          :disabled="!draft.trim()"
          @click="handleAdd"
        >
          保存
        </ElButton>
      </div>
    </div>

    <!-- 全量弹窗 -->
    <ElDialog
      v-model="allVisible"
      :title="`全部详细情况 (${list.length})`"
      width="560px"
      :close-on-click-modal="false"
    >
      <div v-if="list.length" class="order-comment__all">
        <div v-for="item in list" :key="item.id" class="order-comment__item">
          <div class="order-comment__meta">
            <span class="order-comment__name">{{ item.userName || '未知用户' }}</span>
            <span class="order-comment__time">{{ formatTimestamp(item.createdAt) }}</span>
            <ElButton
              v-if="item.userId === currentUserId()"
              link
              type="danger"
              size="small"
              @click="handleDelete(item)"
            >
              删除
            </ElButton>
          </div>
          <div class="order-comment__content">{{ item.content }}</div>
        </div>
      </div>
      <div v-else class="order-comment__empty">暂无详细情况</div>
    </ElDialog>
  </div>
</template>

<script setup lang="ts">
  import { ref, computed, nextTick } from 'vue'
  import { ElMessageBox } from 'element-plus'
  import { EditPen } from '@element-plus/icons-vue'
  import { formatTimestamp } from '@/utils/time'
  import {
    fetchBizOrderDetailList,
    fetchBizOrderDetailAdd,
    fetchBizOrderDetailDelete
  } from '@/api/order/biz-order'
  import { useOrderPerm } from '../useOrderPerm'

  defineOptions({ name: 'OrderComment' })

  const props = defineProps<{
    orderId: number
  }>()

  const { currentUserId } = useOrderPerm()

  const list = ref<Record<string, any>[]>([])
  const draft = ref('')
  const submitting = ref(false)
  const allVisible = ref(false)
  const loading = ref(false)
  /** 输入区是否展开（默认折叠，点击「增加详情」展开） */
  const inputVisible = ref(false)
  const inputRef = ref()

  /** 最新一条（时间正序的最后一条） */
  const latest = computed(() => (list.value.length ? list.value[list.value.length - 1] : null))

  const loadList = async () => {
    if (!props.orderId) return
    loading.value = true
    try {
      list.value = (await fetchBizOrderDetailList(props.orderId))?.list ?? []
    } catch {
      list.value = []
    } finally {
      loading.value = false
    }
  }

  const openAll = () => {
    allVisible.value = true
  }

  const expandInput = () => {
    inputVisible.value = true
    nextTick(() => inputRef.value?.focus?.())
  }

  const collapseInput = () => {
    inputVisible.value = false
    draft.value = ''
  }

  const handleAdd = async () => {
    const content = draft.value.trim()
    if (!content) return
    submitting.value = true
    try {
      await fetchBizOrderDetailAdd({ orderId: props.orderId, content })
      draft.value = ''
      await loadList()
      collapseInput()
    } finally {
      submitting.value = false
    }
  }

  const handleDelete = async (item: Record<string, any>) => {
    try {
      await ElMessageBox.confirm('确定要删除该详细情况吗？删除后无法恢复', '删除确认', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      await fetchBizOrderDetailDelete(item.id)
      await loadList()
    } catch {
      /* 取消或失败均忽略 */
    }
  }

  defineExpose({ loadList })

  watch(
    () => props.orderId,
    () => loadList(),
    { immediate: true }
  )
</script>

<style scoped lang="scss">
  .order-comment {
    &__head {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 8px;
    }
    &__title {
      font-weight: 600;
      color: var(--el-text-color-primary);
    }
    &__latest {
      padding: 8px 10px;
      border: 1px solid var(--el-border-color-lighter);
      border-radius: var(--el-border-radius-base);
      background-color: var(--el-fill-color-light);
      cursor: pointer;
      margin-bottom: 8px;
      transition: border-color 0.2s;
      &:hover {
        border-color: var(--el-color-primary);
      }
    }
    &__meta {
      display: flex;
      align-items: center;
      gap: 8px;
      margin-bottom: 4px;
    }
    &__name {
      font-weight: 500;
      color: var(--el-color-primary);
      font-size: 13px;
    }
    &__time {
      font-size: 12px;
      color: var(--el-text-color-secondary);
    }
    &__content {
      font-size: 13px;
      color: var(--el-text-color-regular);
      white-space: pre-wrap;
      word-break: break-word;
      &--ellipsis {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }
    &__empty {
      padding: 12px 0;
      text-align: center;
      font-size: 13px;
      color: var(--el-text-color-placeholder);
    }
    &__trigger {
      display: flex;
      justify-content: flex-end;
      margin-top: 4px;
    }
    &__input {
      &-actions {
        display: flex;
        justify-content: flex-end;
        margin-top: 6px;
      }
    }
    &__all {
      max-height: 420px;
      overflow-y: auto;
    }
    &__item {
      padding: 10px 0;
      border-bottom: 1px dashed var(--el-border-color-lighter);
      &:last-child {
        border-bottom: none;
      }
    }
  }
</style>
