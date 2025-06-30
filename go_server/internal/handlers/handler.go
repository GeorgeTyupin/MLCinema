package handlers

import (
	"net/http"

<<<<<<< HEAD
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
=======
>>>>>>> f6170b3af6dc0bd2d8d208a044dad324e642ccbb
	"github.com/labstack/echo/v4"
)

func RunServer(c echo.Context) error {
	return c.String(http.StatusOK, "Сервер работает!")
}
<<<<<<< HEAD

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
=======
>>>>>>> f6170b3af6dc0bd2d8d208a044dad324e642ccbb
