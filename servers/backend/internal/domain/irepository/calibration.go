package irepository

import "auto-monitoring/internal/domain"

type ICalibrationRepository interface {
	FindDetailByPhysicalQuantity(domain.PhysicalQuantity) ([]domain.CalibrationDetail, error)
	FindDetailByDeviceUUID(string) ([]domain.CalibrationDetail, error)
	Update(string, bool, string, string) error
}
