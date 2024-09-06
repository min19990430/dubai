package domain

import "time"

type Record struct {
	// StationUUID          string    `json:"station_uuid"`
	DeviceUUID           string    `json:"device_uuid"`
	PhysicalQuantityUUID string    `json:"physical_quantity_uuid"`
	Datetime             time.Time `json:"datetime"`
	Value                float64   `json:"value"`
	Data                 float64   `json:"data"`
	Status               string    `json:"status"`
}
