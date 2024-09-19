package irepository

import "auto-monitoring/internal/domain"

type ILastRepository interface {
	GetStationLast() ([]domain.StationLast, error)
	GetDeviceLast() ([]domain.DeviceLast, error)
}
