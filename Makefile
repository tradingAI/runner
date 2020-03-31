.PHONY: proto build_linux build_darwin clean

up:
	docker-compose -f docker-compose.yml up -d grandet_db minio

down:
	docker-compose -f docker-compose.yml down

proto:
	bash proto.sh

run:
	go run main/client.go

run_docker:
	docker-compose -f docker-compose.yml up -d grandet_db minio tweb

clean:
	cd main && packr2 clean

build_prod_image:
	docker build -f Dockerfile --no-cache -t tradingai/runner:latest .
