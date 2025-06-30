package handlers

import (
	"net/http"

	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/labstack/echo/v4"
)

func RunServer(c echo.Context) error {
	films := make(map[string]string)
	return c.Render(http.StatusOK, "index.html", films)
}

func SearchMovie(c echo.Context) error {

	mockFilms := []models.Film{
		{
			ID:      1,
			Title:   "Интерстеллар",
			Year:    2014,
			Country: "США",
			Genre:   "Научная фантастика",
			Actors: []models.Actor{
				{Name: "Мэттью МакКонахи"},
				{Name: "Энн Хэтэуэй"},
			},
			Description: "Группа исследователей отправляется в космос, чтобы найти новый дом для человечества.",
		},
		{
			ID:      2,
			Title:   "Гарри Поттер и философский камень",
			Year:    2001,
			Country: "Великобритания",
			Genre:   "Фэнтези",
			Actors: []models.Actor{
				{Name: "Дэниел Рэдклифф"},
				{Name: "Эмма Уотсон"},
			},
			Description: "Мальчик узнаёт, что он волшебник, и отправляется учиться в Хогвартс.",
		},
	}

	return c.JSON(http.StatusOK, mockFilms)
}
