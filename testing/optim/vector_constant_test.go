package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"gonum.org/v1/gonum/mat"
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
TestKVector_Eq1
Description:

	This function tests that the Eq() method is properly working for KVector inputs.
*/
func TestKVector_Eq1(t *testing.T) {
	// Constants
	desLength := 10
	var vec1 = optim.KVector(optim.OnesVector(desLength))
	var vec2 = optim.KVector(optim.ZerosVector(desLength))

	// Create Constraint
	constr, err := vec1.Eq(vec2)
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
