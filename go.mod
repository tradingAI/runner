module github.com/tradingAI/runner

go 1.13

replace github.com/tradingAI/proto/gen/go/tweb => ../proto/gen/go/tweb

replace github.com/tradingAI/proto/gen/go/common => ../proto/gen/go/common

replace github.com/tradingAI/proto/gen/go/model => ../proto/gen/go/model

replace github.com/tradingAI/proto/gen/go/scheduler => ../proto/gen/go/scheduler

replace github.com/docker/docker/internal/testutil => gotest.tools/v3 v3.0.0

require (
	docker.io/go-docker v1.0.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker/internal/testutil v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/minio/minio-go/v6 v6.0.52
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/tradingAI/go v0.0.0-20200405140945-1af5566239dc
	github.com/tradingAI/proto/gen/go/common v0.0.0-00010101000000-000000000000
	github.com/tradingAI/proto/gen/go/model v0.0.0-00010101000000-000000000000
	github.com/tradingAI/proto/gen/go/scheduler v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200403120447-c50568487044 // indirect
)
