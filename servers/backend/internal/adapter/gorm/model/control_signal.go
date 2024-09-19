package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type ControlSignal struct {
	UUID     string `gorm:"column:uuid;not null;primary_key;type:char(36)" json:"uuid"`
	Name     string `gorm:"column:name;not null;type:varchar(45)" json:"name"`
	FullName string `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`

	DeviceUUID string `gorm:"column:device_uuid;type:char(36);not null" json:"device_uuid"`
	StatusCode string `gorm:"column:status_code;not null;type:varchar(4)" json:"status_code"`
	IsEnable   bool   `gorm:"column:is_enable;not null;default:0;type:tinyint(4)" json:"-"`
	Priority   int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`
	DataType   string `gorm:"column:data_type;default:'Integer';type:varchar(45)" json:"data_type"`

	Description string `gorm:"column:description;type:text" json:"description"`

	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	SignalValue string     `gorm:"column:signal_value;type:decimal(24,6)" json:"signal_value"`

	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (ControlSignal) TableName() string {
	return "control_signal"
}

func (c ControlSignal) ToDomain() domain.ControlSignal {
	return domain.ControlSignal{
		UUID:        c.UUID,
		Name:        c.Name,
		FullName:    c.FullName,
		DeviceUUID:  c.DeviceUUID,
		StatusCode:  c.StatusCode,
		IsEnable:    c.IsEnable,
		Priority:    c.Priority,
		DataType:    c.DataType,
		Description: c.Description,
		UpdateTime:  c.UpdateTime,
		SignalValue: c.SignalValue,
	}
}

func (c ControlSignal) FromDomain(controlSignal domain.ControlSignal) ControlSignal {
	return ControlSignal{
		UUID:        controlSignal.UUID,
		Name:        controlSignal.Name,
		FullName:    controlSignal.FullName,
		DeviceUUID:  controlSignal.DeviceUUID,
		StatusCode:  controlSignal.StatusCode,
		IsEnable:    controlSignal.IsEnable,
		Priority:    controlSignal.Priority,
		DataType:    controlSignal.DataType,
		Description: controlSignal.Description,
		UpdateTime:  controlSignal.UpdateTime,
		SignalValue: controlSignal.SignalValue,
	}
}
