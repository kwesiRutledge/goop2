package optim

import "gonum.org/v1/gonum/mat"

/*
vector_constant_test.go
Description:
	Creates a vector extension of the constant type K from the original goop.
*/

import (
	"fmt"
)

/*
KVector

	A type which is built on top of the KVector()
	a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
*/
type KVector struct {
	mat.VecDense // Inherit all methods from mat.VecDense
}

/*
NumVars
Description:

	This returns the number of variables in the expression. For constants, this is 0.
*/
func (c KVector) NumVars() int {
	return 0
}

/*
Vars
Description:

	This function returns a slice of the Var ids in the expression. For constants, this is always nil.
*/
func (c KVector) IDs() []uint64 {
	return nil
}

/*
Coeffs
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
func (c KVector) Coeffs() []float64 {
	return nil
}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (c KVector) Constant() mat.Vector {

	return &c.VecDense
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (c KVector) Plus(e VectorExpression) (VectorExpression, error) {
	switch e.(type) {
	case KVector:
		// Cast type
		eAsKVector, _ := e.(KVector)
		// Compute Addition
		var result KVector
		result.AddVec(&c, &eAsKVector)

		return result, nil
	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for expression %v!", e, e)
		return KVector{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This method multiplies the current expression to another and returns the resulting expression.
*/
func (c KVector) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result KVector
	result.ScaleVec(val, &c)

	return result, nil
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (c KVector) LessEq(rhsIn interface{}) (VectorConstraint, error) {

	switch rhsIn.(type) {
	case KVector:
		// Cast type
		rhsAsKVector, _ := rhsIn.(KVector)

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  c,
			RightHandSide: rhsAsKVector,
			Sense:         SenseEqual,
		}, nil

	}

	// Return an error if none of the above types matched.
	return VectorConstraint{}, fmt.Errorf("The input to Eq() (%v) has unexpected type: %T", rhsIn, rhsIn)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (c KVector) GreaterEq(rhsIn interface{}) (VectorConstraint, error) {
	switch rhsIn.(type) {
	case KVector:
		// Cast type
		rhsAsKVector, _ := rhsIn.(KVector)

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  c,
			RightHandSide: rhsAsKVector,
			Sense:         SenseEqual,
		}, nil

	}

	// Return an error if none of the above types matched.
	return VectorConstraint{}, fmt.Errorf("The input to Eq() (%v) has unexpected type: %T", rhsIn, rhsIn)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (c KVector) Eq(rhsIn interface{}) (VectorConstraint, error) {

	switch rhsIn.(type) {
	case KVector:
		// Cast type
		rhsAsKVector, _ := rhsIn.(KVector)

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  c,
			RightHandSide: rhsAsKVector,
			Sense:         SenseEqual,
		}, nil

	}

	// Return an error if none of the above types matched.
	return VectorConstraint{}, fmt.Errorf("The input to Eq() (%v) has unexpected type: %T", rhsIn, rhsIn)

}
