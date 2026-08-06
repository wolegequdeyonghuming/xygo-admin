// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

import { AppRouteRecordRaw } from '@/utils/router'
import { ADMIN_BASE_PATH } from '@/router/routesAlias'
import { frontendRoutes } from '@/router/frontend'

/**
 * 静态路由配置（不需要权限就能访问的路由）
 *
 * 路由隔离设计（对齐 BuildAdmin）：
 *   - 后台管理：/admin 下（动态注册，此处只放登录/错误/catch-all）
 *   - 根路径 / 直接重定向到后台首页，不再加载官网介绍页
 *   - 前台文档/会员页（/docs、/user/*）恢复注册，仅通过手动输入 URL 访问，
 *     门户首页保持下线
 *
 * 注意事项：
 * 1、path、name 不要和动态路由冲突
 * 2、静态路由不管是否登录都可以访问
 */
export const staticRoutes: AppRouteRecordRaw[] = [
  // ===== 根路径：直接进入后台管理首页 =====
  {
    path: '/',
    redirect: `${ADMIN_BASE_PATH}/dashboard/console`
  },

  // ===== 后台管理静态路由（全部在 /admin 下）=====

  // 后台登录页（不需要后台布局包裹）
  {
    path: `${ADMIN_BASE_PATH}/login`,
    name: 'Login',
    component: () => import('@views/backend/auth/login/index.vue'),
    meta: { title: 'menus.login.title', isHideTab: true }
  },
  // 后台注册页
  {
    path: `${ADMIN_BASE_PATH}/register`,
    name: 'Register',
    component: () => import('@views/backend/auth/register/index.vue'),
    meta: { title: 'menus.register.title', isHideTab: true }
  },
  // 后台忘记密码
  {
    path: `${ADMIN_BASE_PATH}/forget-password`,
    name: 'ForgetPassword',
    component: () => import('@views/backend/auth/forget-password/index.vue'),
    meta: { title: 'menus.forgetPassword.title', isHideTab: true }
  },

  // ===== 公共错误页 =====
  {
    path: '/403',
    name: 'Exception403',
    component: () => import('@views/common/exception/403/index.vue'),
    meta: { title: '403', isHideTab: true }
  },
  {
    path: '/500',
    name: 'Exception500',
    component: () => import('@views/common/exception/500/index.vue'),
    meta: { title: '500', isHideTab: true }
  },

  // 后台 outside/iframe 容器
  {
    path: `${ADMIN_BASE_PATH}/outside`,
    component: () => import('@views/backend/index/index.vue'),
    name: 'Outside',
    meta: { title: 'menus.outside.title' },
    children: [
      {
        path: `${ADMIN_BASE_PATH}/outside/iframe/:path`,
        name: 'Iframe',
        component: () => import('@/views/common/outside/Iframe.vue'),
        meta: { title: 'iframe' }
      }
    ]
  },

  // /admin 入口：无论动态路由是否已注册，都先落到默认首页
  {
    path: `${ADMIN_BASE_PATH}`,
    redirect: `${ADMIN_BASE_PATH}/dashboard/console`
  },

  // ===== 前台文档 / 会员页面（原 URL，仅手动输入进入）=====
  ...frontendRoutes,

  // ===== 全局 404 兜底（必须放最后）=====
  {
    path: '/:pathMatch(.*)*',
    name: 'Exception404',
    component: () => import('@views/common/exception/404/index.vue'),
    meta: { title: '404', isHideTab: true }
  }
]
