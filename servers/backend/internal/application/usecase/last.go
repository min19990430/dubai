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

func (lu *LastUsecase) GetLast() ([]domain.Last, error) {
	return lu.last.GetLast()
}
