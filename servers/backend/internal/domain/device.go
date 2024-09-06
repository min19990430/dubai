package domain

import (
	"time"
)

type Device struct {
	UUID        string     `json:"uuid"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	IsEnable    bool       `json:"is_enable"`
	IsConnected bool       `json:"is_connected"`
	Priority    int        `json:"priority"`
	UpdateTime  *time.Time `json:"update_time" `

	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Description string  `json:"description"`

	StationUUID string `json:"station_uuid"`
}

type DeviceStation struct {
	Device
	Station Station `json:"station"`
}
