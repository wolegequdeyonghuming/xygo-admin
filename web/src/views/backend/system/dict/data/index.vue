<!-- 字典数据管理页面 -->
<template>
  <div class="dict-data-page art-full-height">
    <ElRow :gutter="16" class="dict-data-row">
      <ElCol :span="5">
        <ElCard shadow="never" class="dict-type-list">
          <div class="list-header">
            <span class="list-title">字典类型</span>
            <ElInput
              v-model="typeSearch"
              size="small"
              placeholder="搜索"
              clearable
              style="width: 130px"
            />
          </div>
          <ElScrollbar max-height="calc(100vh - 240px)">
            <div
              v-for="item in filteredTypeList"
              :key="item.id"
              class="type-item"
              :class="{ active: currentTypeId === item.id }"
              @click="selectType(item)"
            >
              <span class="type-name">{{ item.name }}</span>
              <span class="type-code">{{ item.type }}</span>
            </div>
            <div v-if="filteredTypeList.length === 0" class="empty-tip">暂无字典类型</div>
          </ElScrollbar>
        </ElCard>
      </ElCol>

      <ElCol :span="19">
        <ElCard class="art-table-card" shadow="never">
          <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
            <template #left>
              <ElButton
                v-auth="'add'"
                :disabled="!currentTypeId"
                @click="showDialog('add')"
                v-ripple
              >
                新增字典数据
              </ElButton>
            </template>
          </ArtTableHeader>

          <ArtTable
            :loading="loading"
            :data="data"
            :columns="columns"
            :pagination="pagination"
            @pagination:size-change="handleSizeChange"
            @pagination:current-change="handleCurrentChange"
          />

          <DictDataDialog
            v-model:visible="dialogVisible"
            :type="dialogType"
            :dict-type-id="currentTypeId"
            :form-data="currentData"
            @submit="handleDialogSubmit"
          />
        </ElCard>
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
  import { ref, computed, onMounted, h } from 'vue'
  import { useTable } from '@/hooks/core/useTable'
  import { ElMessage, ElMessageBox, ElTag } from 'element-plus'
  import {
    fetchDictDataList,
    fetchDeleteDictData,
    fetchDictTypeList,
    type DictDataItem,
    type DictTypeItem
  } from '@/api/backend/system/dict'
  import DictDataDialog from './modules/data-dialog.vue'
  import { DialogType } from '@/types'
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useAuth } from '@/hooks/core/useAuth'

  const { hasAuth } = useAuth()

  defineOptions({ name: 'DictData' })

  const dialogType = ref<DialogType>('add')
  const dialogVisible = ref(false)
  const currentData = ref<Partial<DictDataItem>>({})
  const currentTypeId = ref(1)
  const typeSearch = ref('')
  const typeList = ref<DictTypeItem[]>([])

  const filteredTypeList = computed(() => {
    if (!typeSearch.value) return typeList.value
    return typeList.value.filter(
      (item) => item.name.includes(typeSearch.value) || item.type.includes(typeSearch.value)
    )
  })

  const loadTypeList = async () => {
    const res = await fetchDictTypeList({ page: 1, pageSize: 100, status: 1 })
    typeList.value = res.list || []
    if (typeList.value.length > 0 && !currentTypeId.value) {
      currentTypeId.value = typeList.value[0].id
      await getData()
    }
  }

  const selectType = (item: any) => {
    currentTypeId.value = item.id
    searchParams.dictTypeId = item.id
    getData()
  }

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    searchParams,
    getData,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    core: {
      apiFn: fetchDictDataList,
      apiParams: {
        page: 1,
        pageSize: 20,
        dictTypeId: currentTypeId.value
      },
      paginationKey: {
        current: 'page',
        size: 'pageSize'
      },
      columnsFactory: () => [
        { type: 'index', width: 60, label: '序号' },
        { prop: 'label', label: '字典标签', minWidth: 120 },
        { prop: 'value', label: '字典值', minWidth: 120 },
        { prop: 'cssClass', label: '样式类名', minWidth: 120 },
        {
          prop: 'listClass',
          label: 'Tag样式',
          width: 120,
          align: 'center',
          formatter: (row: DictDataItem) =>
            h(
              ElTag,
              { type: (row.listClass || 'info') as any, size: 'small' },
              () => row.listClass || '默认'
            )
        },
        {
          prop: 'isDefault',
          label: '默认',
          width: 80,
          align: 'center',
          formatter: (row: DictDataItem) => (row.isDefault === 1 ? '是' : '否')
        },
        {
          prop: 'status',
          label: '状态',
          width: 100,
          align: 'center',
          formatter: (row: DictDataItem) =>
            row.status === 1
              ? h(ElTag, { type: 'success', size: 'small' }, () => '启用')
              : h(ElTag, { type: 'danger', size: 'small' }, () => '禁用')
        },
        { prop: 'sort', label: '排序', width: 80, align: 'center' },
        { prop: 'remark', label: '备注', minWidth: 150 },
        {
          prop: 'action',
          label: '操作',
          width: 160,
          fixed: 'right',
          formatter: (row: DictDataItem) =>
            h('div', { class: 'flex items-center gap-1' }, [
              hasAuth('edit') ? h(ArtButtonTable, { type: 'edit', onClick: () => showDialog('edit', row) }) : null,
              hasAuth('delete') ? h(ArtButtonTable, { type: 'delete', onClick: () => handleDelete(row.id) }) : null
            ].filter(Boolean))
        }
      ]
    }
  })

  const showDialog = (type: DialogType, row?: DictDataItem) => {
    dialogType.value = type
    currentData.value = row ? { ...row } : {}
    dialogVisible.value = true
  }

  const handleDialogSubmit = async () => {
    await getData()
    dialogVisible.value = false
  }

  const handleDelete = (id: number) => {
    ElMessageBox.confirm('确定要删除该字典数据吗？', '提示', {
      type: 'warning',
      confirmButtonText: '确定',
      cancelButtonText: '取消'
    })
      .then(async () => {
        await fetchDeleteDictData(id)
        ElMessage.success('删除成功')
        await getData()
      })
      .catch(() => {})
  }

  onMounted(() => {
    loadTypeList()
  })
</script>

<style scoped lang="scss">
  .dict-data-page {
    display: flex;
    flex-direction: column;

    .dict-data-row {
      flex: 1;
      height: auto;

      .el-col:last-child {
        display: flex;
        flex-direction: column;

        .art-table-card {
          flex: 1;
          display: flex;
          flex-direction: column;

          :deep(.el-card__body) {
            flex: 1;
            display: flex;
            flex-direction: column;

            .art-table {
              flex: 1;
            }
          }
        }
      }
    }
  }

  .dict-type-list {
    height: 100%;
    display: flex;
    flex-direction: column;

    .el-scrollbar {
      flex: 1;
    }

    .list-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 12px;
      flex-shrink: 0;

      .list-title {
        font-weight: 600;
        font-size: 14px;
      }
    }

    .type-item {
      padding: 10px 12px;
      border-radius: 4px;
      cursor: pointer;
      display: flex;
      justify-content: space-between;
      align-items: center;

      &:hover {
        background: var(--el-fill-color-light);
      }

      &.active {
        background: var(--el-color-primary-light-9);
        color: var(--el-color-primary);
      }

      .type-name {
        font-size: 14px;
      }

      .type-code {
        font-size: 12px;
        color: var(--el-text-color-secondary);
      }
    }

    .empty-tip {
      text-align: center;
      padding: 20px 0;
      color: var(--el-text-color-secondary);
      font-size: 13px;
    }
  }
</style>
