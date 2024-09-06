package expression

import (
	"errors"

	"auto-monitoring/pkg/expression/waterflow"

	"github.com/expr-lang/expr"
)

func registerOptions() []expr.Option {
	return waterflow.Options
}

func Calculate(formula string, x float64) (float64, error) {
	options := registerOptions()

	program, compileErr := expr.Compile(formula, options...)
	if compileErr != nil {
		return 0, compileErr
	}

	env := map[string]interface{}{
		"X": x,
	}

	output, runErr := expr.Run(program, env)
	if runErr != nil {
		return 0, runErr
	}

	// 檢查output是否為float64
	if val, ok := output.(float64); ok {
		return val, nil
	}
	return 0, errors.New("output is not a float64")
}

func IsTrue(formula string, x float64) (bool, error) {
	options := registerOptions()

	program, compileErr := expr.Compile(formula, options...)
	if compileErr != nil {
		return false, compileErr
	}

	env := map[string]interface{}{
		"X": x,
	}

	output, runErr := expr.Run(program, env)
	if runErr != nil {
		return false, runErr
	}

	// 檢查output是否為bool
	if val, ok := output.(bool); ok {
		return val, nil
	}
	return false, errors.New("output is not a bool")
}
