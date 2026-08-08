-- Migration: 1.5.0
-- Description: 代码生成字段表新增自增与可空标记列

ALTER TABLE `xy_sys_gen_codes_column`
  ADD COLUMN `is_auto_increment` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否自增' AFTER `dict_type`,
  ADD COLUMN `is_nullable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否可空' AFTER `is_auto_increment`;
