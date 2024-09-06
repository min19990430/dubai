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

type DeviceStationUsecase struct {
	deviceStation irepository.IDeviceStationRepository
}

func NewDeviceStationUsecase(deviceStation irepository.IDeviceStationRepository) *DeviceStationUsecase {
	return &DeviceStationUsecase{
		deviceStation: deviceStation,
	}
}

func (dsu *DeviceStationUsecase) List(deviceStation domain.DeviceStation) ([]domain.DeviceStation, error) {
	return dsu.deviceStation.List(deviceStation)
}
