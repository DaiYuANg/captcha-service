# 构建阶段（使用 Go 官方 slim 版本）
FROM golang:1.24-alpine AS builder

# 启用 Go module，关闭 CGO
ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

# 安装 Taskfile
RUN apk add --no-cache curl upx && \
    sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

COPY . .

# 使用 task 或者 go build 构建静态二进制
RUN ./bin/task build || go build -trimpath -ldflags="-s -w" -o dist/captcha-service .

RUN upx --best --lzma dist/captcha-service

FROM gcr.io/distroless/static:nonroot  AS distroless

WORKDIR /app

COPY --from=builder /app/dist/captcha-service /app/captcha-service

USER nonroot:nonroot

ENTRYPOINT ["/app/captcha-service"]

FROM scratch AS scratch

COPY --from=builder /app/dist/captcha-service /captcha-service

ENTRYPOINT ["/captcha-service"]

FROM alpine:latest AS alpine
RUN adduser -D -g '' appuser
WORKDIR /app
COPY --from=builder /app/dist/captcha-service /app/captcha-service
USER appuser
ENTRYPOINT ["/app/captcha-service"]

FROM debian:stable-slim AS debian
WORKDIR /app
COPY --from=builder /app/dist/captcha-service /app/captcha-service
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
RUN chmod +x /app/captcha-service
ENTRYPOINT ["/app/captcha-service"]