.PHONY: proto build_linux build_darwin clean


proto:
	bash proto.sh

run:
	go run main/client.go

test:
	docker-compose -f docker-compose.yml up bazel

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .
