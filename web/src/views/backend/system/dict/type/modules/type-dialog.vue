<!-- 字典类型弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? '新增字典类型' : '编辑字典类型'"
    width="500px"
    align-center
  >
    <ElForm ref="formRef" :model="formState" :rules="rules" label-width="90px">
      <ElFormItem label="字典名称" prop="name">
        <ElInput v-model="formState.name" placeholder="请输入字典名称" />
      </ElFormItem>
      <ElFormItem label="字典标识" prop="type">
        <ElInput v-model="formState.type" placeholder="请输入字典标识（英文，如 gender）" />
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
  import { fetchSaveDictType, type DictTypeItem } from '@/api/backend/system/dict'

  interface Props {
    visible: boolean
    type: string
    formData?: Partial<DictTypeItem>
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
    name: '',
    type: '',
    sort: 0,
    status: 1,
    remark: ''
  })

  const rules: FormRules = {
    name: [
      { required: true, message: '请输入字典名称', trigger: 'blur' },
      { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
    ],
    type: [
      { required: true, message: '请输入字典标识', trigger: 'blur' },
      {
        pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/,
        message: '只能以字母开头，包含字母、数字、下划线',
        trigger: 'blur'
      }
    ]
  }

  const initFormData = () => {
    const isEdit = props.type === 'edit' && props.formData
    const row = props.formData

    if (isEdit && row) {
      formState.id = row.id || 0
      formState.name = row.name || ''
      formState.type = row.type || ''
      formState.sort = row.sort || 0
      formState.status = row.status ?? 1
      formState.remark = row.remark || ''
    } else {
      formState.id = 0
      formState.name = ''
      formState.type = ''
      formState.sort = 0
      formState.status = 1
      formState.remark = ''
    }
  }

  watch(
    () => [props.visible, props.type, props.formData],
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
        await fetchSaveDictType({ ...formState })
        ElMessage.success(dialogType.value === 'add' ? '新增成功' : '编辑成功')
        emit('submit')
        dialogVisible.value = false
      } finally {
        submitLoading.value = false
      }
    })
  }
</script>
