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

func (lu *LastUsecase) GetStationLast(sourceStr string) ([]domain.StationLast, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return nil, err
	}

	return lu.last.GetStationLast(source.String())
}

func (lu *LastUsecase) GetDeviceLast(sourceStr string) ([]domain.DeviceLast, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return nil, err
	}

	return lu.last.GetDeviceLast(source.String())
}
