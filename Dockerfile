FROM golang:1.21.3

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential protobuf-compiler librdkafka-dev -y && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/ktr0731/evans@latest \
    go install github.com/spf13/cobra-cli@latest


