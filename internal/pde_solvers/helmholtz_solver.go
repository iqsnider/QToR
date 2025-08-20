package pde_solvers

import (
	"math"
)

type WaveguideMode struct {
	Eigenvalue float64
	Field      [][]float64
	ModeIndex  [2]int
}

type HelmholtzSolver struct {
	Nx, Ny int     // number of points (Nx*Ny)
	Lx, Ly float64 // domain
	Dx, Dy float64 // point spacing
}

func NewHelmholtzSolver(nx, ny int, lx, ly float64) *HelmholtzSolver {
	return &HelmholtzSolver{
		Nx: nx,
		Ny: ny,
		Lx: lx,
		Ly: ly,
		Dx: lx / float64(nx-1),
		Dy: ly / float64(ny-1),
	}
}

func (hs *HelmholtzSolver) SolveWaveguideModes(numModes int) []WaveguideMode {

	totalInterior := (hs.Nx - 2) * (hs.Ny - 2)
	// pizza
	A := make([][]float64, totalInterior)

	for i := range A {
		A[i] = make([]float64, totalInterior)
	}

	dx2 := hs.Dx * hs.Dx
	dy2 := hs.Dy * hs.Dy

	for j := 1; j < hs.Ny-1; j++ {
		for i := 1; i < hs.Nx-1; i++ {
			row := (j-1)*(hs.Nx-2) + (i - 1)

			//center
			A[row][row] = -2.0/dx2 - 2.0/dy2

			// x neighbors
			if i > 1 {
				col := (j-1)*(hs.Nx-2) + (i - 2)
				A[row][col] = 1.0 / dx2
			}

			if i < hs.Nx-2 {
				col := (j-1)*(hs.Nx-2) + (i)
				A[row][col] = 1.0 / dx2
			}

			// y neighbors
			if j > 1 {
				col := (j-2)*(hs.Nx-2) + (i - 1)
				A[row][col] = 1.0 / dy2
			}

			if j < hs.Ny-2 {
				col := (j)*(hs.Nx-2) + (i - 1)
				A[row][col] = 1.0 / dy2
			}
		}
	}

	modes := hs.powerIterationModes(A, numModes)

	for i := range modes {
		modes[i].Field = hs.vectorToField(modes[i].Field[0])
	}

	return modes
}

func (hs *HelmholtzSolver) powerIterationModes(A [][]float64, numModes int) []WaveguideMode {
	n := len(A)
	modes := make([]WaveguideMode, 0, numModes)

	ACopy := make([][]float64, n)
	for i := range A {
		ACopy[i] = make([]float64, n)
		copy(ACopy[i], A[i])

	}

	for modeIdx := 0; modeIdx < numModes; modeIdx++ {
		eigval, eigvec := hs.powerIteration(ACopy, 1000, 1e-8)

		mode := WaveguideMode{
			Eigenvalue: -eigval,
			Field:      [][]float64{eigvec},
		}

		modes = append(modes, mode)

		hs.deflateMatrix(ACopy, eigvec, eigval)
	}
	return modes
}

func (hs *HelmholtzSolver) powerIteration(A [][]float64, maxIter int, tolerance float64) (float64, []float64) {
	n := len(A)

	x := make([]float64, n)
	for i := range x {
		x[i] = 1.0 + 0.1*math.Sin(float64(i))
	}

	var lambda, lambdaOld float64

	y := make([]float64, n)
	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				y[i] += A[i][j] * x[j]
			}

			numerator, denominator := 0.0, 0.0
			for i := 0; i < n; i++ {
				numerator += x[i] * y[i]
				denominator += x[i] * x[i]
			}

			lambda = numerator / denominator

			// normalize
			norm := math.Sqrt(denominator)
			for i := 0; i < n; i++ {
				x[i] = y[i] / norm
			}

			if iter > 0 && math.Abs(lambda-lambdaOld) < tolerance {
				break
			}
			lambdaOld = lambda
		}
	}
	return lambda, x
}

func (hs *HelmholtzSolver) deflateMatrix(A [][]float64, eigvec []float64, eigval float64) {
	n := len(A)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i][j] -= eigval * eigvec[i] * eigvec[i]
		}
	}
}

func (hs *HelmholtzSolver) vectorToField(vec []float64) [][]float64 {
	field := make([][]float64, hs.Ny)
	for j := range field {
		field[j] = make([]float64, hs.Nx)
	}

	idx := 0
	for j := 1; j < hs.Ny-1; j++ {
		for i := 1; i < hs.Nx-1; i++ {
			if idx < len(vec) {
				field[j][i] = vec[idx]
				idx++
			}
		}
	}

	return field
}
