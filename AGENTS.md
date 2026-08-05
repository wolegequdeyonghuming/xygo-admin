# AGENTS.md

本文件为在此仓库中运行编码代理时提供的指导。所有命令默认从相应子目录执行（见各节）。

## 项目概览

`xygo-admin`：基于 **Vue3 + GoFrame v2** 的通用开源中后台管理框架（RBAC 权限、代码生成、系统监控、消息队列、单二进制部署）。

- `server/`：后端 GoFrame 项目，模块名 `xygo`，Go 1.24。
- `web/`：前端 Vue3 + TypeScript + Element Plus + Pinia + Tailwind（pnpm）。
- `uniapp/`：微信小程序 / 会员端移动应用。
- `docs/`：开发文档（如 `addon-development-guide.md`）。
- `mysql_install.sql` / `pgsql_install.sql`：MySQL/PostgreSQL 初始化脚本。
- `version.json`：在线更新所需版本信息。

### 后端分层结构（server/internal/）

- `api/`：请求/响应结构体（`g.Meta` 定义路由 tag）。
- `internal/controller/`：控制器（薄，仅解析入参并调用 logic）。
- `internal/logic/`：业务逻辑（核心代码，服务实现）。
- `internal/model/`：`entity` / `do` / `input`（DAO 生成与请求模型）。
- `internal/dao/`、`internal/service/`：DAO 与 Service 接口，**由 `gf gen dao/service` 自动生成，勿手改**。
- 聚合文件：`internal/logic/logic.go`、`internal/field`、`internal/crons`、`internal/queues`、`addons` 通过空导入 (`_`) 触发各类注册，改动相关目录常见模块后需同步该聚合导入。

## ⚠️ 首次运行前置（必读）

启动后端前必须完成配置，否则 `g.Cfg().MustGet` 会 panic：

1. 复制配置文件（`config.yaml` 已被 `server/.gitignore` 排除，不在仓库内）：
   - `server/manifest/config/` 下仅有 `config.yaml.example`。
   - 将其复制为 `server/manifest/config/config.yaml`。
2. 导入 `mysql_install.sql` 或 `pgsql_install.sql`，并核对 `config.yaml` 中 `database.default.link` 连接串。
3. 开发期保持 `system.mode: develop`；前端 `dist` 未构建时开发模式不阻塞启动，`product` 模式缺少 `index.html` 会直接启动失败（需先 `pnpm build`）。

## 构建 / 运行 / 测试命令

以下命令均在对应子目录内执行（用 `--workdir` 切换到目标目录，勿用 `cd &&` 拼接）。

### 后端（server/）

```bash
go run main.go            # 启动后端（需先完成上述配置）
go run tools.go           # 交互式工具菜单
go run tools.go migrate up     # 执行数据库迁移
go run tools.go migrate status # 查看迁移状态
go run tools.go addon install <name>  # 安装扩展
go run tools.go gen <table>        # 命令行 CRUD 生成
```

- 编译：`go build -o xygo.exe .`
- 热编译：`gf run main.go`
- Makefile 快捷命令：`make dao`、`make service`、`make logic-import`、`make gen`（一步生成全部）、`make build`、`make run`、`make clean`
- 代码生成基于 `hack/config.yaml` 的 `gfcli` 配置（含数据库连接、`removePrefix: "xy_"`、`jsonCase: CamelLower`）。

### 后端测试（标准 go test）

```bash
go test ./internal/logic/gencodes/ -run TestTableToRelationNameCoreTables -v  # 单个测试
go test ./internal/library/dbdialect/ -v                                      # 某包全部测试
go test ./...                                                                 # 全量
go test ./internal/library/dbdialect/ -run TestPgDialect_CreateTableSQL -v    # 按名过滤
```

- 测试文件放在被测包内，命名 `*_test.go`，使用 table-driven、`t.Fatalf/t.Errorf` 断言风格（参考 `internal/logic/gencodes/generate_test.go`、`internal/library/dbdialect/dialect_test.go`）。

### 前端（web/）

```bash
pnpm install     # 安装依赖（锁定使用 pnpm）
pnpm dev         # 开发服务器（默认 :3006，代理 API）
pnpm build       # 构建（含 vue-tsc --noEmit 类型检查）
pnpm lint        # ESLint
pnpm fix         # eslint --fix
```

## 代码风格规范

### Go（后端）

- **文件头**：多数 Go 文件以一个固定 11 行 `// +--` banner 开头（含版权、MIT License、作者 喜羊羊 `<751300685@qq.com>`）。新增/修改业务文件建议保留此头。
- **导入分组与顺序**：按 stdlib → 第三方（`github.com/gogf/gf/...`）→ 项目内部（`xygo/...`）分组，内部以别名（如 `api "xygo/api/system"`）区分包名冲突；每组空一行。用 `gofmt`/`goimports` 保持格式。
- **GoFrame v2 惯用法**：
  - API 请求/响应结构体内嵌 `g.Meta` 定义路由：`g.Meta \`path:"/health" method:"get" tags:"System" summary:"..."\``。
  - 字段标签顺序常用 `p`(参数名) `d`(默认值) `v`(校验) `json`，响应字段用 `json` + `dc`(描述)。
  - 控制器：`type ControllerV1 struct{}` + `func NewV1() *ControllerV1`；方法签名具名返回 `(res *api.XxxRes, err error)`。
  - 业务服务：`type sXxx struct{}`，`New()` 构造，包内 `init()` 调用 `service.RegisterXxx(New())` 注册。
  - 日志用 `g.Log().Info/Error/Fatalf`；配置用 `g.Cfg().MustGet`；DAO 链式调用 `dao.Xxx.Ctx(ctx).Where(...).OrderAsc(...).Scan(&items)`。
- **错误处理**：返回错误而非 `panic`（顶层除外）。用 `gerror.Wrapf(err, "上下文%v", ...)` 保留上下文，业务错误用 `gerror.New` / `gerror.NewCode(consts.CodeDataNotFound, "中文提示")`；错误文案为中文、面向用户。
- **命名**：类型 PascalCase；服务实现以 `s` 前缀；JSON 字段 lowerCamel（生成器 `jsonCase: CamelLower`）；DAO/表自动生成名去掉 `xy_` 前缀（`removePrefix`）。
- **注释**：使用中文注释，函数注释说明用途/参数/返回；`// 注释` 风格。业务代码可带 `TODO` 注明待办。
- **安全**：不提交密钥（`config.yaml`、`.env` 已被忽略）；JWT secret、数据库密码仅存在于本地配置。

### 前端（web/）

- **Vue3 Composition API**，使用 `<script setup>` 与 TypeScript（`lang="ts"`），`import type` 导入类型。
- **Prettier 配置**（`.prettierrc`）：`printWidth: 100`、`tabWidth: 2`、**无分号**、**单引号**、`trailingComma: "none"`、`vueIndentScriptAndStyle: true`。
- 建议组件/页面遵循现有 `src/` 目录（`api`、`views`、`router`、`store`(Pinia)、`components`、`hooks`）。
- 路由：静态 + 动态加载；状态管理用 Pinia（含持久化插件）。
- 上传、富文本等统一走现有接口与封装（参考 `utils`、`api`）。

### 提交规范（Git Commit）

- 遵循 Conventional Commits，subject 用**中文描述**，常用格式：
  - `feat(v1.x.y): 中文说明` 新增功能
  - `fix(v1.x.y): 中文说明` 修复缺陷
  - 也见 `refactor` / `perf` / `docs` / `chore` / `test` / `build` / `release: vX.Y.Z ...`。
- 参考历史：`fix(v1.4.7): 合并 Pinia 双树并修复前端 TS 配置`、`feat(v1.4.4): 命令行 CRUD 生成器与 dist 更新`、`release: v1.4.8 消息队列 Topic 配置、多 Worker 与延迟重试`。
- 前端经 husky 钩子：`pre-commit` 跑 lint-staged（ESLint + stylelint + prettier），`commit-msg` 跑 commitlint。

### 其他

- 命令环境为 PowerShell（win32）。避免 `cd dir && cmd`，改用工作目录切换。
- 同步数据库结构后，需重新生成以下内容（顺序）：`gf gen dao` → `make logic-import` → `gf gen service`，确保 DAO/Service 与实体一致。