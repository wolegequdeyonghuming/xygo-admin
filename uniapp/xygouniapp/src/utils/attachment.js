// 附件工具：图片/文件识别与打开（小程序端图片与非图片分别处理）
const IMAGE_EXTS = ['png', 'jpg', 'jpeg', 'gif', 'webp', 'bmp', 'svg', 'ico']

// 取扩展名（小写）
export function fileExt(url = '') {
  const seg = String(url).split('/').pop() || ''
  const m = /\.([a-zA-Z0-9]+)(\?.*)?$/.exec(seg)
  return m ? m[1].toLowerCase() : ''
}

// 取文件名（url 末段）
export function fileName(url = '') {
  const seg = String(url || '').split('/').pop()
  return seg || ''
}

// 是否为图片：按 mimetype 优先，其次扩展名
export function isImage(item) {
  const mime = (item.mimetype || item.mime || '').toLowerCase()
  if (mime.startsWith('image/')) return true
  return IMAGE_EXTS.includes(fileExt(item.url || ''))
}

// 打开附件：图片预览；非图片 H5 新开标签、小程序 downloadFile + openDocument
export function openAttachment(item, imageUrls = []) {
  const url = item.url
  if (!url) return
  if (isImage(item)) {
    uni.previewImage({ urls: imageUrls.length ? imageUrls : [url], current: url })
    return
  }
  // #ifdef H5
  window.open(url, '_blank')
  // #endif
  // #ifndef H5
  uni.showLoading({ title: '打开中...' })
  uni.downloadFile({
    url,
    success: (res) => {
      uni.hideLoading()
      uni.openDocument({
        filePath: res.tempFilePath,
        showMenu: true,
        fail: () => uni.showToast({ title: '无法打开该文件', icon: 'none' })
      })
    },
    fail: () => {
      uni.hideLoading()
      uni.showToast({ title: '文件下载失败', icon: 'none' })
    }
  })
  // #endif
}
