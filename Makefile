.PHONY: proto build_linux build_darwin clean


install:
	GO111MODULE=on go mod init && GO111MODULE=on go mod tidy

proto:
	bash proto.sh

run:
	go run main/client.go

test:
	docker-compose -f docker-compose.yml up bazel

build_linux: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr2 build -o client main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .
