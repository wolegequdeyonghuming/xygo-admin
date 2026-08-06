<!-- {{.TableComment}} 搜索栏 -->
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
{{- range .QueryColumns}}
{{- if eq .QueryType "between"}}
{{- if eq .DesignType "date"}}
    {
      label: '{{.Label}}',
      key: '{{.TsName}}Range',
      type: 'daterange',
      props: { type: 'daterange', clearable: true, valueFormat: 'YYYY-MM-DD' }
    },
{{- else if eq .DesignType "time"}}
    {
      label: '{{.Label}}',
      key: '{{.TsName}}Range',
      type: 'timepicker',
      props: { isRange: true, clearable: true, valueFormat: 'HH:mm:ss' }
    },
{{- else if or (eq .DesignType "datetime") (eq .DesignType "timestamp") (eq .FormType "datetime")}}
    {
      label: '{{.Label}}',
      key: '{{.TsName}}Range',
      type: 'datetimerange',
      props: { type: 'datetimerange', clearable: true, valueFormat: 'YYYY-MM-DD HH:mm:ss' }
    },
{{- else if or (eq .DesignType "number") (eq .DesignType "float") (eq .DesignType "weigh") (eq .FormType "inputNumber")}}
    {
      label: '{{.Label}}',
      key: '{{.TsName}}Start',
      span: 8,
      render: () => h('div', { class: 'art-number-range' }, [
        h('input', { type: 'number', value: formFilters.{{.TsName}}Start ?? '', placeholder: '最小值', class: 'art-number-range__input', onInput: (e: any) => formFilters.{{.TsName}}Start = e.target.value === '' ? undefined : Number(e.target.value) }),
        h('span', { class: 'art-number-range__separator' }, '至'),
        h('input', { type: 'number', value: formFilters.{{.TsName}}End ?? '', placeholder: '最大值', class: 'art-number-range__input', onInput: (e: any) => formFilters.{{.TsName}}End = e.target.value === '' ? undefined : Number(e.target.value) }),
      ])
    },
{{- else}}
    {
      label: '{{.Label}}开始',
      key: '{{.TsName}}Start',
      type: 'input',
      props: { clearable: true }
    },
    {
      label: '{{.Label}}结束',
      key: '{{.TsName}}End',
      type: 'input',
      props: { clearable: true }
    },
{{- end}}
{{- else}}
    {
      label: '{{.Label}}',
      key: '{{.TsName}}',
{{- if and (or (eq .DesignType "switch") (eq .DesignType "radio") (eq .DesignType "select") (eq .FormType "select") (eq .FormType "radio")) .DictType}}
      type: 'select',
      props: {
        clearable: true,
        options: dictStore.getDictData('{{.DictType}}').map((i: any) => ({ label: i.label, value: i.value }))
      }
{{- else if and (or (eq .DesignType "switch") (eq .DesignType "radio") (eq .DesignType "select") (eq .FormType "select") (eq .FormType "radio")) .HasOptions}}
      type: 'select',
      props: {
        clearable: true,
        options: [{{range .Options}}{ label: '{{.Label}}', value: {{jsValue .Value}} }, {{end}}]
      }
{{- else if or (eq .DesignType "radio") (eq .DesignType "select") (eq .FormType "select") (eq .FormType "radio")}}
{{- if eq .TsType "number"}}
      type: 'number',
      props: { clearable: true }
{{- else}}
      type: 'input',
      props: { clearable: true }
{{- end}}
{{- else if eq .DesignType "switch"}}
{{- if eq .TsType "number"}}
      type: 'number',
      props: { clearable: true }
{{- else}}
      type: 'input',
      props: { clearable: true }
{{- end}}
{{- else if or (eq .DesignType "date") (eq .DesignType "datetime") (eq .DesignType "timestamp") (eq .FormType "date") (eq .FormType "datetime")}}
{{- if or (eq .DesignType "datetime") (eq .DesignType "timestamp") (eq .FormType "datetime")}}
      type: 'datetime',
      props: { type: 'datetime', clearable: true, valueFormat: 'YYYY-MM-DD HH:mm:ss' }
{{- else}}
      type: 'date',
      props: { type: 'date', clearable: true, valueFormat: 'YYYY-MM-DD' }
{{- end}}
{{- else if eq .DesignType "time"}}
      type: 'timepicker',
      props: { clearable: true, valueFormat: 'HH:mm:ss' }
{{- else if or (eq .DesignType "number") (eq .DesignType "float") (eq .DesignType "weigh") (eq .FormType "inputNumber")}}
      type: 'number',
      props: { clearable: true }
{{- else}}
      type: 'input',
      props: { clearable: true }
{{- end}}
    },
{{- end}}
{{- end}}
{{- if .HasRelations}}
{{- range $rel := .Relations}}
{{- if $rel.FieldConfigs}}
{{- range $fc := $rel.FieldConfigs}}
{{- if $fc.InSearch}}
{{- if or (eq $fc.SearchComponent "daterange") (eq $fc.SearchComponent "datetimerange")}}
    {
      label: '{{$fc.Label}}',
      key: '{{$rel.RelationAlias}}_{{$fc.Field}}Range',
      type: '{{$fc.SearchComponent}}',
      props: { type: '{{$fc.SearchComponent}}', clearable: true{{if eq $fc.SearchComponent "daterange"}}, valueFormat: 'YYYY-MM-DD'{{else}}, valueFormat: 'YYYY-MM-DD HH:mm:ss'{{end}} }
    },
{{- else}}
    {
      label: '{{$fc.Label}}',
      key: '{{$rel.RelationAlias}}_{{$fc.Field}}',
{{- if eq $fc.SearchComponent "select"}}
      type: 'select',
      props: { clearable: true, options: [] }
{{- else if eq $fc.SearchComponent "number"}}
      type: 'number',
      props: { clearable: true }
{{- else if eq $fc.SearchComponent "switch"}}
      type: 'switch',
      props: {}
{{- else if eq $fc.SearchComponent "date"}}
      type: 'date',
      props: { type: 'date', clearable: true, valueFormat: 'YYYY-MM-DD' }
{{- else if eq $fc.SearchComponent "datetime"}}
      type: 'datetime',
      props: { type: 'datetime', clearable: true, valueFormat: 'YYYY-MM-DD HH:mm:ss' }
{{- else if eq $fc.SearchComponent "timepicker"}}
      type: 'timepicker',
      props: { clearable: true, valueFormat: 'HH:mm:ss' }
{{- else if eq $fc.SearchComponent "inputTag"}}
      type: 'inputTag',
      props: { clearable: true }
{{- else if eq $fc.SearchComponent "cascader"}}
      type: 'cascader',
      props: { clearable: true, options: [] }
{{- else if eq $fc.SearchComponent "treeselect"}}
      type: 'treeselect',
      props: { clearable: true, data: [] }
{{- else}}
      type: 'input',
      props: { clearable: true }
{{- end}}
    },
{{- end}}
{{- end}}
{{- end}}
{{- else}}
{{- range $f := $rel.SearchFields}}
    {
      label: '{{$rel.RelationName}}{{$f}}',
      key: '{{$rel.RelationAlias}}_{{$f}}',
      type: 'input',
      props: { clearable: true }
    },
{{- end}}
{{- end}}
{{- end}}
{{- end}}
  ])

  const handleSearch = () => {
    const params: Record<string, any> = { ...formFilters }
{{- range .QueryColumns}}
{{- if eq .QueryType "between"}}
{{- if or (eq .DesignType "date") (eq .DesignType "datetime") (eq .DesignType "timestamp") (eq .DesignType "time") (eq .FormType "date") (eq .FormType "datetime")}}
    {
      const range = formFilters.{{.TsName}}Range
      params.{{.TsName}}Start = Array.isArray(range) && range[0] ? Math.floor(new Date(range[0]).getTime() / 1000) : undefined
      params.{{.TsName}}End = Array.isArray(range) && range[1] ? Math.floor(new Date(range[1] + (range[1].length <= 10 ? ' 23:59:59' : '')).getTime() / 1000) : undefined
      delete params.{{.TsName}}Range
    }
{{- end}}
{{- end}}
{{- end}}
{{- if .HasRelations}}
{{- range $rel := .Relations}}
{{- if $rel.FieldConfigs}}
{{- range $fc := $rel.FieldConfigs}}
{{- if $fc.InSearch}}
{{- if or (eq $fc.SearchComponent "daterange") (eq $fc.SearchComponent "datetimerange")}}
    {
      const range = formFilters.{{$rel.RelationAlias}}_{{$fc.Field}}Range
      params.{{$rel.RelationAlias}}_{{$fc.Field}}Start = Array.isArray(range) && range[0] ? Math.floor(new Date(range[0]).getTime() / 1000) : undefined
      params.{{$rel.RelationAlias}}_{{$fc.Field}}End = Array.isArray(range) && range[1] ? Math.floor(new Date(range[1] + (range[1].length <= 10 ? ' 23:59:59' : '')).getTime() / 1000) : undefined
      delete params.{{$rel.RelationAlias}}_{{$fc.Field}}Range
    }
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
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
