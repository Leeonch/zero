# internal/authentication

放登录、刷新、退出、密码校验和令牌签发等认证业务。相关 HTTP 请求处理放在 `internal/handler/`。

主要对应原 NestJS 的 `src/modules/authentication/`，例如 `authentication.service.ts`、`password-hasher.ts` 和 `token-issuer.service.ts`。
