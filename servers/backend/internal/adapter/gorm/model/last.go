package model

type StationLast struct {
	Station

	PhysicalQuantities []PhysicalQuantity `gorm:"references:UUID;foreignKey:StationUUID"`
}

func (StationLast) TableName() string {
	return "station"
}

type DeviceLast struct {
	Device

	PhysicalQuantities []PhysicalQuantity `gorm:"references:UUID;foreignKey:DeviceUUID"`
}

func (DeviceLast) TableName() string {
	return "device"
}
