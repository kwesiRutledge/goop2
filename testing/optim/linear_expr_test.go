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
	expr := optim.Sum(expr1, expr2, optim.K(constant))

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
