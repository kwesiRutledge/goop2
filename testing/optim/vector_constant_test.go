package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

/*
vector_constant_test.go
Description:
	Tests the new type KVector which represents a constant vector.
*/

/*
TestKVector_Len1
Description:

	This function tests that the Len() method is properly inherited by KVector.
	(Should be inherited from the base type mat.DenseVec)
*/
func TestKVector_Len1(t *testing.T) {
	// Create a KVector
	desLength := 4
	var vec1 = optim.KVector{optim.OnesVector(desLength)}

	if vec1.Len() != desLength {
		t.Errorf("The length of vec1 should be %v, but instead it is %v.", desLength, vec1.Len())
	}
}
