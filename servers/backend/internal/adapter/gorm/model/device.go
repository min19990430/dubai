package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type Device struct {
	UUID        string     `gorm:"column:uuid;type:char(36);primary_key" json:"uuid"`
	ID          string     `gorm:"column:id;type:varchar(45);not null;unique" json:"id"`
	Name        string     `gorm:"column:name;type:varchar(45);not null" json:"name"`
	IsEnable    bool       `gorm:"column:is_enable;type:tinyint;not null;default:0" json:"is_enable"`
	IsConnected bool       `gorm:"column:is_connected;type:tinyint;not null;default:1" json:"is_connected"`
	Priority    int        `gorm:"column:priority;default:0;type:int(11)" json:"priority"`
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime" json:"update_time" `

	Lat         float64 `gorm:"column:lat;type:decimal(10,6);not null" json:"lat"`
	Lon         float64 `gorm:"column:lon;type:decimal(10,6);not null" json:"lon"`
	Description string  `gorm:"column:description;type:text" json:"description"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Device) TableName() string {
	return "device"
}

func (d Device) FromDomain(device domain.Device) Device {
	return Device{
		UUID:        device.UUID,
		ID:          device.ID,
		Name:        device.Name,
		IsEnable:    device.IsEnable,
		IsConnected: device.IsConnected,
		Priority:    device.Priority,
		UpdateTime:  device.UpdateTime,
		Lat:         device.Lat,
		Lon:         device.Lon,
		Description: device.Description,
	}
}

func (d Device) ToDomain() domain.Device {
	return domain.Device{
		UUID:        d.UUID,
		ID:          d.ID,
		Name:        d.Name,
		IsEnable:    d.IsEnable,
		IsConnected: d.IsConnected,
		Priority:    d.Priority,
		UpdateTime:  d.UpdateTime,
		Lat:         d.Lat,
		Lon:         d.Lon,
		Description: d.Description,
	}
}
