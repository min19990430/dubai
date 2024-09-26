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
// B: channelWidth 渠道寬度(m)
// b: notchWidth 堰板缺口寬度(m)
// D: channelDepth 渠道底面到堰頂的垂直距離(m)
// h: head 水頭(m)
func rectangularWaterFlow(channelWidth, notchWidth, channelDepth, head float64) (float64, error) {
	if err := paramsCheck(channelWidth, notchWidth, channelDepth, head); err != nil {
		return 0, err
	}
	K := calculateFlowCoeff(head, channelDepth, channelWidth, notchWidth)
	return calculateFlow(notchWidth, head, K), nil
}

func paramsCheck(channelWidth, notchWidth, channelDepth, head float64) error {
	switch {
	case !(channelWidth >= 0.5 && channelWidth <= 6.3):
		return errors.New(" B must be between 0.5 and 6.3")
	case !(notchWidth >= 0.15 && notchWidth <= 5):
		return errors.New("b must be between 0.1 and 1.5")
	case !(channelDepth >= 0.15 && channelDepth <= 3.5):
		return errors.New(" D must be between 0.15 and 3.5")
	case !((notchWidth*channelDepth)/(channelWidth*channelWidth) >= 0.06):
		return errors.New("b*D/B^2 must be greater than 0.06")
	case !(head >= 0.03 && head <= 0.45*math.Sqrt(notchWidth)):
		return errors.New("h must be between 0.03 and 0.45*sqrt(b)")
	}
	return nil
}

// calculateFlowCoeff 計算K值
// h: head 水頭(m)
// D: channelDepth 渠道底面到堰頂的垂直距離(m)
// B: channelWidth 渠道寬度(m)
// b: notchWidth 堰板缺口寬度(m)
func calculateFlowCoeff(head, channelDepth, channelWidth, notchWidth float64) float64 {
	return 107.1 + 0.177/head + 14.2*(head/channelDepth) - 25.7*math.Sqrt((channelWidth-notchWidth)*head/(channelDepth*channelWidth)) + 2.04*math.Sqrt(channelWidth/channelDepth)
}

// calculateFlow 計算流量
// b: notchWidth 堰板缺口寬度(m)
// h: head 水頭(m)
// K: flowCoeff K值
func calculateFlow(notchWidth, head float64, flowCoeff float64) float64 {
	return flowCoeff * notchWidth * math.Pow(head, 1.5)
}
