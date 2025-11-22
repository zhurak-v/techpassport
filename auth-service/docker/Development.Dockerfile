FROM golang:1.24.5-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

RUN go install github.com/google/wire/cmd/wire@latest

COPY . .

RUN wire ./src/infrastructure/application

RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application ./src/application/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/application /application
RUN chmod +x /application

COPY --from=builder /app/src/infrastructure/database/migrations /app/migrations

EXPOSE 1111

ENTRYPOINT ["/application"]

