<p align="center">
<a href="https://roadrunner.dev" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/roadrunner-server/.github/assets/8040338/e6bde856-4ec6-4a52-bd5b-bfe78736c1ff">
    <img align="center" src="https://github.com/roadrunner-server/.github/assets/8040338/040fb694-1dd3-4865-9d29-8e0748c2c8b8">
  </picture>
</a>
</p>
<p align="center">
 <a href="https://packagist.org/packages/spiral/roadrunner"><img src="https://poser.pugx.org/spiral/roadrunner/version"></a>
	<a href="https://pkg.go.dev/github.com/roadrunner-server/api/v4?tab=doc"><img src="https://godoc.org/github.com/roadrunner-server/api/v4?status.svg"></a>
	<a href="https://github.com/roadrunner-server/api/actions"><img src="https://github.com/roadrunner-server/api/workflows/Linters/badge.svg" alt=""></a>
	<a href="https://goreportcard.com/report/github.com/roadrunner-server/api"><img src="https://goreportcard.com/badge/github.com/roadrunner-server/api"></a>
	<a href="https://discord.gg/TFeEmCs"><img src="https://img.shields.io/badge/discord-chat-magenta.svg"></a>
</p>

# RoadRunner API

This repository contains the **Protocol Buffer definitions** for [RoadRunner](https://roadrunner.dev). These protos are used for external integrations (RPC) and internal communications between RoadRunner plugins.

Generated Go code lives in a separate repository: [`roadrunner-server/api-go`](https://github.com/roadrunner-server/api-go).

## Repository structure

```
roadrunner/api/   — RoadRunner proto definitions (jobs, kv, http, status, etc.)
third_party/api/  — Temporal API submodule (used as a proto dependency)
buf.yaml          — Buf module configuration
buf.gen.yaml      — Buf code generation configuration
```

## Using generated Go packages

Install a package from the [`api-go`](https://github.com/roadrunner-server/api-go) repository:

```bash
go get github.com/roadrunner-server/api-go/v5/build/<module>/<version>
```

Example usage:

```go
package foo

import (
	jobsv1 "github.com/roadrunner-server/api-go/v5/build/jobs/v1"
)

func Push(in *jobsv1.PushRequest, out *jobsv1.Empty) error {
	return nil
}
```

## Auto-generation

Pushing to `master` in this repo triggers a GitHub Actions workflow in [`api-go`](https://github.com/roadrunner-server/api-go) that:

1. Pulls the latest proto definitions via a git submodule.
2. Runs `buf generate` to produce Go code.
3. Commits and pushes the result.

You do not need to run code generation manually — CI handles it automatically.

## Local development

Install [Buf](https://buf.build/docs/installation):

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
```

Lint proto files:

```bash
buf lint
```

Generate code locally (output goes to `build/`):

```bash
buf generate
```

## Centrifugal API

- [API](https://github.com/centrifugal/centrifugo/blob/master/internal/apiproto/api.proto)
- [Proxy](https://github.com/centrifugal/centrifugo/blob/master/internal/proxyproto/proxy.proto)