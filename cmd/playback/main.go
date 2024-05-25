package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	log.Default().Println("Routing...")
	e := echo.New()
	e.Static("/live", "/tmp/hls/live/livetopzera")

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})
	e.Logger.Fatal(e.Start(":8001"))
}
