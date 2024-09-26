package irepository

import "auto-monitoring/internal/domain"

type IPhysicalQuantityEvaluateRepository interface {
	UpdateLast(domain.PhysicalQuantityEvaluate) error
	UpdateFormula(domain.PhysicalQuantityEvaluate) error
}

type IPhysicalQuantityEvaluateDetailRepository interface {
	ListDetail(domain.PhysicalQuantityEvaluate) ([]domain.PhysicalQuantityEvaluateDetail, error)
}
