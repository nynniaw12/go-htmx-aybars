package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/nynniaw12/go-htmx-aybars/templates"
	"github.com/russross/blackfriday/v2"

	"github.com/labstack/echo/v4"

	"github.com/nynniaw12/go-htmx-aybars/helpers"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		fmt.Print(c.Request().URL)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/about", func(c echo.Context) error {
		component := templates.Portfolio()
		fmt.Print(c.Request().URL)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/contact", func(c echo.Context) error {
		component := templates.Contact()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/projects", func(c echo.Context) error {
		component := templates.Projects()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/sampleproj", func(c echo.Context) error {

		fileContents, err := helpers.ReadFromFile("example.md")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return (err)
		}
        extensions := blackfriday.WithExtensions(
			blackfriday.CommonExtensions |
				blackfriday.NoIntraEmphasis |
				blackfriday.FencedCode |
				blackfriday.Tables |
				blackfriday.Strikethrough |
				blackfriday.SpaceHeadings |
				blackfriday.HeadingIDs |
				blackfriday.BackslashLineBreak |
				blackfriday.DefinitionLists |
				blackfriday.Footnotes |
				blackfriday.AutoHeadingIDs) 
		unsafe := blackfriday.Run(fileContents, extensions)
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		return c.HTML(http.StatusOK, string(html))

	})
	e.Static("/static", "static")
	e.Static("/fonts", "fonts")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":1323"))
}
