package handlers

import (
	"net/http"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/database"
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/labstack/echo/v4"
)

func RunServer(c echo.Context) error {
	films := make(map[string]string)
	return c.Render(http.StatusOK, "index.html", films)
}

func SearchMovie(c echo.Context) error {

	var mockFilms = models.Film{ID: 1}
	database.DB.Preload("Actors").First(&mockFilms)

	return c.JSON(http.StatusOK, mockFilms)
}
