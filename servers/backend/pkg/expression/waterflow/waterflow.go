package waterflow

import (
	"math"

	"github.com/expr-lang/expr"
)

const (
	// 坡度最小值
	minGradientSlope = 0.001
)

var Options = []expr.Option{
	Function,
}

var Function = expr.Function(
	"water_flow",
	func(params ...any) (any, error) {
		return waterFlow(params[0].(float64), params[1].(float64), params[2].(float64), params[3].(float64)), nil
	},
)

func waterFlow(waterLevel, radius, coefficient, gradientSlope float64) float64 {
	// waterLevel 不可為0或負數
	if waterLevel <= 0 {
		return 0
	}

	// radius 不可為0或負數
	if radius <= 0 {
		return 0
	}

	// coefficient 不可為0或負數
	if coefficient <= 0 {
		return 0
	}

	// 坡度不可低於0.001
	if gradientSlope <= minGradientSlope {
		gradientSlope = minGradientSlope
	}

	// 水位不可高於半徑，高於半徑則視為滿水位
	if waterLevel > radius {
		waterLevel = radius
	}

	// r = radian / 2
	r := math.Acos((radius - waterLevel) / radius)
	if r == 0 {
		return 0
	}

	// hr 水力半徑 Hydraulic radius
	hr := radius * (r - 0.5*math.Sin(2*r)) / (2 * r)

	// Q = A * V
	v := (1 / coefficient) * math.Pow(hr, 2.0/3.0) * math.Pow(gradientSlope, 0.5)
	// A(角度圓弧面積)
	a := math.Pow(radius, 2)*r - 0.5*math.Pow(radius, 2)*math.Sin(2*r)
	return a * v
}
