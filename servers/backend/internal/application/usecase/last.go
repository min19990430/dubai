package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type LastUsecase struct {
	last irepository.ILastRepository
}

func NewLastUsecase(
	last irepository.ILastRepository,
) *LastUsecase {
	return &LastUsecase{
		last: last,
	}
}

func (lu *LastUsecase) GetStationLast() ([]domain.StationLast, error) {
	return lu.last.GetStationLast()
}

func (lu *LastUsecase) GetDeviceLast() ([]domain.DeviceLast, error) {
	return lu.last.GetDeviceLast()
}
