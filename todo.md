# IAM Service 迁移 TODO

将 `../iam-server` 当前提供的 IAM 能力迁移到 Go，沿用单进程 HTTP + gRPC 的结构。**从用户详情开始，每完成一个能力，就用相同数据对照 NestJS 验收。** 下文管理端路由统一带 `/iam-server` 前缀。

以原仓库的控制器、DTO、实体、`proto/iam/v1/identity.proto` 和 `test/` 为对照。原管理端的部分旧说明仍提到 Cookie Guard，当前控制器和契约测试采用管理端暂不鉴权、客户端直连 `/authentication/*` 的契约。管理端入口沿用现有受限访问边界。

## 1. 用户详情与最小数据接入

- [ ] 在 `internal/config` 和 `internal/svc` 接入 MySQL 配置、连接生命周期和依赖注入。先使用与 NestJS 同结构的测试库和只读账号验证查询，核对 UUID、字段名、可空列与时间格式。
- [ ] 在 `internal/user` 实现详情查询。读取 `users` 的资料、`principals` 的状态和 `user_identifiers` 的用户名，透传请求上下文。
- [ ] 接通 `GET /users/detail?id=<uuid>`。在 HTTP 层校验参数及未知字段，按原契约转换成功与业务错误响应，包括状态码、`code`、`message`、`data`、`timestamp` 和错误的 `path`。
- [ ] 用相同数据对照 NestJS。覆盖合法 ID 返回 200、缺少或非法 ID 返回 400、合法但不存在的 UUID 返回 404、缺少用户名时返回空字符串，以及 `avatarUrl: null`。

## 2. 用户分页与 Identity gRPC

- [ ] 迁移 `GET /users/list`。保持按创建时间降序、默认第 1 页每页 10 条、每页上限 100，并核对空列表、总数和非法分页参数。
- [ ] 从现有 proto 生成 Go 类型，保留包名、服务名、字段编号和 optional 语义。注册 `iam.v1.IdentityService`，让 `GetUser` 复用用户详情查询。
- [ ] 迁移 `BatchGetUsers`。覆盖 1–100 个不重复的 UUID、无效参数、部分或全部缺失、用户及 `missing_ids` 的输入顺序，以及头像字段缺失与时间序列化。

## 3. 用户、身份与凭据写入

- [ ] 核对开始写入所需的现有 MySQL 表结构、外键、唯一约束、排序规则和默认值。在测试库验证与 TypeORM 已有数据互通；若存在实际结构差异，先记录并验证对应的 schema 变更。
- [ ] 迁移用户名及各类标识符的规范化、唯一性和验证语义。复现现有 scrypt 密码凭据格式，用 NestJS 生成的凭据验证 Go，反向验证 Go 生成的凭据。
- [ ] 迁移 `POST /users/create`、`POST /users/update`。验证创建 Principal、User、用户名及密码凭据的事务，覆盖重复用户名冲突、资料字段省略与显式 `null`。
- [ ] 迁移 `GET /identities/list`、`GET /identities/detail` 和 `POST /identities/{create,verify,update-status}`。对照查询筛选、验证时间、重复标识符及批量状态变更。
- [ ] 迁移 `GET /sessions/list`、`POST /sessions/{revoke,revoke-users}`。核对筛选、撤销范围、已撤销会话和关联 Refresh Token 的处理。
- [ ] 迁移 `POST /users/update-status` 和 `POST /credentials/set-password`。核对非活跃状态触发的会话撤销、密码更新触发的凭据与会话联动，以及事务失败时的回滚。

## 4. 应用、OAuth Client 与应用访问

- [ ] 迁移 `GET /applications/{list,detail}` 和 `POST /applications/{create,update-status}`。验证编码唯一性、状态和批量操作结果。
- [ ] 迁移 `GET /oauth-clients/{list,detail}` 和 `POST /oauth-clients/{create,update-status,rotate-secret}`。保持 Public / Confidential、Grant Type、Token 有效期及 Refresh Token 交付配置。核对 Secret 的存储格式和仅在创建或轮换时返回明文的行为。
- [ ] 迁移 `GET /application-access/list` 和 `POST /application-access/{grant,revoke}`。核对目标校验、重复请求、撤销后重新授予及 `authorizationVersion` 更新。

## 5. 菜单、权限与角色

- [ ] 迁移 `GET /menus/{list,detail}` 和 `POST /menus/{create,update,delete}`。核对 Client 隔离、同一 Client 内 Key 唯一、父节点约束、循环检测、父节点省略与置空的差异，以及批量删除的事务行为。
- [ ] 迁移 `GET /permissions/{list,detail}`。权限按 Application 读取，列表全量返回。菜单和权限列表都对照原接口确认参数与响应结构。
- [ ] 迁移 `iam.v1.PermissionService/SyncPermissions`。保留 gRPC metadata 凭据、状态码、有效 Confidential Client 与所属应用验证。覆盖创建、更新、禁用、重新启用、空快照、重复请求、并发串行化、审计和授权版本。
- [ ] 迁移 `iam.v1.MenuService/SyncMenus`。核对后端凭据与目标 Client 归属、公开 `target_client_id`、父节点解析、非法树校验、保留已有 ID、删除缺失项、重试和事务。
- [ ] 迁移 `GET /roles/{list,detail,permissions,menus}` 和 `POST /roles/{create,update,update-status,set-permissions,set-menus,assign-permissions,unassign-permissions}`。核对全量替换、空数组清空、增量修改、应用及 Client 隔离，以及审计与授权版本更新条件。
- [ ] 迁移 `GET /role-assignments/{list,effective-permissions}` 和 `POST /role-assignments/{assign,revoke}`。核对访问资格、有效期、状态、历史记录、有效权限去重及授权版本。
- [ ] 对照有效菜单与权限的计算。覆盖角色状态和应用访问状态、Client 范围、菜单去重与补齐祖先，以及无有效分配时的空列表。

## 6. 客户端认证与令牌

- [ ] 迁移 `GET /.well-known/jwks.json` 与 EdDSA 签发、校验。核对现有密钥派生、`kid`、issuer、audience、claims、过期判断和 HTTP 响应，验证 Go 与 NestJS 可互验签名。
- [ ] 迁移 `POST /authentication/login`。核对标识符匹配、密码校验、用户与 Client 状态、应用访问资格、Public / Confidential Client 规则、会话创建和令牌响应。
- [ ] 迁移 `POST /authentication/refresh`。对照持久化的令牌哈希、Client 绑定、轮换、固定会话到期时间、并发请求和重放后撤销会话的行为。
- [ ] 迁移 `POST /authentication/logout`。核对客户端认证、会话与 Refresh Token 撤销、重复退出及错误响应。
- [ ] 对照 BODY / COOKIE 交付及浏览器行为。验证按 Client 隔离的 Cookie 名称、属性和清除条件，JSON 与可信 Origin 校验，以及允许携带凭据的 CORS。
- [ ] 迁移 `GET /authentication/permissions` 和 `GET /authentication/menus`。以已验证的 Token 和活动会话确定主体及 Client，核对 claims 与会话绑定、当前权限及授权版本、扁平菜单 Key 响应。
- [ ] 使用测试库验证升级前的密码、Client Secret、活动会话和刷新令牌。验证 NestJS 登录后可在 Go 刷新与退出，以及 Go 写入的状态可被 NestJS 继续处理。

## 7. 发布契约与服务替换

- [ ] 迁移原 `GET /` 的响应和 `GET /iam-service/proto/iam/v1/identity.proto` 下载，核对下载内容、内容类型、文件名及缓存响应头。处理当前 Go 的 `GET /hello` 示例，更新入口说明。
- [ ] 提供原有 Swagger / OpenAPI 入口。核对 `/iam-server/api` 的管理端文档、`/iam-service/api` 和 `/iam-service/api-json` 的五个客户端接口，以及现有调用方使用的 operation ID、请求响应结构和安全声明。
- [ ] 核对部署所需的启动参数、MySQL、HTTP / gRPC 监听、签发配置、Origin 和 Cookie 配置。启动 Go 实例验证两个监听端口及退出行为。
- [ ] 在同一份初始数据的独立测试库上，对照两个实现的 HTTP 响应、gRPC 结果和写入副作用；用现有管理端、客户端及业务后端完成集成验证。
- [ ] 用部署环境的数据备份演练切换与回退，核对签发配置和活跃会话能跨实现延续。按调用方切换流量，确保写操作由当前实例处理，验证通过后再停用 NestJS。

## 每项验收

- [ ] 记录原控制器、DTO、服务及相关测试的路径，核对成功、错误、边界值和副作用后再勾选。
- [ ] 对改动的 Go 文件执行 `gofmt`，运行 `go test ./...`。涉及启动、配置或路由时，执行 `go run ./cmd/iam` 并实际调用受影响的 HTTP 或 gRPC 入口。可用 `-f` 指定测试配置。
