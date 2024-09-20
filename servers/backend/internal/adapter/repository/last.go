package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type LastRepository struct {
	gorm *gorm.DB
}

func NewLastRepository(gorm *gorm.DB) irepository.ILastRepository {
	return &LastRepository{
		gorm: gorm,
	}
}

func (l *LastRepository) GetStationLast(source string) ([]domain.StationLast, error) {
	var lastPO []model.StationLast
	err := l.gorm.
		Preload("PhysicalQuantities", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_enable = true").Where("source = ?", source).Order("physical_quantity.priority")
		}).
		Order("priority").
		Find(&lastPO).Error
	if err != nil {
		return nil, err
	}

	var last []domain.StationLast
	for _, v := range lastPO {
		var tempLast domain.StationLast
		tempLast.Station = v.Station.ToDomain()

		for _, pq := range v.PhysicalQuantities {
			tempLast.PhysicalQuantities = append(tempLast.PhysicalQuantities, pq.ToDomain())
		}

		last = append(last, tempLast)
	}
	return last, nil
}

func (l *LastRepository) GetDeviceLast(source string) ([]domain.DeviceLast, error) {
	var lastPO []model.DeviceLast
	err := l.gorm.
		Preload("PhysicalQuantities", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_enable = true").Where("source = ?", source).Order("physical_quantity.priority")
		}).
		Order("priority").
		Find(&lastPO).Error
	if err != nil {
		return nil, err
	}

	var last []domain.DeviceLast
	for _, v := range lastPO {
		var tempLast domain.DeviceLast
		tempLast.Device = v.Device.ToDomain()

		for _, pq := range v.PhysicalQuantities {
			tempLast.PhysicalQuantities = append(tempLast.PhysicalQuantities, pq.ToDomain())
		}

		last = append(last, tempLast)
	}
	return last, nil
}
