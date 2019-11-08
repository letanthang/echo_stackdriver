package middleware

import (
	"github.com/labstack/echo"
)

const HeaderKey = "X-Request-Id"

func DoSomething(h echo.HandlerFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := h(c)
			// c.Response().Header().Set(HeaderKey, id)
			return err
		}
	}
}
