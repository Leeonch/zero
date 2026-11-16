# internal/types

放 go-zero HTTP 请求和响应结构体。目前定义了 `HelloResponse`。从 `.api` 生成代码时，请求和响应类型通常也会生成到这里。

对应原 NestJS 各模块 `src/modules/*/dto/` 中用于 HTTP 的 DTO，例如 `src/modules/user/dto/list-users-query.dto.ts`，以及 `src/common/dto/` 中的公共 HTTP DTO。
