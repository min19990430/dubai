package model

import "auto-monitoring/internal/domain"

type SignalInputMapping struct {
	DeviceUUID                 string `gorm:"column:device_uuid;not null;primary_key;type:char(36)" json:"device_uuid"`
	Type                       string `gorm:"column:type;not null;primary_key;type:varchar(45)" json:"type"`
	Pointer                    int    `gorm:"column:pointer;not null;primary_key;type:int(11)" json:"pointer"`
	TargetPhysicalQuantityUUID string `gorm:"column:target_physical_quantity_uuid;not null;type:char(36)" json:"target_physical_quantity_uuid"`
}

func (SignalInputMapping) TableName() string {
	return "signal_input_mapping"
}

func (sim SignalInputMapping) ToDomain() domain.SignalInputMapping {
	return domain.SignalInputMapping{
		DeviceUUID:                 sim.DeviceUUID,
		Type:                       sim.Type,
		Pointer:                    sim.Pointer,
		TargetPhysicalQuantityUUID: sim.TargetPhysicalQuantityUUID,
	}
}

func (sim SignalInputMapping) FromDomain(signalInputMapping domain.SignalInputMapping) SignalInputMapping {
	return SignalInputMapping{
		DeviceUUID:                 signalInputMapping.DeviceUUID,
		Type:                       signalInputMapping.Type,
		Pointer:                    signalInputMapping.Pointer,
		TargetPhysicalQuantityUUID: signalInputMapping.TargetPhysicalQuantityUUID,
	}
}

type SignalInputMappingDetail struct {
	SignalInputMapping
	PhysicalQuantity PhysicalQuantity `gorm:"references:TargetPhysicalQuantityUUID;foreignKey:UUID"`
}

func (SignalInputMappingDetail) TableName() string {
	return "signal_input_mapping"
}

func (simd SignalInputMappingDetail) ToDomain() domain.SignalInputMappingDetail {
	return domain.SignalInputMappingDetail{
		SignalInputMapping: simd.SignalInputMapping.ToDomain(),
		PhysicalQuantity:   simd.PhysicalQuantity.ToDomain(),
	}
}

func (simd SignalInputMappingDetail) FromDomain(signalInputMappingDetail domain.SignalInputMappingDetail) SignalInputMappingDetail {
	return SignalInputMappingDetail{
		SignalInputMapping: SignalInputMapping{}.FromDomain(signalInputMappingDetail.SignalInputMapping),
		PhysicalQuantity:   PhysicalQuantity{}.FromDomain(signalInputMappingDetail.PhysicalQuantity),
	}
}
