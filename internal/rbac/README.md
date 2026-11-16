# internal/rbac

放角色、权限、角色分配、权限同步和有效权限查询等业务。管理端 HTTP 与业务端 gRPC 可以共用这些规则。

主要对应原 NestJS 的 `src/modules/rbac/`，例如 `rbac.service.ts` 和 `entity/`。`src/iam-server/rbac/` 中的控制器属于 HTTP 入口，`permission-rpc.controller.ts` 属于 gRPC 入口。
