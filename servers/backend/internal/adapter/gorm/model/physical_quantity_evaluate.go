package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type PhysicalQuantityEvaluate struct {
	UUID                       string `gorm:"column:uuid;not null;primary_key;type:char(36)" json:"uuid"`
	PhysicalQuantityUUID       string `gorm:"column:physical_quantity_uuid;not null;type:char(36)" json:"physical_quantity_uuid"`
	TargetPhysicalQuantityUUID string `gorm:"column:target_physical_quantity_uuid;not null;type:char(36)" json:"target_physical_quantity_uuid"`
	Name                       string `gorm:"column:name;not null;type:varchar(45)" json:"name"`
	FullName                   string `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`
	SiUnit                     string `gorm:"column:si_unit;default:'';type:varchar(45)" json:"si_unit"`

	IsEnable bool `gorm:"column:is_enable;not null;default:0;type:tinyint(4)" json:"-"`
	Priority int  `gorm:"column:priority;default:0;type:int(11)" json:"priority"`

	Formula                    string `gorm:"column:formula;type:text" json:"formula"`
	PhysicalQuantityDataType   string `gorm:"column:physical_quantity_data_type;default:'Decimal';type:varchar(45)" json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `gorm:"column:aggregate_calculation_method;default:'avg';type:varchar(45)" json:"aggregate_calculation_method"`

	Description string `gorm:"column:description;type:text" json:"description"`

	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	Value      float64    `gorm:"column:value;type:decimal(24,6)" json:"value"`
	Data       float64    `gorm:"column:data;type:decimal(24,6)" json:"data"`

	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (PhysicalQuantityEvaluate) TableName() string {
	return "physical_quantity_evaluate"
}

func (p PhysicalQuantityEvaluate) ToDomain() domain.PhysicalQuantityEvaluate {
	return domain.PhysicalQuantityEvaluate{
		UUID:                       p.UUID,
		PhysicalQuantityUUID:       p.PhysicalQuantityUUID,
		TargetPhysicalQuantityUUID: p.TargetPhysicalQuantityUUID,
		Name:                       p.Name,
		FullName:                   p.FullName,
		SiUnit:                     p.SiUnit,
		IsEnable:                   p.IsEnable,
		Priority:                   p.Priority,
		Formula:                    p.Formula,
		PhysicalQuantityDataType:   p.PhysicalQuantityDataType,
		AggregateCalculationMethod: p.AggregateCalculationMethod,
		Description:                p.Description,
		UpdateTime:                 p.UpdateTime,
		Value:                      p.Value,
		Data:                       p.Data,
	}
}

func (p PhysicalQuantityEvaluate) FromDomain(domain domain.PhysicalQuantityEvaluate) PhysicalQuantityEvaluate {
	return PhysicalQuantityEvaluate{
		UUID:                       domain.UUID,
		PhysicalQuantityUUID:       domain.PhysicalQuantityUUID,
		TargetPhysicalQuantityUUID: domain.TargetPhysicalQuantityUUID,
		Name:                       domain.Name,
		FullName:                   domain.FullName,
		SiUnit:                     domain.SiUnit,
		IsEnable:                   domain.IsEnable,
		Priority:                   domain.Priority,
		Formula:                    domain.Formula,
		PhysicalQuantityDataType:   domain.PhysicalQuantityDataType,
		AggregateCalculationMethod: domain.AggregateCalculationMethod,
		Description:                domain.Description,
		UpdateTime:                 domain.UpdateTime,
		Value:                      domain.Value,
		Data:                       domain.Data,
	}
}
