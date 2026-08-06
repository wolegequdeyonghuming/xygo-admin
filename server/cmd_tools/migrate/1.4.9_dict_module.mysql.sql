-- Migration: 1.4.9
-- Description: 新增数据字典模块（字典类型 + 字典数据）及菜单权限

CREATE TABLE IF NOT EXISTS `xy_sys_dict_type` (
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100) NOT NULL DEFAULT '' COMMENT '字典名称',
    `type`        varchar(100) NOT NULL DEFAULT '' COMMENT '字典标识',
    `remark`      varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
    `status`      tinyint      NOT NULL DEFAULT 1 COMMENT '状态：0=禁用 1=启用',
    `sort`        int          NOT NULL DEFAULT 0 COMMENT '排序',
    `created_by`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建人ID',
    `updated_by`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新人ID',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建时间（Unix秒）',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新时间（Unix秒）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_dict_type_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典类型';

CREATE TABLE IF NOT EXISTS `xy_sys_dict_data` (
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT,
    `dict_type_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '关联字典类型ID',
    `label`        varchar(100) NOT NULL DEFAULT '' COMMENT '字典标签',
    `value`        varchar(100) NOT NULL DEFAULT '' COMMENT '字典值',
    `css_class`    varchar(100) NOT NULL DEFAULT '' COMMENT '样式类名',
    `list_class`   varchar(100) NOT NULL DEFAULT '' COMMENT '表格样式类名（ElTag type）',
    `is_default`   tinyint      NOT NULL DEFAULT 0 COMMENT '是否默认',
    `status`       tinyint      NOT NULL DEFAULT 1 COMMENT '状态：0=禁用 1=启用',
    `sort`         int          NOT NULL DEFAULT 0 COMMENT '排序',
    `remark`       varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
    `created_by`   bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建人ID',
    `updated_by`   bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新人ID',
    `create_time`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建时间（Unix秒）',
    `update_time`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新时间（Unix秒）',
    PRIMARY KEY (`id`),
    KEY `idx_dict_data_type_id` (`dict_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典数据';

-- 字典管理目录（挂到 系统管理 模块下）
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 1, '字典管理', 'Dict', '/dict', '/index/index', '', 'ri:book-2-line', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 3, 1, '数据字典管理', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'System' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'Dict' AND t.type = 1);

-- 字典类型管理
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 2, '字典类型', 'DictType', 'dict/type', '/system/dict/type', 'sys_dict_type', 'ri:list-settings-line', 0, 1, '', '', '["GET /admin/dict/type/list","GET /admin/dict/type/detail","POST /admin/dict/type/save","POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 1, 1, '字典类型管理', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'Dict' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'DictType' AND t.type = 2);

-- 字典数据管理
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 2, '字典数据', 'DictData', 'dict/data', '/system/dict/data', 'sys_dict_data', 'ri:book-open-line', 0, 1, '', '', '["GET /admin/dict/data/list","POST /admin/dict/data/save","POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 2, 1, '字典数据管理', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'Dict' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'DictData' AND t.type = 2);

-- 字典类型按钮权限
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '新增', 'add', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典类型', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'add');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '编辑', 'edit', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典类型', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'edit');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '删除', 'delete', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典类型', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'delete');

-- 字典数据按钮权限
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '新增', 'add', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典数据', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'add');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '编辑', 'edit', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典数据', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'edit');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '删除', 'delete', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典数据', 0, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'delete');
