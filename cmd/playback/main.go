package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func serveStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")
		filePath := c.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}

		fileStreamPath := filepath.Join("/tmp/hls/", streamName, filePath)
		log.Println("Stream file requested:", fileStreamPath)

		return c.File(fileStreamPath)
	}
}

func main() {
	log.Default().Println("Routing...")
	e := echo.New()
	e.GET("/live/:live", serveStream())
	e.GET("/live/:live/*", serveStream())

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
