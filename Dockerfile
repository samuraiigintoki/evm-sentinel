# Stage 1: Build binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /evm-sentinel cmd/evmsentinel/main.go

# Stage 2: Minimal scratch runtime
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
COPY --from=builder /evm-sentinel /usr/local/bin/evm-sentinel
ENTRYPOINT ["evm-sentinel"]
CMD ["--dir", "/contracts"]
