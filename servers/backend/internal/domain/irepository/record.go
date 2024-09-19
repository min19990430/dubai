package irepository

import (
	"time"

	"auto-monitoring/internal/domain"
)

type IRecordRepository interface {
	Create(tableName string, record domain.Record) error
	CreateMany(tableName string, records []domain.Record) error
	ListMap(start, end time.Time, physicalQuantities []domain.PhysicalQuantity) ([]map[string]string, error)
	List(tableName string, start time.Time, end time.Time) ([]domain.Record, error)
	Last(string, domain.PhysicalQuantity) (domain.Record, error)
}
