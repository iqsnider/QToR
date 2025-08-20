package main

import (
	"fmt"
	"github.com/iqsnider/QToR/internal/pde_solvers"
)

func main() {
	nx := 5
	ny := 5
	lx := 0.02
	ly := 0.02
	solver := pde_solvers.NewHelmholtzSolver(nx, ny, lx, ly)

	numericModes := solver.SolveWaveguideModes(2)
	fmt.Println(numericModes)
}
