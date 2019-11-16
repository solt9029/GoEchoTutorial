package api

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(fasthttp.StatusOK, "hello")
	}
}
