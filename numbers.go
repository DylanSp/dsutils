package dsutils

import "math"

func IntMax(a int, b int) int {
	floatMax := math.Max(float64(a), float64(b))
	return int(floatMax)
}
