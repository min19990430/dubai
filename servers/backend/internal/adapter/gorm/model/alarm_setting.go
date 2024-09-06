package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type AlarmSetting struct {
	UUID                 string `gorm:"primaryKey;column:uuid;type:char(36);not null" json:"uuid,omitempty"`
	PhysicalQuantityUUID string `gorm:"unique;column:physical_quantity_uuid;type:char(36);not null" json:"physical_quantity_uuid,omitempty"`
	Name                 string `gorm:"column:name;type:varchar(45)" json:"name"`
	FullName             string `gorm:"column:full_name;type:varchar(45)" json:"full_name"`
	Priority             int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`

	IsEnable    bool `gorm:"column:is_enable;type:tinyint;not null;default:0" json:"is_enable,omitempty"`
	IsNotify    bool `gorm:"column:is_notify;type:tinyint;not null;default:0"`
	IsActivated bool `gorm:"column:is_activated;type:tinyint;not null;default:0" json:"is_activated,omitempty"`

	BooleanExpression string  `gorm:"column:boolean_expression;type:text"`
	LastOccurValue    float64 `gorm:"column:last_occur_value;type:decimal(24,6)" json:"last_occur_value"`

	LastAlarmOccurTime  *time.Time     `gorm:"column:last_alarm_occur_time;type:datetime" json:"last_alarm_occur_time,omitempty"`
	AlarmContentSetting *string        `gorm:"column:alarm_content_setting;type:text" json:"alarm_content_setting,omitempty"`
	UpdatedAt           *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (AlarmSetting) TableName() string {
	return "alarm_setting"
}

func (as AlarmSetting) FromDomain(alarmSetting domain.AlarmSetting) AlarmSetting {
	return AlarmSetting{
		UUID:                 alarmSetting.UUID,
		PhysicalQuantityUUID: alarmSetting.PhysicalQuantityUUID,
		Name:                 alarmSetting.Name,
		FullName:             alarmSetting.FullName,
		Priority:             alarmSetting.Priority,
		IsEnable:             alarmSetting.IsEnable,
		IsNotify:             alarmSetting.IsNotify,
		IsActivated:          alarmSetting.IsActivated,
		BooleanExpression:    alarmSetting.BooleanExpression,
		LastOccurValue:       alarmSetting.LastOccurValue,
		LastAlarmOccurTime:   alarmSetting.LastAlarmOccurTime,
		AlarmContentSetting:  alarmSetting.AlarmContentSetting,
		UpdatedAt:            alarmSetting.UpdatedAt,
	}
}

func (as AlarmSetting) ToDomain() domain.AlarmSetting {
	return domain.AlarmSetting{
		UUID:                 as.UUID,
		PhysicalQuantityUUID: as.PhysicalQuantityUUID,
		Name:                 as.Name,
		FullName:             as.FullName,
		Priority:             as.Priority,
		IsEnable:             as.IsEnable,
		IsNotify:             as.IsNotify,
		IsActivated:          as.IsActivated,
		BooleanExpression:    as.BooleanExpression,
		LastOccurValue:       as.LastOccurValue,
		LastAlarmOccurTime:   as.LastAlarmOccurTime,
		AlarmContentSetting:  as.AlarmContentSetting,
		UpdatedAt:            as.UpdatedAt,
	}
}

type PQAlarmSettings struct {
	PhysicalQuantity

	AlarmSettings []AlarmSetting `gorm:"references:UUID;foreignKey:PhysicalQuantityUUID"`
}

func (PQAlarmSettings) TableName() string {
	return "physical_quantity"
}
