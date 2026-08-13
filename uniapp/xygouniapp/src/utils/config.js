// 默认同源（H5 走 nginx 反向代理；小程序端在下方用环境变量覆盖）
let BASE_URL = ''
let ASSET_URL = ''

// #ifndef H5
// 小程序端：API 地址由构建环境变量 VITE_API_BASE 注入（见 .env / .env.production），默认本机调试
const MP_API_BASE = import.meta.env.VITE_API_BASE || 'http://127.0.0.1:4096'
BASE_URL = MP_API_BASE
ASSET_URL = MP_API_BASE
// #endif

export default {
  BASE_URL,
  ASSET_URL,
  TOKEN_KEY: 'xygo_staff_token',
  REFRESH_TOKEN_KEY: 'xygo_staff_refresh_token',
  // 系统名称（TODO: 后端系统配置接口就绪后从后台获取，替换此占位值）
  SYS_NAME: '移动业务平台',
  // mock 预览模式：后端 /staff/* 就绪后置为 false，走真实接口
  MOCK_PREVIEW: false
}
