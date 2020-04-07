.PHONY: install update run down test docker_test build_linux build_darwin

install:
	go mod init
	bash proto.sh
	go mod tidy

update:
	go get -u github.com/tradingAI/proto
	bash proto.sh
	go mod tidy

test:
	go test -v ./...

docker_test:
	go test -v ./...

run:
	go run main/main.go

down:
	docker-compose -f docker-compose.yml down

build_linux: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr2 build -o client main/main.go

build_darwin: proto
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 packr2 build -o client main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .
