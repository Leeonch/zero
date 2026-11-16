# internal/config

放配置结构体和配置校验。目前 `config.go` 声明 go-zero HTTP 配置和 `GRPCListenOn` 字段。配置值由入口从 [`etc/iam.yaml`](../../etc/iam.yaml) 加载。

对应原 NestJS 的 `src/app.module.ts` 中 `ConfigModule` 和数据库配置工厂，以及 `src/main.ts` 中通过 `ConfigService` 读取 CORS、gRPC 地址等配置的部分。
