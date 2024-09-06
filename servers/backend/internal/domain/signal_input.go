package domain

import "time"

type SignalInput struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid"`
	Datetime   time.Time `json:"datetime"`
	DI         string    `json:"di"`
	AI         []float64 `json:"ai"`
}
