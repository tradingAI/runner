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
	github.com/NVIDIA/gpu-monitoring-tools v0.0.0-20200403082939-e5b04ac17c03 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker/internal/testutil v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/mholt/archiver/v3 v3.3.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shirou/gopsutil v2.20.3+incompatible
	github.com/stretchr/testify v1.5.1
	github.com/tradingAI/go v0.0.0-20200412172521-d675ba819c87
	github.com/tradingAI/gpu-monitoring-tools v0.0.0-20200414012139-248b8561bf29
	github.com/tradingAI/proto/gen/go/common v0.0.0-00010101000000-000000000000
	github.com/tradingAI/proto/gen/go/model v0.0.0-00010101000000-000000000000
	github.com/tradingAI/proto/gen/go/scheduler v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200413115906-b5235f65be36 // indirect
	google.golang.org/grpc v1.28.1 // indirect
)
