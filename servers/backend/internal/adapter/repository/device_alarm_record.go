package repository

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type DeviceAlarmRecordRepository struct {
	db *gorm.DB
}

func NewDeviceAlarmRecordRepository(db *gorm.DB) irepository.IDeviceAlarmRecordRepository {
	return &DeviceAlarmRecordRepository{db: db}
}

func (d *DeviceAlarmRecordRepository) Create(deviceAlarmRecord domain.DeviceAlarmRecord) error {
	alarmRecordPO := model.DeviceAlarmRecord{}.FromDomain(deviceAlarmRecord)

	return d.db.Create(&alarmRecordPO).Error
}

func (d *DeviceAlarmRecordRepository) List(startTime, endTime time.Time, deviceAlarmRecord domain.DeviceAlarmRecord, reverse bool) ([]domain.DeviceAlarmRecord, error) {
	var alarmRecordPOs []model.DeviceAlarmRecord
	query := d.db.Where("occur_time between ? AND ?", startTime, endTime).
		Where("device_alarm_setting_uuid IN ?", deviceAlarmRecord.DeviceAlarmSettingUUID)

	if reverse {
		query = query.Order("occur_time desc")
	}

	err := query.Find(&alarmRecordPOs).Error
	if err != nil {
		return nil, err
	}

	var deviceAlarmRecords []domain.DeviceAlarmRecord
	for _, alarmRecordPO := range alarmRecordPOs {
		deviceAlarmRecords = append(deviceAlarmRecords, alarmRecordPO.ToDomain())
	}

	return deviceAlarmRecords, nil
}

func (d *DeviceAlarmRecordRepository) ListIn(startTime, endTime time.Time, deviceAlarmSettingUUIDs []string, reverse bool) ([]domain.DeviceAlarmRecord, error) {
	var deviceAlarmRecordPOs []model.DeviceAlarmRecord
	query := d.db.Where("occur_time between ? AND ?", startTime, endTime).
		Where("device_alarm_setting_uuid IN ?", deviceAlarmSettingUUIDs)

	if reverse {
		query = query.Order("occur_time desc")
	}

	err := query.Find(&deviceAlarmRecordPOs).Error
	if err != nil {
		return nil, err
	}

	var deviceAlarmRecords []domain.DeviceAlarmRecord
	for _, deviceAlarmRecordPO := range deviceAlarmRecordPOs {
		deviceAlarmRecords = append(deviceAlarmRecords, deviceAlarmRecordPO.ToDomain())
	}

	return deviceAlarmRecords, nil
}
