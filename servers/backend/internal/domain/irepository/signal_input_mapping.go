package irepository

import "auto-monitoring/internal/domain"

type ISignalInputMappingRepository interface {
	ListByDeviceUUID(deviceUUID string) ([]domain.SignalInputMapping, error)
	Create(domain.SignalInputMapping) error
	Update(domain.SignalInputMapping, domain.SignalInputMapping) error
	Delete(domain.SignalInputMapping) error
}

type ISignalInputMappingDetailRepository interface {
	List(domain.SignalInputMapping) ([]domain.SignalInputMappingDetail, error)
}
