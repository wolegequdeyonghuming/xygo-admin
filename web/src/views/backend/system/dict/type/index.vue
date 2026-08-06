<!-- 字典类型管理页面 -->
<template>
  <div class="dict-type-page art-full-height">
    <ElCard class="art-table-card" shadow="never">
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElButton v-auth="'add'" @click="showDialog('add')" v-ripple>新增字典类型</ElButton>
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

      <DictTypeDialog
        v-model:visible="dialogVisible"
        :type="dialogType"
        :form-data="currentData"
        @submit="handleDialogSubmit"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  import { useTable } from '@/hooks/core/useTable'
  import { ElMessage, ElMessageBox, ElTag } from 'element-plus'
  import {
    fetchDictTypeList,
    fetchDeleteDictType,
    type DictTypeItem
  } from '@/api/backend/system/dict'
  import DictTypeDialog from './modules/type-dialog.vue'
  import { DialogType } from '@/types'

  defineOptions({ name: 'DictType' })

  const dialogType = ref<DialogType>('add')
  const dialogVisible = ref(false)
  const currentData = ref<Partial<DictTypeItem>>({})

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    core: {
      apiFn: fetchDictTypeList,
      apiParams: {
        page: 1,
        pageSize: 20
      },
      paginationKey: {
        current: 'page',
        size: 'pageSize'
      },
      columnsFactory: () => [
        { type: 'index', width: 60, label: '序号' },
        { prop: 'name', label: '字典名称', minWidth: 150 },
        { prop: 'type', label: '字典标识', minWidth: 150 },
        { prop: 'sort', label: '排序', width: 80, align: 'center' },
        {
          prop: 'status',
          label: '状态',
          width: 100,
          align: 'center',
          formatter: (row: DictTypeItem) =>
            row.status === 1
              ? h(ElTag, { type: 'success', size: 'small' }, () => '启用')
              : h(ElTag, { type: 'danger', size: 'small' }, () => '禁用')
        },
        { prop: 'remark', label: '备注', minWidth: 180 },
        {
          prop: 'createTime',
          label: '创建时间',
          width: 180,
          formatter: (row: DictTypeItem) => formatTimestamp(row.createTime)
        },
        {
          prop: 'action',
          label: '操作',
          width: 180,
          fixed: 'right',
          formatter: (row: DictTypeItem) =>
            h('div', { class: 'flex items-center gap-1' }, [
              hasAuth('edit') ? h(ArtButtonTable, { type: 'edit', onClick: () => showDialog('edit', row) }) : null,
              hasAuth('delete') ? h(ArtButtonTable, { type: 'delete', onClick: () => handleDelete(row.id) }) : null
            ].filter(Boolean))
        }
      ]
    }
  })

  const showDialog = (type: DialogType, row?: DictTypeItem) => {
    dialogType.value = type
    currentData.value = row ? { ...row } : {}
    dialogVisible.value = true
  }

  const handleDialogSubmit = async () => {
    await getData()
    dialogVisible.value = false
  }

  const handleDelete = (id: number) => {
    ElMessageBox.confirm('确定要删除该字典类型吗？', '提示', {
      type: 'warning',
      confirmButtonText: '确定',
      cancelButtonText: '取消'
    })
      .then(async () => {
        await fetchDeleteDictType(id)
        ElMessage.success('删除成功')
        await getData()
      })
      .catch(() => {})
  }

  import { h } from 'vue'
  import { formatTimestamp } from '@/utils/time'
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useAuth } from '@/hooks/core/useAuth'

  const { hasAuth } = useAuth()
</script>
