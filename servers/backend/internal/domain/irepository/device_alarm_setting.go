package irepository

import "auto-monitoring/internal/domain"

type IDeviceAlarmSettingRepository interface {
	FindByUUID(uuid string) (domain.DeviceAlarmSetting, error)
	List(domain.DeviceAlarmSetting) ([]domain.DeviceAlarmSetting, error)
	UpdateStatusAndTime(domain.DeviceAlarmSetting) error
	Notify(domain.DeviceAlarmSetting, string) error
}
