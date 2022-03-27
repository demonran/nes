FROM golang:1.17.5-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ARG REGION=CN
ARG VERSION=unknown
ARG GIT_COMMIT_HASH=unknown
ARG BUILD_TIME=unknown

WORKDIR /build
RUN if [ "$REGION" = "CN" ]; then sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories; fi && \
    apk add --no-cache upx

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN if [ "$REGION" = "CN" ]; then export GOPROXY="https://goproxy.cn"; fi && \
    go mod download

COPY . .
RUN go build -ldflags  -a -o nes cmd/server/main.go && \
    upx nes

FROM alpine:3
WORKDIR /app
# current latest version: https://dl-cdn.alpinelinux.org/alpine/v3.14/main/
RUN if [ "$REGION" = "CN" ]; then sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories; fi && \
    apk add --no-cache curl mysql-client # used by docker health check and check mysql data
COPY --from=builder /build/nes .

COPY cmd/server/*.yml /app/
COPY entrypoint.sh /entrypoint.sh

CMD ["/bin/sh", "/entrypoint.sh"]
