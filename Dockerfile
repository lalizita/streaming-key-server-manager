FROM golang:1.22
WORKDIR /app

COPY . /app
RUN go build -o stream-key-manager ./cmd/server/main.go

ENTRYPOINT ["./stream-key-manager"]

EXPOSE 8000