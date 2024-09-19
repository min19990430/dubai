package domain

import (
	"time"
)

type ControlSignal struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`

	DeviceUUID string `json:"device_uuid"`
	StatusCode string `json:"status_code"`
	IsEnable   bool   `json:"-"`
	Priority   int    `json:"priority"`
	DataType   string `json:"data_type"`

	Description string `json:"description"`

	UpdateTime  *time.Time `json:"update_time"`
	SignalValue string     `json:"signal_value"`
}
