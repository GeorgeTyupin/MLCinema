package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/routers"
	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TemplateRenderer struct {
	templates *template.Template
}

var DB *gorm.DB

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	DB = db

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("go_server/templates/*.html")),
	}

	e.Renderer = renderer

	routers.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
