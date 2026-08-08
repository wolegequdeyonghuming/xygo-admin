<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- ArtPersonSelector 人员选择器
  以部门树形式选择 admin_user（排除超管），叶子为用户。
  过滤规则：同类型数组内为并集，不同类型之间为交集。 -->
<template>
  <ElCascader
    v-model="currentValue"
    :options="tree"
    :props="cascaderProps"
    :multiple="multiple"
    :clearable="clearable"
    :placeholder="placeholder"
    :loading="loading"
    :show-all-levels="false"
    filterable
    style="width: 100%"
  />
</template>

<script setup lang="ts">
  import { ref, computed, watch } from 'vue'
  import {
    fetchPersonSelector,
    type PersonSelectorNode
  } from '@/api/backend/system/user'

  interface Props {
    modelValue?: number | number[] | null
    deptIds?: number[]
    roleIds?: number[]
    postIds?: number[]
    multiple?: boolean
    clearable?: boolean
    placeholder?: string
  }

  interface Emits {
    (e: 'update:modelValue', value: number | number[] | null): void
  }

  const props = withDefaults(defineProps<Props>(), {
    modelValue: null,
    deptIds: undefined,
    roleIds: undefined,
    postIds: undefined,
    multiple: false,
    clearable: true,
    placeholder: '请选择人员'
  })

  const emit = defineEmits<Emits>()

  const tree = ref<PersonSelectorNode[]>([])
  const loading = ref(false)

  // 级联配置：叶子为用户(value=userId)，部门节点不可选，值不包含父节点
  const cascaderProps = computed(() => ({
    value: 'value',
    label: 'label',
    children: 'children',
    emitPath: false,
    checkStrictly: false,
    multiple: props.multiple
  }))

  const currentValue = computed({
    get: () => props.modelValue as any,
    set: (val: any) => emit('update:modelValue', val)
  })

  const loadTree = async () => {
    loading.value = true
    try {
      const res = await fetchPersonSelector({
        deptIds: props.deptIds,
        roleIds: props.roleIds,
        postIds: props.postIds
      })
      tree.value = res.list || []
    } catch {
      tree.value = []
    } finally {
      loading.value = false
    }
  }

  // 过滤条件变化时重新加载
  watch(
    () => [props.deptIds, props.roleIds, props.postIds],
    () => {
      loadTree()
    },
    { immediate: true, deep: true }
  )
</script>
