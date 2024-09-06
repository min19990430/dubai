package repository

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmRecordRepository struct {
	gorm *gorm.DB
}

func NewAlarmRecordRepository(gorm *gorm.DB) irepository.IAlarmRecordRepository {
	return &AlarmRecordRepository{gorm: gorm}
}

func (a *AlarmRecordRepository) Create(alarmRecord domain.AlarmRecord) error {
	alarmRecordPO := model.AlarmRecord{}.FromDomain(alarmRecord)

	return a.gorm.Create(&alarmRecordPO).Error
}

func (a *AlarmRecordRepository) List(startTime, endTime time.Time, alarmRecord domain.AlarmRecord, reverse bool) ([]domain.AlarmRecord, error) {
	alarmRecordWherePO := model.AlarmRecord{}.FromDomain(alarmRecord)

	var alarmRecordPOs []model.AlarmRecord

	query := a.gorm.
		Where("occur_time between ? AND ?", startTime, endTime).
		Where(alarmRecordWherePO)

	if reverse {
		query = query.Order("occur_time DESC")
	}

	err := query.Find(&alarmRecordPOs).Error
	if err != nil {
		return nil, err
	}

	var alarmRecords []domain.AlarmRecord
	for _, alarmRecordPO := range alarmRecordPOs {
		alarmRecords = append(alarmRecords, alarmRecordPO.ToDomain())
	}

	return alarmRecords, nil
}

func (a *AlarmRecordRepository) ListIn(startTime, endTime time.Time, alarmSettingUUIDs []string, reverse bool) ([]domain.AlarmRecord, error) {
	var alarmRecordPOs []model.AlarmRecord

	query := a.gorm.
		Where("occur_time between ? AND ?", startTime, endTime).
		Where("alarm_setting_uuid IN ?", alarmSettingUUIDs)

	if reverse {
		query = query.Order("occur_time DESC")
	}

	err := query.Find(&alarmRecordPOs).Error
	if err != nil {
		return nil, err
	}

	var alarmRecords []domain.AlarmRecord
	for _, alarmRecordPO := range alarmRecordPOs {
		alarmRecords = append(alarmRecords, alarmRecordPO.ToDomain())
	}
	return alarmRecords, nil
}
