package routers

import (
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.RunServer)
}
