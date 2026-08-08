<!-- 订单主表 编辑弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="type === 'add' ? '新增订单主表' : '编辑订单主表'"
    width="600px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <!-- 主键隐藏，不在表单中显示 -->
      <ElFormItem label="订单状态" prop="orderStatus">
        <ElSelect v-model="formData.orderStatus" placeholder="请选择订单状态" clearable>
          <ElOption v-for="opt in dictStore.getDictData('order_status')" :key="opt.value" :label="opt.label" :value="Number(opt.value)" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="排单日期" prop="scheduleDate">
        <ElDatePicker v-model="formData.scheduleDate" type="date" value-format="X" placeholder="请选择排单日期" />
      </ElFormItem>
      <ElFormItem label="上门日期" prop="visitDate">
        <ElDatePicker v-model="formData.visitDate" type="date" value-format="X" placeholder="请选择上门日期" />
      </ElFormItem>
      <ElFormItem label="客户姓名" prop="customerName">
        <ElInput v-model="formData.customerName" placeholder="请输入客户姓名" />
      </ElFormItem>
      <ElFormItem label="可联系时间" prop="availableTimeDesc">
        <ElInput v-model="formData.availableTimeDesc" placeholder="请输入可联系时间" />
      </ElFormItem>
      <ElFormItem label="区县" prop="area">
        <ElSelect v-model="formData.area" placeholder="请选择区县" clearable>
          <ElOption v-for="opt in dictStore.getDictData('area')" :key="opt.value" :label="opt.label" :value="Number(opt.value)" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="安装地址" prop="installAddress">
        <ElInput v-model="formData.installAddress" placeholder="请输入安装地址" />
      </ElFormItem>
      <ElFormItem label="联系电话" prop="contactPhone">
        <ElInput v-model="formData.contactPhone" placeholder="请输入联系电话" />
      </ElFormItem>
      <ElFormItem label="预约业务" prop="businessType">
        <ElInput v-model="formData.businessType" placeholder="请输入预约业务" />
      </ElFormItem>
      <ElFormItem label="到期时间" prop="expiryDate">
        <ElInput v-model="formData.expiryDate" placeholder="请输入到期时间" />
      </ElFormItem>
      <ElFormItem label="主卡号码" prop="primaryTelNo">
        <ElInput v-model="formData.primaryTelNo" placeholder="请输入主卡号码" />
      </ElFormItem>
      <ElFormItem label="详细情况" prop="details">
        <ElInput v-model="formData.details" placeholder="请输入详细情况" />
      </ElFormItem>
      <ElFormItem label="话务员" prop="telemarketerId">
        <ElSelect
          v-model="formData.telemarketerId"
          filterable
          remote
          :remote-method="(q: string) => loadTelemarketerOptions(q)"
          placeholder="请选择话务员"
          clearable
          :loading="telemarketerLoading"
        >
          <ElOption
            v-for="opt in telemarketerOptions"
            :key="opt.value"
            :label="opt.label"
            :value="opt.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="收单员" prop="agentId">
        <ElSelect
          v-model="formData.agentId"
          filterable
          remote
          :remote-method="(q: string) => loadAgentOptions(q)"
          placeholder="请选择收单员"
          clearable
          :loading="agentLoading"
        >
          <ElOption
            v-for="opt in agentOptions"
            :key="opt.value"
            :label="opt.label"
            :value="opt.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="收单预约情况" prop="appointmentDesc">
        <ElInput v-model="formData.appointmentDesc" placeholder="请输入收单预约情况" />
      </ElFormItem>
      <ElFormItem label="话务二次回放情况" prop="followUpDesc">
        <ElInput v-model="formData.followUpDesc" placeholder="请输入话务二次回放情况" />
      </ElFormItem>
      <ElFormItem label="成交业务" prop="dealtBusinessType">
        <ElInput v-model="formData.dealtBusinessType" placeholder="请输入成交业务" />
      </ElFormItem>
      <ElFormItem label="携转情况" prop="portingStatus">
        <ElInput v-model="formData.portingStatus" placeholder="请输入携转情况" />
      </ElFormItem>
      <ElFormItem label="客户实际姓名" prop="customerRealName">
        <ElInput v-model="formData.customerRealName" placeholder="请输入客户实际姓名" />
      </ElFormItem>
      <ElFormItem label="客户身份证号" prop="customerIdNumber">
        <ElInputNumber v-model="formData.customerIdNumber" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="实缴额度" prop="paidAmount">
        <ElInputNumber v-model="formData.paidAmount" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="是否乡下单" prop="isRuralOrder">
        <ElRadioGroup v-model="formData.isRuralOrder">
          <ElRadio :value="0">否</ElRadio>
          <ElRadio :value="1">是</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
      <ElFormItem label="新开号码" prop="newPhoneNo">
        <ElInput v-model="formData.newPhoneNo" placeholder="请输入新开号码" />
      </ElFormItem>
      <ElFormItem label="终端串码" prop="deviceSerial">
        <ElInput v-model="formData.deviceSerial" placeholder="请输入终端串码" />
      </ElFormItem>
      <ElFormItem label="宽带账号" prop="broadbandAccount">
        <ElInput v-model="formData.broadbandAccount" placeholder="请输入宽带账号" />
      </ElFormItem>
      <ElFormItem label="是否完工" prop="isCompleted">
        <ElRadioGroup v-model="formData.isCompleted">
          <ElRadio :value="0">未完工</ElRadio>
          <ElRadio :value="1">已完工</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
      <ElFormItem label="话补" prop="subsidyAmount">
        <ElInputNumber v-model="formData.subsidyAmount" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="工号" prop="agencyNo">
        <ElInput v-model="formData.agencyNo" placeholder="请输入工号" />
      </ElFormItem>
      <ElFormItem label="是否纯新增" prop="isNew">
        <ElInput v-model="formData.isNew" placeholder="请输入是否纯新增" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
      </ElFormItem>
      <ElFormItem label="附件" prop="attachmentId">
        <ArtFileSelector v-model="formData.attachmentId" :max-number="5" />
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton type="primary" :loading="loading" @click="handleSubmit">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import type { DialogType } from '@/types'
  import { adminRequest } from '@/utils/http'
  import ArtFileSelector from '@/components/core/forms/art-file-selector/index.vue'
  import ArtIconSelector from '@/components/core/forms/art-icon-selector/index.vue'
  import ArtWangEditor from '@/components/core/forms/art-wang-editor/index.vue'
  import { useDictStore } from '@/store/modules/dict'

  const dictStore = useDictStore()

  const props = defineProps<{
    visible: boolean
    type: DialogType
    editData?: Record<string, any>
  }>()

  const emit = defineEmits<{
    (e: 'update:visible', v: boolean): void
    (e: 'submit', data: Record<string, any>): void
  }>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (val: boolean) => emit('update:visible', val)
  })

  const formRef = ref<FormInstance>()
  const loading = ref(false)

  const defaultForm = (): Record<string, any> => ({
    id: 0,
    orderStatus: 0,
    scheduleDate: '',
    visitDate: '',
    customerName: '',
    availableTimeDesc: '',
    area: 0,
    installAddress: '',
    contactPhone: '',
    businessType: '',
    expiryDate: '',
    primaryTelNo: '',
    details: '',
    telemarketerId: undefined,
    agentId: undefined,
    appointmentDesc: '',
    followUpDesc: '',
    dealtBusinessType: '',
    portingStatus: '',
    customerRealName: '',
    customerIdNumber: 0,
    paidAmount: 0,
    isRuralOrder: 0,
    newPhoneNo: '',
    deviceSerial: '',
    broadbandAccount: '',
    isCompleted: 0,
    subsidyAmount: 0,
    agencyNo: '',
    isNew: '',
    remark: '',
    attachmentId: '',
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    customerName: [{ required: true, message: '客户姓名不能为空', trigger: 'blur' }],
    contactPhone: [{ required: true, message: '联系电话不能为空', trigger: 'blur' }],
  })

  // ==================== 远程下拉选项 ====================
  const telemarketerOptions = ref<{ value: any; label: string }[]>([])
  const telemarketerLoading = ref(false)
  const loadTelemarketerOptions = async (query: string) => {
    telemarketerLoading.value = true
    try {
      const res = await adminRequest.get<any>({
        url: '/user/list',
        params: { pageSize: 50, real_name: query || undefined }
      })
      telemarketerOptions.value = (res.list || []).map((item: any) => ({
        value: item.id,
        label: item.real_name,
      }))
    } catch { /* ignore */ }
    telemarketerLoading.value = false
  }
  const agentOptions = ref<{ value: any; label: string }[]>([])
  const agentLoading = ref(false)
  const loadAgentOptions = async (query: string) => {
    agentLoading.value = true
    try {
      const res = await adminRequest.get<any>({
        url: '/user/list',
        params: { pageSize: 50, real_name: query || undefined }
      })
      agentOptions.value = (res.list || []).map((item: any) => ({
        value: item.id,
        label: item.real_name,
      }))
    } catch { /* ignore */ }
    agentLoading.value = false
  }

  watch(() => props.visible, async (val) => {
    if (val) {
      // 预加载字典数据
      await dictStore.preload(['area', 'order_status'])
      if (props.type === 'edit' && props.editData) {
        Object.assign(formData, props.editData)
        // 编辑时加载已选关联项
        loadTelemarketerOptions('')
        loadAgentOptions('')
      } else {
        Object.assign(formData, defaultForm())
        loadTelemarketerOptions('')
        loadAgentOptions('')
      }
    }
  })

  const handleSubmit = async () => {
    if (!formRef.value) return
    await formRef.value.validate()
    emit('submit', { ...formData })
  }

  const handleClose = () => {
    formRef.value?.resetFields()
    Object.assign(formData, defaultForm())
    dialogVisible.value = false
  }
</script>
