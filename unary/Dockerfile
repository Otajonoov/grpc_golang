FROM golang:1.21.1

WORKDIR /go/src/grpc_golang/unary/server

RUN go mod init grpc_golang/unary/server

ADD . /go/src/grpc_golang/unary/server

RUN go install grpc_golang/unary/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 8080
