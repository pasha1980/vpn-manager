package openvpn

import "vpn-manager/domain/vpn-manager/entity"

type Service struct {
}

func (s *Service) CreateClient(id string) (*entity.Client, error) {

}

func (s *Service) GetClient(id string) (*entity.Client, error) {

}

func (s *Service) DropClient(id string) error {

}

func (s *Service) RenewClient(id string) (*entity.Client, error) {

}

func (s *Service) HealthCheck() bool {

}

func (s *Service) Autofix() error {

}
