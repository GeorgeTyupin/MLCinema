package handlers

import (
	"net/http"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/clients"
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/database"
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/labstack/echo/v4"
)

var mlClient = clients.NewMLClient("http://localhost:5000")

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Film(c echo.Context) error {
	return c.Render(http.StatusOK, "film.html", nil)
}

func SearchMovie(c echo.Context) error {
	query := c.FormValue("query")

	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Введите поисковый запрос",
			"code":  1,
		})
	}

	// Поиск через ML
	films, err := mlClient.SearchMovies(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "ML сервис недоступен",
			"code":  2,
		})
	}

	return c.JSON(http.StatusOK, films)
}

func GetFilms(c echo.Context) error {
	var films []models.Film
	database.DB.Preload("Actors").Preload("Categories").Find(&films)
	return c.JSON(http.StatusOK, films)
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCurrentFilm(c echo.Context) error {
	var film models.Film
	filmID := c.FormValue("film_id")
	database.DB.Preload("Actors").First(&film, filmID)
	return c.JSON(http.StatusOK, film)
}
