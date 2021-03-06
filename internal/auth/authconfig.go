package auth

import (
	"github.com/labstack/echo/v4"
)

func Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add basic auth
		return next(c)
	}
}
