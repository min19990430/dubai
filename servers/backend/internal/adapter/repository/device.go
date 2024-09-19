package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type DeviceRepository struct {
	gorm *gorm.DB
}

func NewDeviceRepository(gorm *gorm.DB) irepository.IDeviceRepository {
	return &DeviceRepository{gorm: gorm}
}

func (d *DeviceRepository) UpdateLastTime(device domain.Device) error {
	devicePO := model.Device{
		UUID:        device.UUID,
		IsConnected: device.IsConnected,
		UpdateTime:  device.UpdateTime,
	}

	return d.gorm.Model(&devicePO).
		Where("uuid = ?", devicePO.UUID).
		Updates(map[string]interface{}{
			"is_connected": devicePO.IsConnected,
			"update_time":  devicePO.UpdateTime,
		}).Error
}

func (d *DeviceRepository) Find(device domain.Device) (domain.Device, error) {
	var devicePO model.Device
	if err := d.gorm.Where(device).First(&devicePO).Error; err != nil {
		return domain.Device{}, err
	}

	return devicePO.ToDomain(), nil
}

func (d *DeviceRepository) FindByUUID(uuid string) (domain.Device, error) {
	return d.Find(domain.Device{UUID: uuid})
}

func (d *DeviceRepository) List(device domain.Device) ([]domain.Device, error) {
	var devicePOs []model.Device

	deviceWherePO := model.Device{}.FromDomain(device)

	err := d.gorm.Where(deviceWherePO).Order("priority").Find(&devicePOs).Error
	if err != nil {
		return nil, err
	}

	var devices []domain.Device
	for _, devicePO := range devicePOs {
		devices = append(devices, devicePO.ToDomain())
	}
	return devices, nil
}

func (d *DeviceRepository) ListIn(device domain.Device, uuids []string) ([]domain.Device, error) {
	var devicePOs []model.Device

	deviceWherePO := model.Device{}.FromDomain(device)

	err := d.gorm.Table("device").
		Where("uuid IN ?", uuids).
		Where(deviceWherePO).
		Order("priority").
		Find(&devicePOs).Error
	if err != nil {
		return nil, err
	}

	var devices []domain.Device
	for _, devicePO := range devicePOs {
		devices = append(devices, devicePO.ToDomain())
	}
	return devices, nil
}

func (d *DeviceRepository) Create(device domain.Device) error {
	devicePO := model.Device{}.FromDomain(device)
	return d.gorm.Create(&devicePO).Error
}

func (d *DeviceRepository) Update(device domain.Device) error {
	devicePO := model.Device{}.FromDomain(device)
	return d.gorm.Model(&devicePO).
		Where("uuid = ?", device.UUID).
		Updates(map[string]interface{}{
			"id":           devicePO.ID,
			"name":         devicePO.Name,
			"is_enable":    devicePO.IsEnable,
			"is_connected": devicePO.IsConnected,
			"priority":     devicePO.Priority,
			"lat":          devicePO.Lat,
			"lon":          devicePO.Lon,
			"description":  devicePO.Description,
		}).Error
}

func (d *DeviceRepository) Delete(device domain.Device) error {
	devicePO := model.Device{}.FromDomain(device)
	return d.gorm.Where("uuid = ?", device.UUID).Delete(&devicePO).Error
}
