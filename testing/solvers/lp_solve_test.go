package solvers_test

import (
	"github.com/kwesiRutledge/goop2/solvers"
	"testing"
)

func TestLPSolve(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t, solvers.NewGurobiSolver())
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t, solvers.NewGurobiSolver())
	})
}
