<!-- 订单分步操作弹窗（录单/派单/预约/收单/回访/完成） -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogTitle"
    width="720px"
    :close-on-click-modal="false"
    @closed="handleClosed"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="130px">
      <ElRow :gutter="16">
        <!-- 1 录单 -->
        <template v-if="step === STEP.RECORD">
          <ElCol :span="12">
            <ElFormItem label="排单日期" prop="scheduleDate">
              <ElDatePicker
                v-model="formData.scheduleDate"
                type="date"
                value-format="YYYY-MM-DD"
                placeholder="请选择排单日期"
                style="width: 100%"
              />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="客户姓名" prop="customerName">
              <ElInput v-model="formData.customerName" placeholder="请输入客户姓名" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="可联系时间" prop="availableTimeDesc">
              <ElInput v-model="formData.availableTimeDesc" placeholder="请输入可联系时间" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="区县" prop="area">
              <ElSelect v-model="formData.area" placeholder="请选择区县" clearable>
                <ElOption
                  v-for="opt in dictStore.getDictData('area')"
                  :key="opt.value"
                  :label="opt.label"
                  :value="Number(opt.value)"
                />
              </ElSelect>
            </ElFormItem>
          </ElCol>
          <ElCol :span="24">
            <ElFormItem label="安装地址" prop="installAddress">
              <ElInput v-model="formData.installAddress" placeholder="请输入安装地址" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="联系电话" prop="contactPhone">
              <ElInput v-model="formData.contactPhone" placeholder="请输入联系电话" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="预约业务" prop="businessType">
              <ElInput v-model="formData.businessType" placeholder="请输入预约业务" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="到期时间" prop="expiryDate">
              <ElInput v-model="formData.expiryDate" placeholder="请输入到期时间" />
            </ElFormItem>
          </ElCol>
        </template>

        <!-- 2 派单 -->
        <ElCol v-if="step === STEP.ASSIGN" :span="24">
          <ElFormItem label="收单员" prop="agentId">
            <ArtPersonSelector v-model="formData.agentId" placeholder="请选择收单员" />
          </ElFormItem>
        </ElCol>

        <!-- 3 预约 -->
        <ElCol v-if="step === STEP.APPOINT" :span="24">
          <ElFormItem label="收单预约情况" prop="appointmentDesc">
            <ElInput
              v-model="formData.appointmentDesc"
              type="textarea"
              :rows="3"
              placeholder="请输入收单预约情况"
            />
          </ElFormItem>
        </ElCol>

        <!-- 4 收单 -->
        <template v-if="step === STEP.VISIT">
          <ElCol :span="12">
            <ElFormItem label="上门日期" prop="visitDate">
              <ElDatePicker
                v-model="formData.visitDate"
                type="date"
                value-format="YYYY-MM-DD"
                placeholder="请选择上门日期"
                style="width: 100%"
              />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="成交业务" prop="dealtBusinessType">
              <ElInput v-model="formData.dealtBusinessType" placeholder="请输入成交业务" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="携转情况" prop="portingStatus">
              <ElInput v-model="formData.portingStatus" placeholder="请输入携转情况" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="客户实际姓名" prop="customerRealName">
              <ElInput v-model="formData.customerRealName" placeholder="请输入客户实际姓名" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="客户身份证号" prop="customerIdNumber">
              <ElInput v-model="formData.customerIdNumber" placeholder="请输入客户身份证号" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="实缴额度" prop="paidAmount">
              <ElInputNumber
                v-model="formData.paidAmount"
                :min="0"
                :precision="2"
                controls-position="right"
                style="width: 100%"
              />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="是否乡下单" prop="isRuralOrder">
              <ElRadioGroup v-model="formData.isRuralOrder">
                <ElRadio :value="0">否</ElRadio>
                <ElRadio :value="1">是</ElRadio>
              </ElRadioGroup>
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="新开号码" prop="newPhoneNo">
              <ElInput v-model="formData.newPhoneNo" placeholder="请输入新开号码" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="终端串码" prop="deviceSerial">
              <ElInput v-model="formData.deviceSerial" placeholder="请输入终端串码" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="话补" prop="subsidyAmount">
              <ElInputNumber
                v-model="formData.subsidyAmount"
                :min="0"
                :precision="2"
                controls-position="right"
                style="width: 100%"
              />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="是否纯新增" prop="isNew">
              <ElInput v-model="formData.isNew" placeholder="请输入是否纯新增" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="24">
            <ElFormItem label="附件" prop="attachmentId">
              <ArtFileSelector v-model="formData.attachmentId" :max-number="5" />
            </ElFormItem>
          </ElCol>
        </template>

        <!-- 5 回访 -->
        <ElCol v-if="step === STEP.FOLLOW" :span="24">
          <ElFormItem label="话务二次回访情况" prop="followUpDesc">
            <ElInput
              v-model="formData.followUpDesc"
              type="textarea"
              :rows="3"
              placeholder="请输入话务二次回访情况"
            />
          </ElFormItem>
        </ElCol>

        <!-- 6 完成 -->
        <template v-if="step === STEP.COMPLETE">
          <ElCol :span="12">
            <ElFormItem label="宽带账号" prop="broadbandAccount">
              <ElInput v-model="formData.broadbandAccount" placeholder="请输入宽带账号" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="工号" prop="agencyNo">
              <ElInput v-model="formData.agencyNo" placeholder="请输入工号" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="12">
            <ElFormItem label="是否完工" prop="isCompleted">
              <ElRadioGroup v-model="formData.isCompleted">
                <ElRadio :value="0">未完工</ElRadio>
                <ElRadio :value="1">已完工</ElRadio>
              </ElRadioGroup>
            </ElFormItem>
          </ElCol>
        </template>
      </ElRow>
    </ElForm>

    <!-- 详细情况（折叠模式：仅显示标题 + 最新一条 + 输入框） -->
    <div v-if="props.orderId > 0" class="order-comment-wrap">
      <OrderComment :order-id="props.orderId" />
    </div>

    <template #footer>
      <ElButton @click="dialogVisible = false">取消</ElButton>
      <ElButton v-if="step === 4" type="warning" :loading="loading" @click="handleSave">
        暂存
      </ElButton>
      <ElButton type="primary" :loading="loading" @click="handleSaveNext">保存并下一步</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import ArtFileSelector from '@/components/core/forms/art-file-selector/index.vue'
  import ArtPersonSelector from '@/components/core/forms/art-person-selector/index.vue'
  import OrderComment from '../components/order-comment.vue'
  import { useDictStore } from '@/store/modules/dict'
  import { fetchBizOrderView } from '@/api/order/biz-order'
  import { ORDER_STEP as STEP } from '../useOrderPerm'

  const dictStore = useDictStore()

  const props = defineProps<{
    visible: boolean
    /** 当前步骤：1 录单 2 派单 3 预约 4 收单 5 回访 6 完成 */
    step: number
    /** 订单 ID（0 表示新增录单） */
    orderId: number
  }>()

  const emit = defineEmits<{
    (e: 'update:visible', v: boolean): void
    (e: 'submit', data: Record<string, any>): void
    (e: 'submitNext', data: Record<string, any>): void
  }>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (val: boolean) => emit('update:visible', val)
  })

  /** 是否编辑已有订单（非新增录单） */
  const isEdit = computed(() => props.orderId > 0)

  const stepTitleMap: Record<number, string> = {
    [STEP.RECORD]: '录单',
    [STEP.ASSIGN]: '派单',
    [STEP.APPOINT]: '预约',
    [STEP.VISIT]: '收单',
    [STEP.FOLLOW]: '回访',
    [STEP.COMPLETE]: '完成'
  }
  const dialogTitle = computed(
    () => `${isEdit.value ? '编辑' : '新增'} - ${stepTitleMap[props.step] || ''}`
  )

  const formRef = ref<FormInstance>()
  const loading = ref(false)

  const defaultForm = (): Record<string, any> => ({
    scheduleDate: '',
    customerName: '',
    availableTimeDesc: '',
    area: undefined,
    installAddress: '',
    contactPhone: '',
    businessType: '',
    expiryDate: '',
    agentId: undefined,
    appointmentDesc: '',
    visitDate: '',
    dealtBusinessType: '',
    portingStatus: '',
    customerRealName: '',
    customerIdNumber: '',
    paidAmount: 0,
    isRuralOrder: 0,
    newPhoneNo: '',
    deviceSerial: '',
    attachmentId: '',
    subsidyAmount: 0,
    isNew: '',
    followUpDesc: '',
    broadbandAccount: '',
    isCompleted: 0,
    agencyNo: ''
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    customerName: [{ required: true, message: '客户姓名不能为空', trigger: 'blur' }],
    contactPhone: [{ required: true, message: '联系电话不能为空', trigger: 'blur' }],
    agentId: [{ required: true, message: '请选择收单员', trigger: 'change' }]
  })

  const resetForm = () => {
    formRef.value?.resetFields()
    Object.assign(formData, defaultForm())
  }

  const assignFormData = (data: Record<string, any>) => {
    const merged: Record<string, any> = { ...data }
    for (const key of ['scheduleDate', 'visitDate']) {
      if (merged[key] && typeof merged[key] === 'string') {
        merged[key] = merged[key].slice(0, 10)
      }
    }
    Object.assign(formData, merged)
  }

  watch(
    () => props.visible,
    async (val) => {
      if (val) {
        resetForm()
        await dictStore.preload(['area'])
        // 编辑已有订单时回显相关字段
        if (props.orderId > 0) {
          try {
            const detail = await fetchBizOrderView(props.orderId)
            assignFormData(detail)
          } catch {
            /* 忽略回显失败 */
          }
        }
      }
    }
  )

  const handleSave = async () => {
    if (!formRef.value) return
    await formRef.value.validate()
    emit('submit', { id: props.orderId, step: props.step, ...formData })
  }

  const handleSaveNext = async () => {
    if (!formRef.value) return
    await formRef.value.validate()
    emit('submitNext', { id: props.orderId, step: props.step, ...formData })
  }

  const handleClosed = () => {
    resetForm()
  }
</script>

<style scoped>
  .order-comment-wrap {
    margin-top: 8px;
    padding-top: 12px;
    border-top: 1px solid var(--el-border-color-lighter);
  }
</style>
