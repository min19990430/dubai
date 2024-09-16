package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type PhysicalQuantity struct {
	UUID       string `gorm:"column:uuid;not null;primary_key;type:char(36)" json:"uuid"`
	Name       string `gorm:"column:name;not null;type:varchar(45)" json:"name"`
	FullName   string `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`
	SiUnit     string `gorm:"column:si_unit;default:'';type:varchar(45)" json:"si_unit"`
	DeviceUUID string `gorm:"column:device_uuid;type:char(36);not null" json:"device_uuid"`

	StatusCode string `gorm:"column:status_code;not null;type:varchar(4)" json:"status_code"`
	IsEnable   bool   `gorm:"column:is_enable;not null;default:0;type:tinyint(4)" json:"-"`
	Priority   int    `gorm:"column:priority;default:0;type:int(11)" json:"priority"`
	Source     string `gorm:"column:source;default:'';type:varchar(45)" json:"source"`

	PhysicalQuantityDataType   string `gorm:"column:physical_quantity_data_type;default:'Decimal';type:varchar(45)" json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `gorm:"column:aggregate_calculation_method;default:'avg';type:varchar(45)" json:"aggregate_calculation_method"`

	CalibrationEnable    bool   `gorm:"column:calibration_enable;type:tinyint;not null;default:0" json:"calibration_enable"`
	CalibrationValue     string `gorm:"column:calibration_value;type:varchar(100);not null;default:'1,2,3,4,5'" json:"calibration_value" `
	CalibrationParameter string `gorm:"column:calibration_parameter;type:varchar(100);not null;default:'1,2,3,4,5'" json:"calibration_parameter" `

	Description string `gorm:"column:description;type:text" json:"description"`

	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	Value      float64    `gorm:"column:value;type:decimal(24,6)" json:"value"`
	Data       float64    `gorm:"column:data;type:decimal(24,6)" json:"data"`

	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (PhysicalQuantity) TableName() string {
	return "physical_quantity"
}

func (p PhysicalQuantity) ToDomain() domain.PhysicalQuantity {
	return domain.PhysicalQuantity{
		UUID:                       p.UUID,
		Name:                       p.Name,
		FullName:                   p.FullName,
		SiUnit:                     p.SiUnit,
		DeviceUUID:                 p.DeviceUUID,
		StatusCode:                 p.StatusCode,
		IsEnable:                   p.IsEnable,
		Priority:                   p.Priority,
		Source:                     p.Source,
		PhysicalQuantityDataType:   p.PhysicalQuantityDataType,
		AggregateCalculationMethod: p.AggregateCalculationMethod,
		CalibrationEnable:          p.CalibrationEnable,
		CalibrationValue:           p.CalibrationValue,
		CalibrationParameter:       p.CalibrationParameter,
		Description:                p.Description,
		UpdateTime:                 p.UpdateTime,
		Value:                      p.Value,
		Data:                       p.Data,
	}
}

func (p PhysicalQuantity) FromDomain(domain domain.PhysicalQuantity) PhysicalQuantity {
	return PhysicalQuantity{
		UUID:                       domain.UUID,
		Name:                       domain.Name,
		FullName:                   domain.FullName,
		SiUnit:                     domain.SiUnit,
		DeviceUUID:                 domain.DeviceUUID,
		StatusCode:                 domain.StatusCode,
		IsEnable:                   domain.IsEnable,
		Priority:                   domain.Priority,
		Source:                     domain.Source,
		PhysicalQuantityDataType:   domain.PhysicalQuantityDataType,
		AggregateCalculationMethod: domain.AggregateCalculationMethod,
		CalibrationEnable:          domain.CalibrationEnable,
		CalibrationValue:           domain.CalibrationValue,
		CalibrationParameter:       domain.CalibrationParameter,
		Description:                domain.Description,
		UpdateTime:                 domain.UpdateTime,
		Value:                      domain.Value,
		Data:                       domain.Data,
	}
}

func (p PhysicalQuantity) ToCalibrationDomain() domain.Calibration {
	return domain.Calibration{
		UUID:                 p.UUID,
		Name:                 p.Name,
		FullName:             p.FullName,
		SiUnit:               p.SiUnit,
		CalibrationEnable:    p.CalibrationEnable,
		CalibrationValue:     p.CalibrationValue,
		CalibrationParameter: p.CalibrationParameter,
		UpdateTime:           p.UpdateTime,
		Value:                p.Value,
		Data:                 p.Data,
		Description:          p.Description,
	}
}

type PhysicalQuantityWithEvaluate struct {
	PhysicalQuantity

	PhysicalQuantityEvaluates []PhysicalQuantityEvaluate `gorm:"references:UUID;foreignKey:PhysicalQuantityUUID"`
}

func (PhysicalQuantityWithEvaluate) TableName() string {
	return "physical_quantity"
}

func (p PhysicalQuantityWithEvaluate) ToDomain() domain.PhysicalQuantityWithEvaluate {
	tempPhysicalQuantityEvaluates := []domain.PhysicalQuantityEvaluate{}
	for _, physicalQuantityEvaluate := range p.PhysicalQuantityEvaluates {
		tempPhysicalQuantityEvaluates = append(tempPhysicalQuantityEvaluates, physicalQuantityEvaluate.ToDomain())
	}

	return domain.PhysicalQuantityWithEvaluate{
		PhysicalQuantity:          p.PhysicalQuantity.ToDomain(),
		PhysicalQuantityEvaluates: tempPhysicalQuantityEvaluates,
	}
}

func (p PhysicalQuantityWithEvaluate) FromDomain(domain domain.PhysicalQuantityWithEvaluate) PhysicalQuantityWithEvaluate {
	tempPhysicalQuantityEvaluates := []PhysicalQuantityEvaluate{}
	for _, physicalQuantityEvaluate := range domain.PhysicalQuantityEvaluates {
		tempPhysicalQuantityEvaluates = append(tempPhysicalQuantityEvaluates, PhysicalQuantityEvaluate{}.FromDomain(physicalQuantityEvaluate))
	}

	return PhysicalQuantityWithEvaluate{
		PhysicalQuantity:          PhysicalQuantity{}.FromDomain(domain.PhysicalQuantity),
		PhysicalQuantityEvaluates: tempPhysicalQuantityEvaluates,
	}
}

type PhysicalQuantityCatchDetail struct {
	PhysicalQuantity

	PhysicalQuantityEvaluates []PhysicalQuantityEvaluate `gorm:"references:UUID;foreignKey:PhysicalQuantityUUID"`

	Device Device `gorm:"references:DeviceUUID;foreignKey:UUID"`

	AlarmSettings []AlarmSetting `gorm:"references:UUID;foreignKey:PhysicalQuantityUUID"`
}

func (PhysicalQuantityCatchDetail) TableName() string {
	return "physical_quantity"
}

func (p PhysicalQuantityCatchDetail) ToDomain() domain.PhysicalQuantityCatchDetail {
	tempAlarmSettings := []domain.AlarmSetting{}
	for _, alarmSetting := range p.AlarmSettings {
		tempAlarmSettings = append(tempAlarmSettings, alarmSetting.ToDomain())
	}

	tempPhysicalQuantityEvaluates := []domain.PhysicalQuantityEvaluate{}
	for _, physicalQuantityEvaluate := range p.PhysicalQuantityEvaluates {
		tempPhysicalQuantityEvaluates = append(tempPhysicalQuantityEvaluates, physicalQuantityEvaluate.ToDomain())
	}

	return domain.PhysicalQuantityCatchDetail{
		PhysicalQuantity:          p.PhysicalQuantity.ToDomain(),
		PhysicalQuantityEvaluates: tempPhysicalQuantityEvaluates,
		Device:                    p.Device.ToDomain(),
		AlarmSettings:             tempAlarmSettings,
	}
}

func (p PhysicalQuantityCatchDetail) FromDomain(domain domain.PhysicalQuantityCatchDetail) PhysicalQuantityCatchDetail {
	tempAlarmSettings := []AlarmSetting{}
	for _, alarmSetting := range domain.AlarmSettings {
		tempAlarmSettings = append(tempAlarmSettings, AlarmSetting{}.FromDomain(alarmSetting))
	}

	tempPhysicalQuantityEvaluates := []PhysicalQuantityEvaluate{}
	for _, physicalQuantityEvaluate := range domain.PhysicalQuantityEvaluates {
		tempPhysicalQuantityEvaluates = append(tempPhysicalQuantityEvaluates, PhysicalQuantityEvaluate{}.FromDomain(physicalQuantityEvaluate))
	}

	return PhysicalQuantityCatchDetail{
		PhysicalQuantity:          PhysicalQuantity{}.FromDomain(domain.PhysicalQuantity),
		PhysicalQuantityEvaluates: tempPhysicalQuantityEvaluates,
		Device:                    Device{}.FromDomain(domain.Device),
		AlarmSettings:             tempAlarmSettings,
	}
}
