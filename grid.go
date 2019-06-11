package gosom

import (
	"fmt"
	"math"
)

// Topology rappresent the different topology avaiable in gosom
type Topology int

const (
	// Hexagonal Rappresent the Hexagonal grid
	Hexagonal Topology = iota
	// Rectangular Rappresent the Rectangular grid
	Rectangular
)

// Grid rappresent a grid of unity of a specified size and topology
type Grid struct {
	Points        [][]float64
	LenX          int
	LenY          int
	Topology      Topology
	Neighbourhood func(float64, float64) float64
	Toroidal      bool
}

// NewGrid create a new Grid
func NewGrid(x int, y int, topology Topology, neighbourhood func(float64, float64) float64, toroidal bool) (*Grid, error) {
	if toroidal == true {
		return nil, fmt.Errorf("Not yet supported")
	}
	somGrid := &Grid{
		LenX:          x,
		LenY:          y,
		Toroidal:      toroidal,
		Points:        make([][]float64, 2),
		Topology:      topology,
		Neighbourhood: neighbourhood,
	}
	somGrid.Points[0] = make([]float64, x*y)
	somGrid.Points[1] = make([]float64, x*y)

	for k := 0; k < len(somGrid.Points[0]); k++ {
		somGrid.Points[0][k] = float64(k%x) + 1
	}
	var z int
	for k := 0; k < len(somGrid.Points[1]); k++ {
		if k%x == 0 {
			z++
		}
		somGrid.Points[1][k] = float64(z)
	}

	if topology == Hexagonal {
		supp := math.Sqrt(3.0) / 2.0
		for i := 0; i < len(somGrid.Points[0]); i++ {
			somGrid.Points[0][i] = somGrid.Points[0][i] + 0.5*float64(int(somGrid.Points[1][i])%2)
			somGrid.Points[1][i] = somGrid.Points[1][i] * supp
		}
	}
	return somGrid, nil
}
