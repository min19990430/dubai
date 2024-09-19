package irepository

import (
	"time"

	"auto-monitoring/internal/domain"
)

type IDeviceAlarmRecordRepository interface {
	Create(domain.DeviceAlarmRecord) error
	List(time.Time, time.Time, domain.DeviceAlarmRecord, bool) ([]domain.DeviceAlarmRecord, error)
	ListIn(time.Time, time.Time, []string, bool) ([]domain.DeviceAlarmRecord, error)
}
