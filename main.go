package main

import (
	"context"
	"github.com/nynniaw12/go-htmx-aybars/templates"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/init", func(c echo.Context) error {
		component :=templates.Portfolio()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Static("/fonts","fonts")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":1323"))
}
