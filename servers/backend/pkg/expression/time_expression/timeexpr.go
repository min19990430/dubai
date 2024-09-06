package timeexpression

import (
	"time"

	"github.com/expr-lang/expr"
)

// 預設的function
// https://expr.medv.io/docs/Language-Definition#date-manipulation
// https://expr.medv.io/docs/Language-Definition#now
// https://expr.medv.io/docs/Language-Definition#durationv
// https://expr.medv.io/docs/Language-Definition#datev-format-timezone

// https://expr.medv.io/docs/Language-Definition#now
// var Now = expr.Function(
// 	"now",
// 	func(params ...any) (any, error) {
// 		return time.Now(), nil
// 	})

var Add = expr.Function(
	"add",
	func(params ...any) (any, error) {
		return params[0].(time.Time).Add(params[1].(time.Duration)), nil
	})

var Sub = expr.Function(
	"sub",
	func(params ...any) (any, error) {
		return params[0].(time.Time).Sub(params[1].(time.Time)), nil
	})

var Before = expr.Function(
	"before",
	func(params ...any) (any, error) {
		return params[0].(time.Time).Before(params[1].(time.Time)), nil
	})

var After = expr.Function(
	"after",
	func(params ...any) (any, error) {
		return params[0].(time.Time).After(params[1].(time.Time)), nil
	})

var Equal = expr.Function(
	"equal",
	func(params ...any) (any, error) {
		return params[0].(time.Time).Equal(params[1].(time.Time)), nil
	})

var Since = expr.Function(
	"since",
	func(params ...any) (any, error) {
		return time.Since(params[0].(time.Time)), nil
	})

var Options = []expr.Option{
	Add,
	Sub,
	Before,
	After,
	Equal,
	Since,
}

func CalculateTime(formula string, inputTime time.Time) (time.Time, error) {
	program, compileErr := expr.Compile(formula, Options...)
	if compileErr != nil {
		return time.Time{}, compileErr
	}

	env := map[string]interface{}{
		"inputTime": inputTime,
	}

	output, runErr := expr.Run(program, env)
	if runErr != nil {
		return time.Time{}, runErr
	}

	// 檢查output是否為time.Time
	if val, ok := output.(time.Time); ok {
		return val, nil
	}
	return time.Time{}, nil
}

func CalculateDuration(formula string, inputTime time.Time) (time.Duration, error) {
	program, compileErr := expr.Compile(formula, Options...)
	if compileErr != nil {
		return 0, compileErr
	}

	env := map[string]interface{}{
		"inputTime": inputTime,
	}

	output, runErr := expr.Run(program, env)
	if runErr != nil {
		return 0, runErr
	}

	// 檢查output是否為time.Time
	if val, ok := output.(time.Duration); ok {
		return val, nil
	}
	return 0, nil
}

func IsTrue(formula string, input map[string]interface{}) (bool, error) {
	program, compileErr := expr.Compile(formula, Options...)
	if compileErr != nil {
		return false, compileErr
	}

	output, runErr := expr.Run(program, input)
	if runErr != nil {
		return false, runErr
	}

	// 檢查output是否為bool
	if val, ok := output.(bool); ok {
		return val, nil
	}
	return false, nil
}
