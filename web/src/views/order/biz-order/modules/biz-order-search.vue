<!-- 订单主表 搜索栏 -->
<template>
  <ArtSearchBar
    v-model="formFilters"
    :items="formItems"
    @reset="handleReset"
    @search="handleSearch"
  />
</template>

<script setup lang="ts">
  import { useDictStore } from '@/store/modules/dict'

  const dictStore = useDictStore()

  const props = defineProps<{ modelValue: Record<string, any> }>()
  const emit = defineEmits<{
    (e: 'update:modelValue', v: Record<string, any>): void
    (e: 'search', params: Record<string, any>): void
    (e: 'reset'): void
  }>()

  const formFilters = reactive({ ...props.modelValue })

  const formItems = computed(() => [
    {
      label: '订单状态',
      key: 'orderStatus',
      type: 'select',
      props: {
        clearable: true,
        options: dictStore.getDictData('order_status').map((i: any) => ({ label: i.label, value: i.value }))
      }
    },
    {
      label: '排单日期',
      key: 'scheduleDateRange',
      type: 'daterange',
      props: { type: 'daterange', clearable: true, valueFormat: 'YYYY-MM-DD' }
    },
    {
      label: '上门日期',
      key: 'visitDateRange',
      type: 'daterange',
      props: { type: 'daterange', clearable: true, valueFormat: 'YYYY-MM-DD' }
    },
    {
      label: '客户姓名',
      key: 'customerName',
      type: 'input',
      props: { clearable: true }
    },
    {
      label: '区县',
      key: 'area',
      type: 'select',
      props: {
        clearable: true,
        options: dictStore.getDictData('area').map((i: any) => ({ label: i.label, value: i.value }))
      }
    },
    {
      label: '安装地址',
      key: 'installAddress',
      type: 'input',
      props: { clearable: true }
    },
    {
      label: '联系电话',
      key: 'contactPhone',
      type: 'input',
      props: { clearable: true }
    },
    {
      label: '预约业务',
      key: 'businessType',
      type: 'input',
      props: { clearable: true }
    },
  ])

  const handleSearch = () => {
    const params: Record<string, any> = { ...formFilters }
    {
      const range = formFilters.scheduleDateRange
      params.scheduleDateStart = Array.isArray(range) && range[0] ? Math.floor(new Date(range[0]).getTime() / 1000) : undefined
      params.scheduleDateEnd = Array.isArray(range) && range[1] ? Math.floor(new Date(range[1] + (range[1].length <= 10 ? ' 23:59:59' : '')).getTime() / 1000) : undefined
      delete params.scheduleDateRange
    }
    {
      const range = formFilters.visitDateRange
      params.visitDateStart = Array.isArray(range) && range[0] ? Math.floor(new Date(range[0]).getTime() / 1000) : undefined
      params.visitDateEnd = Array.isArray(range) && range[1] ? Math.floor(new Date(range[1] + (range[1].length <= 10 ? ' 23:59:59' : '')).getTime() / 1000) : undefined
      delete params.visitDateRange
    }
    emit('update:modelValue', params)
    emit('search', params)
  }

  const handleReset = () => {
    Object.keys(formFilters).forEach(k => (formFilters[k] = undefined))
    emit('update:modelValue', { ...formFilters })
    emit('reset')
  }
</script>
<style lang="scss">
.art-number-range {
  display: inline-flex;
  align-items: center;
  width: 100%;
  height: 32px;
  padding: 0 8px;
  background-color: var(--el-fill-color-blank);
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-base);
  transition: border-color 0.2s;
  &:hover { border-color: var(--el-border-color-hover); }
  &:focus-within { border-color: var(--el-color-primary); }

  &__input {
    flex: 1;
    min-width: 0;
    height: 100%;
    padding: 0 4px;
    font-size: 13px;
    color: var(--el-text-color-regular);
    text-align: center;
    background: transparent;
    border: none;
    outline: none;
    appearance: textfield;
    -moz-appearance: textfield;
    &::-webkit-inner-spin-button,
    &::-webkit-outer-spin-button { appearance: none; margin: 0; }
    &::placeholder { color: var(--el-text-color-placeholder); }
  }

  &__separator {
    flex-shrink: 0;
    padding: 0 6px;
    font-size: 13px;
    color: var(--el-text-color-placeholder);
  }
}
</style>
