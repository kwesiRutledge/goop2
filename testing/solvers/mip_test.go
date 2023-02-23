package solvers_test

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

func solveSimpleMIPModel(t *testing.T, solver optim.Solver) {
	m := optim.NewModel()
	m.ShowLog(false)
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()
	z := m.AddBinaryVariable()

	// Constraint 1
	tempSum, err := x.Plus(y.Mult(2))
	tempSum, _ = tempSum.Plus(z.Mult(3))
	if err != nil {
		t.Errorf("There was an issue creating temporary sum: %v", err)
	}
	tempSumAsSLE, _ := tempSum.(optim.ScalarLinearExpr)
	m.AddConstr(tempSumAsSLE.LessEq(optim.K(4)))

	// Constraint 2
	tempSum, err = x.Plus(y)
	if err != nil {
		t.Errorf("There was an issue creating second sum: %v", err)
	}
	tempSumAsSLE, _ = tempSum.(optim.ScalarLinearExpr)
	m.AddConstr(tempSumAsSLE.GreaterEq(optim.One))

	obj, _ := tempSum.Plus(z.Mult(2))
	objAsSLE, _ := obj.(optim.ScalarLinearExpr)
	m.SetObjective(objAsSLE, optim.SenseMaximize)
	sol, err := m.Optimize(solver)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("x =", sol.Value(x))
	t.Log("y =", sol.Value(y))
	t.Log("z =", sol.Value(z))
}

func solveSumRowsColsModel(t *testing.T, solver optim.Solver) {
	m := optim.NewModel()
	m.ShowLog(false)
	rows := 4
	cols := 4
	vs := m.AddBinaryVariableMatrix(rows, cols)

	for i := 0; i < cols; i++ {
		m.AddConstr(optim.SumCol(vs, i).Eq(optim.One))
	}

	for i := 0; i < rows; i++ {
		m.AddConstr(optim.SumRow(vs, i).Eq(optim.One))
	}

	sol, err := m.Optimize(solver)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prettyPrintVarMatrix(vs, sol))
}

func prettyPrintVarMatrix(vs [][]optim.Variable, sol *optim.Solution) string {
	rows := len(vs)
	cols := len(vs[0])

	matStr := ""
	for i := 0; i < rows; i++ {
		rowStr := ""
		for j := 0; j < cols; j++ {
			if sol.Value(vs[i][j]) > 0.1 {
				rowStr += "1 "
			} else {
				rowStr += "0 "
			}
		}
		matStr += rowStr + "\n"
	}

	return matStr
}
