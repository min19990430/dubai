package valueobject

type Input struct {
	UUID     string  `json:"uuid" binding:"required,uuid"`
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}
