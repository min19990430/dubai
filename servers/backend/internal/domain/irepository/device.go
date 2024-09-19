package irepository

import "auto-monitoring/internal/domain"

type IDeviceRepository interface {
	FindByUUID(string) (domain.Device, error)
	Find(domain.Device) (domain.Device, error)
	List(domain.Device) ([]domain.Device, error)
	ListIn(domain.Device, []string) ([]domain.Device, error)
	UpdateLastTime(domain.Device) error
	Create(domain.Device) error
	Update(domain.Device) error
	Delete(domain.Device) error
}
