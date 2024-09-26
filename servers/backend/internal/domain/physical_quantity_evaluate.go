package domain

import "time"

type PhysicalQuantityEvaluate struct {
	UUID                       string `json:"uuid"`
	PhysicalQuantityUUID       string `json:"physical_quantity_uuid"`
	TargetPhysicalQuantityUUID string `json:"target_physical_quantity_uuid"`
	Name                       string `json:"name"`
	FullName                   string `json:"full_name"`
	SiUnit                     string `json:"si_unit"`

	IsEnable bool `json:"-"`
	Priority int  `json:"priority"`

	FormulaType                string `json:"formula_type"`
	Formula                    string `json:"formula"`
	PhysicalQuantityDataType   string `json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `json:"aggregate_calculation_method"`

	Description string `json:"description"`

	UpdateTime *time.Time `json:"update_time"`
	Value      float64    `json:"value"`
	Data       float64    `json:"data"`
}

type PhysicalQuantityEvaluateDetail struct {
	PhysicalQuantityEvaluate

	PhysicalQuantity       PhysicalQuantity `json:"physical_quantity"`
	TargetPhysicalQuantity PhysicalQuantity `json:"target_physical_quantity"`
}
