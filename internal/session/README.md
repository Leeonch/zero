# internal/session

放 IAM 会话的业务操作，例如会话查询、撤销及按用户批量撤销。

主要对应原 NestJS 的 `src/modules/session/`。其依赖的会话实体目前位于 `src/modules/authentication/entity/session.entity.ts`。管理端入口位于 `src/iam-server/session/session.controller.ts`。
