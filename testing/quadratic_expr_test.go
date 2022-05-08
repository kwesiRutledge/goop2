package testing

import (
	"strings"
	"testing"

	"github.com/kwesiRutledge/goop2"
)

/*
quadratic_expr_test.go
Description:
	Tests some of the basic functions of the quadraticExpr class.
*/

/*
TestQuadraticExpr_NewQuadraticExpr_q01
Description:
	Tests whether or not the function returns two variables for a simple expression.
*/
func TestQuadraticExpr_NewQuadraticExpr_qb01(t *testing.T) {
	// Constants
	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}
	xIndices1 := []uint64{1, 2}

	_, err := goop2.NewQuadraticExpr_qb0(Q1, xIndices1)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

}

/*
TestQuadraticExpr_NewQuadraticExpr_q02
Description:
	Tests whether or not the NewQuadraticExpr_q0() function gracefully fails when given a badly sized Q matrix.
*/
func TestQuadraticExpr_NewQuadraticExpr_qb02(t *testing.T) {
	// Constants
	Q2 := [][]float64{
		[]float64{1.0, 2.0},
	}
	xIndices2 := []uint64{1, 2}

	_, err := goop2.NewQuadraticExpr_qb0(Q2, xIndices2)
	if err == nil {
		t.Errorf("Expected an error, but there was none!")
	}

	if !strings.Contains(err.Error(), "The number of indices was 2 which did not match the first dimension of QIn (1)") {
		t.Errorf("The wrong error was thrown: %v", err)
	}

}

/*
TestQuadraticExpr_NewQuadraticExpr_q03
Description:
	Tests whether or not the NewQuadraticExpr_q0() function gracefully fails when given a badly sized Q matrix.
	(Wrong number of columns)
*/
func TestQuadraticExpr_NewQuadraticExpr_qb03(t *testing.T) {
	// Constants
	Q3 := [][]float64{
		[]float64{1.0},
		[]float64{3.0},
	}
	xIndices3 := []uint64{1, 2}

	_, err := goop2.NewQuadraticExpr_qb0(Q3, xIndices3)
	if err == nil {
		t.Errorf("Expected an error, but there was none!")
	}

	if !strings.Contains(err.Error(), "The number of indices was 2 which did not match the length of QIn's 0th row (1).") {
		t.Errorf("The wrong error was thrown: %v", err)
	}

}

/*
TestQuadraticExpr_NumVars1
Description:
	Tests whether or not the function returns two variables for a simple expression.
*/
func TestQuadraticExpr_NumVars1(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q1, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	// Number of variables for this quadratic expression should be 2
	if qv1.NumVars() != 2 {
		t.Errorf("Expected for 2 variables to be found in quadratic expression; function says %v variables exist.", qv1.NumVars())
	}
}

/*
TestQuadraticExpr_NumVars2
Description:
	Tests whether or not the function returns three variables for a more complex expression.
*/
func TestQuadraticExpr_NumVars2(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)
	v3 := m.AddVar(-10, 10, goop2.Continuous)

	Q2 := [][]float64{
		[]float64{1.0, 2.0, 3.0},
		[]float64{4.0, 5.0, 6.0},
		[]float64{7.0, 8.0, 9.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q2, []uint64{v1.ID, v2.ID, v3.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	// Number of variables for this quadratic expression should be 2
	if qv1.NumVars() != 3 {
		t.Errorf("Expected for 3 variables to be found in quadratic expression; function says %v variables exist.", qv1.NumVars())
	}
}

/*
TestQuadraticExpr_NumVars3
Description:
	Tests whether or not the function returns one variables for a more complex expression.
*/
func TestQuadraticExpr_NumVars3(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)

	Q3 := [][]float64{
		[]float64{2.3},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q3, []uint64{v1.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	// Number of variables for this quadratic expression should be 2
	if qv1.NumVars() != 1 {
		t.Errorf("Expected for 1 variable to be found in quadratic expression; function says %v variables exist.", qv1.NumVars())
	}
}

/*
TestQuadraticExpr_Vars1
Description:
	Tests whether or not the function returns two variables for a simple expression.
*/
func TestQuadraticExpr_Vars1(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q1, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	// Number of variables for this quadratic expression should be 2
	if len(qv1.Vars()) != 2 {
		t.Errorf("Expected for 2 variables to be found in quadratic expression; function says %v variables exist.", len(qv1.Vars()))
	}

	if tempVars := qv1.Vars(); tempVars[0] != v1.ID {
		t.Errorf("Expected for first ID to be %v; received %v.", v1.ID, tempVars[0])
	}

	if tempVars := qv1.Vars(); tempVars[1] != v2.ID {
		t.Errorf("Expected for first ID to be %v; received %v.", v2.ID, tempVars[1])
	}

}

/*
TestQuadraticExpr_Plus1
Description:
	Tests whether or not the function returns one variable index for a more complex expression.
*/
func TestQuadraticExpr_Plus1(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	Q2 := [][]float64{
		[]float64{1.0, 0.0},
		[]float64{0.0, 1.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q1, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	qv2, err := goop2.NewQuadraticExpr_qb0(Q2, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	e3 := qv1.Plus(qv2)

	qv3, ok := e3.(*goop2.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 2 {
		t.Errorf("Expected for 2 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.Q[0][0] != 2.0 {
		t.Errorf("Expected for Q's (0,0)-th element to be 2.0; received %v", qv3.Q[0][0])
	}

	if qv3.Q[0][1] != qv1.Q[0][1] {
		t.Errorf("Expected for Q's (0,1)-th element to be %v; received %v", qv3.Q[0][1], qv1.Q[0][1])
	}

	if qv3.Q[1][0] != qv1.Q[1][0] {
		t.Errorf("Expected for Q's (1,0)-th element to be %v; received %v", qv3.Q[1][0], qv1.Q[1][0])
	}

	if qv3.Q[1][1] != 5.0 {
		t.Errorf("Expected for Q's (1,1)-th element to be 5.0; received %v", qv3.Q[1][1])
	}

}

/*
TestQuadraticExpr_Plus2
Description:
	Tests whether or not the plus function works
	for a sum of a quadratic expression and a linear expression (no id checking done).
*/
func TestQuadraticExpr_Plus2(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q1, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	L2 := []float64{5.0, 6.0}
	le2 := &goop2.LinearExpr{
		L:        L2,
		C:        0.1,
		XIndices: []uint64{v1.ID, v2.ID},
	}

	e3 := qv1.Plus(le2)

	qv3, ok := e3.(*goop2.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 2 {
		t.Errorf("Expected for 2 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.Q[0][0] != 1.0 {
		t.Errorf("Expected for Q's (0,0)-th element to be 2.0; received %v", qv3.Q[0][0])
	}

	if qv3.Q[0][1] != qv1.Q[0][1] {
		t.Errorf("Expected for Q's (0,1)-th element to be %v; received %v", qv3.Q[0][1], qv1.Q[0][1])
	}

	if qv3.Q[1][0] != qv1.Q[1][0] {
		t.Errorf("Expected for Q's (1,0)-th element to be %v; received %v", qv3.Q[1][0], qv1.Q[1][0])
	}

	if qv3.L[1] != le2.L[1] {
		t.Errorf("Expected for L's (1)-th element to be 6.0; received %v", qv3.L[1])
	}

}

/*
TestQuadraticExpr_RewriteInTermsOfIndices1
Description:
	Tests whether or not the rewrite function returns a quadratic expression in three variables when asked.
*/
func TestQuadraticExpr_RewriteInTermsOfIndices1(t *testing.T) {
	// Constants
	m := goop2.NewModel()

	v1 := m.AddVar(-10, 10, goop2.Continuous)
	v2 := m.AddVar(-10, 10, goop2.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	qv1, err := goop2.NewQuadraticExpr_qb0(Q1, []uint64{v1.ID, v2.ID})
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	qvNew, err := qv1.RewriteInTermsOfIndices([]uint64{v1.ID, v2.ID, 200})
	if err != nil {
		t.Errorf("There was an issue rewriting the quadratic expression when there should not have been: %v", err)
	}

	if len(qvNew.Q) != 3 {
		t.Errorf("There were %v rows in the new Q; expected 3", len(qvNew.Q))
	}

	for rowIndex := 0; rowIndex < len(qvNew.Q); rowIndex++ {
		if len(qvNew.Q[rowIndex]) != 3 {
			t.Errorf("There were %v columns in new Q's %v-th row; expected 3", len(qvNew.Q[rowIndex]), rowIndex)
		}
	}

	if len(qvNew.L) != 3 {
		t.Errorf("There were %v elements in the new L; expected 3", len(qvNew.L))
	}

	if qvNew.C != 0.0 {
		t.Errorf("Expected for new C to be 0; received %v", qvNew.C)
	}

}
