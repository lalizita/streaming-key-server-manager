package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lalizita/streaming-key-server-manager/internal/model"
	"github.com/lalizita/streaming-key-server-manager/internal/service"
)

type KeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	KeysService service.KeysService
}

func NewHandler(s service.KeysService) *keysHandler {
	return &keysHandler{
		KeysService: s,
	}
}

func (h *keysHandler) AuthStreamingKey(ctx echo.Context) error {
	body := ctx.Request().Body
	defer body.Close()
	fields, _ := io.ReadAll(body)
	log.Default().Println("Auth...", fields)
	authValues := getKeyValues(fields)

	keys, err := h.KeysService.AuthStreamingKey(authValues.Name, authValues.Key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error findind key")
	}

	if keys.Key == "" {
		log.Default().Println("Forbidden User")
		return ctx.String(http.StatusForbidden, "")
	}

	log.Default().Println("User authenticated")

	newStreamURL := fmt.Sprintf("rtmp://127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to:", newStreamURL)
	return ctx.Redirect(http.StatusFound, newStreamURL)
}

func getKeyValues(s []byte) model.Keys {
	var authValues model.Keys
	pairs := strings.Split(string(s), "&")

	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		key := parts[0]
		value := parts[1]

		if key == "name" {
			s := strings.Split(value, "_")
			authValues.Name = s[0]
			authValues.Key = s[1]
		}
	}
	return authValues
}
