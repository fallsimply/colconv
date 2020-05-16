package colconv

import (
	"math"
)

func min(a float64, b ...float64) float64 {
	var min = a

	for _, num := range b {
		min = math.Min(min, num)
	}

	return min
}

func max(a float64, b ...float64) float64 {
	var max = a

	for _, num := range b {
		max = math.Max(max, num)
	}

	return max
}
