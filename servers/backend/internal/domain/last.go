package domain

type Last struct {
	Device                         `json:"device"`
	Station                        Station                        `json:"station"`
	PhysicalQuantitiesWithEvaluate []PhysicalQuantityWithEvaluate `json:"physical_quantities"`
}
