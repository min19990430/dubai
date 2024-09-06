package irepository

import "auto-monitoring/internal/domain"

type ISignalInputMappingRepository interface {
	ListByDeviceUUID(deviceUUID string) ([]domain.SignalInputMapping, error)
}
