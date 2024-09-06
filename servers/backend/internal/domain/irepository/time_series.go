package irepository

import (
	"time"

	"auto-monitoring/internal/domain"
)

type ITimeSeriesRepository interface {
	Aggregate(time.Time, time.Time, time.Duration, []domain.PhysicalQuantity) (domain.TimeSeries, error)
	AggregateMap(time.Time, time.Time, time.Duration, []domain.PhysicalQuantity) ([]map[string]string, error)
}
