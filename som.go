package gosom

import (
	"fmt"
	"math/rand"
)

const (
	// SumOfSquare rappresent the default distance parameter
	SumOfSquare = iota
)

// Som is the SOM object
type Som struct {
	Grid                   *Grid
	R                      []float64
	DefaultDistance        int
	CodeBooks              [][]float64
	Alpha                  []float64
	Epoch                  int
	NeighbourhoodDistances [][]float64
	distanceFunction       func([]float64, []float64) float64

	// Nr codebooks len(points[0])
}

// NewSom create a new Som object
func NewSom(x int, y int, topology Topology, alpha []float64, epoch int, neighbourhood func(float64, float64) float64, toroidal bool, r float64, distanceFunction func([]float64, []float64) float64) (*Som, error) {
	g, err := NewGrid(x, y, topology, neighbourhood, toroidal)
	if err != nil {
		return nil, err
	}
	var s [][]float64
	if r == 0.0 {
		s = unitDistances(g)
		r = calculateQuantile(s, 2.0/3.0)
	}
	if alpha == nil {
		alpha = []float64{0.05, 0.01}
	}
	if epoch == 0 {
		epoch = 100
	}
	return &Som{
		Grid:                   g,
		R:                      []float64{r, 0},
		DefaultDistance:        SumOfSquare,
		distanceFunction:       distanceFunction,
		Alpha:                  alpha,
		Epoch:                  epoch,
		NeighbourhoodDistances: s,
	}, nil
}

// Train do the actual train of the som
func (s *Som) Train(data [][]float64) error {
	// numobject == len(data)
	// numlayers == --> SE FORKY HA CAPITO GIUSTORKY è 1
	// numCodes == len(s.CodeBooks)
	// totalvars == len(data[0])
	s.initialiseCodeBooks(data)
	var curIter int
	iterations := s.Epoch * len(data)
	// Outer loop: number of iterations
	for i := 0; i < s.Epoch; i++ {
		//  Inner loop: loop over (bootstrapped) objects
		for y := 0; y < len(data); y++ {
			// Select random object
			pos := rand.Intn(len(data))
			//
			// dsupp := fromMatrixToSlice(data)
			// obj := dsupp[pos*len(data[0])]
			// Find best matching unit index and distance
			_, codeBooksPos := s.findBestMatchingUnit(data[pos])
			if codeBooksPos < 0 {
				return fmt.Errorf("can't find codeBooksPos for input number %d", pos)
			}
			//  Linear decays for radius and learning parameter
			tmp := float64(curIter) / float64(iterations)
			threshold := s.R[0] - (s.R[0]-s.R[1])*tmp
			if threshold < 1 {
				threshold = 0.5
			}
			alpha := s.Alpha[0] - (s.Alpha[0]-s.Alpha[1])*tmp
			// Update changes
			csupp := fromMatrixToSlice(s.CodeBooks)
			var distance float64
			for k := 0; k < len(data[pos]); k++ {
				tmp = data[pos][k] - csupp[codeBooksPos*len(data[0])+k]
				distance += tmp * tmp
			}
			// Update all maps
			nsupp := fromMatrixToSlice(s.NeighbourhoodDistances)
			for k := 0; k < len(s.CodeBooks); k++ {
				tmp := s.Grid.Neighbourhood(nsupp[(len(s.CodeBooks)*codeBooksPos)+k], threshold)
				if tmp > 0 {
					for m := 0; m < len(data[0]); m++ {
						s.CodeBooks[k][m] += tmp * alpha * (data[pos][m] - s.CodeBooks[k][m])
					}
				}
			}
			// fmt.Printf("%+v\n", csupp)
			//fromSliceToMatrix(csupp, s.CodeBooks)
			curIter++
		}
	}
	return nil
}

func (s *Som) initialiseCodeBooks(data [][]float64) {
	positions := make(map[int]bool, len(s.Grid.Points[0]))
	for i := 0; i < len(s.Grid.Points[0]); {
		rnd := rand.Intn(len(data))
		exist := positions[rnd]
		if exist {
			continue
		}
		positions[rnd] = true
		i++
	}
	results := make([][]float64, len(positions))
	var i int
	for key := range positions {
		results[i] = data[key]
		i++
	}
	s.CodeBooks = results
}

// findBestMatchingUnit find the BMU
// NOTE: on R code it return also "nind" but we don't get it what this value mean.
func (s *Som) findBestMatchingUnit(data []float64) (float64, int) {
	var bestDistances float64
	var codeBooksPos int
	for i := 0; i < len(s.CodeBooks); i++ {
		distance := s.distanceFunction(s.CodeBooks[i], data)
		if distance < bestDistances || i == 0 {
			bestDistances = distance
			codeBooksPos = i
		}
	}
	return bestDistances, codeBooksPos
}
