package irepository

import "auto-monitoring/internal/domain"

type IAlarmSettingRepository interface {
	UpdateStatusAndTime(domain.AlarmSetting) error
	Notify(domain.AlarmSetting, string) error
	UpdateExpression(string, string) error
	List(domain.AlarmSetting) ([]domain.AlarmSetting, error)
	ListByDeviceUUID(string) ([]domain.PQAlarmSettings, error)
	ListIn(domain.AlarmSetting, []string) ([]domain.AlarmSetting, error)
	ListInPhysicalQuantityUUIDs(domain.AlarmSetting, []string) ([]domain.AlarmSetting, error)
	Create(domain.AlarmSetting) error
	Update(domain.AlarmSetting) error
	Delete(domain.AlarmSetting) error
	ExpressionCheck(domain.AlarmSetting, float64) (bool, error)
}
