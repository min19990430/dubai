package model

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/domain"
)

type Station struct {
	UUID     string `gorm:"column:uuid;primary_key" json:"uuid"`
	ID       string `gorm:"column:id;not null;unique" json:"id"`
	Name     string `gorm:"column:name;NOT NULL" json:"name"`
	Priority int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`
	Address  string `gorm:"column:address" json:"address"`
	IsEnable bool   `gorm:"column:is_enable" json:"is_enable"`

	Lat         float64 `gorm:"column:lat;default:0.0000000" json:"lat"`
	Lon         float64 `gorm:"column:lon;default:0.0000000" json:"lon"`
	Description string  `gorm:"column:description" json:"description"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Station) TableName() string {
	return "station"
}

func (s Station) ToDomain() domain.Station {
	return domain.Station{
		UUID:        s.UUID,
		ID:          s.ID,
		Name:        s.Name,
		Priority:    s.Priority,
		Address:     s.Address,
		IsEnable:    s.IsEnable,
		Lat:         s.Lat,
		Lon:         s.Lon,
		Description: s.Description,
	}
}

func (s Station) FromDomain(station domain.Station) Station {
	return Station{
		UUID:        station.UUID,
		ID:          station.ID,
		Name:        station.Name,
		Priority:    station.Priority,
		Address:     station.Address,
		IsEnable:    station.IsEnable,
		Lat:         station.Lat,
		Lon:         station.Lon,
		Description: station.Description,
	}
}
