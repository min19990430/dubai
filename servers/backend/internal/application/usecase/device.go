package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type DeviceUsecase struct {
	device irepository.IDeviceRepository
}

func NewDeviceUsecase(device irepository.IDeviceRepository) *DeviceUsecase {
	return &DeviceUsecase{
		device: device,
	}
}

func (du *DeviceUsecase) FindByUUID(uuid string) (domain.Device, error) {
	return du.device.FindByUUID(uuid)
}

func (du *DeviceUsecase) List(device domain.Device) ([]domain.Device, error) {
	return du.device.List(device)
}

func (du *DeviceUsecase) Create(device domain.Device) error {
	return du.device.Create(device)
}

func (du *DeviceUsecase) Update(device domain.Device) error {
	return du.device.Update(device)
}

func (du *DeviceUsecase) Delete(key string) error {
	return du.device.Delete(domain.Device{UUID: key})
}
