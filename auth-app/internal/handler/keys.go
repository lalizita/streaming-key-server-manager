package handler

import (
	"fmt"
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
	inputLivename := ctx.Request().PostFormValue("name")
	authValues := extractStreamingKeyValues(inputLivename)
	if authValues == nil {
		log.Println("Invalid input livename format:", inputLivename)
		return ctx.JSON(http.StatusBadRequest, "Invalid input livename format")
	}

	keys, err := h.KeysService.AuthStreamingKey(authValues.Name, authValues.Key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error finding stream key")
	}

	if keys.Key == "" {
		log.Println("Forbidden User, live not found:", authValues.Name)
		return ctx.String(http.StatusForbidden, "")
	}

	log.Println("User authenticated, livename:", authValues.Name)

	// According to nginx-rtmp docs the redirect url must use an IP address
	newStreamURL := fmt.Sprintf("rtmp://127.0.0.1:1935/hls-live/%s", authValues.Name)
	log.Println("Redirecting to:", newStreamURL)

	// Respond with a 302 redirect to the new stream URL
	return ctx.Redirect(http.StatusFound, newStreamURL)
}

func extractStreamingKeyValues(inputLivename string) *model.Keys {
	lastUnderscoreIndex := strings.LastIndex(inputLivename, "_")
	if lastUnderscoreIndex == -1 {
		return nil
	}

	return &model.Keys{
		Name: inputLivename[:lastUnderscoreIndex],
		Key:  inputLivename[lastUnderscoreIndex+1:],
	}
}
