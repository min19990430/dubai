package domain

import "time"

type Calibration struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	SiUnit   string `json:"si_unit"`

	CalibrationEnable    bool   `json:"calibration_enable"`
	CalibrationValue     string `json:"calibration_value" `
	CalibrationParameter string `json:"calibration_parameter" `

	UpdateTime  *time.Time `json:"update_time" `
	Value       float64    `json:"value" `
	Data        float64    `json:"data" `
	Description string     `json:"description"`
}

type CalibrationDetail struct {
	Calibration
	PhysicalQuantity PhysicalQuantity `json:"physical_quantity"`
}
