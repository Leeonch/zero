# internal

放 IAM 服务的私有 Go 包。HTTP 和 gRPC 入口负责协议处理，业务包负责业务规则，`svc` 负责组装它们。

原 NestJS 对应 `src/iam-server/` 中的管理端入口，以及 `src/modules/` 中的客户端入口、RPC 入口和业务实现。迁移时按职责找对应目录。

| Go 目录 | 职责 |
| --- | --- |
| [`config/`](config/README.md)、[`svc/`](svc/README.md) | 配置字段和依赖组装 |
| [`handler/`](handler/README.md)、[`logic/`](logic/README.md)、[`types/`](types/README.md) | go-zero HTTP 请求处理、用例和类型 |
| [`httpapi/`](httpapi/README.md) | 按需放手写的 HTTP 协议辅助代码 |
| [`rpc/`](rpc/README.md) | gRPC 入口 |
| [`authentication/`](authentication/README.md)、[`rbac/`](rbac/README.md)、[`session/`](session/README.md)、[`user/`](user/README.md) | IAM 业务包，当前待迁移 |
