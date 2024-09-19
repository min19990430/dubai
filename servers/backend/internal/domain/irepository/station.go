package irepository

import "auto-monitoring/internal/domain"

type IStationRepository interface {
	FindByUUID(string) (domain.Station, error)
	List(domain.Station) ([]domain.Station, error)
	ListIn(domain.Station, []string) ([]domain.Station, error)
	Create(domain.Station) error
	Update(domain.Station) error
	Delete(domain.Station) error
}
