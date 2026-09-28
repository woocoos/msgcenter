# Changelog

## v0.3.0 (2026-07-15)

### 用户级模板定制

- `MsgTemplate` schema 新增 `user_id` 可选字段，支持用户级模板定制
- 模板选择实现三级回退策略：用户级 > 租户级 > 全局默认
- 前端模板管理页面新增"模板范围"列，显示模板级别（用户/租户/全局）
- 前端模板创建/编辑表单新增用户选择器，支持指定模板适用用户

### 动态附件

- 邮件通知支持通过 Alert annotations 传递动态附件路径（`__attachments__`）
- 支持本地文件路径和 HTTP(S) URL 两种附件来源
- API 阶段自动识别 OSS URL 并解析为本地挂载路径，避免运行时下载
- 配置 `alertManager.mountPaths` 支持按 bucket 映射本地挂载路径

### GraphQL 规范化

- Silence 相关 mutation 重命名：`createSilence` → `createMsgSilence`，`updateSilence` → `updateMsgSilence`，`deleteSilence` → `deleteMsgSilence`
- 统一实体命名规范，所有 mutation 使用 `Msg` 前缀

### 配置变更

- 新增 `alertManager.mountPaths`：bucket 到本地挂载路径的映射（用于动态附件）
- 移除废弃的 `alertManager.storage.path` 配置

## v0.2.0 (2026-07-14)

### 核心重构

#### Dispatcher 并发模型重构
- 使用 `sync.Map` + 多 worker 替代全局互斥锁，提升告警分发并发性能
- 新增 `route_groups.go`，路由分组逻辑独立管理
- 引入 `SlurpAndSubscribe` + OpenTelemetry 追踪，对齐上游 Alertmanager 订阅模式

#### Silence Schema 重构
- `Silence` 重命名为 `MsgSilence`，统一实体命名规范
- 同步上游 Silencer.Mutes 逻辑，新增 Silence 缓存层 (`service/silence/cache.go`)
- 新增 Silence 状态版本索引，修复竞态条件

#### Inhibitor 重构
- 移除 `run.Group` 依赖，简化生命周期管理
- 新增抑制索引 (`service/inhibit/index.go`)，提升匹配效率

### 通知增强

#### Email
- 支持密码文件读取 (`AuthPasswordFile`)，避免明文配置
- 支持隐式 TLS (端口 465 自动切换，可通过 `ForceImplicitTLS` 强制指定)
- 新增邮件线程功能 (`EmailThreading`)，通过 `References`/`In-Reply-To` 头实现邮件会话归组
- 自动生成 `Message-Id`，支持按日期 (`daily`) 或不分日期 (`none`) 分组线程

#### Webhook
- 新增请求超时配置 (`Timeout`)
- 支持 URL 文件读取 (`URLFile`)，敏感信息不入配置
- 支持 URL 模板化，动态生成目标地址
- 响应体 Drain 优化，避免连接复用问题

### 模板函数

新增 12 个模板函数：
- 时间类：`date`、`tz`、`now`、`since`、`humanizeDuration`
- 数据类：`dict`、`list`、`append`
- 其他：`toUpper`、`toLower`、`trimSpace`、`join`、`split`

### 新增模块

| 模块 | 路径 | 说明 |
|------|------|------|
| AlertMarker | `pkg/marker/` | 告警状态追踪，支持 inhibited/silenced 状态查询 |
| ExpiryBucket | `pkg/limit/` | 泛型缓存桶，基于优先队列按过期时间淘汰，支持容量上限 |
| OTel Tracing | `pkg/tracing/` | OpenTelemetry 追踪封装 |

### 代码风格

- `interface{}` 统一替换为 `any`
- import 分组排序规范化（标准库优先）
- `Matcher.UnmarshalJSON` 增强：同时支持 JSON 对象格式和文本格式 (`alertname="value"`)

### 测试

- 新增 Dispatcher、Inhibitor、Silence、Email、Webhook、模板函数的单元测试
- 修复测试套件中的外键约束、空指针等问题
- 补充测试数据：Org ID=1000 及对应 Webhook Channel
