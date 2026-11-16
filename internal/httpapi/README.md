# internal/httpapi

预留给需要手写、可跨 handler 复用的 HTTP 协议代码，例如 Cookie 读写或统一错误响应。HTTP 路由处理函数放在 [`internal/handler/`](../handler/README.md)。目前此目录只有说明文档。

对应原 NestJS 的 `src/common/filters/`、`src/common/interceptors/`，以及 `src/modules/authentication/refresh-token-cookie.ts` 这类 HTTP 协议辅助代码。
