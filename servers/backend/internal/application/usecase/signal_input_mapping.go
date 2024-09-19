package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type SignalInputMappingUsecase struct {
	signalInputMappingRepository irepository.ISignalInputMappingRepository
}

func NewSignalInputMappingUsecase(
	signalInputMappingRepository irepository.ISignalInputMappingRepository,
) *SignalInputMappingUsecase {
	return &SignalInputMappingUsecase{
		signalInputMappingRepository: signalInputMappingRepository,
	}
}

func (su *SignalInputMappingUsecase) ListByDeviceUUID(deviceUUID string) ([]domain.SignalInputMapping, error) {
	return su.signalInputMappingRepository.ListByDeviceUUID(deviceUUID)
}

func (su *SignalInputMappingUsecase) Create(signalInputMapping domain.SignalInputMapping) error {
	return su.signalInputMappingRepository.Create(signalInputMapping)
}

func (su *SignalInputMappingUsecase) Update(old, new domain.SignalInputMapping) error {
	return su.signalInputMappingRepository.Update(old, new)
}

func (su *SignalInputMappingUsecase) Delete(signalInputMapping domain.SignalInputMapping) error {
	return su.signalInputMappingRepository.Delete(signalInputMapping)
}

type SignalInputMappingDetailUsecase struct {
	signalInputMappingDetailRepository irepository.ISignalInputMappingDetailRepository
}

func NewSignalInputMappingDetailUsecase(
	signalInputMappingDetailRepository irepository.ISignalInputMappingDetailRepository,
) *SignalInputMappingDetailUsecase {
	return &SignalInputMappingDetailUsecase{
		signalInputMappingDetailRepository: signalInputMappingDetailRepository,
	}
}

func (su *SignalInputMappingDetailUsecase) List(signalInputMapping domain.SignalInputMapping) ([]domain.SignalInputMappingDetail, error) {
	return su.signalInputMappingDetailRepository.List(signalInputMapping)
}
