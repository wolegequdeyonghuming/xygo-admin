<!-- 订单主表管理 -->
<template>
  <div class="biz-order-page art-full-height">
    <!-- 搜索栏 -->
    <BizOrderSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams" />

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton
              v-if="anyRoleCanStep(STEP.RECORD)"
              type="primary"
              @click="openStepDialog(STEP.RECORD, 0)"
              v-ripple
              >录单</ElButton
            >
            <ElButton v-auth="'add'" @click="showDialog('add')" v-ripple>新增</ElButton>
            <ElButton v-auth="'export'" @click="handleExport" v-ripple>导出</ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

      <!-- 表格 -->
      <ArtTable
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      />
      <!-- 编辑弹窗 -->
      <BizOrderDialog
        v-model:visible="dialogVisible"
        :type="dialogType"
        :edit-data="currentRow"
        @submit="handleDialogSubmit"
      />
      <!-- 分步操作弹窗 -->
      <BizOrderStepDialog
        v-model:visible="stepDialogVisible"
        :step="currentStep"
        :order-id="currentStepOrderId"
        @submit="handleStepSubmit"
        @submitNext="handleStepSubmitNext"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useTable } from '@/hooks/core/useTable'
  import { useAuth } from '@/hooks/core/useAuth'
  import { formatTimestamp } from '@/utils/time'
  import {
    fetchBizOrderList,
    fetchBizOrderEdit,
    fetchBizOrderDelete,
    fetchBizOrderStepEdit,
    fetchBizOrderStepNext
  } from '@/api/order/biz-order'
  import BizOrderSearch from './modules/biz-order-search.vue'
  import BizOrderDialog from './modules/biz-order-dialog.vue'
  import BizOrderStepDialog from './modules/biz-order-step-dialog.vue'
  import { useOrderPerm, ORDER_STEP as STEP } from './useOrderPerm'
  import { ElTag, ElButton, ElMessageBox } from 'element-plus'
  import { useDictStore } from '@/store/modules/dict'
  import DictLabel from '@/components/DictLabel/index.vue'
  const dictStore = useDictStore()
  import { useRouter } from 'vue-router'
  import { DialogType } from '@/types'
  import { ArrowRight, Edit } from '@element-plus/icons-vue'

  defineOptions({ name: 'BizOrder' })
  const { hasAuth } = useAuth()
  const { anyRoleCanStep, canStepFromStatus, canOperateOrder, canDeleteOrder, nextStepForOrder } =
    useOrderPerm()
  const router = useRouter()
  const dialogType = ref<DialogType>('add')
  const dialogVisible = ref(false)
  const currentRow = ref<any>({})

  // 分步弹窗状态
  const stepDialogVisible = ref(false)
  const currentStep = ref<number>(STEP.RECORD)
  const currentStepOrderId = ref(0)

  const searchForm = ref({
    orderStatus: undefined,
    scheduleDateRange: [],
    visitDateRange: [],
    customerName: undefined,
    area: undefined,
    installAddress: undefined,
    contactPhone: undefined,
    businessType: undefined
  })

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    searchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    core: {
      apiFn: fetchBizOrderList,
      apiParams: {
        page: 1,
        pageSize: 20,
        ...searchForm.value
      },
      paginationKey: { current: 'page', size: 'pageSize' },
      columnsFactory: () => [
        {
          prop: 'orderStatus',
          label: '订单状态',
          width: 100,
          align: 'center',
          formatter: (row: any) =>
            h(DictLabel, { dictType: 'order_status', value: row.orderStatus })
        },
        {
          prop: 'scheduleDate',
          label: '排单日期',
          width: 180,
          formatter: (row: any) => formatTimestamp(row.scheduleDate)
        },
        {
          prop: 'visitDate',
          label: '上门日期',
          width: 180,
          formatter: (row: any) => formatTimestamp(row.visitDate)
        },
        {
          prop: 'customerName',
          label: '客户姓名',
          minWidth: 120,
          formatter: (row: any) => row.customerName ?? '-'
        },
        {
          prop: 'availableTimeDesc',
          label: '可联系时间',
          minWidth: 160,
          formatter: (row: any) => row.availableTimeDesc ?? '-'
        },
        {
          prop: 'area',
          label: '区县',
          width: 100,
          align: 'center',
          formatter: (row: any) => h(DictLabel, { dictType: 'area', value: row.area })
        },
        {
          prop: 'installAddress',
          label: '安装地址',
          minWidth: 120,
          formatter: (row: any) => row.installAddress ?? '-'
        },
        {
          prop: 'contactPhone',
          label: '联系电话',
          minWidth: 120,
          formatter: (row: any) => row.contactPhone ?? '-'
        },
        {
          prop: 'businessType',
          label: '预约业务',
          minWidth: 120,
          formatter: (row: any) => row.businessType ?? '-'
        },
        {
          prop: 'expiryDate',
          label: '到期时间',
          minWidth: 120,
          formatter: (row: any) => row.expiryDate ?? '-'
        },
        {
          prop: 'primaryTelNo',
          label: '主卡号码',
          minWidth: 120,
          formatter: (row: any) => row.primaryTelNo ?? '-'
        },
        {
          prop: 'details',
          label: '详细情况',
          minWidth: 120,
          formatter: (row: any) => row.details ?? '-'
        },
        {
          prop: 'telemarketer_real_name',
          label: '话务员',
          minWidth: 120,
          formatter: (row: any) => row.telemarketer_real_name ?? '-'
        },
        {
          prop: 'agent_real_name',
          label: '收单员',
          minWidth: 120,
          formatter: (row: any) => row.agent_real_name ?? '-'
        },
        {
          prop: 'appointmentDesc',
          label: '收单预约情况',
          minWidth: 160,
          formatter: (row: any) => row.appointmentDesc ?? '-'
        },
        {
          prop: 'followUpDesc',
          label: '话务二次回访情况',
          minWidth: 160,
          formatter: (row: any) => row.followUpDesc ?? '-'
        },
        {
          prop: 'dealtBusinessType',
          label: '成交业务',
          minWidth: 120,
          formatter: (row: any) => row.dealtBusinessType ?? '-'
        },
        {
          prop: 'portingStatus',
          label: '携转情况',
          minWidth: 120,
          formatter: (row: any) => row.portingStatus ?? '-'
        },
        {
          prop: 'customerRealName',
          label: '客户实际姓名',
          minWidth: 160,
          formatter: (row: any) => row.customerRealName ?? '-'
        },
        {
          prop: 'customerIdNumber',
          label: '客户身份证号',
          minWidth: 160,
          formatter: (row: any) => row.customerIdNumber ?? '-'
        },
        {
          prop: 'paidAmount',
          label: '实缴额度',
          minWidth: 120,
          formatter: (row: any) => row.paidAmount ?? '-'
        },
        {
          prop: 'isRuralOrder',
          label: '是否乡下单',
          width: 100,
          align: 'center',
          formatter: (row: any) => {
            const map: Record<string, [string, string]> = {
              '0': ['否', 'success'],
              '1': ['是', 'danger']
            }
            const m = map[String(row.isRuralOrder)]
            return m
              ? h(ElTag, { type: m[1] as any, size: 'small' }, () => m[0])
              : h(ElTag, { size: 'small' }, () => String(row.isRuralOrder ?? '-'))
          }
        },
        {
          prop: 'newPhoneNo',
          label: '新开号码',
          minWidth: 120,
          formatter: (row: any) => row.newPhoneNo ?? '-'
        },
        {
          prop: 'deviceSerial',
          label: '终端串码',
          minWidth: 120,
          formatter: (row: any) => row.deviceSerial ?? '-'
        },
        {
          prop: 'broadbandAccount',
          label: '宽带账号',
          minWidth: 120,
          formatter: (row: any) => row.broadbandAccount ?? '-'
        },
        {
          prop: 'isCompleted',
          label: '是否完工',
          width: 100,
          align: 'center',
          formatter: (row: any) => {
            const map: Record<string, [string, string]> = {
              '0': ['未完工', 'success'],
              '1': ['已完工', 'danger']
            }
            const m = map[String(row.isCompleted)]
            return m
              ? h(ElTag, { type: m[1] as any, size: 'small' }, () => m[0])
              : h(ElTag, { size: 'small' }, () => String(row.isCompleted ?? '-'))
          }
        },
        {
          prop: 'subsidyAmount',
          label: '话补',
          minWidth: 100,
          formatter: (row: any) => row.subsidyAmount ?? '-'
        },
        {
          prop: 'agencyNo',
          label: '工号',
          minWidth: 100,
          formatter: (row: any) => row.agencyNo ?? '-'
        },
        {
          prop: 'isNew',
          label: '是否纯新增',
          minWidth: 160,
          formatter: (row: any) => row.isNew ?? '-'
        },
        {
          prop: 'remark',
          label: '备注',
          minWidth: 200,
          formatter: (row: any) => row.remark ?? '-'
        },
        // ---- 关联表展示字段 ----
        {
          prop: 'operation',
          label: '操作',
          width: 360,
          fixed: 'right',
          formatter: (row: any) => buildOperationButtons(row)
        }
      ]
    }
  })

  /** 操作按钮文案（与弹窗标题一致，第4步为「收单」） */
  const stepLabelMap: Record<number, string> = {
    [STEP.RECORD]: '录单',
    [STEP.ASSIGN]: '派单',
    [STEP.APPOINT]: '预约',
    [STEP.VISIT]: '收单',
    [STEP.FOLLOW]: '回访',
    [STEP.COMPLETE]: '完成'
  }

  /** 生成操作列按钮（按角色 + 状态动态渲染） */
  const buildOperationButtons = (row: any) => {
    const buttons: any[] = []
    const status = Number(row.orderStatus)
    // 按状态 + 角色列出全部可执行步骤按钮（录单/派单/预约/收单/回访/完成）
    for (const step of [
      STEP.RECORD,
      STEP.ASSIGN,
      STEP.APPOINT,
      STEP.VISIT,
      STEP.FOLLOW,
      STEP.COMPLETE
    ]) {
      if (canStepFromStatus(status, step) && canOperateOrder(row)) {
        buttons.push(
          h(
            ElButton,
            {
              type: 'primary',
              icon: Edit,
              plain: true,
              class: 'inline-flex items-center justify-center mr-2 rounded-md align-middle',
              onClick: () => openStepDialog(step, row.id)
            },
            () => stepLabelMap[step]
          )
        )
      }
    }
    // 下一步：仅当当前角色能推进时显示，点击只改变状态（不弹窗）
    /*if (nextStepForOrder(row) > 0) {
      buttons.push(
        h(
          ElButton,
          {
            type: 'success',
            icon: ArrowRight,
            plain: true,
            class: 'inline-flex items-center justify-center mr-2 text-sm rounded-md align-middle',
            onClick: () => handleStepNext(row)
          },
          () => '下一步'
        )
      )
    }*/
    // 查看
    buttons.push(h(ArtButtonTable, { type: 'view', onClick: () => handleView(row) }))
    // 管理员/财务文员保留全字段编辑
    if (hasAuth('edit')) {
      buttons.push(h(ArtButtonTable, { type: 'edit', onClick: () => showDialog('edit', row) }))
    }
    // 逻辑删除（角色+状态限制）
    if (canDeleteOrder(row)) {
      buttons.push(h(ArtButtonTable, { type: 'delete', onClick: () => handleDelete(row) }))
    }
    return h('div', { class: 'flex items-center gap-1' }, buttons)
  }

  /** 打开分步弹窗 */
  const openStepDialog = (step: number, orderId: number) => {
    currentStep.value = step
    currentStepOrderId.value = orderId
    nextTick(() => {
      stepDialogVisible.value = true
    })
  }

  /** 分步保存提交 */
  const handleStepSubmit = async (formData: any) => {
    try {
      await fetchBizOrderStepEdit(formData)
      ElMessage.success('保存成功')
      stepDialogVisible.value = false
      await refreshData()
    } catch (e) {
      console.error(e)
    }
  }

  /** 下一步：仅推进订单状态（不弹窗、不保存字段） */
  const handleStepSubmitNext = async (formData: any) => {
    try {
      await fetchBizOrderStepEdit(formData)
      await fetchBizOrderStepNext(formData.id)
      ElMessage.success('保存成功')
      stepDialogVisible.value = false
      await refreshData()
    } catch (e) {
      console.error(e)
    }
  }

  const handleSearch = (params: Record<string, any>) => {
    // 先清空旧搜索值（保留分页参数），再写入新值
    const paramsRecord = searchParams as Record<string, unknown>
    Object.keys(paramsRecord).forEach((key) => {
      if (key !== 'page' && key !== 'pageSize') {
        delete paramsRecord[key]
      }
    })
    // 过滤掉空值，避免后端收到空字符串参数
    for (const [k, v] of Object.entries(params)) {
      if (v !== undefined && v !== null && v !== '') {
        paramsRecord[k] = v
      }
    }
    paramsRecord['page'] = 1 // 搜索时回到第一页
    getData()
  }

  const showDialog = (type: DialogType, row?: any) => {
    dialogType.value = type
    currentRow.value = row || {}
    nextTick(() => {
      dialogVisible.value = true
    })
  }

  const handleView = (row: any) => {
    router.push({ name: 'BizOrderDetail', params: { id: String(row.id) } })
  }

  const handleDelete = async (row: any) => {
    try {
      await ElMessageBox.confirm('确定要删除该记录吗？删除后无法恢复', '删除确认', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      await fetchBizOrderDelete(row.id)
      ElMessage.success('删除成功')
      refreshData()
    } catch (e) {
      if (e !== 'cancel') console.error(e)
    }
  }

  const handleExport = () => {
    ElMessage.info('导出功能开发中')
  }

  const handleDialogSubmit = async (formData: any) => {
    try {
      await fetchBizOrderEdit(formData)
      ElMessage.success(formData.id ? '编辑成功' : '添加成功')
      dialogVisible.value = false
      refreshData()
    } catch (e) {
      console.error(e)
    }
  }

  onMounted(() => {
    dictStore.preload(['area', 'order_status'])
  })
</script>
