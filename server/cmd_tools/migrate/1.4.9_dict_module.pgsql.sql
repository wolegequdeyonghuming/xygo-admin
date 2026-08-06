-- Migration: 1.4.9
-- Description: 新增数据字典模块（字典类型 + 字典数据）及菜单权限

CREATE TABLE IF NOT EXISTS xy_sys_dict_type (
    id          bigserial PRIMARY KEY,
    name        character varying(100) DEFAULT '' NOT NULL,
    type        character varying(100) DEFAULT '' NOT NULL,
    remark      character varying(255) DEFAULT '' NOT NULL,
    status      smallint DEFAULT 1 NOT NULL,
    sort        integer DEFAULT 0 NOT NULL,
    created_by  bigint DEFAULT 0 NOT NULL,
    updated_by  bigint DEFAULT 0 NOT NULL,
    create_time bigint DEFAULT 0 NOT NULL,
    update_time bigint DEFAULT 0 NOT NULL
);

COMMENT ON TABLE  xy_sys_dict_type IS '字典类型';
COMMENT ON COLUMN xy_sys_dict_type.id IS '主键';
COMMENT ON COLUMN xy_sys_dict_type.name IS '字典名称';
COMMENT ON COLUMN xy_sys_dict_type.type IS '字典标识';
COMMENT ON COLUMN xy_sys_dict_type.remark IS '备注';
COMMENT ON COLUMN xy_sys_dict_type.status IS '状态：0=禁用 1=启用';
COMMENT ON COLUMN xy_sys_dict_type.sort IS '排序';
COMMENT ON COLUMN xy_sys_dict_type.created_by IS '创建人ID';
COMMENT ON COLUMN xy_sys_dict_type.updated_by IS '更新人ID';
COMMENT ON COLUMN xy_sys_dict_type.create_time IS '创建时间（Unix秒）';
COMMENT ON COLUMN xy_sys_dict_type.update_time IS '更新时间（Unix秒）';

CREATE UNIQUE INDEX IF NOT EXISTS uk_dict_type_type ON xy_sys_dict_type (type);

CREATE TABLE IF NOT EXISTS xy_sys_dict_data (
    id           bigserial PRIMARY KEY,
    dict_type_id bigint DEFAULT 0 NOT NULL,
    label        character varying(100) DEFAULT '' NOT NULL,
    value        character varying(100) DEFAULT '' NOT NULL,
    css_class    character varying(100) DEFAULT '' NOT NULL,
    list_class   character varying(100) DEFAULT '' NOT NULL,
    is_default   smallint DEFAULT 0 NOT NULL,
    status       smallint DEFAULT 1 NOT NULL,
    sort         integer DEFAULT 0 NOT NULL,
    remark       character varying(255) DEFAULT '' NOT NULL,
    created_by   bigint DEFAULT 0 NOT NULL,
    updated_by   bigint DEFAULT 0 NOT NULL,
    create_time  bigint DEFAULT 0 NOT NULL,
    update_time  bigint DEFAULT 0 NOT NULL
);

COMMENT ON TABLE  xy_sys_dict_data IS '字典数据';
COMMENT ON COLUMN xy_sys_dict_data.id IS '主键';
COMMENT ON COLUMN xy_sys_dict_data.dict_type_id IS '关联字典类型ID';
COMMENT ON COLUMN xy_sys_dict_data.label IS '字典标签';
COMMENT ON COLUMN xy_sys_dict_data.value IS '字典值';
COMMENT ON COLUMN xy_sys_dict_data.css_class IS '样式类名';
COMMENT ON COLUMN xy_sys_dict_data.list_class IS '表格样式类名（ElTag type）';
COMMENT ON COLUMN xy_sys_dict_data.is_default IS '是否默认';
COMMENT ON COLUMN xy_sys_dict_data.status IS '状态：0=禁用 1=启用';
COMMENT ON COLUMN xy_sys_dict_data.sort IS '排序';
COMMENT ON COLUMN xy_sys_dict_data.remark IS '备注';
COMMENT ON COLUMN xy_sys_dict_data.created_by IS '创建人ID';
COMMENT ON COLUMN xy_sys_dict_data.updated_by IS '更新人ID';
COMMENT ON COLUMN xy_sys_dict_data.create_time IS '创建时间（Unix秒）';
COMMENT ON COLUMN xy_sys_dict_data.update_time IS '更新时间（Unix秒）';

CREATE INDEX IF NOT EXISTS idx_dict_data_type_id ON xy_sys_dict_data (dict_type_id);

-- 字典管理目录（挂到 系统管理 模块下）
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 1, '字典管理', 'Dict', '/dict', '/index/index', '', 'ri:book-2-line', 0, 0, '', '', '', 0, 0, 0, '', '', 0, 0, 3, 1, '数据字典管理', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'System' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'Dict' AND t.type = 1);

-- 字典类型管理
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 2, '字典类型', 'DictType', 'dict/type', '/system/dict/type', 'sys_dict_type', 'ri:list-settings-line', 0, 1, '', '', '["GET /admin/dict/type/list","GET /admin/dict/type/detail","POST /admin/dict/type/save","POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 1, 1, '字典类型管理', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'Dict' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'DictType' AND t.type = 2);

-- 字典数据管理
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 2, '字典数据', 'DictData', 'dict/data', '/system/dict/data', 'sys_dict_data', 'ri:book-open-line', 0, 1, '', '', '["GET /admin/dict/data/list","POST /admin/dict/data/save","POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 2, 1, '字典数据管理', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'Dict' AND p.type = 1
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu t WHERE t.name = 'DictData' AND t.type = 2);

-- 字典类型按钮权限
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '新增', 'add', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典类型', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'add');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '编辑', 'edit', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典类型', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'edit');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '删除', 'delete', '', '', 'sys_dict_type', '', 0, 0, '', '', '["POST /admin/dict/type/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典类型', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictType' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'delete');

-- 字典数据按钮权限
INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '新增', 'add', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 1, 1, '新增字典数据', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'add');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '编辑', 'edit', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/save"]', 0, 0, 0, '', '', 0, 0, 2, 1, '编辑字典数据', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'edit');

INSERT INTO xy_admin_menu (parent_id, type, title, name, path, component, resource, icon, hidden, keep_alive, redirect, frame_src, perms, is_frame, affix, show_badge, badge_text, active_path, hide_tab, is_full_page, sort, status, remark, created_by, updated_by, create_time, update_time)
SELECT p.id, 3, '删除', 'delete', '', '', 'sys_dict_data', '', 0, 0, '', '', '["POST /admin/dict/data/delete"]', 0, 0, 0, '', '', 0, 0, 3, 1, '删除字典数据', 0, 0, EXTRACT(EPOCH FROM NOW())::bigint, EXTRACT(EPOCH FROM NOW())::bigint
FROM xy_admin_menu p WHERE p.name = 'DictData' AND p.type = 2
AND NOT EXISTS (SELECT 1 FROM xy_admin_menu sub WHERE sub.parent_id = p.id AND sub.type = 3 AND sub.name = 'delete');
