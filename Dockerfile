# https://github.com/tradingAI/docker/blob/master/bazel.Dockerfile
FROM tradingai/bazel:latest as build

ENV ROOT=/go/src/github.com/tradingAI
ENV PROJECT_PATH=${ROOT}/runner

COPY main ${PROJECT_PATH}/main
COPY client ${PROJECT_PATH}/client
COPY plugins ${PROJECT_PATH}/plugins
COPY Makefile ${PROJECT_PATH}/Makefile
COPY proto.sh ${PROJECT_PATH}/proto.sh
COPY go.mod ${PROJECT_PATH}/go.mod
COPY go.sum ${PROJECT_PATH}/go.sum

WORKDIR ${ROOT}
RUN git clone https://github.com/tradingAI/go.git && \
    git clone https://github.com/tradingAI/proto.git
WORKDIR ${PROJECT_PATH}
RUN make build_linux

# run-time image
FROM alpine

Label maintainer="liuwen.w@qq.com"

ENV PROJECT_PATH=/go/src/github.com/tradingAI/runner

COPY --from=build ${PROJECT_PATH}/main/client /

ENTRYPOINT ["/client"]
