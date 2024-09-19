package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type PhysicalQuantityPresetUsecase struct {
	physicalQuantityPresetRepository irepository.IPhysicalQuantityPresetRepository
}

func NewPhysicalQuantityPresetUsecase(
	physicalQuantityPresetRepository irepository.IPhysicalQuantityPresetRepository,
) *PhysicalQuantityPresetUsecase {
	return &PhysicalQuantityPresetUsecase{
		physicalQuantityPresetRepository: physicalQuantityPresetRepository,
	}
}

func (pu *PhysicalQuantityPresetUsecase) Get(physicalQuantityPreset domain.PhysicalQuantityPreset) (domain.PhysicalQuantityPreset, error) {
	return pu.physicalQuantityPresetRepository.FindOne(physicalQuantityPreset)
}

func (pu *PhysicalQuantityPresetUsecase) List(physicalQuantityPreset domain.PhysicalQuantityPreset) ([]domain.PhysicalQuantityPreset, error) {
	return pu.physicalQuantityPresetRepository.List(physicalQuantityPreset)
}

func (pu *PhysicalQuantityPresetUsecase) Create(physicalQuantityPreset domain.PhysicalQuantityPreset) error {
	return pu.physicalQuantityPresetRepository.Create(physicalQuantityPreset)
}

func (pu *PhysicalQuantityPresetUsecase) Update(physicalQuantityPreset domain.PhysicalQuantityPreset) error {
	return pu.physicalQuantityPresetRepository.Update(physicalQuantityPreset)
}

func (pu *PhysicalQuantityPresetUsecase) Delete(key string) error {
	return pu.physicalQuantityPresetRepository.Delete(domain.PhysicalQuantityPreset{UUID: key})
}
