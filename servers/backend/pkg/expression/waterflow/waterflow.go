package waterflow

import (
	"errors"
	"math"

	"github.com/expr-lang/expr"
)

const (
	// 坡度最小值
	minGradientSlope = 0.001
)

var Options = []expr.Option{
	SemiCircularWaterFlowFunc,
	RectangularWaterFlowFunc,
}

var SemiCircularWaterFlowFunc = expr.Function(
	"SemiCircularWaterFlow",
	func(params ...any) (any, error) {
		return semiCircularWaterFlow(params[0].(float64), params[1].(float64), params[2].(float64), params[3].(float64)), nil
	},
)

func semiCircularWaterFlow(waterLevel, radius, coefficient, gradientSlope float64) float64 {
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

var RectangularWaterFlowFunc = expr.Function(
	"RectangularWaterFlow",
	func(params ...any) (any, error) {
		return rectangularWaterFlow(params[0].(float64), params[1].(float64), params[2].(float64), params[3].(float64))
	},
)

// rectangularWaterFlow 矩形水流量計算
// B: 渠道寬度(m)
// b: 堰板缺口寬度(m)
// D: 渠道底面到堰頂的垂直距離(m)
// h: 水頭(m)
func rectangularWaterFlow(B, b, D, h float64) (float64, error) {
	if err := paramsCheck(B, b, D, h); err != nil {
		return 0, err
	}
	K := calculateK(h, D, B, b)
	return calculateFlow(b, h, K), nil
}

func paramsCheck(B, b, D, h float64) error {
	switch {
	case !(B >= 0.5 && B <= 6.3):
		return errors.New("B must be between 0.5 and 6.3")
	case !(b >= 0.15 && b <= 5):
		return errors.New("b must be between 0.1 and 1.5")
	case !(D >= 0.15 && D <= 3.5):
		return errors.New("D must be between 0.15 and 3.5")
	case !((b*D)/(B*B) >= 0.06):
		return errors.New("b*D/B^2 must be greater than 0.06")
	case !(h >= 0.03 && h <= 0.45*math.Sqrt(b)):
		return errors.New("h must be between 0.03 and 0.45*sqrt(b)")
	}
	return nil
}

// calculateK 計算K值
// h: 水頭(m)
// D: 渠道底面到堰頂的垂直距離(m)
// B: 渠道寬度(m)
// b: 堰板缺口寬度(m)
func calculateK(h, D, B, b float64) float64 {
	return 107.1 + 0.177/h + 14.2*(h/D) - 25.7*math.Sqrt((B-b)*h/(D*B)) + 2.04*math.Sqrt(B/D)
}

// calculateFlow 計算流量
// b: 堰板缺口寬度(m)
// h: 水頭(m)
// K: K值
func calculateFlow(b, h float64, K float64) float64 {
	return K * b * math.Pow(h, 1.5)
}
