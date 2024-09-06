package repository

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	linenotify "auto-monitoring/internal/adapter/line-notify"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmSettingRepository struct {
	gorm       *gorm.DB
	linenotify *linenotify.LineNotify
}

func NewAlarmSettingRepository(gorm *gorm.DB, linenotify *linenotify.LineNotify) irepository.IAlarmSettingRepository {
	return &AlarmSettingRepository{
		gorm:       gorm,
		linenotify: linenotify,
	}
}

func (a *AlarmSettingRepository) UpdateStatusAndTime(alarmSetting domain.AlarmSetting) error {
	alarmSettingPO := model.AlarmSetting{
		UUID:               alarmSetting.UUID,
		IsActivated:        alarmSetting.IsActivated,
		LastOccurValue:     alarmSetting.LastOccurValue,
		LastAlarmOccurTime: alarmSetting.LastAlarmOccurTime,
	}

	return a.gorm.Model(&alarmSettingPO).
		Where("uuid = ?", alarmSettingPO.UUID).
		Updates(map[string]interface{}{
			"is_activated":          alarmSettingPO.IsActivated,
			"last_occur_value":      alarmSettingPO.LastOccurValue,
			"last_alarm_occur_time": alarmSettingPO.LastAlarmOccurTime,
		}).Error
}

func (a *AlarmSettingRepository) Notify(alarmSetting domain.AlarmSetting, content string) error {
	if alarmSetting.LastAlarmOccurTime == nil {
		alarmSetting.LastAlarmOccurTime = new(time.Time)
	}

	a.linenotify.Message.UpdateTime = *alarmSetting.LastAlarmOccurTime
	a.linenotify.Message.ID = ""
	a.linenotify.Message.Name = alarmSetting.Name
	a.linenotify.Message.Event = content

	return a.linenotify.Send()
}

func (a *AlarmSettingRepository) UpdateExpression(uuid string, expression string) error {
	alarmSettingPO := model.AlarmSetting{
		UUID:              uuid,
		BooleanExpression: expression,
	}

	return a.gorm.Model(&alarmSettingPO).
		Where("uuid = ?", alarmSettingPO.UUID).
		Updates(map[string]interface{}{
			"boolean_expression": alarmSettingPO.BooleanExpression,
		}).Error
}

func (a *AlarmSettingRepository) ListByDeviceUUID(deviceUUID string) ([]domain.PQAlarmSettings, error) {
	var pqAlarmSettingsPO []model.PQAlarmSettings

	// !出現PhysicalQuantity操作
	err := a.gorm.Table("physical_quantity").
		Where("physical_quantity.device_uuid = ?", deviceUUID).
		Preload("AlarmSettings", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_enable = true").Order("priority")
		}).
		Find(&pqAlarmSettingsPO).Error
	if err != nil {
		return nil, err
	}

	var pqAlarmSettings []domain.PQAlarmSettings
	for _, pqAlarmSettingPO := range pqAlarmSettingsPO {
		tempAlarmSettings := []domain.AlarmSetting{}
		for _, alarmSetting := range pqAlarmSettingPO.AlarmSettings {
			tempAlarmSettings = append(tempAlarmSettings, alarmSetting.ToDomain())
		}

		pqAlarmSettings = append(pqAlarmSettings, domain.PQAlarmSettings{
			PhysicalQuantity: pqAlarmSettingPO.ToDomain(),
			AlarmSettings:    tempAlarmSettings,
		})
	}

	return pqAlarmSettings, nil
}

func (a *AlarmSettingRepository) ListIn(alarmSetting domain.AlarmSetting, uuids []string) ([]domain.AlarmSetting, error) {
	var alarmSettingsPO []model.AlarmSetting

	alarmSettingWherePO := model.AlarmSetting{}.FromDomain(alarmSetting)

	err := a.gorm.Table("alarm_setting").
		Where("uuid IN ?", uuids).
		Where(alarmSettingWherePO).
		Order("priority").
		Find(&alarmSettingsPO).Error
	if err != nil {
		return nil, err
	}

	var alarmSettings []domain.AlarmSetting
	for _, alarmSettingPO := range alarmSettingsPO {
		alarmSettings = append(alarmSettings, alarmSettingPO.ToDomain())
	}

	return alarmSettings, nil
}

func (a *AlarmSettingRepository) ListInPhysicalQuantityUUIDs(alarmSetting domain.AlarmSetting, physicalQuantityUUIDs []string) ([]domain.AlarmSetting, error) {
	var alarmSettingsPO []model.AlarmSetting

	alarmSettingWherePO := model.AlarmSetting{}.FromDomain(alarmSetting)

	err := a.gorm.Table("alarm_setting").
		Where("physical_quantity_uuid IN (?)", physicalQuantityUUIDs).
		Where(alarmSettingWherePO).
		Order("priority").
		Find(&alarmSettingsPO).Error
	if err != nil {
		return nil, err
	}

	var alarmSettings []domain.AlarmSetting
	for _, alarmSettingPO := range alarmSettingsPO {
		alarmSettings = append(alarmSettings, alarmSettingPO.ToDomain())
	}

	return alarmSettings, nil
}
