package helper

import "math"

func Round(num float64) int {
	return int(math.Round(num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Trunc(num*output) / output
}
