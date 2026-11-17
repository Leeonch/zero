# 第一步：用户详情与最小 MySQL 接入

对应 [迁移总清单的第 1 项](todo.md#1-用户详情与最小数据接入)。使用同一份测试数据，迁移并对照验收 `GET /iam-server/users/detail?id=<uuid>`。沿用现有管理端访问边界和 HTTP、gRPC 单进程结构。用户分页和 Identity gRPC 排在第 2 步。

## 1.1 确认原接口与测试数据

- [x] 从 NestJS 源码确认用户详情的路由、字段和查询语义。
  - **路由与参数**。完整路径为 `GET /iam-server/users/detail?id=<uuid>`。控制器从 query 读取必填的 `id`，DTO 使用未指定版本的 `@IsUUID()`。全局校验拒绝额外查询字段。该控制器及详情方法未声明鉴权 Guard，应用入口也未注册全局 Guard。[管理端控制器](../iam-server/src/iam-server/user/user.controller.ts#L33) · [请求 DTO](../iam-server/src/modules/user/dto/get-user.dto.ts#L4) · [全局校验](../iam-server/src/main.ts#L19)
  - **返回字段**。`data` 包含 `id`、`username`、`displayName`、`avatarUrl`、`status`、`createdAt`、`updatedAt`。其中 `avatarUrl` 可以为 `null`，两个时间字段来自用户记录，`status` 取自关联的 Principal，原样返回 `ACTIVE`、`DISABLED`、`LOCKED` 或 `CLOSED`。[响应 DTO](../iam-server/src/modules/user/dto/user-response.dto.ts#L9) · [字段映射](../iam-server/src/modules/user/user.service.ts#L230) · [状态定义](../iam-server/src/modules/principal/principal-status.ts#L1)
  - **查询语义**。按 `users.principal_id` 查找用户，TypeORM 同时加载关联的 Principal。随后按 `user_identifiers.user_principal_id` 和 `type = 'USERNAME'` 取标识符的原始 `value`，查询不筛选标识符状态。缺少用户名时返回空字符串。用户行不存在时抛出 404，消息为 `用户不存在`。[用户实体](../iam-server/src/modules/user/entity/user.entity.ts#L13) · [用户查找](../iam-server/src/modules/user/user.service.ts#L196) · [用户名查询](../iam-server/src/modules/user/user.service.ts#L213)
- [ ] 准备与 NestJS 同结构的 MySQL 测试库，为 Go 配置只读账号。用 `SHOW CREATE TABLE` 核对 [users](../iam-server/src/modules/user/entity/user.entity.ts)、[principals](../iam-server/src/modules/principal/entity/principal.entity.ts) 和 [user_identifiers](../iam-server/src/modules/identity/entity/user-identifier.entity.ts) 的实际列名、UUID 类型、外键、可空性和时间精度，并确认 NestJS 如何解释库内时间。
- [ ] 准备三类样本：资料完整的用户、只有 EMAIL 或 PHONE 且头像为 NULL 的用户、以及库中不存在的合法 UUID。让其中一个现有用户处于非 ACTIVE 状态，验证状态原样返回。先记录 NestJS 的实际 HTTP 状态和 JSON，尤其是参数校验的 `message`。

## 1.2 接入连接池

- [x] 在 [internal/config](internal/config/README.md) 定义 MySQL 配置，通过 `-f` 指定本地配置，通过 `IAM_MYSQL_PASSWORD` 环境变量传入密码。[配置示例](etc/README.md)
- [x] 在 [internal/svc](internal/svc/README.md) 使用 `database/sql` 和 MySQL 驱动组装进程共享的连接池，供后续用户查询注入。
- [x] 启动时用带超时的 `PingContext` 检查连接，失败时退出。[cmd/iam](cmd/iam/README.md) 在服务正常退出时关闭连接池。
- [ ] 用测试配置验证登录成功。目标端口可达，测试配置通过 TLS 到达 MySQL 登录阶段，服务器返回 `1045 Access denied`。核对密码及该账号对当前连接来源的授权后重试。

## 1.3 实现可复用的用户详情查询

- [ ] 在 [internal/user](internal/user/README.md) 提供接收 `context.Context` 和 ID 的详情查询。用参数化 SQL 按 `users.principal_id` 查找用户并关联 `principals.id`。从 `users` 读取 `display_name`、`avatar_url`、`created_at`、`updated_at`，从 `principals.status` 读取账号状态。
- [ ] 按 `user_identifiers.user_principal_id` 及 `type = 'USERNAME'` 查找用户名，返回原始 `value`。存在非活跃的用户名记录时也按原查询读取，缺少用户名时返回空字符串。查询结果保留头像 NULL 和空字符串的差异。
- [ ] 将用户行不存在映射成可识别的业务错误，独立处理连接、查询和扫描失败。将请求上下文传给数据库，业务返回值保留原始时间类型，供 HTTP 和后续 gRPC 分别序列化。

## 1.4 接通管理端 HTTP 路由

- [ ] 沿用 [cmd/iam](cmd/iam/README.md) 的路由注册方式。由 [internal/handler](internal/handler/README.md) 读取查询参数、按 NestJS 的 `@IsUUID()` 和白名单规则校验，再用 [internal/logic](internal/logic/README.md) 调用用户详情查询。在 [internal/types](internal/types/README.md) 定义 HTTP 结构。
- [ ] 在 [internal/httpapi](internal/httpapi/README.md) 处理详情接口的响应。以 [成功包装](../iam-server/src/common/interceptors/transform.interceptor.ts)、[异常过滤器](../iam-server/src/common/filters/http-exception.filter.ts)和[全局校验配置](../iam-server/src/main.ts)为依据：
  - **200** 返回 `code: 0`、`message: "success"`、`data` 和 `timestamp`。`data` 包含 `id`、`username`、`displayName`、`avatarUrl`、`status`、`createdAt`、`updatedAt`。
  - **400 / 404** 返回 `code: -1`、`message`、`data: null`、`timestamp` 和 `path`。`path` 包含原请求的查询字符串。400 的 `message` 以基线响应为准，404 的 `message` 为 `用户不存在`。
  - `timestamp`、`createdAt`、`updatedAt` 按 NestJS 的 UTC ISO 8601 毫秒格式输出，例如 `2026-01-01T00:00:00.000Z`。

## 1.5 对照验收

- [ ] 对改动的 Go 文件执行 `gofmt`，运行 `go test ./...`。从仓库根目录运行 `go run ./cmd/iam -f <本地测试配置路径>`，实际调用下面的用例，并验证 `GET /hello` 与 gRPC 监听端口。
- [ ] 对两个服务发送相同的请求，逐项核对 HTTP 状态码、JSON 字段、值与类型。响应生成时间 `timestamp` 核对格式和合理范围，库内的 `createdAt` 和 `updatedAt` 核对实际值是否一致。

| 请求或数据条件 | 预期结果 |
| --- | --- |
| 已存在的 UUID，含用户名和头像 | 200，资料、用户名、状态与日期和 NestJS 一致 |
| 已存在的 UUID，只有 EMAIL 或 PHONE 且头像为 NULL | 200，`username` 为 `""`，`avatarUrl` 为 `null` |
| 非 ACTIVE 用户的合法 UUID | 200，保留主体原有状态 |
| 缺少 `id`、`id=` 或非法 UUID | 400，校验消息与错误结构和 NestJS 一致 |
| 额外查询字段，如 `?id=<uuid>&extra=1` | 400，`path` 带原查询字符串 |
| 重复 `id`，以及大小写或其他 UUID 边界输入 | 分别记录 NestJS 的结果并核对 Go 的状态码与响应 |
| 合法但不存在的 UUID | 404，`message` 为 `用户不存在` |

- [ ] 记录两边的启动命令、样本条件、请求和结果。全部核对通过后，勾选总清单第 1 项。
