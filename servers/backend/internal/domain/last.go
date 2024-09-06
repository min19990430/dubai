package domain

type Last struct {
	Device             `json:"device"`
	Station            Station            `json:"station"`
	PhysicalQuantities []PhysicalQuantity `json:"physical_quantities"`
}
