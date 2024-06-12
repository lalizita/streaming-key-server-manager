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
	log.Default().Println("Auth...")
	body := ctx.Request().Body
	defer body.Close()
	fields, _ := io.ReadAll(body)
	log.Println("=====", string(fields))
	authValues := getKeyValues(fields)

	keys, err := h.KeysService.AuthStreamingKey(authValues.Name, authValues.Key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error findind key")
	}

	if keys.Key != "" {
		log.Default().Println("User authenticated")

		log.Println("Name:", authValues.Name)
		log.Println("Addr:", authValues.Addr)

		// According to the nginx rtmp module the redirect url must use an IP address
		// If not set the query parameter "name", at least in my tests, the name of the
		// directory created for hls is not changed
		newStreamURL := fmt.Sprintf("rtmp://%s:1935/hls-live/%s?name=%s", authValues.Addr, authValues.Name, authValues.Name)
		log.Println("Redirecting to:", newStreamURL)

		// Respond with a 302 redirect to the new stream URL
		return ctx.Redirect(http.StatusFound, newStreamURL)
	}

	log.Default().Println("Forbidden User")
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
			// TODO: Fix the following considering that the name can also contain an underscore
			s := strings.Split(value, "_")
			authValues.Name = s[0]
			authValues.Key = s[1]
		}
		if key == "addr" {
			authValues.Addr = value
		}
	}
	return authValues
}
