package domain

type PhysicalQuantityPreset struct {
	UUID                       string `json:"uuid"`
	Priority                   int    `json:"priority"`
	Name                       string `json:"name"`
	FullName                   string `json:"full_name"`
	SiUnit                     string `json:"si_unit"`
	StatusCode                 string `json:"status_code"`
	PhysicalQuantityDataType   string `json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `json:"aggregate_calculation_method"`
	CalibrationEnable          bool   `json:"calibration_enable"`
	CalibrationValue           string `json:"calibration_value"`
	CalibrationParameter       string `json:"calibration_parameter"`
	Description                string `json:"description"`
}
