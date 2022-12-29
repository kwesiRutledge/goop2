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
type KVector mat.VecDense // Inherit all methods from mat.VecDense

/*
Len

	Computes the length of the KVector given.
*/
func (kv KVector) Len() int {
	kvAsVector := mat.VecDense(kv)
	return kvAsVector.Len()
}

/*
At
Description:

	This function returns the value at the k index.
*/
func (kv KVector) At(i int) float64 {
	return kv.AtVec(i)
}

/*
AtVec
Description:

	This function returns the value at the k index.
*/
func (kv KVector) AtVec(i int) float64 {
	kvAsVector := mat.VecDense(kv)
	return kvAsVector.AtVec(i)
}

/*
NumVars
Description:

	This returns the number of variables in the expression. For constants, this is 0.
*/
func (kv KVector) NumVars() int {
	return 0
}

/*
Vars
Description:

	This function returns a slice of the Var ids in the expression. For constants, this is always nil.
*/
func (kv KVector) IDs() []uint64 {
	return nil
}

/*
Coeffs
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
func (kv KVector) LinearCoeff() mat.Matrix {
	return Identity(kv.Len())
}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (kv KVector) Constant() mat.Vector {
	kvAsVector := mat.VecDense(kv)

	return &kvAsVector
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (kv KVector) Plus(e VectorExpression) (VectorExpression, error) {
	switch e.(type) {
	case KVector:
		// Cast type
		eAsKVector, _ := e.(KVector)

		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		eAsVec := mat.VecDense(eAsKVector)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVector(result), nil
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
func (kv KVector) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result mat.VecDense
	kvAsVec := mat.VecDense(kv)
	result.ScaleVec(val, &kvAsVec)

	return KVector(result), nil
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kv KVector) LessEq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kv KVector) GreaterEq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kv KVector) Eq(rhsIn interface{}) (VectorConstraint, error) {
	return kv.Comparison(rhsIn, SenseEqual)
}

func (kv KVector) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	switch rhs.(type) {
	case KVector:
		// Cast type
		rhsAsKVector, _ := rhs.(KVector)

		// Check Lengths
		if kv.Len() != rhsAsKVector.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The left hand side's dimension (%v) and the left hand side's dimension (%v) do not match!",
					kv.Len(),
					rhsAsKVector.Len(),
				)
		}

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsAsKVector,
			Sense:         sense,
		}, nil
	case VarVector:
		// Cast type
		rhsAsVV, _ := rhs.(VarVector)

		// Return constraint
		return rhsAsVV.Comparison(kv, sense)
	case VectorLinearExpr:
		// Cast Type
		rhsAsVLE, _ := rhs.(VectorLinearExpr)

		// Return constraint
		return rhsAsVLE.Comparison(kv, sense)
	default:
		// Return an error
		return VectorConstraint{}, fmt.Errorf("The input to KVector's '%v' comparison (%v) has unexpected type: %T", sense, rhs, rhs)

	}
}
