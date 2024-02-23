package main

import (
	"context"
	"fmt"

	"github.com/nynniaw12/go-htmx-aybars/templates"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
        fmt.Print(c.Request().URL);
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/about", func(c echo.Context) error {
		component :=templates.Portfolio()
        fmt.Print(c.Request().URL);
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/contact", func(c echo.Context) error {
		component :=templates.Contact()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/library", func(c echo.Context) error {
		component :=templates.Library()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/site", func(c echo.Context) error {
		component :=templates.Site()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Static("/fonts","fonts")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":1323"))
}
