package model

import "auto-monitoring/internal/domain"

type Last struct {
	Device
	Station                        Station                        `gorm:"references:StationUUID;foreignKey:UUID"`
	PhysicalQuantitiesWithEvaluate []PhysicalQuantityWithEvaluate `gorm:"references:UUID;foreignKey:DeviceUUID"`
}

func (Last) TableName() string {
	return "device"
}

func (l Last) ToDomain() domain.Last {
	tempPhysicalQuantities := []domain.PhysicalQuantityWithEvaluate{}
	for _, physicalQuantity := range l.PhysicalQuantitiesWithEvaluate {
		tempPhysicalQuantities = append(tempPhysicalQuantities, physicalQuantity.ToDomain())
	}

	return domain.Last{
		Device:                         l.Device.ToDomain(),
		Station:                        l.Station.ToDomain(),
		PhysicalQuantitiesWithEvaluate: tempPhysicalQuantities,
	}
}

func (l Last) FromDomain(domain domain.Last) Last {
	tempPhysicalQuantities := []PhysicalQuantityWithEvaluate{}
	for _, physicalQuantity := range domain.PhysicalQuantitiesWithEvaluate {
		tempPhysicalQuantities = append(tempPhysicalQuantities, PhysicalQuantityWithEvaluate{}.FromDomain(physicalQuantity))
	}

	return Last{
		Device:                         Device{}.FromDomain(domain.Device),
		Station:                        Station{}.FromDomain(domain.Station),
		PhysicalQuantitiesWithEvaluate: tempPhysicalQuantities,
	}
}
