package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/auth", func(c echo.Context) error {
		log.Default().Println("Running auth...")
		body := c.Request().Body
		defer body.Close()

		fields, _ := io.ReadAll(body)
		fmt.Println(string(fields))

		return c.String(http.StatusOK, "WORKING")
	})
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
