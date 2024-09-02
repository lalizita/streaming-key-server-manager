package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
)

func serveStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")
		filePath := c.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}

		fileStreamPath := filepath.Join("/hls/live/", streamName, filePath)
		log.Default().Println("Stream file requested:", fileStreamPath)

		return c.File(fileStreamPath)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/live/:live/*", serveStream())
	e.Logger.Fatal(e.Start(":8001"))
}
