package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type PhysicalQuantityPresetRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityPresetRepository(gorm *gorm.DB) irepository.IPhysicalQuantityPresetRepository {
	return &PhysicalQuantityPresetRepository{gorm: gorm}
}

func (pqp *PhysicalQuantityPresetRepository) FindOne(physicalQuantityPreset domain.PhysicalQuantityPreset) (domain.PhysicalQuantityPreset, error) {
	physicalQuantityPresetPO := model.PhysicalQuantityPreset{}.FromDomain(physicalQuantityPreset)

	err := pqp.gorm.Where(physicalQuantityPresetPO).First(&physicalQuantityPresetPO).Error
	if err != nil {
		return domain.PhysicalQuantityPreset{}, err
	}

	return physicalQuantityPresetPO.ToDomain(), nil
}

func (pqp *PhysicalQuantityPresetRepository) List(physicalQuantityPreset domain.PhysicalQuantityPreset) ([]domain.PhysicalQuantityPreset, error) {
	var physicalQuantityPresets []model.PhysicalQuantityPreset

	physicalQuantityPresetWherePO := model.PhysicalQuantityPreset{}.FromDomain(physicalQuantityPreset)

	err := pqp.gorm.Where(physicalQuantityPresetWherePO).Order("priority").
		Find(&physicalQuantityPresets).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantityPresetsDomain []domain.PhysicalQuantityPreset
	for _, physicalQuantityPreset := range physicalQuantityPresets {
		physicalQuantityPresetsDomain = append(physicalQuantityPresetsDomain, physicalQuantityPreset.ToDomain())
	}

	return physicalQuantityPresetsDomain, err
}

func (pqp *PhysicalQuantityPresetRepository) Create(physicalQuantityPreset domain.PhysicalQuantityPreset) error {
	physicalQuantityPresetPO := model.PhysicalQuantityPreset{}.FromDomain(physicalQuantityPreset)

	return pqp.gorm.Create(&physicalQuantityPresetPO).Error
}

func (pqp *PhysicalQuantityPresetRepository) Update(physicalQuantityPreset domain.PhysicalQuantityPreset) error {
	physicalQuantityPresetPO := model.PhysicalQuantityPreset{}.FromDomain(physicalQuantityPreset)

	return pqp.gorm.Model(&physicalQuantityPresetPO).
		Where("uuid = ?", physicalQuantityPresetPO.UUID).
		Updates(map[string]interface{}{
			"uuid":                         physicalQuantityPresetPO.UUID,
			"priority":                     physicalQuantityPresetPO.Priority,
			"name":                         physicalQuantityPresetPO.Name,
			"full_name":                    physicalQuantityPresetPO.FullName,
			"si_unit":                      physicalQuantityPresetPO.SiUnit,
			"status_code":                  physicalQuantityPresetPO.StatusCode,
			"physical_quantity_data_type":  physicalQuantityPresetPO.PhysicalQuantityDataType,
			"aggregate_calculation_method": physicalQuantityPresetPO.AggregateCalculationMethod,
			"calibration_enable":           physicalQuantityPresetPO.CalibrationEnable,
			"calibration_value":            physicalQuantityPresetPO.CalibrationValue,
			"calibration_parameter":        physicalQuantityPresetPO.CalibrationParameter,
			"description":                  physicalQuantityPresetPO.Description,
		}).Error
}

func (pqp *PhysicalQuantityPresetRepository) Delete(physicalQuantityPreset domain.PhysicalQuantityPreset) error {
	physicalQuantityPresetPO := model.PhysicalQuantityPreset{}.FromDomain(physicalQuantityPreset)

	return pqp.gorm.Delete(&physicalQuantityPresetPO).Error
}
