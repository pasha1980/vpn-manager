package vpn_manager

import (
	apiError "vpn-manager/domain/infrastructure/error"
	"vpn-manager/domain/openvpn"
	"vpn-manager/domain/vpn-manager/entity"
)

type VPNServiceManager interface {
	CreateClient(id string) (*entity.Client, error)
	GetClient(id string) (*entity.Client, error)
	DropClient(id string) error
	RenewClient(id string) (*entity.Client, error)

	HealthCheck() bool
	Autofix() error
}

var managers = map[string]VPNServiceManager{
	"openvpn": &openvpn.Service{},
}

func getService(name string) (VPNServiceManager, error) {
	service := managers[name]
	if service == nil {
		return nil, apiError.NewNotFoundError("Service "+name+" not found", nil)
	}

	return service, nil
}
