package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
)

type ControlSignalRepository struct {
	gorm *gorm.DB
}

func NewControlSignalRepository(gorm *gorm.DB) *ControlSignalRepository {
	return &ControlSignalRepository{
		gorm: gorm,
	}
}

func (c *ControlSignalRepository) ListIn(controlSignal domain.ControlSignal, uuid []string) ([]domain.ControlSignal, error) {
	var controlSignalPOs []model.ControlSignal

	controlSignalWherePO := model.ControlSignal{}.FromDomain(controlSignal)

	err := c.gorm.Table("control_signal").
		Where("uuid IN ?", uuid).
		Where(controlSignalWherePO).
		Find(&controlSignalPOs).Error
	if err != nil {
		return nil, err
	}

	var controlSignals []domain.ControlSignal
	for _, controlSignalPO := range controlSignalPOs {
		controlSignals = append(controlSignals, controlSignalPO.ToDomain())
	}

	return controlSignals, nil
}

func (c *ControlSignalRepository) UpdateSignalValue(uuid string, signalValue string) error {
	controlSignalPO := model.ControlSignal{
		UUID:        uuid,
		SignalValue: signalValue,
	}

	return c.gorm.Model(&controlSignalPO).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{
			"signal_value": controlSignalPO.SignalValue,
		}).Error
}
