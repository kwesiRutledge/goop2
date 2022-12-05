package optim_test

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
TestVectorLinearExpression_Check1
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is valid.
*/
func TestVectorLinearExpression_Check1(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	L1 := mat.NewDense(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err != nil {
		t.Errorf("The vector linear expression was supposed to be valid, but received an error: %v", err)
	}
}

/*
TestVectorLinearExpression_Check2
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is NOT valid. L is too big in rows.
*/
func TestVectorLinearExpression_Check2(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	L1 := mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimension of L (%v x %v) and C (length %v) do not match!", nL, mL, c1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #2, instead received %v", err)
	}
}

/*
TestVectorLinearExpression_Check3
Description:

	This test will evaluate whether or not the linear expression that has been given is valid.
	In this case, the VectorLinearExpression is NOT valid. L is too big in columns.
*/
func TestVectorLinearExpression_Check3(t *testing.T) {
	m := optim.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// Create Vector Variables
	vv1 := optim.VarVector{
		Elements: []optim.Var{*x, *y},
	}

	L1 := mat.NewDense(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	c1 := mat.NewVecDense(2, []float64{5.0, 6.0})

	// Use these to create expression.
	ve1 := optim.VectorLinearExpr{
		vv1, L1, c1,
	}

	// ve1 should pass all checks.
	err := ve1.Check()
	if err == nil {
		t.Errorf("The vector linear expression was supposed to be invalid, but received no errors!")
	}

	nL, mL := L1.Dims()
	if err.Error() != fmt.Sprintf("Dimensions of L (%v x %v) and x (length %v) do not match appropriately.", nL, mL, vv1.Len()) {
		t.Errorf("The vector linear expression was supposed to have dimension error #1, instead received %v", err)
	}
}
