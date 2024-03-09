package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/lalizita/streaming-key-server-manager/db"
	"github.com/lalizita/streaming-key-server-manager/internal/handler"
	"github.com/lalizita/streaming-key-server-manager/internal/repository"
	"github.com/lalizita/streaming-key-server-manager/internal/service"
)

func main() {
	//conectar o banco
	db, err := db.OpenConn()
	if err != nil {
		log.Fatalf("Error connect database")
	}

	//init
	keyRepository := repository.NewKeysRepository(db)
	keysService := service.NewKeysService(keyRepository)
	keysHandler := handler.NewHandler(keysService)

	e := echo.New()
	e.GET("/:url_key", keysHandler.GetStreamingKey)
	e.Logger.Fatal(e.Start(":8000"))
}
