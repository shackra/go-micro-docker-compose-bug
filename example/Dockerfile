FROM golang:1.14-alpine3.11

RUN apk --no-cache add unzip curl git protobuf build-base

# Get the grpc-go source from GitHub
ENV GRPC_GO_VER 1.29.1
RUN mkdir -p /tmp/grpc-go && \
        curl -L https://github.com/grpc/grpc-go/archive/v${GRPC_GO_VER}.zip > /tmp/grpc-go/grpc-go.zip && \
        cd /tmp/grpc-go && \
        unzip grpc-go.zip && \
        mkdir -p /go/src/google.golang.org/grpc/ && \
        cp -r /tmp/grpc-go/grpc-go-${GRPC_GO_VER}/* /go/src/google.golang.org/grpc/

RUN mkdir -p /tmp/protobuf-release && \
        curl -L https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip > /tmp/protobuf-release/protobuf.zip && \
        cd /tmp/protobuf-release && \
        unzip protobuf.zip && \
        cp bin/protoc /usr/local/bin/

# Install protoc-gen-go
ENV GO_PROTOBUF_VER 1.4.2
RUN GO111MODULE=on go get -u github.com/golang/protobuf/protoc-gen-go@v${GO_PROTOBUF_VER} && \
        GO111MODULE=on go get github.com/micro/micro/v2/cmd/protoc-gen-micro@v2.9.2 && \
        GO111MODULE=on go get -u github.com/golang/protobuf/proto

WORKDIR /app

COPY . .

RUN make proto && make build

ENTRYPOINT [ "/app/example-service" ]
