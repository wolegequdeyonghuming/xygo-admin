-- ========================================================
-- 升级脚本：字典模块（1.4.9）
-- 用于已部署系统新增字典功能
-- 全新安装请使用 mysql_install.sql
-- ========================================================

-- --------------------------------------------------------
-- 表的结构 `xy_sys_dict_type`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `xy_sys_dict_type` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `type` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典标识',
  `remark` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：0=禁用 1=启用',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `create_time` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间（Unix秒）',
  `update_time` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间（Unix秒）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_dict_type_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典类型';

-- --------------------------------------------------------
-- 表的结构 `xy_sys_dict_data`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `xy_sys_dict_data` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dict_type_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '关联字典类型ID',
  `label` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典标签',
  `value` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典值',
  `css_class` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '样式类名',
  `list_class` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表格样式类名（ElTag type）',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：0=禁用 1=启用',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `create_time` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间（Unix秒）',
  `update_time` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间（Unix秒）',
  PRIMARY KEY (`id`),
  KEY `idx_dict_data_type_id` (`dict_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典数据';

-- --------------------------------------------------------
-- 字典类型种子数据
-- --------------------------------------------------------
INSERT INTO `xy_sys_dict_type` (`id`, `name`, `type`, `remark`, `status`, `sort`, `created_by`, `updated_by`, `create_time`, `update_time`) VALUES
(1, '性别', 'gender', '会员性别字典', 1, 0, 0, 0, 1770000000, 1770000000),
(2, '状态', 'common_status', '通用启用/禁用状态', 1, 1, 0, 0, 1770000000, 1770000000);

-- --------------------------------------------------------
-- 字典数据种子数据
-- --------------------------------------------------------
INSERT INTO `xy_sys_dict_data` (`id`, `dict_type_id`, `label`, `value`, `css_class`, `list_class`, `is_default`, `status`, `sort`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`) VALUES
(1, 1, '男', '1', '', 'primary', 1, 1, 0, '', 0, 0, 1770000000, 1770000000),
(2, 1, '女', '2', '', 'danger', 0, 1, 1, '', 0, 0, 1770000000, 1770000000),
(3, 1, '未知', '0', '', 'info', 0, 1, 2, '', 0, 0, 1770000000, 1770000000),
(4, 2, '启用', '1', '', 'success', 1, 1, 0, '', 0, 0, 1770000000, 1770000000),
(5, 2, '禁用', '0', '', 'danger', 0, 1, 1, '', 0, 0, 1770000000, 1770000000);

-- --------------------------------------------------------
-- 字典管理菜单（挂到 Auth 权限管理模块下）
-- --------------------------------------------------------
INSERT INTO `xy_admin_menu` (`id`, `parent_id`, `type`, `title`, `name`, `path`, `component`, `resource`, `icon`, `hidden`, `keep_alive`, `redirect`, `frame_src`, `perms`, `is_frame`, `affix`, `show_badge`, `badge_text`, `active_path`, `hide_tab`, `is_full_page`, `sort`, `status`, `remark`, `created_by`, `updated_by`, `create_time`, `update_time`) VALUES
(900, 60, 1, '字典管理', 'Dict', '/dict', '/index/index', '', 'ri:book-2-line', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 3, 1, '数据字典管理', 0, 0, 1770000000, 1770000000),
(901, 900, 2, '字典类型', 'DictType', 'dict/type', '/system/dict/type', 'sys_dict_type', 'ri:list-settings-line', 0, 1, '', '', '["GET /admin/dict/type/list","GET /admin/dict/type/detail","POST /admin/dict/type/save","POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 1, 1, '字典类型管理', 0, 0, 1770000000, 1770000000),
(902, 900, 2, '字典数据', 'DictData', 'dict/data', '/system/dict/data', 'sys_dict_data', 'ri:book-open-line', 0, 1, '', '', '["GET /admin/dict/data/list","POST /admin/dict/data/save","POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 2, 1, '字典数据管理', 0, 0, 1770000000, 1770000000),
(903, 901, 3, '新增', 'add', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典类型', 0, 0, 1770000000, 1770000000),
(904, 901, 3, '编辑', 'edit', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典类型', 0, 0, 1770000000, 1770000000),
(905, 901, 3, '删除', 'delete', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典类型', 0, 0, 1770000000, 1770000000),
(906, 902, 3, '新增', 'add', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典数据', 0, 0, 1770000000, 1770000000),
(907, 902, 3, '编辑', 'edit', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典数据', 0, 0, 1770000000, 1770000000),
(908, 902, 3, '删除', 'delete', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典数据', 0, 0, 1770000000, 1770000000);
