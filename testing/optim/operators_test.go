package optim

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

/*
TestOperators_LessEq1
Description:

	Tests whether or not a good LessEq comparison successfully is built and contains the right variables.
*/
func TestOperators_LessEq1(t *testing.T) {
	// Constants
	desLength := 5

	m := optim.NewModel()
	var vec1 = m.AddVariableVector(desLength)
	var vec2 = optim.OnesVector(desLength)

	// Algorithm
	constr, err := optim.LessEq(vec1, vec2)
	if err != nil {
		t.Errorf("There was an issue compusing the LessEq comparison: %v", err)
	}

	vecConstr, ok := constr.(optim.VectorConstraint)
	if !ok {
		t.Errorf("expected constraint to be a Vector constraint, but it was really of type %T.", constr)
	}

	lhs := vecConstr.LeftHandSide
	lhsAsVarVector, ok := lhs.(optim.VarVector)
	if !ok {
		t.Errorf("The left hand side was expected to be a VarVector, but instead it was %T.", lhs)
	}

	for varIndex := 0; varIndex < vec1.Len(); varIndex++ {
		vec1_i := vec1.AtVec(varIndex)
		lhs_i := lhsAsVarVector.AtVec(varIndex)
		if vec1_i.ID != lhs_i.ID {
			t.Errorf(
				"vec1's %v-th element (%v) is different from left hand side's %v-th element (%v).",
				varIndex,
				vec1_i,
				varIndex,
				lhs_i,
			)
		}

	}

}
