package vpn_manager

import (
	"github.com/labstack/echo/v4"
	"vpn-manager/domain/vpn-manager/dto"
)

func CreateClient(c echo.Context) error {
	serviceParameter := c.Param("service")
	service, err := getService(serviceParameter)
	if err != nil {
		return err
	}

	id := c.Param("id")

	client, err := service.CreateClient(id)
	if err != nil {
		return err
	}

	clientDto, err := dto.ClientToDTO(*client)
	if err != nil {
		return err
	}

	return c.JSON(200, clientDto)
}

func DropClient(c echo.Context) error {
	serviceParameter := c.Param("service")
	service, err := getService(serviceParameter)
	if err != nil {
		return err
	}

	id := c.Param("id")

	err = service.DropClient(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

func RenewClient(c echo.Context) error {
	serviceParameter := c.Param("service")
	service, err := getService(serviceParameter)
	if err != nil {
		return err
	}

	id := c.Param("id")

	client, err := service.RenewClient(id)
	if err != nil {
		return err
	}

	clientDto, err := dto.ClientToDTO(*client)
	if err != nil {
		return err
	}

	return c.JSON(200, clientDto)
}

func CheckStatus(c echo.Context) error {

	var statusDTO dto.ServiceStatus

	for name, manager := range managers {
		statusDTO.Service[name] = manager.HealthCheck()
	}

	return c.JSON(200, statusDTO)
}
