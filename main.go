package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/labstack/echo"
	"github.com/sergioglesio/streaming-key-server-manager/config/db"
	"github.com/sergioglesio/streaming-key-server-manager/internal/handler"
	"github.com/sergioglesio/streaming-key-server-manager/internal/repository"
	"github.com/sergioglesio/streaming-key-server-manager/internal/service"
)

var (
	// Variável global para controlar a permissão de anúncios
	adsEnabled bool
	// Mutex para sincronização segura de variáveis
	mu sync.Mutex
)

func main() {
	db, err := db.OpenConn()
	if err != nil {
		log.Fatalf("Error connect databse main L15")
	}

	// Init app
	keysRepository := repository.NewKeyRepository(db)
	keysService := service.NewKeysService(keysRepository)
	keysHandler := handler.NewHandler(keysService)

	log.Default().Println("Routing...")
	e := echo.New()

	e.POST("/auth", keysHandler.AuthStreamingKey)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})

	// Endpoint para controlar a permissão de anúncios
	e.POST("/toggle-ads", toggleAdsHandler)

	// Endpoint para executar o script insert-ads.sh
	e.POST("/insert-ads", insertAdsHandler)

	e.Logger.Fatal(e.Start(":8000"))
}

// toggleAdsHandler altera o status de adsEnabled com uma requisição POST
func toggleAdsHandler(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	adsEnabled = !adsEnabled
	log.Printf("Ads are now %v\n", adsEnabled)
	return c.String(http.StatusOK, fmt.Sprintf("Ads are now %v", adsEnabled))
}

// insertAdsHandler executa o script insert-ads.sh
func insertAdsHandler(c echo.Context) error {
	cmd := exec.Command("/hls/ads/insert-ads.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run insert-ads.sh: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to insert ads")
	}
	log.Printf("Output: %s", output)
	return c.String(http.StatusOK, "Ads inserted successfully")
}
