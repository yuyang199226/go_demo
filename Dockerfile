FROM golang:1.13.8 AS buildimage

ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN mkdir -p /go/src/code.aibee.cn/mlp/store_business_engine
COPY . /go/src/code.aibee.cn/mlp/store_business_engine/

WORKDIR $GOPATH/src/code.aibee.cn/mlp/store_business_engine
RUN make install


FROM ubuntu:16.04

COPY --from=buildimage /go/bin/store_business_engine ./
COPY config/yaml ./config/yaml
COPY test/resource ./test/resource

RUN apt-get -qq update && apt-get install -y curl vim iputils-ping

# RUN wget -nv https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip && \
#    unzip -o protoc-3.3.0-linux-x86_64.zip -d /usr/local bin/protoc && \
#    rm -f protoc-3.3.0-linux-x86_64.zip

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get install -y tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata

#RUN go get -d -u github.com/golang/protobuf/protoc-gen-go && \
#    git -C /go/src/github.com/golang/protobuf checkout "v1.2.0" && \
#    go install github.com/golang/protobuf/protoc-gen-go







