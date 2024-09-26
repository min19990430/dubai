package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmSettingUsecase struct {
	alarmSetting irepository.IAlarmSettingRepository

	physicalQuantity irepository.IPhysicalQuantityRepository
}

func NewAlarmSettingUsecase(alarmSetting irepository.IAlarmSettingRepository, physicalQuantity irepository.IPhysicalQuantityRepository) *AlarmSettingUsecase {
	return &AlarmSettingUsecase{
		alarmSetting:     alarmSetting,
		physicalQuantity: physicalQuantity,
	}
}

func (asu *AlarmSettingUsecase) TestExpression(BooleanExpression string, value float64) (bool, error) {
	return asu.alarmSetting.ExpressionCheck(domain.AlarmSetting{BooleanExpression: BooleanExpression}, value)
}

func (asu *AlarmSettingUsecase) UpdateExpression(uuid, expression string) error {
	return asu.alarmSetting.UpdateExpression(uuid, expression)
}

func (asu *AlarmSettingUsecase) ListByDeviceUUID(uuid string) ([]domain.PQAlarmSettings, error) {
	pqAlarmSettings, err := asu.alarmSetting.ListByDeviceUUID(uuid)
	if err != nil {
		return nil, err
	}

	var pqAlarmSettingsResult []domain.PQAlarmSettings
	for _, pqAlarmSetting := range pqAlarmSettings {
		if len(pqAlarmSetting.AlarmSettings) == 0 {
			continue
		}
		pqAlarmSettingsResult = append(pqAlarmSettingsResult, pqAlarmSetting)
	}

	return pqAlarmSettingsResult, nil
}

func (asu *AlarmSettingUsecase) ListByStationUUID(stationUUID string) ([]domain.PQAlarmSettings, error) {
	physicalQuantities, err := asu.physicalQuantity.List(domain.PhysicalQuantity{StationUUID: stationUUID})
	if err != nil {
		return nil, err
	}

	var pqAlarmSettings []domain.PQAlarmSettings
	for _, physicalQuantity := range physicalQuantities {
		pqAlarmSetting, listErr := asu.alarmSetting.List(domain.AlarmSetting{PhysicalQuantityUUID: physicalQuantity.UUID})
		if listErr != nil {
			return nil, listErr
		}
		if len(pqAlarmSetting) == 0 {
			continue
		}
		pqAlarmSettings = append(pqAlarmSettings,
			domain.PQAlarmSettings{
				PhysicalQuantity: physicalQuantity,
				AlarmSettings:    pqAlarmSetting,
			})
	}

	return pqAlarmSettings, nil
}

func (asu *AlarmSettingUsecase) Create(alarmSetting domain.AlarmSetting) error {
	return asu.alarmSetting.Create(alarmSetting)
}

func (asu *AlarmSettingUsecase) Update(alarmSetting domain.AlarmSetting) error {
	return asu.alarmSetting.Update(alarmSetting)
}

func (asu *AlarmSettingUsecase) Delete(key string) error {
	return asu.alarmSetting.Delete(domain.AlarmSetting{UUID: key})
}
