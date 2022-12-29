package optim_test

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
vector_constant_test.go
Description:
	Tests the new type KVector which represents a constant vector.
*/

/*
TestKVector_At1
Description:

	This test verifies whether or not a 1 is retrieved when we create a KVector
	using OnesVector().
*/
func TestKVector_At1(t *testing.T) {
	// Create a KVector
	desLength := 4
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	targetIndex := 2

	if vec1.At(targetIndex) != 1.0 {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.At(targetIndex), 1.0)
	}
}

/*
TestKVector_At2
Description:

	This test verifies whether or not an arbitrary number is retrieved when we create a KVector
	using NewVecDense().
*/
func TestKVector_At2(t *testing.T) {
	// Create a KVector
	vec1Elts := []float64{1.0, 3.0, 5.0, 7.0, 9.0}
	var vec1 = optim.KVector(*mat.NewVecDense(5, vec1Elts))
	targetIndex := 3

	if vec1.At(targetIndex) != vec1Elts[targetIndex] {
		t.Errorf("vec1[%v] = %v; expected %v.", targetIndex, vec1.At(targetIndex), vec1Elts[targetIndex])
	}
}

/*
TestKVector_Len1
Description:

	This function tests that the Len() method works.
	(Should be inherited from the base type mat.DenseVec)
*/
func TestKVector_Len1(t *testing.T) {
	// Create a KVector
	desLength := 4
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVector_Len2
Description:

	This function tests that the Len() method is properly inherited by KVector.
*/
func TestKVector_Len2(t *testing.T) {
	// Create a KVector
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}

/*
TestKVector_Comparison1
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
*/
func TestKVector_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength))

	// Create Constraint
	constr, err := vec1.Comparison(vec2, optim.SenseEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison2
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Uses SenseLessThanEqual.
	Comparison of:
	- KVector
	- VarVector
*/
func TestKVector_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVarVector(desLength)

	// Create Constraint
	constr, err := vec1.Comparison(vec2, optim.SenseLessThanEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison3
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVector
	- VectorLinearExpression
*/
func TestKVector_Comparison3(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVarVector(desLength)

	L1 := optim.Identity(desLength)
	c1 := optim.OnesVector(desLength)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, &c1,
	}

	// Create Constraint
	constr, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}

	if constr.LeftHandSide.Len() != vec1.Len() {
		t.Errorf(
			"Expected left hand side (length %v) to have same length as vec1 (length %v).",
			constr.LeftHandSide.Len(),
			vec1.Len(),
		)
	}
}

/*
TestKVector_Comparison4
Description:

	This function tests that the Comparison() method is properly working for KVector inputs.
	Input is bad (dimension of linear vector expression is different from constant vector) and error should be thrown.
	Uses SenseGreaterThanEqual.
	Comparison of:
	- KVector
	- VectorLinearExpression
*/
func TestKVector_Comparison4(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = m.AddVarVector(desLength - 1)

	L1 := optim.Identity(desLength - 1)
	c1 := optim.OnesVector(desLength - 1)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, &c1,
	}

	// Create Constraint
	_, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if strings.Contains(err.Error(), fmt.Sprintf("The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!", vec1.Len(), ve1.Len())) {
		t.Errorf("There was an issue creating equality constraint between vec1 and vec2: %v", err)
	}
}
