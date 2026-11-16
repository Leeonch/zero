# internal/rpc

放 gRPC 服务的注册和方法实现。方法负责读取 metadata、转换 Protobuf 请求和响应，并调用 `user`、`rbac` 等业务包。当前监听在 `cmd/iam/main.go` 启动，业务方法尚未迁移。

对应原 NestJS 的 `src/modules/user/identity-rpc.controller.ts`、`src/modules/rbac/permission-rpc.controller.ts` 和 `src/modules/application/menu-rpc.controller.ts`。原 RPC 契约在 `proto/iam/v1/identity.proto`。
