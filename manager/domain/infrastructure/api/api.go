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

	privateApi := api.Group("/api", authMiddleware)
	privateApi.GET("/status", vpn_manager.CheckStatus)

	serviceApi := privateApi.Group("/:service")
	serviceApi.POST("/client/:id", vpn_manager.CreateClient)
	//serviceApi.PUT("/client/:id/renew", vpn_manager.RenewClient)
	serviceApi.DELETE("/client/:id", vpn_manager.DropClient)

	return api.Start(config.Envs.HttpAddress)
}
