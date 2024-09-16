package repository

import (
	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"

	"gorm.io/gorm"
)

type PhysicalQuantityEvaluateRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityEvaluateRepository(gorm *gorm.DB) irepository.IPhysicalQuantityEvaluateRepository {
	return &PhysicalQuantityEvaluateRepository{gorm: gorm}
}

func (pqe *PhysicalQuantityEvaluateRepository) UpdateLast(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate) error {
	physicalQuantityEvaluatePO := model.PhysicalQuantityEvaluate{
		UUID:       physicalQuantityEvaluate.UUID,
		UpdateTime: physicalQuantityEvaluate.UpdateTime,
		Value:      physicalQuantityEvaluate.Value,
		Data:       physicalQuantityEvaluate.Data,
	}

	return pqe.gorm.Model(&physicalQuantityEvaluatePO).
		Where("uuid = ?", physicalQuantityEvaluatePO.UUID).
		Updates(map[string]interface{}{
			"update_time": physicalQuantityEvaluatePO.UpdateTime,
			"value":       physicalQuantityEvaluatePO.Value,
			"data":        physicalQuantityEvaluatePO.Data,
		}).Error
}
