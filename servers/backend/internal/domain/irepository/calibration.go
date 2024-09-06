package irepository

import "auto-monitoring/internal/domain"

type ICalibrationRepository interface {
	FindDetailByDeviceUUID(string) ([]domain.CalibrationDetail, error)
	Update(string, bool, string, string) error
}
