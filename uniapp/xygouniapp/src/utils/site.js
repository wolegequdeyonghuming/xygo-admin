// 站点名称（从 /site/index 读取 site_name，缓存供登录页/首页共用）
import { ref } from 'vue'
import { get } from '@/utils/request'

export const siteName = ref('')

export async function loadSiteName() {
  if (siteName.value) return siteName.value
  try {
    const data = await get('/site/index')
    if (data && data.siteName) {
      siteName.value = data.siteName
      // H5：同步浏览器标签页标题
      // #ifdef H5
      if (typeof document !== 'undefined') {
        document.title = data.siteName
      }
      // #endif
    }
  } catch (e) {
    // 忽略，使用兜底文案
  }
  return siteName.value
}
