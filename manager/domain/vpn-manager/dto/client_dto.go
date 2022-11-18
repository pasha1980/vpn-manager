package dto

import (
	"encoding/base64"
	"os"
	"vpn-manager/domain/vpn-manager/entity"
)

type ClientDTO struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
	Client   string `json:"client"`
}

func ClientToDTO(client entity.Client) (*ClientDTO, error) {
	var dto ClientDTO
	dto.ID = client.ID
	dto.FileName = client.ConfigFileName

	data, err := os.ReadFile(client.ConfigPath)
	if err != nil {
		return nil, err
	}
	dto.Client = base64.StdEncoding.EncodeToString(data)
	return &dto, nil
}
