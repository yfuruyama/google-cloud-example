FROM golang:1.13 AS builder
WORKDIR /go/src/github.com/yfuruyama/google-cloud-example/cloud_run/grpc_unary
ADD . .
# build server
RUN cd server && CGO_ENABLED=0 go build -o server && cd ../
# build client
RUN cd client && CGO_ENABLED=0 go build -o client && cd ../

FROM alpine:latest
RUN apk update && apk add ca-certificates && update-ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/yfuruyama/google-cloud-example/cloud_run/grpc_unary/server/server /server
COPY --from=builder /go/src/github.com/yfuruyama/google-cloud-example/cloud_run/grpc_unary/client/client /client
ENV PORT 8080
CMD ["/server"]
