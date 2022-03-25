package solvers

import (
	"testing"
)

func TestLPSolve(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t, NewGurobiSolver())
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t, NewGurobiSolver())
	})
}
