package handler

import (
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
	log.Default().Println("Auth...")
	body := ctx.Request().Body
	defer body.Close()
	fields, _ := io.ReadAll(body)
	authValues := getKeyValues(fields)

	keys, err := h.KeysService.AuthStreamingKey(authValues.Name, authValues.Key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error findind key")
	}

	if keys.Key != "" {
		return ctx.String(http.StatusOK, "Ok")
	}

	return ctx.String(http.StatusForbidden, "")
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
