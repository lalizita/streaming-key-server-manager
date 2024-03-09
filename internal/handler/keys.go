package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lalizita/streaming-key-server-manager/internal/service"
)

type KeysHandler interface {
	GetStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	KeysService service.KeysService
}

func NewHandler(s service.KeysService) *keysHandler {
	return &keysHandler{
		KeysService: s,
	}
}

func (h *keysHandler) GetStreamingKey(ctx echo.Context) error {
	urlAndKey := ctx.Param("url_key")
	str := strings.Split(urlAndKey, "_")

	keys, err := h.KeysService.GetStreamingKey(str[0], str[1])
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error findind key")
	}

	return ctx.JSON(http.StatusOK, keys)
}
