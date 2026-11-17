# IAM 服务

本项目以学习 Go 为目的，同时将现有的 NestJS `iam-server` 迁移到 Go。优先采用简明、符合 Go 惯例的写法，方便阅读和学习。按功能逐步迁移，并对照原服务验证行为。

本仓库包含一个 Go 服务，在同一进程中运行 go-zero HTTP 和 gRPC。Go 模块路径为 `example.com/iam-service`。当前的 HTTP 示例为 `GET /hello`，gRPC 监听器尚未注册业务方法。

## 仓库布局

- `cmd/iam/` 是可执行程序的入口，负责读取配置、组装依赖，并启动两个监听器。
- `etc/iam.yaml` 保存运行配置，`internal/config/` 定义对应的 Go 配置类型。
- `internal/svc/` 组装服务的共享依赖。
- `internal/handler/` 处理 HTTP 请求和响应，`internal/logic/` 编排 HTTP 用例，`internal/types/` 定义 HTTP 请求和响应类型。
- `internal/httpapi/` 可用于存放共用的 HTTP 协议辅助代码，`internal/rpc/` 存放 gRPC 服务实现。
- `internal/authentication/`、`internal/rbac/`、`internal/session/` 和 `internal/user/` 用于存放 IAM 业务代码。根目录下的 `user/` 包目前供 Hello World 示例使用。

## 实现与迁移

- 在 HTTP 或 gRPC 层处理各自协议的请求解析和响应。将可以复用的 IAM 规则放在 `internal/` 下负责该业务的包中，供两种协议的入口调用。
- 通过构造函数显式传入共享依赖。在调用业务逻辑和访问数据时传递请求的上下文。
- 迁移功能时，先阅读相关目录的 README，再查看原 NestJS 仓库 `C:\Users\vc007\Desktop\github\iam-server`（相对于本仓库为 `../iam-server`）的源码，核对已有行为和公开契约。实现时遵循本仓库的 Go 包布局。
- 除非任务要求修改契约，否则保留现有调用方可观察到的行为。迁移 IAM 功能时，根据原实现核对相关的身份认证、授权和会话行为。

## 开发与验证

在仓库根目录运行以下命令。

```sh
go run ./cmd/iam
go test ./...
```

使用 `gofmt` 格式化改动过的 Go 文件。修改启动流程、配置或路由时，启动服务并验证受影响的入口。默认监听地址定义在 `etc/iam.yaml` 中，可以通过 `-f` 指定其他配置文件。
