package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/backend_skyshi/src/routers"
)

func main() {
	e, err := routers.New()

	if err != nil {
		log.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to API TODO")
	})

	e.Logger.Fatal(e.Start(":3030"))
}
