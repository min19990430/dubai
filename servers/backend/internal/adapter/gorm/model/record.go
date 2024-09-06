package model

import (
	"time"

	"auto-monitoring/internal/domain"
)

type Record struct {
	// StationUUID          string    `gorm:"column:station_uuid;type:char(36)" json:"station_uuid"`
	DeviceUUID           string    `gorm:"column:device_uuid;type:char(36)" json:"device_uuid"`
	PhysicalQuantityUUID string    `gorm:"column:physical_quantity_uuid;primary_key;not null;type:char(36)" json:"physical_quantity_uuid"`
	Datetime             time.Time `gorm:"column:datetime;not null;primary_key;type:datetime" json:"datetime"`
	Value                float64   `gorm:"column:value;default:0.000000;type:decimal(24,6)" json:"value"`
	Data                 float64   `gorm:"column:data;default:0.000000;type:decimal(24,6)" json:"data"`
	Status               string    `gorm:"column:status;not null;type:varchar(4)" json:"status"`
}

func (Record) TableName() string {
	return "records"
}

func (r Record) FromDomain(domain domain.Record) Record {
	return Record{
		// StationUUID:          domain.StationUUID,
		DeviceUUID:           domain.DeviceUUID,
		PhysicalQuantityUUID: domain.PhysicalQuantityUUID,
		Datetime:             domain.Datetime,
		Value:                domain.Value,
		Data:                 domain.Data,
		Status:               domain.Status,
	}
}

func (r Record) ToDomain() domain.Record {
	return domain.Record{
		// StationUUID:          r.StationUUID,
		DeviceUUID:           r.DeviceUUID,
		PhysicalQuantityUUID: r.PhysicalQuantityUUID,
		Datetime:             r.Datetime,
		Value:                r.Value,
		Data:                 r.Data,
		Status:               r.Status,
	}
}
