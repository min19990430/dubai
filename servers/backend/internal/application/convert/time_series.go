package convert

import "auto-monitoring/internal/domain"

type TimeSeriesColumn struct{}

func (TimeSeriesColumn) FromPhysicalQuantity(pq domain.PhysicalQuantity) domain.TimeSeriesColumn {
	return domain.TimeSeriesColumn{
		UUID:        pq.UUID,
		Name:        pq.Name,
		FullName:    pq.FullName,
		SiUnit:      pq.SiUnit,
		StatusCode:  pq.StatusCode,
		Priority:    pq.Priority,
		Description: pq.Description,
	}
}

func (TimeSeriesColumn) CreatePhysicalQuantityStatus(pq domain.PhysicalQuantity) domain.TimeSeriesColumn {
	return domain.TimeSeriesColumn{
		UUID:        pq.UUID,
		Name:        pq.Name + "_status",
		FullName:    pq.FullName + " Status",
		SiUnit:      "",
		StatusCode:  "",
		Priority:    0,
		Description: pq.Name + "_status",
	}
}

func (TimeSeriesColumn) CreateTimeColumn() domain.TimeSeriesColumn {
	return domain.TimeSeriesColumn{
		UUID:        "",
		Name:        "time",
		FullName:    "time",
		SiUnit:      "",
		StatusCode:  "",
		Priority:    0,
		Description: "",
	}
}
