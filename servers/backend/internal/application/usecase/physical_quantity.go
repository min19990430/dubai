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

func (pu *PhysicalQuantityUsecase) FindByUUID(uuid string) (domain.PhysicalQuantity, error) {
	return pu.physicalQuantityRepository.FindByUUID(uuid)
}

func (pu *PhysicalQuantityUsecase) UpdateStatus(physicalQuantity domain.PhysicalQuantity, statusCode string) error {
	return pu.physicalQuantityRepository.UpdateStatus(physicalQuantity, statusCode)
}

func (pu *PhysicalQuantityUsecase) List(physicalQuantity domain.PhysicalQuantity) ([]domain.PhysicalQuantity, error) {
	return pu.physicalQuantityRepository.List(physicalQuantity)
}

func (pu *PhysicalQuantityUsecase) Create(physicalQuantity domain.PhysicalQuantity) error {
	return pu.physicalQuantityRepository.Create(physicalQuantity)
}

func (pu *PhysicalQuantityUsecase) Update(physicalQuantity domain.PhysicalQuantity) error {
	return pu.physicalQuantityRepository.Update(physicalQuantity)
}

func (pu *PhysicalQuantityUsecase) Delete(key string) error {
	return pu.physicalQuantityRepository.Delete(domain.PhysicalQuantity{UUID: key})
}
