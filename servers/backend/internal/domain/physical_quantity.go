package domain

import "time"

type PhysicalQuantity struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	SiUnit     string `json:"si_unit"`
	DeviceUUID string `json:"device_uuid"`

	StatusCode string `json:"status_code"`
	IsEnable   bool   `json:"-"`
	Priority   int    `json:"priority"`
	Source     string `json:"source"`

	PhysicalQuantityDataType   string `json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `json:"aggregate_calculation_method"`

	CalibrationEnable    bool   `json:"calibration_enable"`
	CalibrationValue     string `json:"calibration_value" `
	CalibrationParameter string `json:"calibration_parameter" `

	Description string `json:"description"`

	UpdateTime *time.Time `json:"update_time"`
	Value      float64    `json:"value"`
	Data       float64    `json:"data"`
}

type PhysicalQuantityWithEvaluate struct {
	PhysicalQuantity

	PhysicalQuantityEvaluates []PhysicalQuantityEvaluate `json:"physical_quantity_evaluates"`
}

type PhysicalQuantityCatchDetail struct {
	PhysicalQuantity `json:"physical_quantity"`

	PhysicalQuantityEvaluates []PhysicalQuantityEvaluate `json:"physical_quantity_evaluates"`

	Device Device `gorm:"references:DeviceUUID;foreignKey:UUID" json:"device"`

	AlarmSettings []AlarmSetting `gorm:"references:UUID;foreignKey:PhysicalQuantityUUID" json:"alarm_settings"`
}
