.PHONY: install update run proto down test docker_test build_linux build_darwin

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
	go test -v ./...

run:
	go run main/main.go

proto:
	bash proto.sh

down:
	docker-compose -f docker-compose.yml down

build_linux: proto
	GOOS=linux GOARCH=amd64 go build -o client main/main.go

build_darwin: proto
	GOOS=darwin GOARCH=amd64 go build -o client main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .

rm_run:
	docker stop 123456789
	docker rm 123456789
	go run main/main.go
