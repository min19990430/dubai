package domain

type SignalInputMapping struct {
	DeviceUUID                 string `json:"device_uuid"`
	Type                       string `json:"type"`
	Pointer                    int    `json:"pointer"`
	TargetPhysicalQuantityUUID string `json:"target_physical_quantity_uuid"`
}

type SignalInputMappingDetail struct {
	SignalInputMapping
	PhysicalQuantity PhysicalQuantity `json:"physical_quantity"`
}
