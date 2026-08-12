let BASE_URL = 'http://127.0.0.1:4096'

// #ifdef H5
BASE_URL = ''
// #endif

// 图片等静态资源源（H5 下走后端直连，不依赖 vite 代理；小程序用同一后端地址）
const ASSET_URL = 'http://127.0.0.1:4096'

export default {
  BASE_URL,
  ASSET_URL,
  TOKEN_KEY: 'xygo_staff_token',
  REFRESH_TOKEN_KEY: 'xygo_staff_refresh_token',
  // 系统名称（TODO: 后端系统配置接口就绪后从后台获取，替换此占位值）
  SYS_NAME: '收单员工作台',
  // mock 预览模式：后端 /staff/* 就绪后置为 false，走真实接口
  MOCK_PREVIEW: false
}
