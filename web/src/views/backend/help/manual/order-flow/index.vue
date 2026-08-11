<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 用户操作手册 · 订单全流程处理 -->
<template>
  <ManualViewer :title="title" :description="description" :markdown="body" />
</template>

<script setup lang="ts">
  import ManualViewer from '../manual-viewer.vue'
  import markdown from './order-flow.md?raw'

  defineOptions({ name: 'ManualOrderFlow' })

  const title = '用户操作手册 · 订单全流程处理'
  const description =
    '适用对象：使用订单业务的日常用户（话务员、收单员、收单员管理员、财务文员、管理员）'
  const body = markdown.replace(/^#\s+[^\n]+\n+/, '')
</script>
