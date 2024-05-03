package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lalizita/streaming-key-server-manager/config/db"
	"github.com/lalizita/streaming-key-server-manager/config/env"
	"github.com/lalizita/streaming-key-server-manager/internal/handler"
	"github.com/lalizita/streaming-key-server-manager/internal/repository"
	"github.com/lalizita/streaming-key-server-manager/internal/service"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	ctx := context.Background()

	var envConfig env.EnvConfig
	if err := envconfig.Process(ctx, &envConfig); err != nil {
		log.Fatal(err)
	}

	db, err := db.OpenConn(envConfig)
	if err != nil {
		log.Fatalf("Error connect database")
	}

	//init
	keyRepository := repository.NewKeysRepository(db)
	keysService := service.NewKeysService(keyRepository)
	keysHandler := handler.NewHandler(keysService)

	log.Default().Println("Routing...")
	e := echo.New()
	e.POST("/auth", keysHandler.AuthStreamingKey)

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "working")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
