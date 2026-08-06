// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

/**
 * 字典缓存状态管理
 *
 * 提供字典数据的统一获取、缓存与去重
 *
 * @module store/modules/dict
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { adminRequest } from '@/utils/http'

export interface DictItem {
  label: string
  value: string
  cssClass: string
  listClass: string
  isDefault: number
}

const SITE_DICT_URL = '/site/dict/data'

export const useDictStore = defineStore('dictStore', () => {
  const cache = ref<Record<string, DictItem[]>>({})
  const loading = ref<Record<string, boolean>>({})
  const promises = ref<Record<string, Promise<DictItem[]>>>({})

  /**
   * 加载单个字典类型数据（带缓存与去重）
   */
  const fetchDictData = async (dictType: string): Promise<DictItem[]> => {
    if (!dictType) return []

    if (cache.value[dictType]) {
      return cache.value[dictType]
    }

    if (dictType in promises.value) {
      return promises.value[dictType]
    }

    loading.value[dictType] = true

    const promise: Promise<DictItem[]> = adminRequest
      .get<{ list: DictItem[] }>({
        url: SITE_DICT_URL,
        params: { type: dictType }
      })
      .then((res) => {
        cache.value[dictType] = res.list || []
        return res.list || []
      })
      .finally(() => {
        loading.value[dictType] = false
        delete promises.value[dictType]
      })

    promises.value[dictType] = promise
    return promise
  }

  /**
   * 批量预加载多个字典
   */
  const preload = async (dictTypes: string[]): Promise<void> => {
    const unique = [...new Set(dictTypes)].filter((t) => !!t)
    await Promise.all(unique.map((t) => fetchDictData(t)))
  }

  /**
   * 获取字典数据列表（需先 fetchDictData 或 preload）
   */
  const getDictData = (dictType: string): DictItem[] => {
    return cache.value[dictType] || []
  }

  /**
   * 获取字典选项列表（用于 select/radio）
   */
  const getOptions = (dictType: string) => {
    return computed(() =>
      (cache.value[dictType] || []).map((item) => ({
        label: item.label,
        value: item.value
      }))
    )
  }

  /**
   * 根据值获取标签
   */
  const getLabel = (dictType: string, value: string | number): string => {
    if (value === undefined || value === null) return ''
    const list = cache.value[dictType] || []
    const found = list.find((item) => String(item.value) === String(value))
    return found ? found.label : String(value)
  }

  /**
   * 根据值获取 Tag 类型
   */
  const getTagType = (dictType: string, value: string | number): string => {
    if (value === undefined || value === null) return ''
    const list = cache.value[dictType] || []
    const found = list.find((item) => String(item.value) === String(value))
    return found?.listClass || 'info'
  }

  /**
   * 清空缓存
   */
  const clearCache = (): void => {
    cache.value = {}
    loading.value = {}
    promises.value = {}
  }

  /**
   * 检查字典是否已加载
   */
  const isLoaded = (dictType: string): boolean => {
    return !!cache.value[dictType]
  }

  return {
    cache,
    loading,
    fetchDictData,
    preload,
    getDictData,
    getOptions,
    getLabel,
    getTagType,
    clearCache,
    isLoaded
  }
})
