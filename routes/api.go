package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/solt9029/GoEchoTutorial/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/hello", api.Hello())
		g.GET("/popper", api.Popper())
	}
}
