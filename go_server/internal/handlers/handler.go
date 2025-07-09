package handlers

import (
	"net/http"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/database"
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func SearchMovie(c echo.Context) error {

	var mockFilms = models.Film{ID: 1}
	database.DB.Preload("Actors").First(&mockFilms)

	return c.JSON(http.StatusOK, mockFilms)
}

func GetFilms(c echo.Context) error {
	var films []models.Film

	database.DB.Preload("Actors").Find(&films)

	return c.JSON(http.StatusOK, films)
}
