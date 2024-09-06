package repository

import (
	"errors"

	bigcacheTool "auto-monitoring/internal/adapter/bigcache"
	"auto-monitoring/internal/adapter/gorm/model"

	"github.com/allegro/bigcache/v3"
	"gorm.io/gorm"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

const cachePrefix = "SignalInputMapping:"

type SignalInputMappingRepository struct {
	gorm *gorm.DB
	c    *bigcache.BigCache
}

func NewSignalInputMappingRepository(gorm *gorm.DB, c *bigcache.BigCache) irepository.ISignalInputMappingRepository {
	return &SignalInputMappingRepository{
		gorm: gorm,
		c:    c,
	}
}

func (simr *SignalInputMappingRepository) ListByDeviceUUID(deviceUUID string) ([]domain.SignalInputMapping, error) {
	var signalInputMappingsPO []model.SignalInputMapping
	var signalInputMappings []domain.SignalInputMapping
	var cacheKey = simr.groupName(deviceUUID)

	result, getErr := bigcacheTool.Get(simr.c, cacheKey)
	if getErr != nil {
		findErr := simr.gorm.Where("device_uuid = ?", deviceUUID).Find(&signalInputMappingsPO).Error
		if findErr != nil {
			return nil, findErr
		}

		setErr := bigcacheTool.Set(simr.c, cacheKey, signalInputMappingsPO)
		if setErr != nil {
			return nil, setErr
		}

		for _, signalInputMappingPO := range signalInputMappingsPO {
			signalInputMappings = append(signalInputMappings, signalInputMappingPO.ToDomain())
		}

		return signalInputMappings, nil
	}

	if result != nil {
		sim, ok := result.([]model.SignalInputMapping)
		if !ok {
			return nil, errors.New("failed to convert result to []domain.SignalInputMapping")
		}
		signalInputMappingsPO = sim
	} else {
		findErr := simr.gorm.Where("device_uuid = ?", deviceUUID).Find(&signalInputMappingsPO).Error
		if findErr != nil {
			return nil, findErr
		}

		setErr := bigcacheTool.Set(simr.c, cacheKey, signalInputMappingsPO)
		if setErr != nil {
			return nil, setErr
		}
	}

	for _, signalInputMappingPO := range signalInputMappingsPO {
		signalInputMappings = append(signalInputMappings, signalInputMappingPO.ToDomain())
	}
	return signalInputMappings, nil
}

func (SignalInputMappingRepository) groupName(name string) string {
	return cachePrefix + name
}
