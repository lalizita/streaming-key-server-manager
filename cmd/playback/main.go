package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// serveStream handles the streaming requests
func serveStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")

		dirEntries, err := os.ReadDir("/tmp/hls")
		if err != nil {
			log.Printf("Failed to read directory: %v", err)
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		// Find the directory matching the pattern "nameoflive_*"
		var streamURL string
		for _, entry := range dirEntries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), streamName+"_") {
				streamPath := filepath.Join("/tmp/hls", entry.Name(), c.Param("*"))
				streamURL = streamPath
				break
			}
		}

		if streamURL == "" {
			return c.String(http.StatusNotFound, "Stream not found")
		}

		return c.File(streamURL)
	}
}

func main() {
	e := echo.New()
	log.Println("Routing...")

	e.GET("/live/:live", serveStream())
	e.GET("/live/:live/*", serveStream())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "working")
	})

	if err := e.Start(":8001"); err != nil {
		e.Logger.Fatal("Shutting down the server: ", err)
	}
}
