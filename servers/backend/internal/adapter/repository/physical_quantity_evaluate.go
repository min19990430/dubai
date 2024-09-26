package repository

import (
	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (pqe *PhysicalQuantityEvaluateRepository) UpdateFormula(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate) error {
	physicalQuantityEvaluatePO := model.PhysicalQuantityEvaluate{
		UUID:        physicalQuantityEvaluate.UUID,
		FormulaType: physicalQuantityEvaluate.FormulaType,
		Formula:     physicalQuantityEvaluate.Formula,
	}

	return pqe.gorm.Model(&physicalQuantityEvaluatePO).
		Where("uuid = ?", physicalQuantityEvaluatePO.UUID).
		Updates(map[string]interface{}{
			"formula_type": physicalQuantityEvaluatePO.FormulaType,
			"formula":      physicalQuantityEvaluatePO.Formula,
		}).Error
}

type PhysicalQuantityEvaluateDetailRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityEvaluateDetailRepository(gorm *gorm.DB) irepository.IPhysicalQuantityEvaluateDetailRepository {
	return &PhysicalQuantityEvaluateDetailRepository{gorm: gorm}
}

func (pqed *PhysicalQuantityEvaluateDetailRepository) ListDetail(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate) ([]domain.PhysicalQuantityEvaluateDetail, error) {
	physicalQuantityEvaluateWherePO := model.PhysicalQuantityEvaluate{}.FromDomain(physicalQuantityEvaluate)

	var physicalQuantityEvaluateDetailPOs []model.PhysicalQuantityEvaluateDetail
	err := pqed.gorm.Preload(clause.Associations).
		Where(physicalQuantityEvaluateWherePO).
		Order("priority").
		Find(&physicalQuantityEvaluateDetailPOs).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantityEvaluateDetails []domain.PhysicalQuantityEvaluateDetail
	for _, physicalQuantityEvaluateDetailPO := range physicalQuantityEvaluateDetailPOs {
		physicalQuantityEvaluateDetails = append(physicalQuantityEvaluateDetails, physicalQuantityEvaluateDetailPO.ToDomain())
	}

	return physicalQuantityEvaluateDetails, err
}
