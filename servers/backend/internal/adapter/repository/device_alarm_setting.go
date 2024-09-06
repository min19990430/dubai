package repository

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	linenotify "auto-monitoring/internal/adapter/line-notify"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type DeviceAlarmSettingRepository struct {
	db         *gorm.DB
	linenotify *linenotify.LineNotify
}

func NewDeviceAlarmSettingRepository(db *gorm.DB, linenotify *linenotify.LineNotify) irepository.IDeviceAlarmSettingRepository {
	return &DeviceAlarmSettingRepository{db: db, linenotify: linenotify}
}

func (das *DeviceAlarmSettingRepository) FindByUUID(uuid string) (domain.DeviceAlarmSetting, error) {
	var deviceAlarmSettingPO model.DeviceAlarmSetting
	err := das.db.Where("uuid = ?", uuid).First(&deviceAlarmSettingPO).Error
	if err != nil {
		return domain.DeviceAlarmSetting{}, err
	}
	return deviceAlarmSettingPO.ToDomain(), nil
}

func (das *DeviceAlarmSettingRepository) List(deviceAlarmSetting domain.DeviceAlarmSetting) ([]domain.DeviceAlarmSetting, error) {
	var deviceAlarmSettingPOs []model.DeviceAlarmSetting

	deviceAlarmSettingWherePO := model.DeviceAlarmSetting{}.FromDomain(deviceAlarmSetting)

	err := das.db.Where(deviceAlarmSettingWherePO).Find(&deviceAlarmSettingPOs).Error
	if err != nil {
		return nil, err
	}

	var deviceAlarmSettings []domain.DeviceAlarmSetting
	for _, deviceAlarmSetting := range deviceAlarmSettingPOs {
		deviceAlarmSettings = append(deviceAlarmSettings, deviceAlarmSetting.ToDomain())
	}
	return deviceAlarmSettings, nil
}

func (das *DeviceAlarmSettingRepository) UpdateStatusAndTime(deviceAlarmSetting domain.DeviceAlarmSetting) error {
	deviceAlarmSettingPO := model.DeviceAlarmSetting{
		UUID:          deviceAlarmSetting.UUID,
		IsActivated:   deviceAlarmSetting.IsActivated,
		LastOccurTime: deviceAlarmSetting.LastOccurTime,
	}

	return das.db.Model(&deviceAlarmSettingPO).
		Where("uuid = ?", deviceAlarmSettingPO.UUID).
		Updates(map[string]interface{}{
			"is_activated":    deviceAlarmSettingPO.IsActivated,
			"last_occur_time": deviceAlarmSettingPO.LastOccurTime,
		}).Error
}

func (das *DeviceAlarmSettingRepository) Notify(deviceAlarmSetting domain.DeviceAlarmSetting, content string) error {
	if deviceAlarmSetting.LastOccurTime == nil {
		deviceAlarmSetting.LastOccurTime = new(time.Time)
	}

	das.linenotify.Message.UpdateTime = *deviceAlarmSetting.LastOccurTime
	das.linenotify.Message.ID = ""
	das.linenotify.Message.Name = deviceAlarmSetting.Name
	das.linenotify.Message.Event = content

	return das.linenotify.Send()
}
