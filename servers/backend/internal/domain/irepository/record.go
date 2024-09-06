package irepository

import (
	"time"

	"auto-monitoring/internal/domain"
)

type IRecordRepository interface {
	// Create(domain.Record) error
	CreateMany(tableName string, records []domain.Record) error
	ListMap(start, end time.Time, tableName string, physicalQuantities []domain.PhysicalQuantity) ([]map[string]string, error)
	List(time.Time, time.Time, string, domain.PhysicalQuantity) ([]domain.Record, error)
	Last(string, domain.PhysicalQuantity) (domain.Record, error)
}
