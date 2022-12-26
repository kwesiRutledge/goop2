package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

/*
constraint_test.go
Description:
	Tests for all functions and objects defined in the constraint.go file.
*/

/*
TestConstraint_IsConstraint1
Description:

	This test verifies if a scalar constraint is properly detected by IsConstraint.
*/
func TestConstraint_IsConstraint1(t *testing.T) {
	// Constants
	m := optim.NewModel()

	// Create a scalar constraint.

	lhs0 := optim.One
	x := m.AddBinaryVar()

	scalarConstr0 := optim.Eq(lhs0, x)

	if !optim.IsConstraint(scalarConstr0) {
		t.Errorf("The scalar constraint is not implementing a Constraint() interface!")
	}
}

/*
TestConstraint_IsConstraint2
Description:

	This test verifies if a vector constraint is properly detected by IsConstraint.
*/
func TestConstraint_IsConstraint2(t *testing.T) {
	// Constants
	m := optim.NewModel()

	// Create a scalar constraint.

	lhs0 := optim.OnesVector(4)
	x := m.AddVar(0, 3.0, optim.Continuous)
	vv1 := optim.VarVector{
		Elements: []optim.Var{x, x, x, x},
	}

	scalarConstr0 := optim.Eq(lhs0, vv1)

	if !optim.IsConstraint(scalarConstr0) {
		t.Errorf("The scalar constraint is not implementing a Constraint() interface!")
	}
}
