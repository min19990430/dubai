package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type CalibrationRepository struct {
	gorm *gorm.DB
}

func NewCalibrationRepository(gorm *gorm.DB) irepository.ICalibrationRepository {
	return &CalibrationRepository{
		gorm: gorm,
	}
}

func (c *CalibrationRepository) FindDetailByPhysicalQuantity(physicalQuantity domain.PhysicalQuantity) ([]domain.CalibrationDetail, error) {
	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	// ! 出現PhysicalQuantity操作
	var physicalQuantities []model.PhysicalQuantity
	err := c.gorm.Table("physical_quantity").
		Where(physicalQuantityWherePO).
		Order("priority").
		Find(&physicalQuantities).Error
	if err != nil {
		return nil, err
	}

	var calibrationDetails []domain.CalibrationDetail
	for _, physicalQuantity := range physicalQuantities {
		calibrationDetails = append(calibrationDetails,
			domain.CalibrationDetail{
				Calibration:      physicalQuantity.ToCalibrationDomain(),
				PhysicalQuantity: physicalQuantity.ToDomain(),
			})
	}

	return calibrationDetails, err
}

func (c *CalibrationRepository) FindDetailByDeviceUUID(deviceUUID string) ([]domain.CalibrationDetail, error) {
	// ! 出現PhysicalQuantity操作
	var physicalQuantities []model.PhysicalQuantity

	err := c.gorm.Table("physical_quantity").
		Where("device_uuid = ?", deviceUUID).
		Order("priority").
		Find(&physicalQuantities).Error
	if err != nil {
		return nil, err
	}

	var calibrationDetails []domain.CalibrationDetail
	for _, physicalQuantity := range physicalQuantities {
		calibrationDetails = append(calibrationDetails,
			domain.CalibrationDetail{
				Calibration:      physicalQuantity.ToCalibrationDomain(),
				PhysicalQuantity: physicalQuantity.ToDomain(),
			})
	}

	return calibrationDetails, err
}

func (c *CalibrationRepository) Update(uuid string, calibrationEnable bool, calibrationValue string, calibrationParameter string) error {
	return c.gorm.Table("physical_quantity").
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"calibration_enable":    calibrationEnable,
			"calibration_value":     calibrationValue,
			"calibration_parameter": calibrationParameter,
		}).Error
}
