package dto

type ServiceStatus struct {
	Service map[string]bool `json:"service"`
	Version string          `json:"version"`
}
