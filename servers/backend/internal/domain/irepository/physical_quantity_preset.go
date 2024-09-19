package irepository

import "auto-monitoring/internal/domain"

type IPhysicalQuantityPresetRepository interface {
	FindOne(domain.PhysicalQuantityPreset) (domain.PhysicalQuantityPreset, error)
	List(domain.PhysicalQuantityPreset) ([]domain.PhysicalQuantityPreset, error)
	Create(domain.PhysicalQuantityPreset) error
	Update(domain.PhysicalQuantityPreset) error
	Delete(domain.PhysicalQuantityPreset) error
}
