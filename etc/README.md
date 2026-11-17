# etc

放服务运行时读取的配置文件。目前 `iam.yaml` 配置服务名称、HTTP 地址和端口，以及 gRPC 监听地址。部署时可以用 `-f` 指定其他配置文件。

需要连接 MySQL 时，将 `iam.yaml` 复制为本地配置，并加入以下配置。默认端口为 `3306`，默认启用 TLS 并验证服务器证书。

```yaml
MySQL:
  Host: db.example.com
  Port: 3306
  User: iam_reader
  Database: iam_test
```

测试环境使用未受本机信任的证书时，可以在本地配置中设置 `TLSSkipVerify: true`。连接仍使用 TLS，此设置会跳过服务器证书校验。

在 PowerShell 中将密码读入当前会话的环境变量，再用本地配置启动。

```powershell
$env:IAM_MYSQL_PASSWORD = Read-Host -MaskInput 'MySQL 密码'
go run ./cmd/iam -f <本地配置路径>
```

对应原 NestJS 的 `env/` 及 `src/app.module.ts` 中的 `ConfigModule` 加载方式。Go 侧的配置字段定义在 [`internal/config/`](../internal/config/README.md)。
