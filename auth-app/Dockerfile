FROM golang:1.22 as builder

WORKDIR /app

# Define build arguments for GOOS and GOARCH with default values
ARG GOOS=linux
ARG GOARCH=amd64

# Set environment variables for Go to use these arguments
ENV GOOS=${GOOS}
ENV GOARCH=${GOARCH}

# Copy and install dependencies first before building
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project and build with mounted cache
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 go build -o stream-key-manager ./cmd/server/main.go


# Final stage
FROM scratch

COPY --from=builder /app/stream-key-manager /stream-key-manager

CMD ["/stream-key-manager"]

EXPOSE 8000