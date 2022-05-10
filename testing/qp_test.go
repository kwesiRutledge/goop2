package testing

/*
qp_test.go
Description:
	This file is meant to contain multiple examples of creating quadratic programs using the goop2
	module.
*/

import (
	"fmt"
	"testing"

	"github.com/kwesiRutledge/goop2"
	"github.com/kwesiRutledge/goop2/solvers"
)

/*
TestQP1
Description:
	Create a simple quadratic program with quadratic objective and no constraints on the optimization variable x.
*/
func TestQP1(t *testing.T) {
	// Constants

	// Create Model
	m := goop2.NewModel()

	// Create Optimization Variables
	m.ShowLog(true)
	x := m.AddVar(-10, 10, goop2.Continuous)
	y := m.AddVar(-10, 10, goop2.Continuous)

	// Create Objective
	Q1 := [][]float64{
		[]float64{1.0, 0.0},
		[]float64{0.0, 1.0},
	}
	L1 := []float64{-6.0, -4.0}
	C1 := -9.0 - 4.0
	qe1, err := goop2.NewQuadraticExpr(Q1, L1, C1, []uint64{x.ID, y.ID})

	m.SetObjective(qe1, goop2.SenseMinimize)
	sol, err := m.Optimize(solvers.NewGurobiSolver())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v", x)

	if sol.Value(x) != 3.0 {
		t.Errorf("Expected for the optimal value of x to be 3.0; received %v", sol.Value(x))
	}

	if sol.Value(y) != 2.0 {
		t.Errorf("Expected for the optimal value of y to be 2.0; received %v", sol.Value(y))
	}

}
