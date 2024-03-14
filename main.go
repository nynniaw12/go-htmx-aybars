package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	// "github.com/microcosm-cc/bluemonday"
	// "github.com/russross/blackfriday/v2"

	"github.com/nynniaw12/go-htmx-aybars/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bytes"
	"regexp"

	"github.com/microcosm-cc/bluemonday"

	"github.com/nynniaw12/go-htmx-aybars/helpers"
	"github.com/yuin/goldmark"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Check if the directory exists
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

		fileContents, err := helpers.ReadFromFile("projects/example.md")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return (err)
		}
		parsedFileContents := parseMarkdown(string(fileContents), c)

		md := goldmark.New()

		var buf bytes.Buffer
		if err := md.Convert([]byte(parsedFileContents), &buf); err != nil {
			panic(err)
		}

		html := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())

        // fix this here seperate into a templ
		formattedHTML := `
    <div class="mx-[20px] mt-[20px] flex flex-col">
        %s
    </div>
`
		styledHTML := strings.ReplaceAll(string(html), "<h1>", "<h1 class=\"text-4xl font-bold mb-4\">")
		styledHTML = strings.ReplaceAll(styledHTML, "<h2>", "<h2 class=\"text-2xl font-bold mb-2\">")
		styledHTML = strings.ReplaceAll(styledHTML, "<img", "<img class=\"max-w-full max-h-96 mx-auto\"")

		return c.HTML(http.StatusOK, fmt.Sprintf(formattedHTML, styledHTML))
	})
	e.Static("/static", "static")
	e.Static("/fonts", "fonts")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":1323"))
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
