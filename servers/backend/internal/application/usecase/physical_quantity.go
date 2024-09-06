package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type PhysicalQuantityUsecase struct {
	physicalQuantityRepository irepository.IPhysicalQuantityRepository
}

func NewPhysicalQuantityUsecase(
	physicalQuantityRepository irepository.IPhysicalQuantityRepository,
) *PhysicalQuantityUsecase {
	return &PhysicalQuantityUsecase{
		physicalQuantityRepository: physicalQuantityRepository,
	}
}

func (pu *PhysicalQuantityUsecase) UpdateStatus(physicalQuantity domain.PhysicalQuantity, statusCode string) error {
	return pu.physicalQuantityRepository.UpdateStatus(physicalQuantity, statusCode)
}
