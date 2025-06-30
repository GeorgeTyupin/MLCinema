package routers

import (
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.RunServer)
<<<<<<< HEAD
	e.POST("/", handlers.SearchMovie)
=======
>>>>>>> f6170b3af6dc0bd2d8d208a044dad324e642ccbb
}
