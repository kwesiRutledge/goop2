package testing

import (
	"github.com/kwesiRutledge/goop2/optim"
	"testing"
)

func TestDot(t *testing.T) {
	N := 10
	m := optim.NewModel()
	xs := m.AddBinaryVarVector(N)
	coeffs := make([]float64, N)

	for i := 0; i < N; i++ {
		coeffs[i] = float64(i + 1)
	}

	expr := optim.Dot(xs, coeffs)

	for i, coeff := range expr.Coeffs() {
		if coeffs[i] != coeff {
			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
		}
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}

func TestDotPanic(t *testing.T) {
	N := 10
	m := optim.NewModel()
	xs := m.AddBinaryVarVector(N)
	coeffs := make([]float64, N-1)

	for i := 0; i < N-1; i++ {
		coeffs[i] = float64(i + 1)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Coeff size mismatch: Code did not panic")
		}
	}()

	optim.Dot(xs, coeffs)
}

func TestSumVars(t *testing.T) {
	numVars := 3
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()
	expr := optim.SumVars(x, y, z)

	for _, coeff := range expr.Coeffs() {
		if coeff != 1 {
			t.Errorf("Coeff mismatch: %v != 1", coeff)
		}
	}

	if expr.NumVars() != numVars {
		t.Errorf("NumVars mismatch: %v != %v", expr.NumVars(), numVars)
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}

/*
TestUtil_Identity1
Description:

	Create identity matrix of dimension 1 (scalar?).
*/
func TestUtil_Identity1(t *testing.T) {
	// Constants
	n := 1
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}

/*
TestUtil_Identity2
Description:

	Create identity matrix of dimension 10.
*/
func TestUtil_Identity2(t *testing.T) {
	// Constants
	n := 10
	// Algorithm
	identMat1 := optim.Identity(n)

	nX, nY := identMat1.Dims()
	if (nX != n) || (nY != n) {
		t.Errorf("The identity matrix created has dimension %v x %v; Expected %v x %v.",
			nX, nY,
			n, n,
		)
	}

}
