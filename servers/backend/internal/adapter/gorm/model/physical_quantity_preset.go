package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type PhysicalQuantityPreset struct {
	UUID       string `gorm:"column:uuid;not null;primary_key;type:char(36)" json:"uuid"`
	Priority   int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`
	Name       string `gorm:"column:name;not null;type:varchar(45)" json:"name"`
	FullName   string `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`
	SiUnit     string `gorm:"column:si_unit;default:'';type:varchar(45)" json:"si_unit"`
	StatusCode string `gorm:"column:status_code;not null;type:varchar(4)" json:"status_code"`

	PhysicalQuantityDataType   string `gorm:"column:physical_quantity_data_type;default:'Decimal';type:varchar(45)" json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `gorm:"column:aggregate_calculation_method;default:'avg';type:varchar(45)" json:"aggregate_calculation_method"`

	CalibrationEnable    bool   `gorm:"column:calibration_enable;type:tinyint;not null;default:0" json:"calibration_enable"`
	CalibrationValue     string `gorm:"column:calibration_value;type:varchar(100);not null;default:'1,2,3,4,5'" json:"calibration_value" `
	CalibrationParameter string `gorm:"column:calibration_parameter;type:varchar(100);not null;default:'1,2,3,4,5'" json:"calibration_parameter" `

	Description string `gorm:"column:description;type:text" json:"description"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (PhysicalQuantityPreset) TableName() string {
	return "physical_quantity_preset"
}

func (p PhysicalQuantityPreset) ToDomain() domain.PhysicalQuantityPreset {
	return domain.PhysicalQuantityPreset{
		UUID:                       p.UUID,
		Priority:                   p.Priority,
		Name:                       p.Name,
		FullName:                   p.FullName,
		SiUnit:                     p.SiUnit,
		StatusCode:                 p.StatusCode,
		PhysicalQuantityDataType:   p.PhysicalQuantityDataType,
		AggregateCalculationMethod: p.AggregateCalculationMethod,
		CalibrationEnable:          p.CalibrationEnable,
		CalibrationValue:           p.CalibrationValue,
		CalibrationParameter:       p.CalibrationParameter,
		Description:                p.Description,
	}
}

func (p PhysicalQuantityPreset) FromDomain(domain domain.PhysicalQuantityPreset) PhysicalQuantityPreset {
	return PhysicalQuantityPreset{
		UUID:                       domain.UUID,
		Priority:                   domain.Priority,
		Name:                       domain.Name,
		FullName:                   domain.FullName,
		SiUnit:                     domain.SiUnit,
		StatusCode:                 domain.StatusCode,
		PhysicalQuantityDataType:   domain.PhysicalQuantityDataType,
		AggregateCalculationMethod: domain.AggregateCalculationMethod,
		CalibrationEnable:          domain.CalibrationEnable,
		CalibrationValue:           domain.CalibrationValue,
		CalibrationParameter:       domain.CalibrationParameter,
		Description:                domain.Description,
	}
}
