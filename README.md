<p align="center">
 <a href="https://roadrunner.dev" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://user-images.githubusercontent.com/7326800/205905278-3899e2c8-5c15-4347-820b-a8ea4c5ba2d7.png">
    <img align="center" src="https://user-images.githubusercontent.com/796136/50286124-6f7f3780-046f-11e9-9f45-e8fedd4f786d.png">
  </picture>
</a>
</p>
<p align="center">
 <a href="https://packagist.org/packages/roadrunner-server/roadrunner"><img src="https://poser.pugx.org/roadrunner-server/roadrunner/version"></a>
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

# Building API

- Install buf:  `go install github.com/bufbuild/buf/cmd/buf@latest`.
- In the repository root run: `buf generate --debug`
