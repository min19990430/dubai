package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type StationUsecase struct {
	station irepository.IStationRepository
}

func NewStationUsecase(station irepository.IStationRepository) *StationUsecase {
	return &StationUsecase{
		station: station,
	}
}

func (su *StationUsecase) FindByUUID(uuid string) (domain.Station, error) {
	return su.station.FindByUUID(uuid)
}

func (su *StationUsecase) List(station domain.Station) ([]domain.Station, error) {
	return su.station.List(station)
}

func (su *StationUsecase) Create(station domain.Station) error {
	return su.station.Create(station)
}

func (su *StationUsecase) Update(station domain.Station) error {
	return su.station.Update(station)
}

func (su *StationUsecase) Delete(key string) error {
	return su.station.Delete(domain.Station{UUID: key})
}
