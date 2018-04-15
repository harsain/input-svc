package model

type HealthCheck struct {
	Status    string `json:status`
	CreatedAt string `json:"created_at"`
}
