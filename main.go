package main

import (
	"fmt"
	"github.com/hoto/template-go-web-server/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	config.ParseArgsAndFlags()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

type Response struct {
	Message string `json:"name"`
}

func hello(c echo.Context) error {
	r := &Response{
		Message: "Hello world!",
	}
	return c.JSON(http.StatusOK, r)
}
