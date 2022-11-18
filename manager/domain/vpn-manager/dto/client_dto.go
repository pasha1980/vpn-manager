package dto

import (
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
	dto.FileName = client.FileName

	data, err := os.ReadFile(client.FilePath)
	if err != nil {
		return nil, err
	}
	dto.Client = string(data)
	return &dto, nil
}
