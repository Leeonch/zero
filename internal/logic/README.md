# internal/logic

放 go-zero HTTP 用例的编排代码。一个用例可以接收请求参数，调用用户、认证或权限等业务包，再整理返回值。目前有 Hello World 用例。

对应原 NestJS 控制器中调用 Service、组合结果的部分，例如 `src/iam-server/user/user.controller.ts`。原 `src/modules/user/user.service.ts` 中可由 HTTP 和 gRPC 共同使用的业务规则归 [`internal/user/`](../user/README.md)。
