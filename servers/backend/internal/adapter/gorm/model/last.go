package model

type Last struct {
	Device
	Station            Station            `gorm:"references:StationUUID;foreignKey:UUID"`
	PhysicalQuantities []PhysicalQuantity `gorm:"references:UUID;foreignKey:DeviceUUID"`
}

func (Last) TableName() string {
	return "device"
}
