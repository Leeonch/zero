# IAM Service

这是 IAM 服务的 Go 项目。当前可从仓库根目录执行 `go run ./cmd/iam`。HTTP 暂时提供 `GET /hello`，gRPC 已监听，业务方法待迁移。

## 目录

- [`cmd/`](cmd/README.md) 放可执行程序，当前入口在 [`cmd/iam/`](cmd/iam/README.md)。

- [`etc/`](etc/README.md) 放运行配置。

- [`internal/`](internal/README.md) 放服务的私有代码，包括协议入口、依赖组装和 IAM 业务包。

- [`user/`](user/README.md) 是当前 Hello World 示例用的 Go 包。

原 NestJS 项目的启动文件是 `src/main.ts`，业务集中在 `src/modules/`，管理端 HTTP 入口集中在 `src/iam-server/`。各目录的 README 说明了对应代码的位置。
