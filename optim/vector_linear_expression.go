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
func (vle VectorLinearExpr) Check() error {
	// Extract the dimension of the vector x
	m := vle.X.Length()
	nL, mL := vle.L.Dims()
	nC := vle.C.Len()

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
IDs
Description:

	Returns the goop2 ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpr) IDs() []uint64 {
	return vle.X.IDs()
}

/*
NumVars
Description:

	Returns the goop2 ID of each variable in the current vector linear expression.
*/
func (vle VectorLinearExpr) NumVars() int {
	return len(vle.IDs())
}

/*
LinearCoeff
Description:

	Returns the matrix which is applied as a coefficient to the vector X in our expression.
*/
func (vle VectorLinearExpr) LinearCoeff() mat.Matrix {

	return vle.L
}

/*
Constant
Description:

	Returns the vector which is given as an offset vector in the linear expression represented by v
	(the c in the above expression).
*/
func (vle VectorLinearExpr) Constant() mat.Vector {

	return vle.C
}

/*
GreaterEq
Description:

	Creates a VectorConstraint that declares vle is greater than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpr) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	// Constant

	// Algorithm
	switch rhs.(type) {
	case KVector:
		return VectorConstraint{}, fmt.Errorf("Unimplemented.")
	}

	return VectorConstraint{}, fmt.Errorf("This place should never be reached!")
}

/*
LessEq
Description:

	Creates a VectorConstraint that declares vle is less than or equal to the value to the right hand side rhs.
*/
func (vle VectorLinearExpr) LessEq(rhs interface{}) (VectorConstraint, error) {
	// Constant

	// Algorithm
	switch rhs.(type) {
	case KVector:
		return VectorConstraint{}, fmt.Errorf("Unimplemented.")
	}

	return VectorConstraint{}, fmt.Errorf("This place should never be reached!")
}

/*
Mult
Description:

	Returns an expression which scales every dimension of the vector linear expression by the input.
*/
func (vle VectorLinearExpr) Mult(c float64) (VectorExpression, error) {
	return vle, fmt.Errorf("The multiplication method has not yet been implemented!")
}

/*
Plus
Description:

	Returns an expression which adds the expression e to the vector linear expression at hand.
*/
func (vle VectorLinearExpr) Plus(e VectorExpression) (VectorExpression, error) {
	return vle, fmt.Errorf("The addition method has not yet been implemented!")
}

/*
LessEq
Description:

	Returns a constraint between the current vector linear expression and the input given
	as the right hand side.
*/
//func (v VectorLinearExpr) LessEq(rhsIn interface{}) (VectorConstraint, error) {
//	// Output depends on the input type
//	switch rhsIn.(type) {
//	case K:
//		// Constant on right hand side.
//		rhsK, _ := rhsIn.(K)
//
//		lhsDim, _ := v.L.Dims()
//
//		onesVec := OnesVector(lhsDim)
//		var rhs KVector
//		rhs.ScaleVec(rhsK.float64, onesVec)
//
//		// Create new VectorExpression
//		return VectorConstraint{
//			LeftHandSide:  v,
//			RightHandSide: rhs,
//			Sense:         SenseLessThanEqual,
//		}, nil
//	}
//
//	return nil, fmt.Errorf("Unexpected type of right hand side %v: %T", rhsIn, rhsIn)
//}

/*
Eq
Description:

	Creates a constraint between the current vector linear expression v and the
	rhs given by rhs.
*/
func (vle VectorLinearExpr) Eq(rhs interface{}) (VectorConstraint, error) {
	// Constants

	// Check Input
	err := vle.Check()
	if err != nil {
		return VectorConstraint{}, fmt.Errorf(
			"There was an issue in the provided vector linear expression %v: %v",
			vle, err,
		)
	}

	// Algorithm
	switch rhs.(type) {
	case KVector:
		rhsAsKVector, _ := rhs.(KVector)
		// Check length of input and output.
		if rhsAsKVector.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsAsKVector.Len(),
				)
		}
		return VectorConstraint{vle, rhsAsKVector, SenseEqual}, nil
	case mat.VecDense:
		rhsAsVecDense, _ := rhs.(mat.VecDense)
		return vle.Eq(KVector(rhsAsVecDense))
	case VectorLinearExpr:
		rhsAsVLE, _ := rhs.(VectorLinearExpr)
		// Check length of input and output.
		if rhsAsVLE.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsAsVLE.Len(),
				)
		}
		return VectorConstraint{vle, rhsAsVLE, SenseEqual}, nil
	case VarVector:
		rhsAsVV, _ := rhs.(VarVector)
		// Check length of input and output.
		if rhsAsVV.Len() != vle.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two vector inputs to Eq() must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					vle.Len(),
					rhsAsVV.Len(),
				)
		}
		return VectorConstraint{vle, rhsAsVV, SenseEqual}, nil

	default:
		return VectorConstraint{}, fmt.Errorf("The comparison of vector linear expression %v with object of type %T is not currently supported.", vle, rhs)
	}
}

/*
Len
Description:

	The size of the constraint.
*/
func (vle VectorLinearExpr) Len() int {
	// Constants

	// Algorithm
	return vle.C.Len()
}
