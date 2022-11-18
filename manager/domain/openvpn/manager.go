package openvpn

import (
	"os"
	"os/exec"
	"strings"
	"vpn-manager/config"
	apiError "vpn-manager/domain/infrastructure/error"
	"vpn-manager/domain/vpn-manager/entity"
)

type Service struct {
}

func (s *Service) CreateClient(id string) (*entity.Client, error) {
	existingClient, err := s.GetClient(id)
	if err != nil {
		return nil, err
	}

	if existingClient != nil {
		if existingClient.IsActive {
			return existingClient, nil
		} else {
			return nil, apiError.NewBadRequestError("Client is removed", nil)
		}
	}

	output, err := s.executeScript("new-client", id)
	if err != nil {
		return nil, err
	}

	client := entity.Client{
		ID:             id,
		ConfigPath:     output,
		ConfigFileName: id + ".ovpn",
		IsActive:       true,
	}

	return &client, nil
}

func (s *Service) GetClient(id string) (*entity.Client, error) {
	dirData, err := os.ReadDir(config.Envs.OpenvpnDataDir + "/clients")
	if err != nil {
		return nil, err
	}

	for _, element := range dirData {
		if element.IsDir() {
			continue
		}

		if element.Name() == id+".ovpn" {
			client := entity.Client{
				ID:             id,
				ConfigPath:     config.Envs.OpenvpnDataDir + "/clients/" + element.Name(),
				ConfigFileName: element.Name(),
				IsActive:       true,
			}
			return &client, nil
		}
	}

	dirData, err = os.ReadDir(config.Envs.OpenvpnDataDir + "/removed")
	if err != nil {
		return nil, err
	}

	for _, element := range dirData {
		if element.IsDir() {
			continue
		}

		if element.Name() == id+".ovpn" {
			client := entity.Client{
				ID:             id,
				ConfigPath:     config.Envs.OpenvpnDataDir + "/removed/" + element.Name(),
				ConfigFileName: element.Name(),
				IsActive:       false,
			}
			return &client, nil
		}
	}

	return nil, nil
}

func (s *Service) DropClient(id string) error {
	existingClient, err := s.GetClient(id)
	if err != nil {
		return err
	}

	if existingClient == nil {
		return nil
	}

	if !existingClient.IsActive {
		return nil
	}

	_, err = s.executeScript("remove-client", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RenewClient(id string) (*entity.Client, error) {
	existingClient, err := s.GetClient(id)
	if err != nil {
		return nil, err
	}

	if existingClient == nil {
		return nil, apiError.NewNotFoundError("Client not found", nil)
	}

	if existingClient != nil && existingClient.IsActive {
		return existingClient, nil
	}

	output, err := s.executeScript("renew-client", id)
	if err != nil {
		return nil, err
	}

	client := entity.Client{
		ID:             id,
		ConfigPath:     output,
		ConfigFileName: id + ".ovpn",
		IsActive:       true,
	}

	return &client, nil
}

func (s *Service) HealthCheck() bool {
	output, err := s.executeScript("health-check")
	if err != nil {
		return false
	}

	if output == "OK" {
		return true
	}

	return false
}

func (s *Service) Autofix() error {
	_, err := s.executeScript("autofix")
	return err
}

func (s *Service) executeScript(script string, parameters ...string) (string, error) {
	scriptFile := config.Envs.OpenvpnScriptDir + "/" + script + ".sh"
	cmd := exec.Command(scriptFile, strings.Join(parameters, " "))
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	outputStr := string(output)
	outputStr = strings.Replace(outputStr, "\n", "", -1)

	return outputStr, nil
}
