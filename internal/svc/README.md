# internal/svc

放 go-zero 的 `ServiceContext` 和依赖组装代码。当前只注入配置，迁移时在此创建并传入数据库连接、业务服务等依赖。

对应原 NestJS 的 `src/app.module.ts` 以及 `src/modules/` 下各 `*.module.ts` 中注册和连接 Provider 的部分。Go 通过构造函数显式传递依赖。
