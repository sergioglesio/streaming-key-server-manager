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

type IKeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	keysService service.IKeyService
}

func NewHandler(serv service.IKeyService) IKeysHandler {
	return &keysHandler{
		keysService: serv,
	}
}
func (kh *keysHandler) AuthStreamingKey(c echo.Context) error {
	log.Default().Println("Running auth...")
	body := c.Request().Body
	defer body.Close()

	fields, _ := io.ReadAll(body)
	passedKeyValue := getStreamKey(fields)

	keys, err := kh.keysService.AuthStreamingKey(passedKeyValue.Name, passedKeyValue.Key)
	if err != nil {
		return c.String(http.StatusBadRequest, "problem with streaming key")
	}

	if keys.Key != "" {
		log.Default().Println("User authenticated!")
		return c.String(http.StatusOK, "OK")
	}

	return c.String(http.StatusForbidden, "Forbidden")
}

func getStreamKey(s []byte) model.Keys {
	var authValues model.Keys

	pairs := strings.Split(string(s), "&")

	for _, pair := range pairs {
		splitPair := strings.Split(pair, "=")
		key := splitPair[0]
		value := splitPair[1]

		if key == "name" {
			allPassedValues := strings.Split(value, "_")
			authValues.Name = allPassedValues[0]
			authValues.Key = allPassedValues[1]
		}
	}
	return authValues

}
