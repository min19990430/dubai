package convert

import (
	"errors"
	"strconv"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type SignalInputToInput struct {
	signalInputMapping irepository.ISignalInputMappingRepository
}

func NewSignalInputToInput(signalInputMappingCache irepository.ISignalInputMappingRepository) *SignalInputToInput {
	return &SignalInputToInput{
		signalInputMapping: signalInputMappingCache,
	}
}

func (siti *SignalInputToInput) SignalInputToInput(signalInputs []domain.SignalInput) ([]domain.InputsWithDeviceUUID, error) {
	var inputsWithDeviceUUID []domain.InputsWithDeviceUUID

	for _, signalInput := range signalInputs {
		si, listErr := siti.signalInputMapping.ListByDeviceUUID(signalInput.DeviceUUID)
		if listErr != nil {
			return nil, listErr
		}

		var inputs []domain.Input
		for _, s := range si {
			switch s.Type {
			case "DI":
				// DI string ex."01000110000100..."
				// can't out of pointer
				if s.Pointer < len(signalInput.DI) && s.Pointer >= 0 {
					signal, parseErr := strconv.ParseFloat(string(signalInput.DI[s.Pointer]), 64)
					if parseErr != nil {
						return nil, parseErr
					}

					inputs = append(inputs, domain.Input{
						UUID:     s.TargetPhysicalQuantityUUID,
						Datetime: signalInput.Datetime,
						Value:    signal,
					})
				}
			case "AI":
				// AI []float64 ex.[4.5,15.8...]
				// can't out of pointer
				if s.Pointer < len(signalInput.AI) && s.Pointer >= 0 {
					inputs = append(inputs, domain.Input{
						UUID:     s.TargetPhysicalQuantityUUID,
						Datetime: signalInput.Datetime,
						Value:    signalInput.AI[s.Pointer],
					})
				}
			default:
				return nil, errors.New("type not found")
			}
		}
		inputsWithDeviceUUID = append(inputsWithDeviceUUID, domain.InputsWithDeviceUUID{
			DeviceUUID: signalInput.DeviceUUID,
			Inputs:     inputs,
		})
	}
	return inputsWithDeviceUUID, nil
}
