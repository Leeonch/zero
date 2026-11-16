# etc

放服务运行时读取的配置文件。目前 `iam.yaml` 配置服务名称、HTTP 地址和端口，以及 gRPC 监听地址。部署时可以用 `-f` 指定其他配置文件。

对应原 NestJS 的 `env/` 及 `src/app.module.ts` 中的 `ConfigModule` 加载方式。Go 侧的配置字段定义在 [`internal/config/`](../internal/config/README.md)。
