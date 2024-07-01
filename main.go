package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sergioglesio/streaming-key-server-manager/config/db"
	"github.com/sergioglesio/streaming-key-server-manager/internal/handler"
	"github.com/sergioglesio/streaming-key-server-manager/internal/repository"
	"github.com/sergioglesio/streaming-key-server-manager/internal/service"
)

func main() {
	db, err := db.OpenConn()
	if err != nil {
		log.Fatalf("Error connect databse main L15")
	}

	//init app
	keysRepository := repository.NewKeyRepository(db)
	keysService := service.NewKeysService(keysRepository)
	keysHandler := handler.NewHandler(keysService)

	log.Default().Println("Routing...")
	e := echo.New()

	e.POST("/auth", keysHandler.AuthStreamingKey)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
