FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o metrics-api ./cmd/api
RUN go build -o metrics-collector ./cmd/collector


FROM alpine:3.18 AS metrics-api
WORKDIR /app
COPY --from=builder /app/metrics-api /metrics-api

# API 서버 포트
EXPOSE 8080
CMD ["/metrics-api"]

FROM alpine:3.18  AS metrics-collector
WORKDIR /app
COPY --from=builder /app/metrics-collector /metrics-collector

CMD ["/metrics-collector"]

