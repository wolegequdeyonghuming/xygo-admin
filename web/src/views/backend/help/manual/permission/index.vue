<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 用户操作手册 · 权限管理 -->
<template>
  <ManualViewer :title="title" :description="description" :markdown="body" />
</template>

<script setup lang="ts">
  import ManualViewer from '../manual-viewer.vue'
  import markdown from './permission.md?raw'

  defineOptions({ name: 'ManualPermission' })

  const title = '用户操作手册 · 权限管理'
  const description = '适用对象：普通管理员（负责创建用户、分配角色、维护菜单/部门/岗位）'
  const body = markdown.replace(/^#\s+[^\n]+\n+/, '')
</script>
