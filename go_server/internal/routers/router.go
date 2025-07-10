package routers

import (
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.Index)
	e.POST("/", handlers.SearchMovie)
	e.GET("/film", handlers.Film)
	e.POST("/api/get-films", handlers.GetFilms)
	e.POST("/api/get-categories", handlers.GetCategories)
	e.POST("/api/get-current-film", handlers.GetCurrentFilm)
}
