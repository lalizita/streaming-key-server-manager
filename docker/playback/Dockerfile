FROM golang:1.22-alpine as builder

# RUN apk add --no-cache tar

WORKDIR /app

# Define build arguments for GOOS and GOARCH with default values
ARG GOOS=linux
ARG GOARCH=amd64

# Set environment variables for Go to use these arguments
ENV GOOS=${GOOS}
ENV GOARCH=${GOARCH}

COPY . .
RUN CGO_ENABLED=0 go build -o playback ./cmd/playback/main.go

# Download ffmpeg binary
RUN wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-${GOARCH}-static.tar.xz && \
    mkdir ffmpeg && \
    tar -xf ffmpeg-release-* -C ffmpeg --strip-components 1

FROM scratch

COPY --from=builder /app/playback /playback
COPY --from=builder /app/ffmpeg/ffmpeg /usr/bin/
COPY --from=builder /app/ffmpeg/ffprobe /usr/bin/

EXPOSE 8001

ENTRYPOINT ["./playback"]