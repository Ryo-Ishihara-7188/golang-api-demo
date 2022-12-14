package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go Echo")
	})
	e.Logger.Fatal(e.Start(":8080"))
}