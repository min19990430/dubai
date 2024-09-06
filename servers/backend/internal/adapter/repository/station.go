package repository

import (
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type StationRepository struct {
	gorm *gorm.DB
}

func NewStationRepository(gorm *gorm.DB) irepository.IStationRepository {
	return &StationRepository{gorm: gorm}
}

func (s *StationRepository) FindByUUID(uuid string) (domain.Station, error) {
	var stationPO model.Station

	err := s.gorm.Where("uuid = ?", uuid).First(&stationPO).Error
	if err != nil {
		return domain.Station{}, err
	}

	return stationPO.ToDomain(), nil
}

func (s *StationRepository) List(station domain.Station) ([]domain.Station, error) {
	var stationWherePO = model.Station{}.FromDomain(station)
	var stationPOs []model.Station

	err := s.gorm.Where(stationWherePO).Find(&stationPOs).Error
	if err != nil {
		return nil, err
	}

	var stations []domain.Station
	for _, stationPO := range stationPOs {
		stations = append(stations, stationPO.ToDomain())
	}
	return stations, nil
}

func (s *StationRepository) ListIn(station domain.Station, uuids []string) ([]domain.Station, error) {
	var stationPOs []model.Station

	stationWherePO := model.Station{}.FromDomain(station)

	err := s.gorm.
		Where("uuid IN ?", uuids).
		Where(stationWherePO).
		Find(&stationPOs).Error
	if err != nil {
		return nil, err
	}

	var stations []domain.Station
	for _, stationPO := range stationPOs {
		stations = append(stations, stationPO.ToDomain())
	}
	return stations, nil
}
