package main

import (
	"html/template"
	"io"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/database"
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/routers"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	database.InitDB()

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("go_server/templates/*.html")),
	}

	e.Renderer = renderer

	routers.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
