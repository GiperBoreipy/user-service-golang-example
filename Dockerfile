FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/user_service ./cmd


FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata && adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /out/user_service /app/user_service

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/user_service"]
