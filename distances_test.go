package gosom

import (
	"testing"
)

func TestSumOfSquareDistance(t *testing.T) {
	ris := SumOfSquareDistance([]float64{0.1, 0.2, 0.3, 0.4, 0.5}, []float64{0.1, 0.5, 0.8, 0.9, 1})
	if ris != 0.84 {
		t.Fatalf("Error expected %v but having %v", 0.84, ris)
	}
}

func TestEuclideanDistance(t *testing.T) {
	ris := EuclideanDistance([]float64{0.1, 0.2, 0.3, 0.4, 0.5}, []float64{0.1, 0.5, 0.8, 0.9, 1})
	if ris != 0.916515138991168 {
		t.Fatalf("Error expected %v but having %v", 0.916515138991168, ris)
	}
}

func TestManhattanDistance(t *testing.T) {
	ris := ManhattanDistance([]float64{0.1, 0.2, 0.3, 0.4, 0.5}, []float64{0.1, 0.5, 0.8, 0.9, 1})
	if ris != 1.8 {
		t.Fatalf("Error expected %v but having %v", 1.8, ris)
	}
}

func TestTanimotoDistance(t *testing.T) {
	ris := TanimotoDistance([]float64{0.1, 0.2, 0.3, 0.4, 0.5}, []float64{0.1, 0.5, 0.8, 0.9, 1})
	if ris != 0.8 {
		t.Fatalf("Error expected %v but having %v", 0.8, ris)
	}
}
