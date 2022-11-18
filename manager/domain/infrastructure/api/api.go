package api

import (
	"github.com/labstack/echo/v4"
	"vpn-manager/config"
)

func InitHttp() error {
	server := echo.New()
	server.HTTPErrorHandler = NewErrorHandler()

	server.GET("/ping", func(c echo.Context) error {
		return c.String(200, "PONG")
	})

	// todo

	return server.Start(config.Envs.HttpAddress)
}
