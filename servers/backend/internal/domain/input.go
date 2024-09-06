package domain

import "time"

type Input struct {
	UUID     string    `json:"uuid" binding:"required,uuid"`
	Datetime time.Time `json:"datetime"`
	Value    float64   `json:"value"`
}

type InputsWithDeviceUUID struct {
	DeviceUUID string  `json:"device_uuid" binding:"required,uuid"`
	Inputs     []Input `json:"inputs"`
}
