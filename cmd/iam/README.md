# cmd/iam

`main.go` 是 IAM 进程的入口。它读取配置，组装依赖，注册 HTTP 路由，并启动 HTTP 和 gRPC 监听。从仓库根目录执行 `go run ./cmd/iam`。

对应原 NestJS 的 `src/main.ts`。原 `src/app.module.ts` 中注册模块和创建依赖的工作，在 Go 侧由入口显式调用 [`internal/svc/`](../../internal/svc/README.md) 完成。
