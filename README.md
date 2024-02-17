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

To install and use generated packages:
```bash
go get github.com/roadrunner-server/api/v4/build/<API_NAME>/v1
```

The Proto API is used for external integrations, mostly for RPC or as internal communications. For example:
```go
package foo

import (
	jobsv1 "github.com/roadrunner-server/api/v4/build/jobs/v1"
)

func Push(in *jobsv1.PushRequest, out *jobsv1.Empty) error {
	return nil
}
```

# Centrifugal API
- [API](https://github.com/centrifugal/centrifugo/blob/master/internal/apiproto/api.proto)
- [Proxy](https://github.com/centrifugal/centrifugo/blob/master/internal/proxyproto/proxy.proto)

# Building API

- Install buf:  `go install github.com/bufbuild/buf/cmd/buf@latest`.
- In the repository root run: `buf generate --debug`
