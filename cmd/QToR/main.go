package main

import (
	"fmt"
	"github.com/iqsnider/QToR/internal/pde_solvers"
)

func main() {
	nx := 100
	ny := 100
	lx := 0.02
	ly := 0.02
	solver := pde_solvers.NewHelmholtzSolver(nx, ny, lx, ly)
	fmt.Println(solver)
}
