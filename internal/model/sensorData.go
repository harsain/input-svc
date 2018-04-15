package model

import (
	"time"
)

type SensorData struct {
	ID            string    `json:"id"`
	SensorID      string    `json:"sensor_id"`
	Density       float64   `json:"density"`
	Speed         float64   `json:"speed"`
	UnitOfSpeed   string    `json:"unit_of_speed"`
	UnitOfDensity string    `json:"unit_of_density"`
	Timestamp     time.Time `json:"timestamp"`
	CreatedAt     string    `json:"created_at"`
}
