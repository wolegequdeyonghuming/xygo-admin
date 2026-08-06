# 数据字典功能设计方案

## 1. 现状分析

系统**没有**数据字典功能。现有相关代码：
- `sys_gen_codes_column.dict_type` 列存在但未使用（前端 UI 已注释："字典类型已移除，使用字段注释中的静态选项映射"）
- 代码生成器仅支持**静态** `dict-options`（注释解析 `"状态:0=成功,1=失败"` 或设计器手动配置），编译时嵌入 `.vue`
- `sys_config` 中 `group=dictionary` 的配置行是**配置分组字典**（管理 config_group 的 UI 元数据），与数据字典无关

## 2. 数据库表

### `xy_sys_dict_type` — 字典类型

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint PK | 主键 |
| name | varchar(100) | 字典名称（如"性别"） |
| type | varchar(100) | 字典标识（如"gender"），唯一索引 |
| remark | varchar(255) | 备注 |
| status | tinyint(1) | 状态：0=禁用 1=启用 |
| sort | int | 排序 |
| created_by | bigint | 创建人ID |
| updated_by | bigint | 更新人ID |
| create_time | bigint | 创建时间（Unix秒） |
| update_time | bigint | 更新时间（Unix秒） |

### `xy_sys_dict_data` — 字典数据

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint PK | 主键 |
| dict_type_id | bigint | 关联字典类型ID |
| label | varchar(100) | 字典标签（如"男"） |
| value | varchar(100) | 字典值（如"1"） |
| css_class | varchar(100) | 样式类名 |
| list_class | varchar(100) | 表格样式类名（ElTag type: success/danger/warning/info/primary） |
| is_default | tinyint(1) | 是否默认 |
| status | tinyint(1) | 状态：0=禁用 1=启用 |
| sort | int | 排序 |
| remark | varchar(255) | 备注 |
| created_by | bigint | 创建人ID |
| updated_by | bigint | 更新人ID |
| create_time | bigint | 创建时间（Unix秒） |
| update_time | bigint | 更新时间（Unix秒） |

## 3. 后端模块

### 文件结构

```
server/
├── api/admin/admin_dict.go              # 请求/响应结构体 (g.Meta 路由)
├── internal/model/input/adminin/dict.go # 输入模型
├── internal/controller/admin/dict.go    # 薄控制器
├── internal/logic/dict/dict.go          # 业务逻辑
├── internal/service/dict.go             # 服务接口
├── internal/model/entity/sys_dict_*.go  # gf gen dao 生成
├── internal/model/do/sys_dict_*.go      # gf gen dao 生成
├── internal/dao/sys_dict_*.go           # gf gen dao 生成
└── cmd_tools/migrate/1.4.9_dict_module.mysql.sql  # 迁移 SQL
```

### API 路由

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin/dict/type/list` | 字典类型列表（分页） |
| GET | `/admin/dict/type/detail` | 字典类型详情 |
| POST | `/admin/dict/type/save` | 保存字典类型（新增/编辑） |
| POST | `/admin/dict/type/delete` | 删除字典类型 |
| GET | `/admin/dict/data/list` | 字典数据列表（按类型ID） |
| POST | `/admin/dict/data/save` | 保存字典数据 |
| POST | `/admin/dict/data/delete` | 删除字典数据 |
| GET | `/site/dict/data` | **公开接口**：按 type 标识返回字典数据，无需登录 |

公开接口 `/site/dict/data?type=gender` 返回 `{list: [{label:"男",value:"1",listClass:"primary"},{label:"女",value:"0",listClass:"danger"}]}`。

### 生成步骤

1. 创建迁移 SQL → 执行 `go run tools.go migrate up`
2. `gf gen dao`（生成 entity/do/dao）
3. 手写 api → model/input → controller → logic → service 接口
4. 编辑 `internal/logic/logic.go` 添加 `_ "xygo/internal/logic/dict"`
5. `gf gen service`（生成 service 实现绑定）
6. `go build` 验证

## 4. 前端

### 文件结构

```
web/src/
├── api/backend/dict.ts                          # 管理 API 封装
├── store/modules/dict.ts                        # DictStore（Pinia）
├── components/DictLabel/index.vue               # DictLabel 组件
├── views/backend/system/dict/type/index.vue     # 字典类型管理页
└── views/backend/system/dict/data/index.vue     # 字典数据管理页
```

### DictStore（Pinia）

- **缓存**：`cache: Record<string, DictItem[]>` — dictType → 数据列表
- **请求去重**：同一 `dictType` 只发一次请求（`promises` 去重）
- **getters**：`getDictData(type)`、`getLabel(type, value)`、`getTagType(type, value)`、`getOptions(type)`
- **actions**：`fetchDictData(type)`、`preload(types[])`（批量预加载）、`clearCache()`

### DictLabel 组件

接收 `dictType` + `value` props，自动从 store 映射标签并渲染为 `ElTag`：

```vue
<ElTag v-if="label" :type="tagType" :size="size">{{ label }}</ElTag>
<span v-else>{{ fallback }}</span>
```

Props: `dictType: string`, `value: string | number`, `size?: 'small' | 'default' | 'large'`, `fallback?: string`

### 管理页面

- **字典类型管理**：标准 CRUD 表格（名称/标识/状态/排序/备注），dialog 表单
- **字典数据管理**：左侧字典类型列表（点击切换），右侧数据表格（标签/值/样式/状态/排序），dialog 表单

## 5. 代码生成器集成

### TplColumn 新增字段

```go
DictType string // 字典类型标识，如 "gender"，空表示无字典
```

`buildTplColumn` 中从 `col.DictType` 赋值。

### 生成模板改造

**表格列渲染（`web_index.vue.tpl`）**：当 `.DictType` 非空时，`formatter` 改用 DictLabel 组件

**表单字段（`web_dialog.vue.tpl`）**：当 `.DictType` 非空时，ElRadioGroup/ElSelect 选项改为 `v-for="opt in dictStore.getOptions('xxx')"`

**搜索表单（`web_search.vue.tpl`）**：同上，ElSelect 选项来源改为 store

**`<script setup>` 注入**：收集所有列的 DictType，生成 `dictStore.preload([...])` 调用

### 前端代码生成器 UI

恢复 `field-property-panel.vue` 中的 `dictType` 选择器，下拉框列出所有启用字典类型（调用 `/admin/dict/type/list`）。

## 6. 菜单与权限

挂到 Auth 权限管理模块（id=140）下：

```
权限管理(140)
├── 部门管理(141)
├── 岗位管理(142)
├── 字典管理(Dict)  ← 新增 type=1
│   ├── 字典类型(DictType)  ← 新增 type=2
│   │   ├── 新增(add)  type=3
│   │   ├── 编辑(edit) type=3
│   │   └── 删除(delete) type=3
│   └── 字典数据(DictData)  ← 新增 type=2
│       ├── 新增(add)  type=3
│       ├── 编辑(edit) type=3
│       └── 删除(delete) type=3
```

菜单使用 idempotent INSERT...SELECT 模式（跟随 SMS 模块迁移风格），不硬编码 ID。

## 7. 实现顺序

1. 创建迁移 SQL（`1.4.9_dict_module.mysql.sql` + `.pgsql.sql`）
2. 执行迁移 → `gf gen dao` 生成 entity/do/dao
3. 编写后端：api → input → controller → logic → service
4. 更新 `logic.go` 聚合导入 + `go build` 验证
5. 前端：DictStore + DictLabel 组件 + 管理页面 + API 封装
6. 代码生成器集成：TplColumn + 模板改造 + UI 恢复 dictType
7. 更新 `mysql_install.sql` / `pgsql_install.sql` 种子数据
8. `pnpm build` 验证前端构建