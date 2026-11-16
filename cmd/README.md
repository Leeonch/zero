# cmd

放可执行程序的入口。当前只有 [`iam/`](iam/README.md)，负责启动 IAM 服务。

对应原 NestJS 的 `src/main.ts` 所在的应用入口。NestJS 把入口放在 `src/` 中，Go 在这里用 `cmd/` 集中放入口。
