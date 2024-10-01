package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type PhysicalQuantityRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityRepository(gorm *gorm.DB) irepository.IPhysicalQuantityRepository {
	return &PhysicalQuantityRepository{gorm: gorm}
}

func (pq *PhysicalQuantityRepository) FindByUUID(uuid string) (domain.PhysicalQuantity, error) {
	var physicalQuantityPO model.PhysicalQuantity
	err := pq.gorm.Where("uuid = ?", uuid).First(&physicalQuantityPO).Error
	if err != nil {
		return domain.PhysicalQuantity{}, err
	}

	return physicalQuantityPO.ToDomain(), nil
}

func (pq *PhysicalQuantityRepository) List(physicalQuantity domain.PhysicalQuantity) ([]domain.PhysicalQuantity, error) {
	var physicalQuantities []model.PhysicalQuantity

	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	err := pq.gorm.Where(physicalQuantityWherePO).Order("priority").
		Find(&physicalQuantities).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantitiesDomain []domain.PhysicalQuantity
	for _, physicalQuantity := range physicalQuantities {
		physicalQuantitiesDomain = append(physicalQuantitiesDomain, physicalQuantity.ToDomain())
	}

	return physicalQuantitiesDomain, err
}

func (pq *PhysicalQuantityRepository) ListIn(physicalQuantity domain.PhysicalQuantity, uuids []string) ([]domain.PhysicalQuantity, error) {
	var physicalQuantitiesPO []model.PhysicalQuantity

	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	err := pq.gorm.
		Where("uuid IN ?", uuids).
		Where(physicalQuantityWherePO).
		Order("priority").
		Find(&physicalQuantitiesPO).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantitiesDomain []domain.PhysicalQuantity
	for _, physicalQuantityPO := range physicalQuantitiesPO {
		physicalQuantitiesDomain = append(physicalQuantitiesDomain, physicalQuantityPO.ToDomain())
	}

	return physicalQuantitiesDomain, err
}

func (pq *PhysicalQuantityRepository) ListInDeviceUUIDs(physicalQuantity domain.PhysicalQuantity, deviceUUIDs []string) ([]domain.PhysicalQuantity, error) {
	var physicalQuantitiesPO []model.PhysicalQuantity

	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	err := pq.gorm.
		Where("device_uuid IN (?)", deviceUUIDs).
		Where(physicalQuantityWherePO).
		Order("priority").
		Find(&physicalQuantitiesPO).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantitiesDomain []domain.PhysicalQuantity
	for _, physicalQuantityPO := range physicalQuantitiesPO {
		physicalQuantitiesDomain = append(physicalQuantitiesDomain, physicalQuantityPO.ToDomain())
	}

	return physicalQuantitiesDomain, err
}

func (pq *PhysicalQuantityRepository) UpdateLast(physicalQuantity domain.PhysicalQuantity) error {
	physicalQuantityPO := model.PhysicalQuantity{
		UUID:       physicalQuantity.UUID,
		StatusCode: physicalQuantity.StatusCode,
		UpdateTime: physicalQuantity.UpdateTime,
		Value:      physicalQuantity.Value,
		Data:       physicalQuantity.Data,
	}

	return pq.gorm.Model(&physicalQuantityPO).
		Where("uuid = ?", physicalQuantityPO.UUID).
		Updates(map[string]interface{}{
			"status_code": physicalQuantityPO.StatusCode,
			"update_time": physicalQuantityPO.UpdateTime,
			"value":       physicalQuantityPO.Value,
			"data":        physicalQuantityPO.Data,
		}).Error
}

func (pq *PhysicalQuantityRepository) Create(physicalQuantity domain.PhysicalQuantity) error {
	physicalQuantityPO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	return pq.gorm.Create(&physicalQuantityPO).Error
}

func (pq *PhysicalQuantityRepository) Update(physicalQuantity domain.PhysicalQuantity) error {
	physicalQuantityPO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	return pq.gorm.Model(&physicalQuantityPO).
		Where("uuid = ?", physicalQuantityPO.UUID).
		Updates(map[string]interface{}{
			"name":                         physicalQuantityPO.Name,
			"full_name":                    physicalQuantityPO.FullName,
			"si_unit":                      physicalQuantityPO.SiUnit,
			"device_uuid":                  physicalQuantityPO.DeviceUUID,
			"station_uuid":                 physicalQuantityPO.StationUUID,
			"status_code":                  physicalQuantityPO.StatusCode,
			"priority":                     physicalQuantityPO.Priority,
			"physical_quantity_data_type":  physicalQuantityPO.PhysicalQuantityDataType,
			"aggregate_calculation_method": physicalQuantityPO.AggregateCalculationMethod,
			"calibration_enable":           physicalQuantityPO.CalibrationEnable,
			"calibration_value":            physicalQuantityPO.CalibrationValue,
			"calibration_parameter":        physicalQuantityPO.CalibrationParameter,
			"description":                  physicalQuantityPO.Description,
		}).Error
}

func (pq *PhysicalQuantityRepository) Delete(physicalQuantity domain.PhysicalQuantity) error {
	physicalQuantityPO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	return pq.gorm.Delete(&physicalQuantityPO).Error
}

type PhysicalQuantityWithEvaluateRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityWithEvaluateRepository(gorm *gorm.DB) irepository.IPhysicalQuantityWithEvaluateRepository {
	return &PhysicalQuantityWithEvaluateRepository{gorm: gorm}
}

func (pqwe *PhysicalQuantityWithEvaluateRepository) List(physicalQuantity domain.PhysicalQuantity) ([]domain.PhysicalQuantityWithEvaluate, error) {
	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	var physicalQuantityWithEvaluates []model.PhysicalQuantityWithEvaluate
	err := pqwe.gorm.Preload(clause.Associations).
		Where(physicalQuantityWherePO).
		Order("priority").
		Find(&physicalQuantityWithEvaluates).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantityWithEvaluatesDomain []domain.PhysicalQuantityWithEvaluate
	for _, physicalQuantityWithEvaluate := range physicalQuantityWithEvaluates {
		physicalQuantityWithEvaluatesDomain = append(physicalQuantityWithEvaluatesDomain, physicalQuantityWithEvaluate.ToDomain())
	}

	return physicalQuantityWithEvaluatesDomain, nil
}

type PhysicalQuantityCatchDetailRepository struct {
	gorm *gorm.DB
}

func NewPhysicalQuantityCatchDetailRepository(gorm *gorm.DB) irepository.IPhysicalQuantityCatchDetailRepository {
	return &PhysicalQuantityCatchDetailRepository{gorm: gorm}
}

func (pqcd *PhysicalQuantityCatchDetailRepository) FindOne(physicalQuantityCatchDetail domain.PhysicalQuantityCatchDetail) (domain.PhysicalQuantityCatchDetail, error) {
	physicalQuantityCatchDetailWherePO := model.PhysicalQuantityCatchDetail{}.FromDomain(physicalQuantityCatchDetail)

	var physicalQuantityCatchDetailPO model.PhysicalQuantityCatchDetail
	err := pqcd.gorm.Preload(clause.Associations).
		Where(physicalQuantityCatchDetailWherePO).
		Order("priority").
		First(&physicalQuantityCatchDetailPO).Error
	if err != nil {
		return domain.PhysicalQuantityCatchDetail{}, err
	}

	return physicalQuantityCatchDetailPO.ToDomain(), nil
}

func (pqcd *PhysicalQuantityCatchDetailRepository) List(physicalQuantity domain.PhysicalQuantity) ([]domain.PhysicalQuantityCatchDetail, error) {
	physicalQuantityWherePO := model.PhysicalQuantity{}.FromDomain(physicalQuantity)

	var physicalQuantityCatchDetails []model.PhysicalQuantityCatchDetail
	err := pqcd.gorm.Preload(clause.Associations).
		Where(physicalQuantityWherePO).
		Order("priority").
		Find(&physicalQuantityCatchDetails).Error
	if err != nil {
		return nil, err
	}

	var physicalQuantityCatchDetailsDomain []domain.PhysicalQuantityCatchDetail
	for _, physicalQuantityCatchDetail := range physicalQuantityCatchDetails {
		physicalQuantityCatchDetailsDomain = append(physicalQuantityCatchDetailsDomain, physicalQuantityCatchDetail.ToDomain())
	}

	return physicalQuantityCatchDetailsDomain, err
}

func (pq *PhysicalQuantityRepository) UpdateStatus(physicalQuantity domain.PhysicalQuantity, statusCode string) error {
	physicalQuantityPO := model.PhysicalQuantity{
		UUID:       physicalQuantity.UUID,
		StatusCode: statusCode,
	}

	update := pq.gorm.Model(&physicalQuantityPO).
		Where("uuid = ?", physicalQuantityPO.UUID).
		Update("status_code", physicalQuantityPO.StatusCode)
	if update.Error != nil {
		return update.Error
	}
	if update.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
