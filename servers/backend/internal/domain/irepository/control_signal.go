package irepository

import "auto-monitoring/internal/domain"

type IControlSignalRepository interface {
	ListIn(domain.ControlSignal, []string) ([]domain.ControlSignal, error)
	UpdateSignalValue(string, string) error
}
