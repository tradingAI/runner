# https://github.com/tradingAI/docker/blob/master/bazel.Dockerfile
FROM tradingai/bazel:latest

ENV CODE_DIR /root/github.com/

ARG BUILD_TIME
ENV BUILD_TIME=${BUILD_TIME}

# install tenvs
WORKDIR  $CODE_DIR
RUN cd $CODE_DIR && rm -rf tenvs
RUN git clone https://github.com/tradingAI/tenvs.git
# Clean up pycache and pyc files
RUN cd $CODE_DIR/tenvs && rm -rf __pycache__ && \
    find . -name "*.pyc" -delete && \
    pip install -r requirements.txt && \
    pip install -e .

RUN rm -rf /root/.cache/pip \
    && find / -type d -name __pycache__ -exec rm -r {} \+

WORKDIR $CODE_DIR/tenvs

ARG TUSHARE_TOKEN
ENV TUSHARE_TOKEN=${TUSHARE_TOKEN}
RUN export TUSHARE_TOKEN=$TUSHARE_TOKEN

CMD /bin/bash
