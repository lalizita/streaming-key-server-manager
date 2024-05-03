FROM golang:1.22 as builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o stream-key-manager ./cmd/server/main.go

FROM scratch
COPY --from=builder /app/stream-key-manager /stream-key-manager
CMD ["/stream-key-manager"]

EXPOSE 8000