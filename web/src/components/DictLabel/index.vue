<template>
  <ElTag v-if="tag && label" :type="tagType as any" :size="size as any">
    {{ label }}
  </ElTag>
  <span v-else-if="!tag && label">{{ label }}</span>
  <span v-else>{{ fallback }}</span>
</template>

<script setup lang="ts">
  import { ref, watch, onMounted } from 'vue'
  import { ElTag } from 'element-plus'
  import { useDictStore } from '@/store/modules/dict'

  const props = defineProps<{
    dictType: string
    value?: string | number
    size?: string
    fallback?: string
    tag?: boolean
  }>()

  const dictStore = useDictStore()
  const label = ref('')
  const tagType = ref('')
  const loaded = ref(false)

  const resolveDict = async () => {
    if (!props.dictType) {
      label.value = String(props.value ?? props.fallback ?? '')
      loaded.value = true
      return
    }

    if (!dictStore.isLoaded(props.dictType)) {
      await dictStore.fetchDictData(props.dictType)
    }

    if (props.value === undefined || props.value === null || props.value === '') {
      label.value = props.fallback ?? '-'
      tagType.value = ''
    } else {
      label.value = dictStore.getLabel(props.dictType, props.value)
      tagType.value = dictStore.getTagType(props.dictType, props.value)
    }
    loaded.value = true
  }

  onMounted(() => {
    resolveDict()
  })

  watch(
    () => [props.dictType, props.value],
    () => {
      resolveDict()
    }
  )
</script>
