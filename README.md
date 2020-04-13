![Test](https://github.com/tradingAI/runner/workflows/Test/badge.svg)
![Docker](https://github.com/tradingAI/runner/workflows/Docker/badge.svg)
# runner
Job runner

**[Design](https://github.com/tradingAI/scheduler/blob/master/docs/README.md)**

## 快速开始
- golang 1.13+
- export GO111MODULE=on
- install: `make intall`
- test: `make test`

## 部署(Inprogress)
- mac os 目前只能支持binary + shell方式部署, [why?](docs/mac_machine_info.md)
    - 下载binary
    - run.sh(TODO:wen)
- linux
    - 配置环境变量:(TODO:wen)
    - docker-compose -f starter/docker-compose-prod up -d
