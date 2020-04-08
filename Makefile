.PHONY: install update run proto down test docker_test build_linux build_darwin run_prod

install:
	go mod init
	bash proto.sh
	go mod tidy

update:
	go get -u -v google.golang.org/grpc
	go get -u github.com/tradingAI/proto
	bash proto.sh
	go mod tidy

test:
	go test -v ./...

docker_test:
	docker-compose up test

run:
	go run main/main.go

proto:
	bash proto.sh

down:
	docker-compose down --remove-orphans

build_linux: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o runner main/main.go

build_darwin: proto
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o runner main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .

rm_run:
	docker stop 123456789
	docker rm 123456789
	go run main/main.go

run_prod:
	docker pull registry.cn-hangzhou.aliyuncs.com/tradingai/runner:latest
	docker-compose -f starter/docker-compose-prod.yml up
