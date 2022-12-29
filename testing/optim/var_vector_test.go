package optim_test

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"strings"
	"testing"
)

func TestVarVector_Length1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y},
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
		Elements: []optim.Var{x, y, x, y, x},
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
		Elements: []optim.Var{x, y},
	}

	extractedV := vv1.At(1)
	if extractedV != y {
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
		Elements: []optim.Var{x, y},
	}

	extractedV := vv1.At(1)
	extractedV.ID = 100

	if extractedV == y {
		t.Errorf("Expected for extracted variable, %v, to be DIFFERENT from %v. They were the same!", extractedV, y)
	}
}

/*
TestVarVector_VariableIDs1
Description:

	This test will check to see if 2 unique ids in a VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVector_VariableIDs1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y},
	}

	extractedIDs := vv1.IDs()
	// Check to see that x and y have ids in extractedIDs
	if foundIndex, _ := optim.FindInSlice(x.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}

	if foundIndex, _ := optim.FindInSlice(y.ID, extractedIDs); foundIndex == -1 {
		t.Errorf("x was not found in vv1, but it should have been!")
	}
}

/*
TestVarVector_VariableIDs2
Description:

	This test will check to see if a single unique id in a large VariableVector object will be returned correctly when
	the VariableIDs method is called.
*/
func TestVarVector_VariableIDs2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, x, x},
	}

	extractedIDs := vv1.IDs()
	// Check to see that only x has ids in extractedIDs
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
TestVarVector_Constant1
Description:

	This test verifies that the constant method returns an all zero vector for any varvector object.
*/
func TestVarVector_Constant1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y},
	}

	extractedConstant := vv1.Constant()

	// Check to see that all elts in the vector are zero.
	for eltIndex := 0; eltIndex < vv1.Len(); eltIndex++ {
		constElt := extractedConstant.AtVec(eltIndex)
		if constElt != 0.0 {
			t.Errorf("Constant vector at index %v is %v; not 0.", eltIndex, constElt)
		}
	}
}

/*
TestVarVector_Constant2
Description:

	This test verifies that the constant method returns an all zero vector for any varvector object.
	This one will be extremely long.
*/
func TestVarVector_Constant2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	extractedConstant := vv1.Constant()

	// Check to see that all elts in the vector are zero.
	for eltIndex := 0; eltIndex < vv1.Len(); eltIndex++ {
		constElt := extractedConstant.AtVec(eltIndex)
		if constElt != 0.0 {
			t.Errorf("Constant vector at index %v is %v; not 0.", eltIndex, constElt)
		}
	}
}

/*

 */

/*
TestVarVector_Eq1
Description:

	This test verifies that the Eq method works between a varvector and another object.
*/
func TestVarVector_Eq1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	zerosAsVecDense := optim.ZerosVector(vv1.Len())
	zerosAsKVector := optim.KVector(zerosAsVecDense)

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(zerosAsKVector)
	if err != nil {
		t.Errorf("There was an issue creating an equality constraint between vv1 and the zero vector.")
	}
}

/*
TestVarVector_Eq2
Description:

	This test verifies that the Eq method works between a varvector and another object.
	Comparison should be between var vector and an unsupported type.
*/
func TestVarVector_Eq2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	badRHS := false

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(badRHS)
	expectedError := fmt.Sprintf("The Eq() method for VarVector is not implemented yet for type %T!", badRHS)
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error \"%v\"; received \"%v\"", expectedError, err)
	}
}

/*
TestVarVector_Eq2
Description:

	This test verifies that the Eq method works between a varvector and another var vector.
*/
func TestVarVector_Eq3(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variable
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y},
	}

	vv2 := optim.VarVector{
		Elements: []optim.Var{y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x, y, x},
	}

	// Verify that constraint can be created with no issues.
	_, err := vv1.Eq(vv2)
	if err != nil {
		t.Errorf("There was an error creating equality constraint between the two varvectors: %v", err)
	}
}

/*
TestVarVector_Comparison1
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
*/
func TestVarVector_Comparison1(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	var vec1 = m.AddVarVector(desLength)
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

/*
TestVarVector_Comparison2
Description:

	Tests how well the comparison function works with a VectorLinearExpression comparison.
	Valid comparison of
*/
func TestVarVector_Comparison2(t *testing.T) {
	// Constants
	desLength := 10
	m := optim.NewModel()
	var vec1 = m.AddVarVector(desLength)
	var vec2 = m.AddVarVector(desLength)

	L1 := optim.Identity(desLength)
	c1 := optim.OnesVector(desLength)

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vec2, L1, &c1,
	}

	// Create Constraint
	_, err := vec1.Comparison(ve1, optim.SenseGreaterThanEqual)
	if err != nil {
		t.Errorf("There was an error computing a comparison for operator >=: %v", err)
	}
}
