package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
	"auto-monitoring/pkg/expression"
)

type PhysicalQuantityEvaluateUsecase struct {
	physicalQuantityEvaluateRepository irepository.IPhysicalQuantityEvaluateRepository
}

func NewPhysicalQuantityEvaluateUsecase(
	physicalQuantityEvaluateRepository irepository.IPhysicalQuantityEvaluateRepository,
) *PhysicalQuantityEvaluateUsecase {
	return &PhysicalQuantityEvaluateUsecase{
		physicalQuantityEvaluateRepository: physicalQuantityEvaluateRepository,
	}
}

func (pqu *PhysicalQuantityEvaluateUsecase) Evaluate(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate, value float64) (float64, error) {
	return expression.Calculate(physicalQuantityEvaluate.Formula, value)
}

func (pqu *PhysicalQuantityEvaluateUsecase) UpdateFormula(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate, formulaType string, params map[string]any) error {
	formulaTypeCode, err := FormulaTypeFromString(formulaType)
	if err != nil {
		return err
	}

	formulaTypeFunc, err := FormulaTypeFactory{}.CreateFormulaTypeFunc(formulaTypeCode)
	if err != nil {
		return err
	}

	formula, err := formulaTypeFunc.SettingFormula(params)
	if err != nil {
		return err
	}

	newPhysicalQuantityEvaluate := domain.PhysicalQuantityEvaluate{
		UUID:        physicalQuantityEvaluate.UUID,
		FormulaType: formulaTypeCode.String(),
		Formula:     formula,
	}

	return pqu.physicalQuantityEvaluateRepository.UpdateFormula(newPhysicalQuantityEvaluate)
}

type PhysicalQuantityEvaluateDetailUsecase struct {
	physicalQuantityRepository               irepository.IPhysicalQuantityRepository
	physicalQuantityEvaluateDetailRepository irepository.IPhysicalQuantityEvaluateDetailRepository
}

func NewPhysicalQuantityEvaluateDetailUsecase(
	physicalQuantityRepository irepository.IPhysicalQuantityRepository,
	physicalQuantityEvaluateDetailRepository irepository.IPhysicalQuantityEvaluateDetailRepository,
) *PhysicalQuantityEvaluateDetailUsecase {
	return &PhysicalQuantityEvaluateDetailUsecase{
		physicalQuantityRepository:               physicalQuantityRepository,
		physicalQuantityEvaluateDetailRepository: physicalQuantityEvaluateDetailRepository,
	}
}

func (pqdu *PhysicalQuantityEvaluateDetailUsecase) ListByPhysicalQuantity(physicalQuantity domain.PhysicalQuantity) ([]domain.PhysicalQuantityEvaluateDetail, error) {
	physicalQuantityList, err := pqdu.physicalQuantityRepository.List(physicalQuantity)
	if err != nil {
		return nil, err
	}

	var physicalQuantityEvaluateDetailList []domain.PhysicalQuantityEvaluateDetail
	for _, physicalQuantity := range physicalQuantityList {
		physicalQuantityEvaluateDetail, err := pqdu.physicalQuantityEvaluateDetailRepository.ListDetail(
			domain.PhysicalQuantityEvaluate{
				PhysicalQuantityUUID: physicalQuantity.UUID,
			})
		if err != nil {
			return nil, err
		}

		if len(physicalQuantityEvaluateDetail) > 0 {
			physicalQuantityEvaluateDetailList = append(physicalQuantityEvaluateDetailList, physicalQuantityEvaluateDetail...)
		}
	}

	return physicalQuantityEvaluateDetailList, nil
}

func (pqdu *PhysicalQuantityEvaluateDetailUsecase) ListDetail(physicalQuantityEvaluate domain.PhysicalQuantityEvaluate) ([]domain.PhysicalQuantityEvaluateDetail, error) {
	return pqdu.physicalQuantityEvaluateDetailRepository.ListDetail(physicalQuantityEvaluate)
}
