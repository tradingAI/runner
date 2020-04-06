module github.com/tradingAI/runner

go 1.13

replace github.com/tradingAI/proto/gen/go/tweb => ../proto/gen/go/tweb

replace github.com/tradingAI/proto/gen/go/common => ../proto/gen/go/common

replace github.com/tradingAI/proto/gen/go/model => ../proto/gen/go/model

replace github.com/tradingAI/proto/gen/go/scheduler => ../proto/gen/go/scheduler

replace github.com/docker/docker/internal/testutil => gotest.tools/v3 v3.0.0

require (
	4d63.com/gochecknoinits v0.0.0-20200108094044-eb73b47b9fc4 // indirect
	docker.io/go-docker v1.0.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/alecthomas/gocyclo v0.0.0-20150208221726-aa8f8b160214 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker/internal/testutil v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gordonklaus/ineffassign v0.0.0-20200309095847-7953dde2c7bf // indirect
	github.com/jgautheron/goconst v0.0.0-20200227150835-cda7ea3bf591 // indirect
	github.com/mdempsky/maligned v0.0.0-20180708014732-6e39bd26a8c8 // indirect
	github.com/mibk/dupl v1.0.0 // indirect
	github.com/minio/minio-go/v6 v6.0.52
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opennota/check v0.0.0-20180911053232-0c771f5545ff // indirect
	github.com/securego/gosec v0.0.0-20200401082031-e946c8c39989 // indirect
	github.com/tradingAI/go v0.0.0-20200405140945-1af5566239dc
	github.com/tradingAI/proto/gen/go/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/tradingAI/proto/gen/go/scheduler v0.0.0-00010101000000-000000000000
	github.com/walle/lll v1.0.1 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	google.golang.org/genproto v0.0.0-20200403120447-c50568487044 // indirect
	mvdan.cc/interfacer v0.0.0-20180901003855-c20040233aed // indirect
	mvdan.cc/lint v0.0.0-20170908181259-adc824a0674b // indirect
)
