package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func serveStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		streamName := c.Param("live")

		dirEntries, err := os.ReadDir("/tmp/hls")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Iterate over directory entries and print directory names
		streamURL := ""
		for _, entry := range dirEntries {
			if entry.IsDir() {
				fmt.Println(entry.Name())
				streamPath, err := findStreamDirectory(entry.Name(), streamName)
				fmt.Println("=========>", streamPath)
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}

				streamURL = fmt.Sprintf("/tmp/hls/%s", streamPath)
				break
			}
		}
		fmt.Println(streamURL)

		return c.File(streamURL)
	}
}

func findLiveStream(c echo.Context) error {
	streamName := c.Param("live")

	dirEntries, err := os.ReadDir("/tmp/hls")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Iterate over directory entries and print directory names
	streamURL := ""
	for _, entry := range dirEntries {
		if entry.IsDir() {
			fmt.Println(entry.Name())
			streamPath, err := findStreamDirectory(entry.Name(), streamName)
			fmt.Println("=========>", streamPath)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			streamURL = fmt.Sprintf("/tmp/hls/%s", streamPath)
			break
		}
	}
	fmt.Println(streamURL)
	return c.String(http.StatusOK, streamURL)
}

func main() {
	log.Default().Println("Routing...")
	e := echo.New()
	// e.Static("/live", "/tmp/hls/live/livetopzera")
	e.GET("/live/:live", serveStream())
	e.GET("/live/:live/*", serveStream())

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})
	e.Logger.Fatal(e.Start(":8001"))
}

func findStreamDirectory(dirName, name string) (string, error) {
	// Split the input string based on '_'
	parts := strings.Split(dirName, "_")

	if len(parts) == 0 || parts[0] == "" {
		return "", errors.New("stream directory not found")
	}

	if parts[0] == name {
		return dirName, nil
	}
	return "", errors.New("error finding stream")
}
