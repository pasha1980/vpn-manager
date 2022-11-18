package api

import (
	"github.com/labstack/echo/v4"
	"vpn-manager/config"
	vpn_manager "vpn-manager/domain/vpn-manager"
)

func InitHttp() error {
	api := echo.New()
	api.HTTPErrorHandler = NewErrorHandler()

	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "PONG")
	})

	privateApi := api.Group("/api/:service", authMiddleware)

	privateApi.POST("/client/:id/new", vpn_manager.CreateClient)
	privateApi.PUT("/client/:id/renew", vpn_manager.RenewClient)
	privateApi.DELETE("/client/:id", vpn_manager.DropClient)
	privateApi.GET("/status", vpn_manager.CheckStatus)

	return api.Start(config.Envs.HttpAddress)
}
