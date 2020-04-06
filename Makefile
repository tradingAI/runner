.PHONY: install update proto run down test, docker_test build_linux build_darwin clean

install:
	go mod init
	go mod tidy

update:
	go mod tidy

test:
	go test -v ./...

proto:
	bash proto.sh

run:
	docker-compose -f docker-compose.yml up runner

down:
	docker-compose -f docker-compose.yml down

docker_test:
	docker-compose up test

build_linux: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 packr2 build -o client main/main.go

build_darwin: proto
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 packr2 build -o client main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .
