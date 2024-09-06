package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmSettingUsecase struct {
	alarmSetting irepository.IAlarmSettingRepository

	device irepository.IDeviceRepository
}

func NewAlarmSettingUsecase(alarmSetting irepository.IAlarmSettingRepository, device irepository.IDeviceRepository) *AlarmSettingUsecase {
	return &AlarmSettingUsecase{
		alarmSetting: alarmSetting,
		device:       device,
	}
}

func (asu *AlarmSettingUsecase) UpdateExpression(uuid, expression string) error {
	return asu.alarmSetting.UpdateExpression(uuid, expression)
}

func (asu *AlarmSettingUsecase) ListByDeviceUUID(uuid string) ([]domain.PQAlarmSettings, error) {
	return asu.alarmSetting.ListByDeviceUUID(uuid)
}

func (asu *AlarmSettingUsecase) ListByStationUUID(uuid string) ([]domain.PQAlarmSettings, error) {
	device, err := asu.device.List(domain.Device{StationUUID: uuid})
	if err != nil {
		return nil, err
	}

	var deviceUUIDs []string
	for _, d := range device {
		deviceUUIDs = append(deviceUUIDs, d.UUID)
	}

	var pqAlarmSettings []domain.PQAlarmSettings
	for _, deviceUUID := range deviceUUIDs {
		pqAlarmSetting, listErr := asu.alarmSetting.ListByDeviceUUID(deviceUUID)
		if listErr != nil {
			return nil, listErr
		}
		pqAlarmSettings = append(pqAlarmSettings, pqAlarmSetting...)
	}

	return pqAlarmSettings, nil
}
