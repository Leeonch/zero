# IAM Service

This project is for learning Go through the migration of the existing NestJS `iam-server` to Go. Favor simple, idiomatic Go so the implementation is easy to read and learn from. Migrate one capability at a time and verify its behavior against the original service.

This repository contains a single Go service that runs go-zero HTTP and gRPC in the same process. The module path is `example.com/iam-service`. The current HTTP example is `GET /hello`, and the gRPC listener has no registered application methods yet.

## Repository layout

- `cmd/iam/` is the executable entry point. It loads configuration, wires dependencies, and starts both listeners.
- `etc/iam.yaml` holds runtime settings. `internal/config/` defines the corresponding Go configuration types.
- `internal/svc/` assembles shared dependencies for the service.
- `internal/handler/` handles HTTP requests and responses. `internal/logic/` coordinates HTTP use cases, and `internal/types/` defines HTTP request and response types.
- `internal/httpapi/` is available for shared HTTP protocol helpers. `internal/rpc/` is where gRPC service implementations belong.
- `internal/authentication/`, `internal/rbac/`, `internal/session/`, and `internal/user/` are reserved for IAM business code. The top-level `user/` package currently serves the Hello World example.

## Implementation and migration

- Keep transport-specific parsing and response handling in the HTTP or gRPC layer. Put reusable IAM rules in the owning business package under `internal/` so both transports can call them.
- Pass shared dependencies explicitly through constructors. Propagate request contexts through business calls and data access.
- When migrating a capability, read the relevant directory README and inspect the original NestJS repository at `C:\Users\vc007\Desktop\github\iam-server` (`../iam-server` relative to this repository). Use its source to verify existing behavior and public contracts. Follow this repository's Go package layout.
- Preserve existing caller-facing behavior unless the task calls for a contract change. For IAM migrations, check the relevant authentication, authorization, and session behavior against the source implementation.

## Development and verification

Run these commands from the repository root:

```sh
go run ./cmd/iam
go test ./...
```

Use `gofmt` on changed Go files. For changes to startup, configuration, or routing, start the service and exercise the affected entry point. The default addresses are configured in `etc/iam.yaml`, and `-f` selects an alternate configuration file.
