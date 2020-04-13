.PHONY: install update run proto down test dtest build_linux build_darwin prod

install:
	go mod init
	bash proto.sh
	go mod tidy

update:
	go get -u -v google.golang.org/grpc
	go get -u github.com/tradingAI/proto
	go get -u github.com/tradingAI/go
	bash proto.sh
	go mod tidy

test:
	rm -rf /tmp/runner/data/models/*
	rm -rf /tmp/runner/data/tensorboards/*
	cp -R ./runner/testdata/upload/model/* /tmp/runner/data/models/
	cp -R ./runner/testdata/upload/tensorboard/* /tmp/runner/data/tensorboards/
	go test ./...

# docker test
dtest:
	# docker-compose down
	rm -rf /tmp/runner/minio/data/*
	docker-compose up test

run:
	go run main/main.go

proto:
	bash proto.sh

down:
	docker-compose down --remove-orphans

build_linux: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main/runner main/main.go

build_darwin: proto
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o main/runner main/main.go

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .

rm_run:
	docker stop 123456789
	docker rm 123456789
	go run main/main.go

prod:
	docker pull registry.cn-hangzhou.aliyuncs.com/tradingai/runner:latest
	docker-compose -f starter/docker-compose-prod.yml up
