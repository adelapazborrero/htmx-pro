package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemmplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count  int
	Count2 int
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	count := Count{Count: 0, Count2: 0}
	e.Renderer = NewTemmplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.POST("/count2", func(c echo.Context) error {
		count.Count2++
		return c.Render(200, "count2", count)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
