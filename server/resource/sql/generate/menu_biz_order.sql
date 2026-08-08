-- 订单主表 菜单 SQL
-- 上级菜单ID: 0
-- ======= 顶级模式：创建目录(type=1) + 页面(type=2) + 按钮(type=3) =======

-- 1. 创建目录
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (0, 1, '订单主表', 'BizOrderDir', '/order/biz-order', '', '', 'mdi:order-bool-descending', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 2, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @parentId = LAST_INSERT_ID();

-- 2. 创建页面菜单
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@parentId, 2, '订单主表列表', 'BizOrder', 'order/biz-order', '/order/biz-order/index', 'biz_order', '', 0, 1, '', '', '["GET /admin/biz-order/list"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

SET @pageId = LAST_INSERT_ID();

-- 3. 创建按钮权限（根据选项按需生成）
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '查看订单主表', 'BizOrderView', '', '', '', '', 0, 0, '', '', '["GET /admin/biz-order/view"]', 0, 0, 0, '', '', 0, 0, 1, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
-- 详情页路由（隐藏页面，与列表页同级，active_path 高亮列表页）
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@parentId, 2, '订单主表详情', 'BizOrderDetail', 'order/biz-order/detail', '/order/biz-order/detail/index', '', '', 1, 0, '', '', '["GET /admin/biz-order/view"]', 0, 0, 0, '', '/order/biz-order', 0, 0, 0, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '新增订单主表', 'BizOrderAdd', '', '', '', '', 0, 0, '', '', '["POST /admin/biz-order/edit"]', 0, 0, 0, '', '', 0, 0, 2, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '编辑订单主表', 'BizOrderEdit', '', '', '', '', 0, 0, '', '', '["POST /admin/biz-order/edit","GET /admin/biz-order/view"]', 0, 0, 0, '', '', 0, 0, 3, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '删除订单主表', 'BizOrderDelete', '', '', '', '', 0, 0, '', '', '["POST /admin/biz-order/delete"]', 0, 0, 0, '', '', 0, 0, 4, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO `xy_admin_menu` (`parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`)
VALUES (@pageId, 3, '导出订单主表', 'BizOrderExport', '', '', '', '', 0, 0, '', '', '["GET /admin/biz-order/export"]', 0, 0, 0, '', '', 0, 0, 5, 1, '', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
