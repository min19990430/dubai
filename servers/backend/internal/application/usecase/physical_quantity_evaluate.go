package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
	"auto-monitoring/pkg/expression"
)

type PhysicalQuantityEvaluateUsecase struct {
	physicalQuantityEvaluateRepository irepository.IPhysicalQuantityEvaluateRepository
}

func NewPhysicalQuantityEvaluateUsecase(
	physicalQuantityEvaluateRepository irepository.IPhysicalQuantityEvaluateRepository,
) *PhysicalQuantityEvaluateUsecase {
	return &PhysicalQuantityEvaluateUsecase{
		physicalQuantityEvaluateRepository: physicalQuantityEvaluateRepository,
	}
}

func (pqu *PhysicalQuantityEvaluateUsecase) Evaluate(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate, value float64) (float64, error) {
	return expression.Calculate(physicalQuantityEvaluate.Formula, value)
}
