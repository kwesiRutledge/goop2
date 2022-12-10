package optim

/*
vector_linear_expression.go
Description:

*/

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// VectorLinearExpr represents a linear general expression of the form
//
//	L' * x + C
//
// where L is an n x m matrix of coefficients that matches the dimension of x, the vector of variables
// and C is a constant vector
type VectorLinearExpr struct {
	X VarVector
	L mat.Matrix // Matrix of coefficients. Should match the dimensions of XIndices
	C mat.Vector
}

/*
Check
Description:

	Checks to see if the VectorLinearExpression is well-defined.
*/
func (v VectorLinearExpr) Check() error {
	// Extract the dimension of the vector x
	m := v.X.Length()
	nL, mL := v.L.Dims()
	nC := v.C.Len()

	// Compare the length of vector x with the appropriate dimension of L
	if m != mL {
		return fmt.Errorf("Dimensions of L (%v x %v) and x (length %v) do not match appropriately.", nL, mL, m)
	}

	// Compare the size of the matrix L with the vector C that it will be compared to.
	if nC != nL {
		return fmt.Errorf("Dimension of L (%v x %v) and C (length %v) do not match!", nL, mL, nC)
	}

	// If all other checks passed, then the VectorLinearExpression seems valid.
	return nil
}

/*
VariableIDs
Description:

	Returns the goop2 ID of each variable in the current vector linear expression.
*/
func (v VectorLinearExpr) VariableIDs() []uint64 {
	return v.X.IDs()
}

/*
Coeffs
Description:

	Returns a list of coefficients that correspond to each one of the elements of the coefficient matrix.
*/
func (v VectorLinearExpr) Coeffs() []float64 {

	var coeffsOut []float64
	nRows, nCols := v.L.Dims()
	for rowIndex := 0; rowIndex < nRows; rowIndex++ {
		for colIndex := 0; colIndex < nCols; colIndex++ {
			coeffsOut = append(coeffsOut, v.L.At(rowIndex, colIndex))
		}
	}

	return coeffsOut
}
