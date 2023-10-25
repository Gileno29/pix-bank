FROM golang:1.15
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLE=1

RUN apt-get update && \
    apt-get install build-essencial protobuf-compiler librdkafka-dev -y && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get google.golang.org/protobuf/cmd/proto-gen-go && \
    go get github.com/spf13/cobra/cobra && \
    wget https://github.com/spf13/cobra/cobra && \
    wget https://github.com/ktr8731/evans/realeases/download/8.9.1/evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    mv evans ../bin && rm -f evans_linux_amd64.tar.gz

CMD ["tail", "-f", "/dev/null"]
