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


/*
 Navicat Premium Dump SQL

 Source Server         : 81.71.150.152-8
 Source Server Type    : MySQL
 Source Server Version : 80036 (8.0.36)
 Source Host           : 81.71.150.152:9798
 Source Schema         : xygonew

 Target Server Type    : MySQL
 Target Server Version : 80036 (8.0.36)
 File Encoding         : 65001

 Date: 14/08/2026 16:53:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for xy_biz_order_details
-- ----------------------------
DROP TABLE IF EXISTS `xy_biz_order_details`;
CREATE TABLE `xy_biz_order_details`  (
  `id` bigint NOT NULL COMMENT '主键',
  `order_id` bigint NOT NULL COMMENT '订单id',
  `created_at` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint NULL DEFAULT NULL COMMENT '修改时间',
  `user_id` bigint NULL DEFAULT NULL COMMENT '操作人',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `is_deleted` tinyint NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单详细情况描述表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for xy_biz_order
-- ----------------------------
DROP TABLE IF EXISTS `xy_biz_order`;
CREATE TABLE `xy_biz_order`  (
  `id` bigint NOT NULL COMMENT '主键',
  `order_no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `schedule_date` datetime NULL DEFAULT NULL COMMENT '排单日期',
  `visit_date` datetime NULL DEFAULT NULL COMMENT '上门日期',
  `customer_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户姓名',
  `available_time_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '可联系时间',
  `area` int NULL DEFAULT NULL COMMENT '区县（字典）',
  `install_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '安装地址',
  `contact_phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '联系电话',
  `business_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '预约业务',
  `expiry_date` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '到期时间',
  `primary_tel_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主卡号码',
  `details` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '详细情况（关联）',
  `telemarketer_id` bigint NULL DEFAULT NULL COMMENT '话务员id',
  `agent_id` bigint NULL DEFAULT NULL COMMENT '收单员id',
  `appointment_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收单预约情况',
  `follow_up_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '话务二次回访情况',
  `dealt_business_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '成交业务',
  `porting_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '携转情况',
  `customer_real_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户实际姓名',
  `customer_id_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户身份证号',
  `paid_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '实缴额度（元）',
  `is_rural_order` tinyint NULL DEFAULT 0 COMMENT '是否乡下单',
  `new_phone_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '新开号码',
  `device_serial` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '终端串码',
  `broadband_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '宽带账号',
  `is_completed` tinyint NULL DEFAULT 0 COMMENT '是否完工',
  `subsidy_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '话补',
  `agency_no` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '工号',
  `is_new` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否纯新增',
  `created_at` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` tinyint NULL DEFAULT 0 COMMENT '是否删除',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `order_status` int NULL DEFAULT NULL COMMENT '订单状态',
  `attachment_id` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '附件id',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_order_no`(`order_no` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单主表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
