package testing

import (
	"testing"

	"github.com/kwesiRutledge/goop2"
)

/*
quadratic_expr_test.go
Description:
	Tests some of the basic functions of the quadraticExpr class.
*/

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

	qv1 := goop2.NewQuadraticExpr(1.5, v1.ID, v2.ID)

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

	qv1 := goop2.NewQuadraticExpr(1.5, v1.ID, v2.ID)
	qv1.Coefficients = append(qv1.Coefficients, 3)
	qv1.VariablePairs = append(qv1.VariablePairs, [2]uint64{v1.ID, v3.ID})

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

	qv1 := goop2.NewQuadraticExpr(1.5, v1.ID, v1.ID)

	// Number of variables for this quadratic expression should be 2
	if qv1.NumVars() != 1 {
		t.Errorf("Expected for 1 variable to be found in quadratic expression; function says %v variables exist.", qv1.NumVars())
	}
}
