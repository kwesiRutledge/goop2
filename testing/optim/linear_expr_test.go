package optim_test

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

func TestLinearExprCoeffsAndConstant(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// 2 * x + 4 * y - 5
	coeffs := []float64{2, 4}
	constant := -5.0
	expr1, err := x.Mult(coeffs[0])
	if err != nil {
		t.Errorf("There was an error computing the first multiplication: %v", err)
	}
	expr2, err := y.Mult(coeffs[1])
	if err != nil {
		t.Errorf("There was an error computing the second multiplication: %v", err)
	}

	expr, err := optim.Sum(expr1, expr2, optim.K(constant))
	if err != nil {
		t.Errorf("There was an issue computing the Sum of the expressions: %v", err)
	}

	for i, coeff := range expr.Coeffs() {
		if coeffs[i] != coeff {
			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
		}
	}

	if expr.Constant() != constant {
		t.Errorf("Constant mismatch: %v != %v", expr.Constant(), constant)
	}
}

//func TestLinearExprCoeffsAndConstant(t *testing.T) {
//	m := optim.NewModel()
//	x := m.AddBinaryVar()
//	y := m.AddBinaryVar()
//
//	// 2 * x + 4 * y - 5
//	coeffs := []float64{2, 4}
//	constant := -5.0
//	expr := optim.Sum(x.Mult(coeffs[0]), y.Mult(coeffs[1]), optim.K(constant))
//
//	for i, coeff := range expr.Coeffs() {
//		if coeffs[i] != coeff {
//			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
//		}
//	}
//
//	if expr.Constant() != constant {
//		t.Errorf("Constant mismatch: %v != %v", expr.Constant(), constant)
//	}
//}

/*
TestScalarLinearExpr_Plus1
Description:

	This function should test the Plus method of ScalarLinearExpr for a very nice case.
	Two SLEs with the SAME varvector and simple constants.
*/
func TestScalarLinearExpr_Plus1(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	L2 := optim.OnesVector(2)
	L2.ScaleVec(3.0, &L2)
	c2 := 5.0

	m := optim.NewModel()
	vv1 := m.AddVarVector(2)

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv1,
	}

	sle2 := optim.ScalarLinearExpr{
		L: L2, C: c2, X: vv1,
	}

	// Algorithm
	sle3, err := sle1.Plus(&sle2)
	if err != nil {
		t.Errorf("There was an issue computing the product of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(*optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	for dimIndex := 0; dimIndex < 2; dimIndex++ {
		if sle3AsSLE.L.AtVec(dimIndex) != sle1.L.AtVec(dimIndex)+sle2.L.AtVec(dimIndex) {
			t.Errorf(
				"Sum failed for L at index %v; %v != %v + %v",
				dimIndex,
				sle3AsSLE.L.AtVec(dimIndex),
				sle1.L.AtVec(dimIndex),
				sle2.L.AtVec(dimIndex),
			)
		}
	}
}

/*
TestScalarLinearExpr_Plus2
Description:

	This function should test the Plus method of ScalarLinearExpr for a very nice case.
	Two SLEs with very similar varvector objects simple constants.
*/
func TestScalarLinearExpr_Plus2(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	L2 := optim.OnesVector(2)
	L2.ScaleVec(3.0, &L2)
	c2 := 5.0

	m := optim.NewModel()
	vv1 := m.AddVarVector(3)

	vv2 := optim.VarVector{
		vv1.Elements[:2],
	}
	vv3 := optim.VarVector{
		vv1.Elements[1:],
	}

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv2,
	}

	sle2 := optim.ScalarLinearExpr{
		L: L2, C: c2, X: vv3,
	}

	// Algorithm
	sle3, err := sle1.Plus(&sle2)
	if err != nil {
		t.Errorf("There was an issue computing the product of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(*optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	// Check that dimension of new expression has three d X
	if (sle3AsSLE.X.Len() != 3) || (sle1.X.Len() != 2) || (sle2.X.Len() != 2) {
		t.Errorf("The ScalarLinearExpression created by this sum should have dimension three even though the original two had dimension 2.")
	}

	for XIndex, elt := range sle3AsSLE.X.Elements {
		switch elt.ID {
		case vv2.Elements[0].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L1.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L1.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		case vv3.Elements[0].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L1.AtVec(0)+L2.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L1.AtVec(0)+L2.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		case vv3.Elements[1].ID:
			if sle3AsSLE.L.AtVec(XIndex) != L2.AtVec(0) {
				t.Errorf(
					"The variable with ID %v is expected to have coefficient %v; received %v",
					elt.ID,
					L2.AtVec(0),
					sle3AsSLE.L.AtVec(XIndex),
				)
			}
		default:
			t.Errorf("Unexpected ID received! %v", elt.ID)
		}

	}
}

/*
TestScalarLinearExpr_Plus3
Description:

	This function should test the Plus method of ScalarLinearExpr for the case of (SLE + K).
*/
func TestScalarLinearExpr_Plus3(t *testing.T) {
	// Constants
	L1 := optim.OnesVector(2)
	c1 := 2.0

	K1 := optim.K(5)

	m := optim.NewModel()
	vv1 := m.AddVarVector(2)

	// Create sle's
	sle1 := optim.ScalarLinearExpr{
		L: L1, C: c1, X: vv1,
	}

	// Algorithm
	sle3, err := sle1.Plus(K1)
	if err != nil {
		t.Errorf("There was an issue computing the product of sle1 and sle2: %v", err)
	}

	sle3AsSLE, ok1 := sle3.(*optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("Expected the addition of ScalarLinearExpr with another ScalarLinearExpr to create another ScalarLinearExpr. Received %T.", sle3)
	}

	if sle3AsSLE.C != sle1.C+float64(K1) {
		t.Errorf(
			"Expected for the new SLE's constant to be equal to the sum of Kq and c1. %v != %v + %v",
			sle3AsSLE.C,
			sle1.C,
			K1,
		)
	}

	for LIndex := 0; LIndex < sle3AsSLE.L.Len(); LIndex++ {
		if sle3AsSLE.L.AtVec(LIndex) != sle1.L.AtVec(LIndex) {
			t.Errorf(
				"The linear vector multiplying X was expected to be the same for sle1 and sle3, but sle3[%v] = %v != %v = sle1[%v]",
				LIndex,
				sle3AsSLE.L.AtVec(LIndex),
				LIndex,
				sle1.L.AtVec(LIndex),
			)
		}
	}
}
