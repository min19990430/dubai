package domain

// type Last struct {
// 	Device                         `json:"device"`
// 	Station                        Station                        `json:"station"`
// 	PhysicalQuantitiesWithEvaluate []PhysicalQuantityWithEvaluate `json:"physical_quantities"`
// }

type StationLast struct {
	Station            `json:"station"`
	PhysicalQuantities []PhysicalQuantity `json:"physical_quantities"`
}

type DeviceLast struct {
	Device             `json:"device"`
	PhysicalQuantities []PhysicalQuantity `json:"physical_quantities"`
}
