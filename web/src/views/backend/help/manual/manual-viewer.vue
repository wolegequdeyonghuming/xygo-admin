<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 通用用户手册查看组件：渲染 Markdown 正文 + 右侧章节目录 -->
<template>
  <main class="manual-page px-4 pb-4 pt-4">
    <div class="max-w-[1400px] mx-auto flex gap-4 lg:gap-6 items-start">
      <!-- 正文 -->
      <article class="flex-1 min-w-0">
        <div
          class="bg-white/60 backdrop-blur-lg rounded-[24px] shadow-clay-card border border-[#d1d9e6]/40 p-8 lg:p-10"
        >
          <header class="mb-8 pb-6 border-b border-[#d1d9e6]/40">
            <h1
              class="font-heading font-black text-2xl lg:text-3xl text-clay-foreground mb-3 leading-tight"
            >
              {{ title }}
            </h1>
            <p v-if="description" class="mt-2 text-sm text-clay-muted leading-relaxed">{{
              description
            }}</p>
          </header>

          <div class="doc-markdown-body">
            <MdPreview
              :model-value="markdown"
              :editor-id="previewId"
              preview-theme="github"
              code-theme="atom"
              :noMermaid="true"
              :noKatex="true"
            />
          </div>
        </div>
      </article>

      <!-- 右侧章节目录 -->
      <aside
        class="hidden xl:flex flex-col w-48 flex-shrink-0 sticky top-[88px] self-start max-h-[calc(100vh-100px)] overflow-y-auto sidebar-scroll"
      >
        <div
          class="bg-white/60 backdrop-blur-lg rounded-[24px] shadow-clay-card border border-[#d1d9e6]/40 p-4"
        >
          <h4 class="text-[10px] font-black text-clay-muted uppercase tracking-[0.1em] mb-3 px-2"
            >本章目录</h4
          >
          <MdCatalog :editor-id="previewId" :scroll-element="scrollElement" class="doc-toc" />
        </div>
      </aside>
    </div>
  </main>
</template>

<script setup lang="ts">
  import { MdPreview, MdCatalog } from 'md-editor-v3'
  import 'md-editor-v3/lib/preview.css'

  defineOptions({ name: 'ManualViewer' })

  withDefaults(
    defineProps<{
      title: string
      markdown: string
      description?: string
    }>(),
    {
      description: ''
    }
  )

  const previewId = 'manual-preview'
  const scrollElement = document.documentElement
</script>

<style lang="scss" scoped>
  .text-clay-foreground {
    color: #32325d;
  }
  .text-clay-muted {
    color: #8898aa;
  }
  .font-heading {
    font-family: 'Nunito', 'PingFang SC', sans-serif;
  }

  .shadow-clay-card {
    box-shadow:
      16px 16px 32px rgba(165, 175, 190, 0.3),
      -10px -10px 24px rgba(255, 255, 255, 0.9),
      inset 6px 6px 12px rgba(90, 141, 238, 0.03),
      inset -6px -6px 12px rgba(255, 255, 255, 1);
  }

  .sidebar-scroll {
    scrollbar-width: thin;
    scrollbar-color: transparent transparent;
    transition: scrollbar-color 0.3s;

    &:hover {
      scrollbar-color: rgba(165, 175, 190, 0.35) transparent;
    }
    &::-webkit-scrollbar {
      width: 4px;
    }
    &::-webkit-scrollbar-track {
      background: transparent;
    }
    &::-webkit-scrollbar-thumb {
      background: transparent;
      border-radius: 99px;
    }
    &:hover::-webkit-scrollbar-thumb {
      background: rgba(165, 175, 190, 0.4);
    }
    &::-webkit-scrollbar-thumb:hover {
      background: rgba(130, 145, 165, 0.5);
    }
  }

  /* ===== Markdown 内容样式 ===== */
  .doc-markdown-body {
    max-width: 100%;
    overflow: hidden;
    word-break: break-word;

    :deep(.md-editor) {
      background: transparent;
    }
    :deep(.md-editor-preview-wrapper) {
      padding: 0;
      max-width: 100%;
      overflow: hidden;
    }
    :deep(.md-editor-preview) {
      font-size: 15px;
      line-height: 1.8;
      color: #32325d;
      max-width: 100%;
      overflow-wrap: break-word;
      word-break: break-word;

      h1,
      h2,
      h3,
      h4,
      h5,
      h6 {
        font-family: 'Nunito', 'PingFang SC', sans-serif;
        font-weight: 800;
        color: #32325d;
        margin-top: 1.5em;
        margin-bottom: 0.5em;
      }
      h1 {
        font-size: 1.75em;
      }
      h2 {
        font-size: 1.4em;
        padding-bottom: 0.3em;
        border-bottom: 2px solid #eef2f7;
      }
      h3 {
        font-size: 1.15em;
      }
      p {
        max-width: 75ch;
      }
      ul,
      ol {
        max-width: 75ch;
      }
      a {
        color: #5a8dee;
        text-decoration: none;
        &:hover {
          text-decoration: underline;
        }
      }
      code:not([class*='language-']) {
        background: #f0f3f8;
        padding: 2px 6px;
        border-radius: 6px;
        font-size: 0.9em;
        color: #e74c3c;
      }
      pre {
        border-radius: 12px;
        overflow-x: auto;
        max-width: 100%;
      }
      blockquote {
        border-left: 4px solid #5a8dee;
        background: #f8faff;
        padding: 12px 16px;
        border-radius: 0 12px 12px 0;
        color: #8898aa;
        max-width: 75ch;
      }
      table {
        border-collapse: collapse;
        width: 100%;
        max-width: 100%;
        overflow-x: auto;
        display: block;
        th,
        td {
          border: 1px solid #eef2f7;
          padding: 8px 12px;
        }
        th {
          background: #f8faff;
          font-weight: 700;
        }
      }
      img {
        border-radius: 12px;
        max-width: 100%;
      }
    }
  }

  /* ===== TOC 样式 ===== */
  .doc-toc {
    :deep(.md-editor-catalog-link) {
      font-size: 12px;
      font-weight: 600;
      color: #8898aa;
      padding: 4px 8px;
      border-radius: 8px;
      transition: all 0.2s;
      border-left: 2px solid transparent;
      &:hover {
        color: #5a8dee;
        background: rgba(90, 141, 238, 0.05);
      }
      &.md-editor-catalog-active {
        color: #5a8dee;
        background: rgba(90, 141, 238, 0.08);
        border-left-color: #5a8dee;
      }
    }
  }
</style>
