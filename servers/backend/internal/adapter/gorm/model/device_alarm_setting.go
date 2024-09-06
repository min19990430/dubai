package model

import (
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/domain"
)

type DeviceAlarmSetting struct {
	UUID       string `gorm:"primaryKey;column:uuid;type:char(36);not null" json:"uuid,omitempty"`
	DeviceUUID string `gorm:"column:device_uuid;type:char(36);not null" json:"device_uuid"`
	Name       string `gorm:"column:name;type:varchar(45);not null" json:"name,omitempty"`
	FullName   string `gorm:"column:full_name;type:varchar(45);not null" json:"full_name,omitempty"`
	Priority   int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`

	BooleanExpression string `gorm:"column:boolean_expression;type:text"`

	IsEnable    bool `gorm:"column:is_enable;type:tinyint;not null;default:0" json:"is_enable,omitempty"`
	IsNotify    bool `gorm:"column:is_notify;type:tinyint;not null;default:0"`
	IsActivated bool `gorm:"column:is_activated;type:tinyint;not null;default:0" json:"is_activated,omitempty"`

	LastOccurTime *time.Time `gorm:"column:last_occur_time;type:datetime" json:"last_occur_time,omitempty"`

	ContentSetting *string `gorm:"column:content_setting;type:text" json:"content_setting,omitempty"`

	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (DeviceAlarmSetting) TableName() string {
	return "device_alarm_settings"
}

func (r DeviceAlarmSetting) FromDomain(domain domain.DeviceAlarmSetting) DeviceAlarmSetting {
	return DeviceAlarmSetting{
		UUID:              domain.UUID,
		DeviceUUID:        domain.DeviceUUID,
		Name:              domain.Name,
		FullName:          domain.FullName,
		Priority:          domain.Priority,
		BooleanExpression: domain.BooleanExpression,
		IsEnable:          domain.IsEnable,
		IsNotify:          domain.IsNotify,
		IsActivated:       domain.IsActivated,
		LastOccurTime:     domain.LastOccurTime,
		ContentSetting:    domain.ContentSetting,
	}
}

func (r DeviceAlarmSetting) ToDomain() domain.DeviceAlarmSetting {
	return domain.DeviceAlarmSetting{
		UUID:              r.UUID,
		DeviceUUID:        r.DeviceUUID,
		Name:              r.Name,
		FullName:          r.FullName,
		Priority:          r.Priority,
		BooleanExpression: r.BooleanExpression,
		IsEnable:          r.IsEnable,
		IsNotify:          r.IsNotify,
		IsActivated:       r.IsActivated,
		LastOccurTime:     r.LastOccurTime,
		ContentSetting:    r.ContentSetting,
	}
}
