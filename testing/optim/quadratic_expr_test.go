package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
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

	// Create inputs for NeqQuadraticExpr
	Q_vectorized := append(Q1[0], Q1[1]...)
	Q1_as_mat := mat.NewDense(2, 2, Q_vectorized)

	var vv = optim.VarVector{}
	for _, tempId := range xIndices1 {
		vv.Elements = append(vv.Elements, optim.Variable{ID: tempId})
	}

	// Algorithm
	_, err := optim.NewQuadraticExpr_qb0(*Q1_as_mat, vv)
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

	// Create Inputs for NewQuadraticExpr_qb0
	Q2_vectorized := Q2[0]
	Q2_as_mat := mat.NewDense(1, 2, Q2_vectorized)

	var vv = optim.VarVector{}
	for _, tempId := range xIndices2 {
		vv.Elements = append(vv.Elements, optim.Variable{ID: tempId})
	}

	// Algorithm
	_, err := optim.NewQuadraticExpr_qb0(*Q2_as_mat, vv)
	if err == nil {
		t.Errorf("Expected an error, but there was none!")
	}

	if !strings.Contains(err.Error(), "The number of indices was 2 which did not match the number of rows in QIn (1)") {
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

	// Create Inputs for NewQuadraticExpr_qb0
	Q3_vectorized := append(Q3[0], Q3[1]...)
	Q3_as_mat := mat.NewDense(1, 2, Q3_vectorized)

	var vv = optim.VarVector{}
	for _, tempId := range xIndices3 {
		vv.Elements = append(vv.Elements, optim.Variable{ID: tempId})
	}

	// Algorithm
	_, err := optim.NewQuadraticExpr_qb0(*Q3_as_mat, vv)
	if err == nil {
		t.Errorf("Expected an error, but there was none!")
	}

	if !strings.Contains(err.Error(), "The number of indices was 2 which did not match the number of rows in QIn (1).") {
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	// Create inputs for NeqQuadraticExpr
	Q_vectorized := append(Q1[0], Q1[1]...)
	Q1_as_mat := mat.NewDense(2, 2, Q_vectorized)

	// Algorithm
	qv1, err := optim.NewQuadraticExpr_qb0(*Q1_as_mat, optim.VarVector{[]optim.Variable{v1, v2}})
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q2 := [][]float64{
		[]float64{1.0, 2.0, 3.0},
		[]float64{4.0, 5.0, 6.0},
		[]float64{7.0, 8.0, 9.0},
	}

	// Create Inputs for NewQuadraticExpr_qb0
	Q2_vectorized := append(Q2[0], Q2[1]...)
	Q2_vectorized = append(Q2_vectorized, Q2[2]...)
	Q2_as_mat := mat.NewDense(3, 3, Q2_vectorized)

	// Algorithm
	qv1, err := optim.NewQuadraticExpr_qb0(
		*Q2_as_mat,
		optim.VarVector{
			[]optim.Variable{v1, v2, v3},
		})
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q3 := [][]float64{
		[]float64{2.3},
	}
	// Create Inputs for NewQuadraticExpr_qb0
	Q3_vectorized := Q3[0]
	Q3_as_mat := mat.NewDense(1, 1, Q3_vectorized)

	// Algorithm
	qv1, err := optim.NewQuadraticExpr_qb0(
		*Q3_as_mat,
		optim.VarVector{
			[]optim.Variable{v1},
		},
	)
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1 := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	// Create inputs for NeqQuadraticExpr
	Q_vectorized := append(Q1[0], Q1[1]...)
	Q1_as_mat := mat.NewDense(2, 2, Q_vectorized)

	// Algorithm
	qv1, err := optim.NewQuadraticExpr_qb0(
		*Q1_as_mat,
		optim.VarVector{
			[]optim.Variable{v1, v2},
		},
	)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	// Number of variables for this quadratic expression should be 2
	if len(qv1.IDs()) != 2 {
		t.Errorf("Expected for 2 variables to be found in quadratic expression; function says %v variables exist.", len(qv1.IDs()))
	}

	if tempVars := qv1.IDs(); tempVars[0] != v1.ID {
		t.Errorf("Expected for first ID to be %v; received %v.", v1.ID, tempVars[0])
	}

	if tempVars := qv1.IDs(); tempVars[1] != v2.ID {
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	Q2_aoa := [][]float64{
		[]float64{1.0, 0.0},
		[]float64{0.0, 1.0},
	}
	// Convert array of arrays to matrices
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	Q2_vals := append(Q2_aoa[0], Q2_aoa[1]...)
	Q2 := *mat.NewDense(2, 2, Q2_vals)

	vv := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Algorithm
	qv1, err := optim.NewQuadraticExpr_qb0(Q1, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	qv2, err := optim.NewQuadraticExpr_qb0(Q2, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	e3, err := qv1.Plus(qv2)
	if err != nil {
		t.Errorf("Received an error when computing Plus(): %v", err)
	}

	qv3, ok := e3.(*optim.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 2 {
		t.Errorf("Expected for 2 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.Q.At(0, 0) != 2.0 {
		t.Errorf("Expected for Q's (0,0)-th element to be 2.0; received %v", qv3.Q.At(0, 0))
	}

	if qv3.Q.At(0, 1) != qv1.Q.At(0, 1) {
		t.Errorf("Expected for Q's (0,1)-th element to be %v; received %v", qv3.Q.At(0, 1), qv1.Q.At(0, 1))
	}

	if qv3.Q.At(1, 0) != qv1.Q.At(1, 0) {
		t.Errorf("Expected for Q's (1,0)-th element to be %v; received %v", qv3.Q.At(1, 0), qv1.Q.At(1, 0))
	}

	if qv3.Q.At(1, 1) != 5.0 {
		t.Errorf("Expected for Q's (1,1)-th element to be 5.0; received %v", qv3.Q.At(1, 1))
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
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}
	// Converting Q to matrices
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	vv := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Algorithm / Tests
	qv1, err := optim.NewQuadraticExpr_qb0(Q1, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	L2_a := []float64{5.0, 6.0}
	L2 := *mat.NewVecDense(2, L2_a)
	le2 := optim.ScalarLinearExpr{
		L: L2,
		C: 0.1,
		X: vv,
	}

	e3, err := qv1.Plus(le2)
	if err != nil {
		t.Errorf("There was an issue adding together the two expressions: %v", err)
	}

	qv3, ok := e3.(*optim.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 2 {
		t.Errorf("Expected for 2 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.Q.At(0, 0) != 1.0 {
		t.Errorf("Expected for Q's (0,0)-th element to be 2.0; received %v", qv3.Q.At(0, 0))
	}

	if qv3.Q.At(0, 1) != qv1.Q.At(0, 1) {
		t.Errorf("Expected for Q's (0,1)-th element to be %v; received %v", qv3.Q.At(0, 1), qv1.Q.At(0, 1))
	}

	if qv3.Q.At(1, 0) != qv1.Q.At(1, 0) {
		t.Errorf("Expected for Q's (1,0)-th element to be %v; received %v", qv3.Q.At(1, 0), qv1.Q.At(1, 0))
	}

	if qv3.L.AtVec(1) != le2.L.AtVec(1) {
		t.Errorf("Expected for L's (1)-th element to be 6.0; received %v", qv3.L.AtVec(1))
	}

}

/*
TestQuadraticExpr_Plus3
Description:

	Tests whether or not the Plus() function works for two quadratic expressions containing
	slightly different variables.
*/
func TestQuadraticExpr_Plus3(t *testing.T) {
	// Constants
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0, 3.0},
		[]float64{4.0, 5.0, 6.0},
		[]float64{7.0, 8.0, 9.0},
	}

	Q2_aoa := [][]float64{
		[]float64{10.0, 11.0, 12.0},
		[]float64{13.0, 14.0, 15.0},
		[]float64{16.0, 17.0, 18.0},
	}

	// Converting Arrays of arrays to matrices
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1_vals = append(Q1_vals, Q1_aoa[2]...)
	Q1 := *mat.NewDense(3, 3, Q1_vals)

	Q2_vals := append(Q2_aoa[0], Q2_aoa[1]...)
	Q2_vals = append(Q2_vals, Q2_aoa[2]...)
	Q2 := *mat.NewDense(3, 3, Q2_vals)

	vv := optim.VarVector{
		[]optim.Variable{v1, v2, v3},
	}

	// Algorithm / Testing
	qv1, err := optim.NewQuadraticExpr_qb0(Q1, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	qv2, err := optim.NewQuadraticExpr_qb0(Q2, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	e3, err := qv1.Plus(qv2)
	if err != nil {
		t.Errorf("There was an issue adding qv1 and qv2: %v", err)
	}

	qv3, ok := e3.(*optim.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 3 {
		t.Errorf("Expected for 3 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.Q.At(0, 0) != qv1.Q.At(0, 0)+qv2.Q.At(0, 0) {
		t.Errorf("Expected for Q's (0,0)-th element to be %v; received %v", qv1.Q.At(0, 0)+qv2.Q.At(0, 0), qv3.Q.At(0, 0))
	}

	if qv3.Q.At(1, 1) != qv1.Q.At(1, 1)+qv2.Q.At(1, 1) {
		t.Errorf("Expected for Q's (1,1)-th element to be %v; received %v", qv1.Q.At(1, 1)+qv2.Q.At(1, 1), qv3.Q.At(1, 1))
	}

	if qv3.Q.At(2, 2) != qv1.Q.At(2, 2)+qv2.Q.At(2, 2) {
		t.Errorf("Expected for Q's (2,2)-th element to be %v; received %v", qv1.Q.At(2, 2)+qv2.Q.At(2, 2), qv3.Q.At(2, 2))
	}

}

/*
TestQuadraticExpr_Plus4
Description:

	Tests whether or not the Plus() function works for a quadratic expression and a linear one containing
	slightly different variables.
*/
func TestQuadraticExpr_Plus4(t *testing.T) {
	// Constants
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	L1_a := []float64{1.0, 7.0}

	C1 := 3.14

	// Preparing constants for NewQuadraticExpr
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	L1 := *mat.NewVecDense(2, L1_a)

	vv1 := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Quantities for Second Expression
	L2 := *mat.NewVecDense(2, []float64{2.0, 11.0})
	C2 := 1.25

	vv2 := optim.VarVector{
		[]optim.Variable{v2, v3},
	}

	// Algorithm
	qe1, err := optim.NewQuadraticExpr(Q1, L1, C1, vv1)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	le2 := optim.ScalarLinearExpr{
		L: L2,
		C: C2,
		X: vv2,
	}

	e3, err := qe1.Plus(le2)
	if err != nil {
		t.Errorf("There was an issue adding qe1 and le2: %v", err)
	}

	qv3, ok := e3.(*optim.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 3 {
		t.Errorf("Expected for 3 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.L.AtVec(0) != qe1.L.AtVec(0) {
		t.Errorf("Expected for L's 0-th element to be 1.0; received %v", qv3.L.AtVec(0))
	}

	if qv3.L.AtVec(1) != qe1.L.AtVec(1)+le2.L.AtVec(0) {
		t.Errorf("Expected for L's 1-th element to be 5.0; received %v", qv3.L.AtVec(1))
	}

	if qv3.L.AtVec(2) != le2.L.AtVec(1) {
		t.Errorf("Expected for L's 2-th element to be 11.0; received %v", qv3.L.AtVec(2))
	}

	if qv3.C != qe1.C+le2.C {
		t.Errorf("Expected for constant of final quadratic expression to be %v; received %v", qe1.C+le2.C, qv3.C)
	}

}

/*
TestQuadraticExpr_Plus5
Description:

	Tests whether or not the Plus() function works for a quadratic expression and a constant one.
*/
func TestQuadraticExpr_Plus5(t *testing.T) {
	// Constants
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	L1_a := []float64{1.0, 7.0}

	C1 := 3.14

	// Preparing constants for NewQuadraticExpr
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	L1 := *mat.NewVecDense(2, L1_a)

	vv1 := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Quantities for Second Expression
	K2 := optim.K(1.25)

	// Algorithm
	qe1, err := optim.NewQuadraticExpr(Q1, L1, C1, vv1)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	e3, err := qe1.Plus(K2)
	if err != nil {
		t.Errorf("There was an issue adding qe1 and le2: %v", err)
	}

	qv3, ok := e3.(*optim.QuadraticExpr)
	if !ok {
		t.Errorf("Unable to convert expression to Quadratic Expression.")
	}

	// Number of variables for this quadratic expression should be 2
	if qv3.NumVars() != 2 {
		t.Errorf("Expected for 3 variable to be found in quadratic expression; function says %v variables exist.", qv3.NumVars())
	}

	if qv3.L.AtVec(0) != qe1.L.AtVec(0) {
		t.Errorf("Expected for L's 0-th element to be %v; received %v", qe1.L.AtVec(0), qv3.L.AtVec(0))
	}

	if qv3.L.AtVec(1) != qe1.L.AtVec(1) {
		t.Errorf("Expected for L's 1-th element to be %v; received %v", qe1.L.AtVec(1), qv3.L.AtVec(1))
	}

	if qv3.C != qe1.C+float64(K2) {
		t.Errorf("Expected for constant of final quadratic expression to be %v; received %v", qe1.C, qv3.C)
	}

}

/*
TestQuadraticExpr_RewriteInTermsOfIndices1
Description:

	Tests whether or not the rewrite function returns a quadratic expression in three variables when asked.
*/
func TestQuadraticExpr_RewriteInTermsOfIndices1(t *testing.T) {
	// Constants
	m := optim.NewModel()

	v1 := m.AddVariableClassic(-10, 10, optim.Continuous)
	v2 := m.AddVariableClassic(-10, 10, optim.Continuous)

	Q1_aoa := [][]float64{
		[]float64{1.0, 2.0},
		[]float64{3.0, 4.0},
	}

	// Prepare variables for use with NewQuadraticExpr_qb0
	Q1_vals := append(Q1_aoa[0], Q1_aoa[1]...)
	Q1 := *mat.NewDense(2, 2, Q1_vals)

	vv := optim.VarVector{
		[]optim.Variable{v1, v2},
	}

	// Algorithm/Test
	qv1, err := optim.NewQuadraticExpr_qb0(Q1, vv)
	if err != nil {
		t.Errorf("There was an issue creating a basic quadratic expression: %v", err)
	}

	v3 := m.AddVariableClassic(-10, 10, optim.Continuous)
	vv2 := optim.VarVector{
		[]optim.Variable{v1, v2, v3},
	}

	qvNew, err := qv1.RewriteInTermsOf(vv2)
	if err != nil {
		t.Errorf("There was an issue rewriting the quadratic expression when there should not have been: %v", err)
	}

	if n_rows, _ := qvNew.Q.Dims(); n_rows != 3 {
		t.Errorf("There were %v rows in the new Q; expected 3", n_rows)
	}

	if _, n_cols := qvNew.Q.Dims(); n_cols != 3 {
		t.Errorf("There were %v columns in the new Q; expected 3", n_cols)
	}

	if qvNew.L.Len() != 3 {
		t.Errorf("There were %v elements in the new L; expected 3", qvNew.L.Len())
	}

	if qvNew.C != 0.0 {
		t.Errorf("Expected for new C to be 0; received %v", qvNew.C)
	}

}
