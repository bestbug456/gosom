package gosom

import "testing"

func TestRectangularGrid(t *testing.T) {
	grid, _ := NewGrid(4, 5, Rectangular, nil, false)

	X := []float64{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	Y := []float64{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}

	for i := 0; i < len(X); i++ {
		if grid.Points[0][i] != X[i] {
			t.Fatalf("Expected %v but having %v", X[i], grid.Points[0][i])
		}
	}

	for i := 0; i < len(Y); i++ {
		if grid.Points[1][i] != Y[i] {
			t.Fatalf("Expected %v but having %v", Y[i], grid.Points[1][i])
		}
	}
}

func TestHexagonalGrid(t *testing.T) {
	grid, _ := NewGrid(4, 5, Hexagonal, nil, false)

	X := []float64{1.5, 2.5, 3.5, 4.5, 1.0, 2.0, 3.0, 4.0, 1.5, 2.5, 3.5, 4.5, 1.0, 2.0, 3.0, 4.0, 1.5, 2.5, 3.5, 4.5}
	Y := []float64{
		0.8660254037844386, 0.8660254037844386, 0.8660254037844386, 0.8660254037844386,
		1.7320508075688772, 1.7320508075688772, 1.7320508075688772, 1.7320508075688772,
		2.598076211353316, 2.598076211353316, 2.598076211353316, 2.598076211353316,
		3.4641016151377544, 3.4641016151377544, 3.4641016151377544, 3.4641016151377544,
		4.330127018922193, 4.330127018922193, 4.330127018922193, 4.330127018922193}

	for i := 0; i < len(X); i++ {
		if grid.Points[0][i] != X[i] {
			t.Fatalf("Expected %v but having %v", X[i], grid.Points[0][i])
		}
	}

	for i := 0; i < len(Y); i++ {
		if grid.Points[1][i] != Y[i] {
			t.Fatalf("Expected %v but having %v", Y[i], grid.Points[1][i])
		}
	}
}
