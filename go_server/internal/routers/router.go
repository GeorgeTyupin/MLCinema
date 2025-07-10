package routers

import (
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.Index)
	e.POST("/", handlers.SearchMovie)
	e.GET("/film", handlers.Film)

	api := e.Group("/api")

	api.POST("/get-films", handlers.GetFilms)
	api.POST("/get-categories", handlers.GetCategories)
	api.POST("/get-current-film", handlers.GetCurrentFilm)
}
