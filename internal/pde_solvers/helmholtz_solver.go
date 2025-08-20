package pde_solvers

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
