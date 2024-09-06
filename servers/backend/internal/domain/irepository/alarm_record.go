package irepository

import (
	"time"

	"auto-monitoring/internal/domain"
)

type IAlarmRecordRepository interface {
	Create(domain.AlarmRecord) error
	List(time.Time, time.Time, domain.AlarmRecord, bool) ([]domain.AlarmRecord, error)
	ListIn(time.Time, time.Time, []string, bool) ([]domain.AlarmRecord, error)
}
