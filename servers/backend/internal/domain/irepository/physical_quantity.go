package irepository

import "auto-monitoring/internal/domain"

type IPhysicalQuantityCatchDetailRepository interface {
	FindOne(domain.PhysicalQuantityCatchDetail) (domain.PhysicalQuantityCatchDetail, error)
	List(domain.PhysicalQuantity) ([]domain.PhysicalQuantityCatchDetail, error)
}

type IPhysicalQuantityWithEvaluateRepository interface {
	List(domain.PhysicalQuantity) ([]domain.PhysicalQuantityWithEvaluate, error)
}

type IPhysicalQuantityRepository interface {
	FindByUUID(string) (domain.PhysicalQuantity, error)
	List(domain.PhysicalQuantity) ([]domain.PhysicalQuantity, error)
	ListIn(domain.PhysicalQuantity, []string) ([]domain.PhysicalQuantity, error)
	ListInDeviceUUIDs(domain.PhysicalQuantity, []string) ([]domain.PhysicalQuantity, error)
	UpdateLast(domain.PhysicalQuantity) error
	UpdateStatus(domain.PhysicalQuantity, string) error
	Create(domain.PhysicalQuantity) error
	Update(domain.PhysicalQuantity) error
	Delete(domain.PhysicalQuantity) error
}
