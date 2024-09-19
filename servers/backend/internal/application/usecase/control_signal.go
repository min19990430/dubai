package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type ControlSignalUsecase struct {
	ControlSignalRepository irepository.IControlSignalRepository
}

func NewControlSignalUsecase(controlSignalRepository irepository.IControlSignalRepository) *ControlSignalUsecase {
	return &ControlSignalUsecase{
		ControlSignalRepository: controlSignalRepository,
	}
}

func (cu *ControlSignalUsecase) ListIn(uuid []string) ([]domain.ControlSignal, error) {
	return cu.ControlSignalRepository.ListIn(domain.ControlSignal{}, uuid)
}

func (cu *ControlSignalUsecase) UpdateSignalValue(uuid string, value string) error {
	return cu.ControlSignalRepository.UpdateSignalValue(uuid, value)
}
