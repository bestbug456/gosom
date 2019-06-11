package gosom

import (
	"testing"
)

func TestUnitDistances(t *testing.T) {
	grid, _ := NewGrid(2, 3, Rectangular, nil, false)
	ris := unitDistances(grid)
	expected := [][]float64{
		[]float64{0, 1, 1, 1, 2, 2},
		[]float64{1, 0, 1, 1, 2, 2},
		[]float64{1, 1, 0, 1, 1, 1},
		[]float64{1, 1, 1, 0, 1, 1},
		[]float64{2, 2, 1, 1, 0, 1},
		[]float64{2, 2, 1, 1, 1, 0},
	}
	for i := 0; i < len(expected); i++ {
		for y := 0; y < len(expected[i]); y++ {
			if ris[i][y] != expected[i][y] {
				t.Fatalf("Error: element in position %d-%d should be %v but is %v (%+v)", i, y, expected[i][y], ris[i][y], ris)
			}
		}
	}
}

func TestUnitDistancesHexagonal(t *testing.T) {
	grid, _ := NewGrid(2, 3, Hexagonal, nil, false)
	ris := unitDistances(grid)

	expected := [][]float64{
		[]float64{0, 1, 0.9999999999999999, 0.9999999999999999, 1.7320508075688774, 2},
		[]float64{1, 0, 1.7320508075688772, 0.9999999999999999, 2, 1.7320508075688774},
		[]float64{0.9999999999999999, 1.7320508075688772, 0, 1, 1.0000000000000002, 1.7320508075688774},
		[]float64{0.9999999999999999, 0.9999999999999999, 1, 0, 1.0000000000000002, 1.0000000000000002},
		[]float64{1.7320508075688774, 2, 1.0000000000000002, 1.0000000000000002, 0, 1},
		[]float64{2, 1.7320508075688774, 1.7320508075688774, 1.0000000000000002, 1, 0},
	}

	for i := 0; i < len(expected); i++ {
		for y := 0; y < len(expected[i]); y++ {
			if ris[i][y] != expected[i][y] {
				t.Fatalf("Error: element in position %d-%d should be %v but is %v (%+v)", i, y, expected[i][y], ris[i][y], ris)
			}
		}
	}
}

func TestQuantile(t *testing.T) {
	grid, _ := NewGrid(2, 3, Hexagonal, nil, false)
	ris := unitDistances(grid)
	p := calculateQuantile(ris, 2.0/3.0)
	if p != 1.2440169358562918 {
		t.Fatalf("Error: expected %v but having %v", 1.2440169358562918, p)
	}
}
