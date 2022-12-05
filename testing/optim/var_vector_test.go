package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

func TestVarVector_Length1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	if vv1.Length() != 2 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 2)
	}

}

/*
TestVarVector_Length2
Description:

	Tests that a larger vector variable (contains 5 elements) properly returns the right length.
*/
func TestVarVector_Length2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y, *x, *y, *x},
	}

	if vv1.Length() != 5 {
		t.Errorf("The length of vv1 was %v; expected %v", vv1.Length(), 5)
	}

}

/*
TestVarVector_At1
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
*/
func TestVarVector_At1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	extractedV := vv1.At(1)
	if extractedV != *y {
		t.Errorf("Expected for extracted variable, %v, to be the same as %v. They were different!", extractedV, y)
	}
}

/*
TestVarVector_At2
Description:

	Tests whether or not we can properly retrieve an element from a given vector.
	Makes sure that if we change the extracted vector, it does not effect the element saved in the slice.
*/
func TestVarVector_At2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	extractedV := vv1.At(1)
	extractedV.ID = 100

	if extractedV == *y {
		t.Errorf("Expected for extracted variable, %v, to be DIFFERENT from %v. They were the same!", extractedV, y)
	}
}
