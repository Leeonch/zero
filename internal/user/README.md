# internal/user

放正式的 IAM 用户业务，例如用户创建、查询、状态更新和数据访问。HTTP 管理接口与 gRPC 用户查询都可以调用这个包。

主要对应原 NestJS 的 `src/modules/user/user.service.ts`、`src/modules/user/entity/` 等业务文件。原 `src/iam-server/user/user.controller.ts` 和 `identity-rpc.controller.ts` 分别对应 HTTP 和 gRPC 入口。
