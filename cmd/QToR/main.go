package main

import (
	"fmt"
	"github.com/iqsnider/QToR/internal/pde_solvers"
)

func printMode(mode pde_solvers.WaveguideMode, modeNum int) {
	fmt.Printf("\n=== Mode %d (Eigenvalue: %.6f) ===\n", modeNum, mode.Eigenvalue)
	field := mode.Field
	for j := len(field) - 1; j >= 0; j-- {
		for i := 0; i < len(field[j]); i++ {
			val := field[j][i]
			if val > 0.5 {
				fmt.Print(" + ")
			} else if val < -0.5 {
				fmt.Print(" - ")
			} else if val > 0.1 {
				fmt.Print(" . ")
			} else if val < -0.1 {
				fmt.Print(" , ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func main() {
	nx := 10
	ny := 10
	lx := 0.02
	ly := 0.02
	solver := pde_solvers.NewHelmholtzSolver(nx, ny, lx, ly)

	numericModes := solver.SolveWaveguideModes(2)
	for i, mode := range numericModes {
		printMode(mode, i+1)
	}
}
