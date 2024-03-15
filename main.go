package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/nynniaw12/go-htmx-aybars/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bytes"
	"regexp"

	"github.com/microcosm-cc/bluemonday"

	"github.com/nynniaw12/go-htmx-aybars/helpers"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	directoryPath := "projects"

	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}

	e.Static("/dir", directoryPath)

	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		fmt.Print(c.Request().URL)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.POST("/mailit", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		message := c.FormValue("message")

		resp, _, err := helpers.SendMail(message, email, name)

		if err != nil {
            fmt.Print(err)
			// return c.String(http.StatusInternalServerError, "Error sending email")
		}

		return c.String(http.StatusOK, resp)
	})
	e.GET("/about", func(c echo.Context) error {
		component := templates.Portfolio()
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
	e.GET("/juxta", func(c echo.Context) error {

		formattedHTML, err := convertMarkdown("projects/juxta.md", c)
		if err != nil {
			fmt.Println("Error converting Markdown:", err)
			return err
		}

		return c.HTML(http.StatusOK, formattedHTML)
	})
	e.GET("/website", func(c echo.Context) error {

		formattedHTML, err := convertMarkdown("projects/website.md", c)
		if err != nil {
			fmt.Println("Error converting Markdown:", err)
			return err
		}

		return c.HTML(http.StatusOK, formattedHTML)
	})
	e.GET("/dotfiles", func(c echo.Context) error {

		formattedHTML, err := convertMarkdown("projects/dotfiles.md", c)
		if err != nil {
			fmt.Println("Error converting Markdown:", err)
			return err
		}

		return c.HTML(http.StatusOK, formattedHTML)
	})
	e.Static("/static", "static")
	e.Static("/fonts", "fonts")
	e.Static("/css", "css")
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":8080"))
}

func parseMarkdown(content string, c echo.Context) string {
	re := regexp.MustCompile(`!\[\[([\w\s\.\-]+)\]\]`)

	currentURL := c.Scheme() + "://" + c.Request().Host

	parsedContent := re.ReplaceAllStringFunc(content, func(match string) string {
		imageName := re.FindStringSubmatch(match)[1]
		imageLink := fmt.Sprintf("![%s](%s/dir/%s)", imageName, currentURL, imageName)
		return imageLink
	})

	return parsedContent
}

func convertMarkdown(filename string, c echo.Context) (string, error) {
	fileContents, err := helpers.ReadFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	parsedFileContents := parseMarkdown(string(fileContents), c)

	md := goldmark.New(
		goldmark.WithExtensions(
			meta.New(meta.WithTable()),
			extension.Table,
			extension.Typographer,
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(parsedFileContents), &buf); err != nil {
		panic(err)
	}

	html := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())

	formattedHTML := `
    <div class="space-y-4 text-white flex flex-col">
        %s
    </div>
`
	styledHTML := strings.ReplaceAll(string(html), "<h1>", "<h1 class=\"text-3xl font-bold mb-4\">")
	styledHTML = strings.ReplaceAll(styledHTML, "<h2>", "<h2 class=\"text-xl font-bold mb-2\">")
	styledHTML = strings.ReplaceAll(styledHTML, "<img", "<img class=\"max-w-full h-auto max-h-80 mx-auto\"")
	styledHTML = strings.ReplaceAll(styledHTML, "<table", "<table class=\"text-white table-auto divide-y divide-gray-700 divide-x divide-gray-700 border-2 border-gray-700\"")
	styledHTML = strings.ReplaceAll(styledHTML, "<p>", "<p class=\"mb-4\">")

	return fmt.Sprintf(formattedHTML, styledHTML), nil

}
