package gosom

import (
	"math"
)

// Gaussian do the gaussian
func Gaussian(distance float64, r float64) float64 {
	return math.Exp(-(distance * distance) / (2 * r * r))
}

// Bubble is the generalisation of the lambda function
func Bubble(distance float64, r float64) float64 {
	if distance <= r {
		return 1
	}
	return 0
}
