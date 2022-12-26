package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
var_vector.go
Description:
	The VarVector type will represent a
*/

/*
VarVector
Description:

	Represnts a variable in a optimization problem. The variable is
*/
type VarVector struct {
	Elements []Var
}

// =========
// Functions
// =========

/*
Length
Description:

	Returns the length of the vector of optimization variables.
*/
func (vv VarVector) Length() int {
	return len(vv.Elements)
}

/*
Len
Description:

	This function is created to mirror the GoNum Vector API. Does the same thing as Length.
*/
func (vv VarVector) Len() int {
	return vv.Length()
}

/*
At
Description:

	Mirrors the gonum api for vectors. This extracts the element of the variable vector at the index x.
*/
func (vv VarVector) At(x int) Var {
	return vv.Elements[x]
}

/*
IDs
Description:

	Returns the unique indices
*/
func (vv VarVector) IDs() []uint64 {
	// Algorithm
	var IDSlice []uint64

	for _, elt := range vv.Elements {
		IDSlice = append(IDSlice, elt.ID)
	}

	return Unique(IDSlice)

}

/*
NumVars
Description:

	The number of unique variables inside the variable vector.
*/
func (vv VarVector) NumVars() int {
	return len(vv.IDs())
}

/*
Constant
Description:

	Returns an all zeros vector as output from the method.
*/
func (vv VarVector) Constant() mat.Vector {
	zerosOut := ZerosVector(vv.Len())
	return &zerosOut
}

/*
LinearCoeff
Description:

	Returns the matrix which is multiplied by Variables to get the current "expression".
	For a single vector, this is an identity matrix.
*/
func (vv VarVector) LinearCoeff() mat.Matrix {
	return Identity(vv.Len())
}

/*
Plus
Description:

	This member function computes the addition of the receiver vector var with the
	incoming vector expression ve.
*/
func (vv VarVector) Plus(ve VectorExpression) (VectorExpression, error) {
	return vv, fmt.Errorf("The Plus() method for VarVector is not implemented yet!")
}

/*
Mult
Description:

	This member function computest the multiplication of the receiver vector var with some
	incoming vector expression (may result in quadratic?).
*/
func (vv VarVector) Mult(c float64) (VectorExpression, error) {
	return vv, fmt.Errorf("The Mult() method for VarVector is not implemented yet!")
}

/*
LessEq
Description:

	This method creates a less than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) LessEq(rhs interface{}) (VectorConstraint, error) {
	return VectorConstraint{}, fmt.Errorf("The LessEq() method for VarVector is not implemented yet!")
}

/*
GreaterEq
Description:

	This method creates a greater than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) GreaterEq(rhs interface{}) (VectorConstraint, error) {
	return VectorConstraint{}, fmt.Errorf("The GreaterEq() method for VarVector is not implemented yet!")
}

/*
Eq
Description:

	This method creates an equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VarVector) Eq(rhs interface{}) (VectorConstraint, error) {
	// Constants

	// Algorithm
	return vv.Comparison(rhs, SenseEqual)

}

/*
Comparison
Description:

	This method creates a constraint of type sense between
	the receiver (as left hand side) and rhs (as right hand side) if both are valid.
*/
func (vv VarVector) Comparison(rhs interface{}, sense ConstrSense) (VectorConstraint, error) {
	// Constants

	// Algorithm
	switch rhs.(type) {
	case KVector:
		// Cast type
		rhsAsKVector, _ := rhs.(KVector)

		// Check length of input and output.
		if vv.Len() != rhsAsKVector.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vv.Len(),
					rhsAsKVector.Len(),
				)
		}
		return VectorConstraint{vv, rhsAsKVector, sense}, nil
	case mat.VecDense:
		// Cast Type
		rhsAsVecDense, _ := rhs.(mat.VecDense)
		rhsAsKVector := KVector(rhsAsVecDense)

		return vv.Comparison(rhsAsKVector, sense)

	case VarVector:
		// Cast Type
		rhsAsVV, _ := rhs.(VarVector)

		// Check length of input and output.
		if vv.Len() != rhsAsVV.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The two inputs to comparison '%v' must have the same dimension, but #1 has dimension %v and #2 has dimension %v!",
					sense,
					vv.Len(),
					rhsAsVV.Len(),
				)
		}
		// Do Computation
		return VectorConstraint{vv, rhsAsVV, sense}, nil

	default:
		return VectorConstraint{}, fmt.Errorf("The Eq() method for VarVector is not implemented yet for type %T!", rhs)
	}
}
