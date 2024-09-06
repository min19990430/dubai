package model

import "auto-monitoring/internal/domain"

type Station struct {
	UUID     string `gorm:"column:uuid;primary_key" json:"uuid"`
	ID       string `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name;NOT NULL" json:"name"`
	Address  string `gorm:"column:address" json:"address"`
	IsEnable bool   `gorm:"column:is_enable" json:"is_enable"`

	Lat         float64 `gorm:"column:lat;default:0.0000000" json:"lat"`
	Lon         float64 `gorm:"column:lon;default:0.0000000" json:"lon"`
	Description string  `gorm:"column:description" json:"description"`
}

func (Station) TableName() string {
	return "station"
}

func (s Station) ToDomain() domain.Station {
	return domain.Station{
		UUID:        s.UUID,
		ID:          s.ID,
		Name:        s.Name,
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
		Address:     station.Address,
		IsEnable:    station.IsEnable,
		Lat:         station.Lat,
		Lon:         station.Lon,
		Description: station.Description,
	}
}
