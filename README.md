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

## Checklist(TODO, add some checks in regtest)
- [ ] Job正常运行完，应清除runner上所有job的信息，包括日志(ELK上已经保存了日志)
    - evals results
    - infers results
    - logs
    - models
    - progress_bars
    - shells
    - tensorboard events
- [ ] Job运行时内部error
- [ ] 手动kill container
- [ ] 手动kill runner
- [ ] graceful stop runner
