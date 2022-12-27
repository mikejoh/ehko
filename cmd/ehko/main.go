package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikejoh/ehko"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "<pre><h3>ehko</h3></pre>")
	})
	e.POST("/alerts", ehko.Alerts)
	e.GET("/responder/:code", ehko.Responder)

	e.Logger.Fatal(e.Start(":5001"))
}