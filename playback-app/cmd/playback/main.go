package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/streaming-key-server-manager/config/env"
	"github.com/sethvargo/go-envconfig"
)

func serveStream(env env.EnvConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")
		filePath := c.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}

		fileStreamPath := filepath.Join(env.StreamBaseFilePath, streamName, filePath)
		log.Println("Stream file requested:", fileStreamPath)

		return c.File(fileStreamPath)
	}
}

func main() {
	ctx := context.Background()

	var envConfig env.EnvConfig
	if err := envconfig.Process(ctx, &envConfig); err != nil {
		log.Fatal(err)
	}

	log.Println("Routing...")
	e := echo.New()
	e.GET("/live/:live", serveStream(envConfig))
	e.GET("/live/:live/*", serveStream(envConfig))

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
