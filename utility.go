package gosom

import (
	"math"
	"sort"
)

func unitDistances(g *Grid) [][]float64 {

	if g.Toroidal == false {
		if g.Topology == Hexagonal {
			return calculateMatrixDistances(g, false)
		}
		return calculateMatrixDistances(g, true)
	}

	// TBD
	return nil
}

func calculateMatrixDistances(g *Grid, floor bool) [][]float64 {
	ris := make([][]float64, len(g.Points[1]))
	for i := 0; i < len(ris); i++ {
		if len(ris[i]) == 0 {
			ris[i] = make([]float64, len(g.Points[0]))
		}
		for y := 0; y < len(ris[i]); y++ {
			p := []float64{g.Points[0][i], g.Points[1][i]}
			q := []float64{g.Points[0][y], g.Points[1][y]}
			ris[i][y] = calculateDistanceBetween2Points(p, q, floor)
		}
	}
	return ris
}

func calculateDistanceBetween2Points(p, q []float64, floor bool) float64 {
	if floor == false {
		return math.Sqrt(math.Pow(p[0]-q[0], 2) + math.Pow(p[1]-q[1], 2))
	}
	return float64(int(math.Sqrt(math.Pow(p[0]-q[0], 2) + math.Pow(p[1]-q[1], 2))))
}

func calculateQuantile(matrix [][]float64, probability float64) float64 {
	sorted := make([]float64, len(matrix)*len(matrix[0]))
	var counter int
	for i := 0; i < len(matrix); i++ {
		for y := 0; y < len(matrix[i]); y++ {
			sorted[counter] = matrix[i][y]
			counter++

		}
	}
	sort.Float64s(sorted)

	pos := (float64(len(sorted)) - 1) * probability
	posFloor := math.Floor(pos)
	posCeiling := math.Ceil(pos)
	h := pos - posFloor
	return ((1 - h) * sorted[int(posFloor)]) + (h * sorted[int(posCeiling)])
}

func fromMatrixToSlice(data [][]float64) []float64 {
	ris := make([]float64, len(data)*len(data[0]))
	var pos int
	for i := 0; i < len(data[0]); i++ {
		for y := 0; y < len(data); y++ {
			ris[pos] = data[y][i]
			pos++
		}
	}
	return ris
}

func fromSliceToMatrix(data []float64, matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for y := 0; y < len(matrix[0]); y++ {
			matrix[i][y] = data[(i*len(matrix[0]) + y)]
		}
	}
}
