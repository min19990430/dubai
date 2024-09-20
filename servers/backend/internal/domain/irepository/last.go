package irepository

import "auto-monitoring/internal/domain"

type ILastRepository interface {
	GetStationLast(source string) ([]domain.StationLast, error)
	GetDeviceLast(source string) ([]domain.DeviceLast, error)
}
