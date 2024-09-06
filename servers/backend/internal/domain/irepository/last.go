package irepository

import "auto-monitoring/internal/domain"

type ILastRepository interface {
	GetLast() ([]domain.Last, error)
}
