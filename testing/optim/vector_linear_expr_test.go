package optim_test

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
TestVectorLinearExpression_Check1
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is valid.
*/
func TestVectorLinearExpression_Check1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := mat.NewDense(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was supposed to be valid, but received an error: %v", err)
	}
}

/*
TestVectorLinearExpression_Check2
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is NOT valid. L is too big in rows.
*/
func TestVectorLinearExpression_Check2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimension of L (%v x %v) and C (length %v) do not match!", nL, mL, c1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #2, instead received %v", err)
	}
}

/*
TestVectorLinearExpression_Check3
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is NOT valid. L is too big in columns.
*/
func TestVectorLinearExpression_Check3(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := mat.NewDense(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimensions of L (%v x %v) and x (length %v) do not match appropriately.", nL, mL, vv1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #1, instead received %v", err)
	}
}

/*
TestVectorLinearExpression_VariableIDs1
Description:

	This test the VariableIDs() method when a variable vector with 2 unique vectors.
*/
func TestVectorLinearExpression_VariableIDs1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := mat.NewVecDense(3, []float64{5.0, 6.0, 7.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	extractedIDs := ve1.IDs()
	// Check to see that x and y have ids in extractedIDs
	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVectorLinearExpression_VariableIDs2
Description:

	This test the VariableIDs() method works for a variable vector with 1 unique vectors.
*/
func TestVectorLinearExpression_VariableIDs2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	L1 := mat.NewDense(2, 4, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	extractedIDs := ve1.IDs()
	// Check to see that x has id in extractedIDs (y should not be there)
	if len(extractedIDs) != 1 {
		t.Errorf("There is only one unique variable ID and yet %v IDs were returned by the IDs() method.", len(extractedIDs))
	}

	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex != -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVectorLinearExpression_Coeffs1
Description:

	This test the Coeffs() method which should return the matrix's elements in a prescribed order.
*/
func TestVectorLinearExpression_Coeffs1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	L1 := mat.NewDense(3, 2, LElts)
	c1 := mat.NewVecDense(3, []float64{5.0, 6.0, 7.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	m_L, n_L := L1.Dims()
	extractedCoeffMat := ve1.LinearCoeff()
	for rowIndex := 0; rowIndex < m_L; rowIndex++ {
		for colIndex := 0; colIndex < n_L; colIndex++ {
			// Compare each element of the matrix
			if L1.At(rowIndex, colIndex) != extractedCoeffMat.At(rowIndex, colIndex) {
				t.Errorf(
					"The extracted coefficient at index %v,%v (%v) is not the same as the given one (%v).",
					rowIndex, colIndex,
					extractedCoeffMat.At(rowIndex, colIndex),
					L1.At(rowIndex, colIndex),
				)
			}
		}
	}

}

/*
TestVectorLinearExpression_Coeffs2
Description:

	This test the Coeffs() method which should return the matrix's elements in a prescribed order.
*/
func TestVectorLinearExpression_Coeffs2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, x, x, x},
	}

	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	L1 := mat.NewDense(2, 4, LElts)
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	m_L, n_L := L1.Dims()
	extractedCoeffMat := ve1.LinearCoeff()
	for rowIndex := 0; rowIndex < m_L; rowIndex++ {
		for colIndex := 0; colIndex < n_L; colIndex++ {
			// Compare each element of the matrix
			if L1.At(rowIndex, colIndex) != extractedCoeffMat.At(rowIndex, colIndex) {
				t.Errorf(
					"The extracted coefficient at index %v,%v (%v) is not the same as the given one (%v).",
					rowIndex, colIndex,
					extractedCoeffMat.At(rowIndex, colIndex),
					L1.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestVectorLinearExpression_LessEq1
Description:
	This tests that the less than or equal to command works with a constant input.
*/
//func TestVectorLinearExpression_LessEq1(t *testing.T) {
//	// Constants
//	m := optim.NewModel()
//	x := m.AddBinaryVariable()
//
//	// Create Vector Variables
//	vv1 := optim.VarVector{
//		Elements: []optim.Variable{*x, *x, *x, *x},
//	}
//
//	LElts := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
//	L1 := mat.NewDense(2, 4, LElts)
//	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})
//
//	// Use these to create expression.
//	ve1 := optim.VectorLinearExpr{
//		vv1, L1, c1,
//	}
//
//	err := ve1.Check()
//	if err != nil {
//		t.Errorf("The vector linear expression was invalid! %v", err)
//	}
//
//	// Algorithm
//	constr1, err := ve1.LessEq(2.0)
//	if err != nil {
//		t.Errorf("There was an error computing the constraint ve1 <= 2.0: %v", err)
//	}
//
//	if constr1.LeftHandSide != ve1 {
//		t.Errorf("The left hand side (%v) should be the same as ve1 (%v).", constr1.LeftHandSide, ve1)
//	}
//
//}

/*
TestVectorLinearExpression_Eq1
Description:

	Tests whether or not an equality constraint between a ones vector and a standard vector variable works well.
	Eq comparison between:
	- Vector Linear Expression, and
	- mat.VecDense
*/
func TestVectorLinearExpression_Eq1(t *testing.T) {
	// Constants
	m := optim.NewModel()

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpr{
		vv1,
		optim.Identity(2),
		&c,
	}

	ones0 := optim.OnesVector(2)

	// Create Constraint
	constr, err := vle1.Eq(ones0)
	if err != nil {
		t.Errorf("There was a problem creating the vector constraint using Eq(): %v", err)
	}

	n_R := 2
	for rowIndex := 0; rowIndex < n_R; rowIndex++ {
		if constr.LeftHandSide.Constant().AtVec(rowIndex) != vle1.Constant().AtVec(rowIndex) {
			t.Errorf(
				"The constraint's left hand side has constant value %v at index %v; expected %v!",
				constr.LeftHandSide.Constant().AtVec(rowIndex),
				rowIndex,
				vle1.Constant().AtVec(rowIndex),
			)
		}
	}

}

/*
TestVectorLinearExpression_Eq2
Description:

	Tests whether or not an equality constraint between a bool and a proper vector variable leads to an error.
	Eq comparison between:
	- Vector Linear Expression, and
	- bool
*/
func TestVectorLinearExpression_Eq2(t *testing.T) {
	// Constants
	m := optim.NewModel()

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpr{
		vv1,
		optim.Identity(2),
		&c,
	}

	badRHS := false

	// Create Constraint
	_, err := vle1.Eq(badRHS)
	if !strings.Contains(err.Error(), fmt.Sprintf("vector linear expression %v with object of type %T is not currently supported.", vle1, badRHS)) {
		t.Errorf(
			"Expected an error containing \"vector linear expression %v with object of type %T is not currently supported\"; instead received %v",
			vle1,
			badRHS,
			err,
		)
	}

}

/*
TestVectorLinearExpression_Eq3
Description:

	Tests whether or not an equality constraint between a KVector and a proper vector variable leads to an error.
	Eq comparison between:
	- Vector Linear Expression, and
	- KVector
*/
func TestVectorLinearExpression_Eq3(t *testing.T) {
	// Constants
	m := optim.NewModel()

	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}
	c := optim.ZerosVector(2)
	vle1 := optim.VectorLinearExpr{
		vv1,
		optim.Identity(2),
		&c,
	}

	onesVec1 := optim.OnesVector(2)
	onesVec2 := optim.KVector(onesVec1)

	// Create Constraint
	vectorConstraint, err := vle1.Eq(onesVec2)
	if err != nil {
		t.Errorf(
			"There was an issue creating a constraint between %v and %v: %v",
			vle1,
			onesVec2,
			err,
		)
	}

	if vectorConstraint.LeftHandSide.Len() != onesVec2.Len() {
		t.Errorf("The length of lhs (%v) and rhs (%v) should be the same!", vle1.Len(), onesVec2.Len())
	}

}

/*
TestVectorLinearExpression_Eq4
Description:

	This test will evaluate how well the Eq() method for the vector of linear constraints works.
	Creates a simple two-dimensional constraint.
	Eq comparison between:
	- Vector Linear Expression, and
	- VarVector
*/
func TestVectorLinearExpression_Eq4(t *testing.T) {
	m := optim.NewModel()
	dimX := 2
	x := m.AddVariableVector(dimX)

	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		x, L1, &c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Create equality comparison.
	_, err = ve1.Eq(x)
	if err != nil {
		t.Errorf("There was an issue creating the equality constraint")
	}

}

/*
TestVectorLinearExpression_Eq5
Description:

	This test will evaluate how well the Eq() method for the vector of linear constraints works.
	Creates a simple two-dimensional constraint.
	Eq comparison between:
	- Vector Linear Expression, and
	- Vector Linear Expression
*/
func TestVectorLinearExpression_Eq5(t *testing.T) {
	m := optim.NewModel()
	dimX := 2
	x := m.AddVariableVector(dimX)

	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		x, L1, &c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// Create equality comparison.
	_, err = ve1.Eq(ve1)
	if err != nil {
		t.Errorf("There was an issue creating the equality constraint")
	}

}

/*
TestVectorLinearExpression_Len1
Description:

	This test will evaluate how well the Len() method for the vector of linear constraints works.
	A constraint between two vectors of length 2
*/
func TestVectorLinearExpression_Len1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y},
	}

	L1 := mat.NewDense(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was supposed to be valid, but received an error: %v", err)
	}

	if ve1.Len() != 2 {
		t.Errorf("Len() of vector linear expression was %v; expeted 2", ve1.Len())
	}
}

/*
TestVectorLinearExpression_Len2
Description:

	This test will evaluate how well the Len() method for the vector of linear constraints works.
	A constraint between two vectors of length 10
*/
func TestVectorLinearExpression_Len2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVariable()
	y := m.AddBinaryVariable()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Variable{x, y, x, y, x, y, x, y, x, y},
	}

	dimX := 10
	L1 := optim.Identity(dimX)
	c1 := optim.OnesVector(dimX)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, &c1,
	}

	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was invalid! %v", err)
	}

	// ve1 should pass all checks.
	if ve1.Len() != dimX {
		t.Errorf("Len() of vector linear expression was %v; expeted %v", ve1.Len(), dimX)
	}

}
