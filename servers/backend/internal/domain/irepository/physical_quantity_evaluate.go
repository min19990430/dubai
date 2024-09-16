package irepository

import "auto-monitoring/internal/domain"

type IPhysicalQuantityEvaluateRepository interface {
	UpdateLast(domain.PhysicalQuantityEvaluate) error
}
