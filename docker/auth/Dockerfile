FROM golang:1.22-alpine as builder

WORKDIR /app

# Define build arguments for GOOS and GOARCH with default values
ARG GOOS=linux
ARG GOARCH=amd64

# Set environment variables for Go to use these arguments
ENV GOOS=${GOOS}
ENV GOARCH=${GOARCH}

COPY . .

RUN CGO_ENABLED=0 go build -o stream-key-manager ./cmd/server/main.go

FROM scratch

COPY --from=builder /app/stream-key-manager /stream-key-manager

CMD ["/stream-key-manager"]

EXPOSE 8000