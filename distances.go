package gosom

import (
	"math"
)

// SumOfSquareDistance calculate the sum of square distance
func SumOfSquareDistance(data, codes []float64) float64 {
	var result float64
	for i := 0; i < len(data); i++ {
		result += (data[i] - codes[i]) * (data[i] - codes[i])
	}
	return result
}

// EuclideanDistance calculate the euclidean distance
func EuclideanDistance(data, codes []float64) float64 {
	return math.Sqrt(SumOfSquareDistance(data, codes))
}

// ManhattanDistance calculate the Manhattan distance
func ManhattanDistance(data, codes []float64) float64 {
	var result float64
	for i := 0; i < len(data); i++ {
		result += math.Abs((data[i] - codes[i]))
	}
	return result
}

// TanimotoDistance calculate the Tanimoto distance
func TanimotoDistance(data, codes []float64) float64 {
	var supp float64
	for i := 0; i < len(data); i++ {
		if (data[i] > 0.5 && codes[i] < 0.5) || (data[i] <= 0.5 && codes[i] >= 0.5) {
			supp++
		}
	}
	return supp / float64(len(data))
}
