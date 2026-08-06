<!-- 字典数据弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? '新增字典数据' : '编辑字典数据'"
    width="600px"
    align-center
  >
    <ElForm ref="formRef" :model="formState" :rules="rules" label-width="100px">
      <ElFormItem label="字典标签" prop="label">
        <ElInput v-model="formState.label" placeholder="请输入字典标签" />
      </ElFormItem>
      <ElFormItem label="字典值" prop="value">
        <ElInput v-model="formState.value" placeholder="请输入字典值" />
      </ElFormItem>
      <ElFormItem label="样式类名" prop="cssClass">
        <ElInput v-model="formState.cssClass" placeholder="请输入样式类名（自定义CSS）" />
      </ElFormItem>
      <ElFormItem label="Tag样式" prop="listClass">
        <ElSelect v-model="formState.listClass" placeholder="请选择Tag样式" style="width: 100%">
          <ElOption label="默认" value="" />
          <ElOption label="成功(success)" value="success" />
          <ElOption label="警告(warning)" value="warning" />
          <ElOption label="危险(danger)" value="danger" />
          <ElOption label="信息(info)" value="info" />
          <ElOption label="主要(primary)" value="primary" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="是否默认" prop="isDefault">
        <ElSwitch v-model="formState.isDefault" :active-value="1" :inactive-value="0" />
      </ElFormItem>
      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="formState.sort" :min="0" controls-position="right" style="width: 100%" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElSwitch v-model="formState.status" :active-value="1" :inactive-value="0" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="formState.remark" type="textarea" :rows="3" placeholder="请输入备注" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="dialog-footer">
        <ElButton @click="dialogVisible = false">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit" :loading="submitLoading">确定</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { ref, reactive, computed, watch, nextTick } from 'vue'
  import type { FormInstance, FormRules } from 'element-plus'
  import { ElMessage } from 'element-plus'
  import { fetchSaveDictData, type DictDataItem } from '@/api/backend/system/dict'

  interface Props {
    visible: boolean
    type: string
    dictTypeId: number
    formData?: Partial<DictDataItem>
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
  })

  const dialogType = computed(() => props.type)
  const submitLoading = ref(false)
  const formRef = ref<FormInstance>()

  const formState = reactive({
    id: 0,
    dictTypeId: 0,
    label: '',
    value: '',
    cssClass: '',
    listClass: '',
    isDefault: 0,
    sort: 0,
    status: 1,
    remark: ''
  })

  const rules: FormRules = {
    label: [
      { required: true, message: '请输入字典标签', trigger: 'blur' },
      { max: 100, message: '最多 100 个字符', trigger: 'blur' }
    ],
    value: [
      { required: true, message: '请输入字典值', trigger: 'blur' },
      { max: 100, message: '最多 100 个字符', trigger: 'blur' }
    ]
  }

  const initFormData = () => {
    const isEdit = props.type === 'edit' && props.formData
    const row = props.formData

    if (isEdit && row) {
      formState.id = row.id || 0
      formState.dictTypeId = row.dictTypeId || props.dictTypeId
      formState.label = row.label || ''
      formState.value = row.value || ''
      formState.cssClass = row.cssClass || ''
      formState.listClass = row.listClass || ''
      formState.isDefault = row.isDefault ?? 0
      formState.sort = row.sort || 0
      formState.status = row.status ?? 1
      formState.remark = row.remark || ''
    } else {
      formState.id = 0
      formState.dictTypeId = props.dictTypeId
      formState.label = ''
      formState.value = ''
      formState.cssClass = ''
      formState.listClass = ''
      formState.isDefault = 0
      formState.sort = 0
      formState.status = 1
      formState.remark = ''
    }
  }

  watch(
    () => [props.visible, props.type, props.formData, props.dictTypeId],
    ([visible]) => {
      if (visible) {
        initFormData()
        nextTick(() => {
          formRef.value?.clearValidate()
        })
      }
    },
    { immediate: true }
  )

  const handleSubmit = async () => {
    if (!formRef.value) return

    await formRef.value.validate(async (valid) => {
      if (!valid) return

      submitLoading.value = true
      try {
        await fetchSaveDictData({ ...formState })
        ElMessage.success(dialogType.value === 'add' ? '新增成功' : '编辑成功')
        emit('submit')
        dialogVisible.value = false
      } finally {
        submitLoading.value = false
      }
    })
  }
</script>
