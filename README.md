<p align="center">
 <img src="https://user-images.githubusercontent.com/796136/50286124-6f7f3780-046f-11e9-9f45-e8fedd4f786d.png" height="75px" alt="RoadRunner">
</p>
<p align="center">
 <a href="https://packagist.org/packages/spiral/roadrunner"><img src="https://poser.pugx.org/spiral/roadrunner/version"></a>
	<a href="https://pkg.go.dev/github.com/roadrunner-server/api/v2?tab=doc"><img src="https://godoc.org/github.com/roadrunner-server/api/v2?status.svg"></a>
	<a href="https://github.com/roadrunner-server/api/actions"><img src="https://github.com/roadrunner-server/api/workflows/Linters/badge.svg" alt=""></a>
	<a href="https://goreportcard.com/report/github.com/roadrunner-server/api"><img src="https://goreportcard.com/badge/github.com/roadrunner-server/api"></a>
	<a href="https://lgtm.com/projects/g/roadrunner-server/api/alerts/"><img alt="Total alerts" src="https://img.shields.io/lgtm/alerts/g/roadrunner-server/api.svg?logo=lgtm&logoWidth=18"/></a>
	<a href="https://discord.gg/TFeEmCs"><img src="https://img.shields.io/badge/discord-chat-magenta.svg"></a>
</p>

# RoadRunner API

## ALL protobuf API located here: [Buf](https://buf.build/roadrunner-server/api)  
To install and use generated packages:
```bash
go get go.buf.build/protocolbuffers/go/roadrunner-server/api
```

Proto API used for the external integrations for the RPC (mostly) or as the internal communications. For example:

```go
package foo

import (
	jobsv1 "go.buf.build/protocolbuffers/go/roadrunner-server/api"
)

func Push(in *jobsv1.PushRequest, out *jobsv1.Empty) error {
	return nil
}
```

Here is the method used in the RR to accept the `Job` from the external system.

You may also navigate to the [`DOCS`](https://buf.build/roadrunner-server/api/docs) tab and inspect the full API. 
