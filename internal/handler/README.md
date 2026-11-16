# internal/handler

放 go-zero HTTP handler，负责解析请求、调用 `logic`、写入响应。目前这里有 `GET /hello` 的 handler，路由在 `cmd/iam/main.go` 注册。

对应原 NestJS 的 `src/iam-server/**/*.controller.ts` 和 `src/modules/authentication/authentication.controller.ts` 中处理 HTTP 请求的部分。
