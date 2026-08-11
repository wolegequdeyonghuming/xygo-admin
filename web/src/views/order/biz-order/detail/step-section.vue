<!-- 详情页步骤分区卡片 -->
<template>
  <ElCard shadow="never" class="step-section">
    <template #header>
      <span class="font-semibold">{{ title }}</span>
    </template>
    <slot />
  </ElCard>
</template>

<script setup lang="ts">
  defineOptions({ name: 'StepSection' })

  defineProps<{
    title: string
  }>()
</script>

<style scoped>
  .step-section {
    margin-top: 16px;
    :deep(.el-card__body) {
      padding: 16px;
    }
    :deep(.el-card__header) {
      padding: 12px 16px;
    }
  }
</style>
