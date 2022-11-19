package api

import (
	"github.com/labstack/echo/v4"
	"vpn-manager/domain/infrastructure/auth"
	apiError "vpn-manager/domain/infrastructure/error"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if len(token) == 0 {
			token = c.QueryParam("token")
		}

		if !auth.CheckApiToken(token) {
			return apiError.NewAccessDeniedError("Incorrect token")
		}

		return next(c)
	}
}
