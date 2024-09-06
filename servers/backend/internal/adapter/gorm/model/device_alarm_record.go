package model

import (
	"time"

	"auto-monitoring/internal/domain"
)

type DeviceAlarmRecord struct {
	ID                     uint      `gorm:"column:id;not null;primary_key;auto_increment;type:bigint(20)" json:"-"`
	DeviceAlarmSettingUUID string    `gorm:"column:device_alarm_setting_uuid;type:char(36)" json:"device_alarm_setting_uuid"`
	Name                   string    `gorm:"column:name;not null;index;type:varchar(45)" json:"device_name"`
	FullName               string    `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`
	OccurTime              time.Time `gorm:"column:occur_time;not null;index;type:datetime" json:"occur_time"`
	Content                string    `gorm:"column:content;not null;type:text" json:"content"`
}

func (DeviceAlarmRecord) TableName() string {
	return "device_alarm_records"
}

func (r DeviceAlarmRecord) FromDomain(domain domain.DeviceAlarmRecord) DeviceAlarmRecord {
	return DeviceAlarmRecord{
		DeviceAlarmSettingUUID: domain.DeviceAlarmSettingUUID,
		Name:                   domain.Name,
		FullName:               domain.FullName,
		OccurTime:              domain.OccurTime,
		Content:                domain.Content,
	}
}

func (r DeviceAlarmRecord) ToDomain() domain.DeviceAlarmRecord {
	return domain.DeviceAlarmRecord{
		DeviceAlarmSettingUUID: r.DeviceAlarmSettingUUID,
		Name:                   r.Name,
		FullName:               r.FullName,
		OccurTime:              r.OccurTime,
		Content:                r.Content,
	}
}
