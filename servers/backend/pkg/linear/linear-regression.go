package linear

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func avg(xs []float64) float64 {
	var r float64
	for _, x := range xs {
		r += x
	}
	r /= float64(len(xs))
	return r
}

// alpha = avg(y) - (avg(x) * beta)
func alpha(x, y []float64) float64 {
	return avg(y) - (avg(x) * beta(x, y))
}

// beta = sigma( (x - avg(x) * (y - avg(y)) ) / sigma( (x - avg(x)) ^ 2 )
func beta(x, y []float64) float64 {
	var m, d float64
	avgX, avgY := avg(x), avg(y)
	for i := 0; i < len(x); i++ {
		m += (x[i] - avgX) * (y[i] - avgY)
		d += math.Pow((x[i] - avgX), 2)
	}
	if d == 0 {
		d = 1
	}
	return (m / d)
}

// LinearRegression line regression formula y = alpha + (beta * x)
func linearRegression(x, y []float64, value float64) float64 {
	return alpha(x, y) + (beta(x, y) * value)
}

func ComputeLinearRegression(value, parameter string, input float64) float64 {
	var valueArray []float64
	valueStrArray := strings.Split(value, ",")
	for _, Str := range valueStrArray {
		temp, _ := strconv.ParseFloat(Str, 64)
		valueArray = append(valueArray, temp)
	}
	var parameterArray []float64
	parameterStrArray := strings.Split(parameter, ",")
	for _, Str := range parameterStrArray {
		temp, _ := strconv.ParseFloat(Str, 64)
		parameterArray = append(parameterArray, temp)
	}
	return linearRegression(parameterArray, valueArray, input)
}

func ComputeTwoPointLinearRegression(value, parameter string, input float64) float64 {
	points, _ := parsePoints(parameter, value)

	closest1, closest2 := findClosestPoints(points, input)

	return linearRegression(
		[]float64{closest1.X, closest2.X},
		[]float64{closest1.Y, closest2.Y},
		input)
}

type Point struct {
	X float64
	Y float64
}

type PointsByX []Point

func (p PointsByX) Len() int           { return len(p) }
func (p PointsByX) Less(i, j int) bool { return p[i].X < p[j].X }
func (p PointsByX) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func parsePoints(xStr, yStr string) ([]Point, error) {
	xValues := strings.Split(xStr, ",")
	yValues := strings.Split(yStr, ",")

	if len(xValues) != len(yValues) {
		return nil, fmt.Errorf("X軸組與Y軸組的數量不一致")
	}

	points := make([]Point, len(xValues))

	for i := 0; i < len(xValues); i++ {
		x, err := strconv.ParseFloat(xValues[i], 64)
		if err != nil {
			return nil, fmt.Errorf("解析X軸數值時發生錯誤: %w", err)
		}

		y, err := strconv.ParseFloat(yValues[i], 64)
		if err != nil {
			return nil, fmt.Errorf("解析Y軸數值時發生錯誤: %w", err)
		}

		points[i] = Point{X: x, Y: y}
	}

	return points, nil
}

func findClosestPoints(points []Point, x float64) (Point, Point) {
	sort.Sort(PointsByX(points))            // 按照 X 軸值進行排序
	points = removeDuplicateElement(points) // 移除重複的點

	// 處理不足兩個點的情況
	if len(points) <= 1 {
		return points[0], points[0]
	}

	if x < points[0].X {
		return points[0], points[1]
	}

	if x > points[len(points)-1].X {
		return points[len(points)-2], points[len(points)-1]
	}

	// 這裡的演算法是從左到右，找到第一個比 x 大的點，然後往前一個點就是最接近 x 的點
	for i := 0; i < len(points)-1; i++ {
		if points[i].X <= x && x <= points[i+1].X {
			return points[i], points[i+1]
		}
	}

	return Point{}, Point{}
}

func removeDuplicateElement(addrs []Point) []Point {
	result := make([]Point, 0, len(addrs))
	temp := map[float64]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item.X]; !ok {
			temp[item.X] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
