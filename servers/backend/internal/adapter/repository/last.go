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

func (l *LastRepository) GetLast() ([]domain.Last, error) {
	var lastPO []model.Last
	err := l.gorm.Table("last").Table("device").
		Preload("Station").
		Preload("PhysicalQuantities", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_enable = true").Order("physical_quantity.priority")
		}).
		Where("is_enable = true").
		Order("priority").
		Find(&lastPO).Error
	if err != nil {
		return nil, err
	}

	var last []domain.Last
	for _, v := range lastPO {
		var tempLast domain.Last
		tempLast.Station = v.Station.ToDomain()

		tempLast.Device = v.Device.ToDomain()

		for _, pq := range v.PhysicalQuantities {
			tempLast.PhysicalQuantities = append(tempLast.PhysicalQuantities, pq.ToDomain())
		}

		last = append(last, tempLast)
	}
	return last, nil
}
