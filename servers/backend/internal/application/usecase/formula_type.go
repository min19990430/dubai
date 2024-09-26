package usecase

import (
	"errors"
	"fmt"
	"math"
)

type FormulaType int

const (
	Normal                FormulaType = iota
	SemiCircularWaterFlow             // 半圓水流
	RectangularWaterFlow              // 矩形水流
)

var formulaTypeMap = map[FormulaType]string{
	Normal:                "Normal",
	SemiCircularWaterFlow: "SemiCircularWaterFlow",
	RectangularWaterFlow:  "RectangularWaterFlow",
}

func (f FormulaType) String() string {
	return formulaTypeMap[f]
}

func (f FormulaType) IsValid() bool {
	_, ok := formulaTypeMap[f]
	return ok
}

func FormulaTypeFromString(s string) (FormulaType, error) {
	for k, v := range formulaTypeMap {
		if s == v {
			return k, nil
		}
	}

	return -1, errors.New(fmt.Sprintf("invalid formula type: %s", s))
}

type FormulaTypeFactory struct{}

func (f FormulaTypeFactory) CreateFormulaTypeFunc(formulaTypeCode FormulaType) (FormulaTypeFunc, error) {
	switch formulaTypeCode {
	case Normal:
		return NormalFormula{}, nil
	case SemiCircularWaterFlow:
		return SemiCircularWaterFlowFormula{}, nil
	case RectangularWaterFlow:
		return RectangularWaterFlowFormula{}, nil
	default:
		return nil, fmt.Errorf("invalid formula type: %d", formulaTypeCode)
	}
}

type FormulaTypeFunc interface {
	SettingFormula(params map[string]any) (string, error)
}

type NormalFormula struct{}

func (n NormalFormula) SettingFormula(params map[string]any) (string, error) {
	// 檢查參數是否正確
	if _, ok := params["formula"]; !ok {
		return "", errors.New("missing formula")
	}

	// 檢查formula是否為string
	if formula, ok := params["formula"].(string); ok {
		return formula, nil
	}

	return "", errors.New("formula is not a string")
}

type SemiCircularWaterFlowFormula struct{}

func (s SemiCircularWaterFlowFormula) SettingFormula(params map[string]any) (string, error) {
	if _, ok := params["radius"]; !ok {
		return "", errors.New("missing radius")
	}

	radius, ok := params["radius"].(float64)
	if ok {
		if radius <= 0 {
			return "", errors.New("radius must be greater than 0")
		}
	} else {
		return "", errors.New("radius is not a float64")
	}

	if _, ok := params["coefficient"]; !ok {
		return "", errors.New("missing coefficient")
	}

	coefficient, ok := params["coefficient"].(float64)
	if ok {
		if coefficient <= 0 {
			return "", errors.New("coefficient must be greater than 0")
		}
	} else {
		return "", errors.New("coefficient is not a float64")
	}

	if _, ok := params["gradient_slope"]; !ok {
		return "", errors.New("missing gradient_slope")
	}

	gradientSlope, ok := params["gradient_slope"].(float64)
	if ok {
		if gradientSlope <= 0 {
			return "", errors.New("gradient_slope must be greater than 0")
		}
	} else {
		return "", errors.New("gradient_slope is not a float64")
	}

	return fmt.Sprintf("RectangularWaterFlow(X/100, %f, %f, %f)*3600", radius, coefficient, gradientSlope), nil
}

type RectangularWaterFlowFormula struct{}

func (r RectangularWaterFlowFormula) SettingFormula(params map[string]any) (string, error) {
	B, err := validateParam(params, "B")
	if err != nil {
		return "", err
	}

	b, err := validateParam(params, "b")
	if err != nil {
		return "", err
	}

	D, err := validateParam(params, "D")
	if err != nil {
		return "", err
	}

	_, ok := params["N"]
	if !ok {
		return "", errors.New("missing N")
	}

	floatN, ok := params["N"].(float64)
	if !ok {
		return "", errors.New("N is not an number")
	}

	// 檢查 N 是否有小數點
	if floatN != math.Floor(floatN) {
		return "", errors.New("N should be an integer without decimal points")
	}

	N := int(floatN)

	return fmt.Sprintf("X>=0.03&&X<=0.45*%.2f^0.5?RectangularWaterFlow(%.2f, %.2f, %.2f, X)*60*%d:0.0", b, B, b, D, N), nil
}

func validateParam(params map[string]any, key string) (float64, error) {
	if _, ok := params[key]; !ok {
		return 0, fmt.Errorf("missing %s", key)
	}

	value, ok := params[key].(float64)
	if !ok {
		return 0, fmt.Errorf("%s is not a float64", key)
	}

	if value <= 0 {
		return 0, fmt.Errorf("%s must be greater than 0", key)
	}

	return value, nil
}
