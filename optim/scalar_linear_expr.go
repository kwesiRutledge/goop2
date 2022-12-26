package optim

import "fmt"

// ScalarLinearExpr represents a linear general expression of the form
//
//	L' * x + C
//
// where L is a vector of coefficients that matches the dimension of x, the vector of variables
// variables and C is a constant
type ScalarLinearExpr struct {
	XIndices []uint64
	L        []float64 // Vector of coefficients. Should match the dimensions of XIndices
	C        float64
}

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewLinearExpr(c float64) ScalarExpression {
	return &ScalarLinearExpr{C: c}
}

// NumVars returns the number of variables in the expression
func (e *ScalarLinearExpr) NumVars() int {
	return len(e.XIndices)
}

// Vars returns a slice of the Var ids in the expression
func (e *ScalarLinearExpr) IDs() []uint64 {
	return e.XIndices
}

// Coeffs returns a slice of the coefficients in the expression
func (e *ScalarLinearExpr) Coeffs() []float64 {
	return e.L
}

// Constant returns the constant additive value in the expression
func (e *ScalarLinearExpr) Constant() float64 {
	return e.C
}

// Plus adds the current expression to another and returns the resulting
// expression
func (e *ScalarLinearExpr) Plus(other ScalarExpression) ScalarExpression {
	e.XIndices = append(e.XIndices, other.IDs()...)
	e.L = append(e.L, other.Coeffs()...)
	e.C += other.Constant()
	return e
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (e *ScalarLinearExpr) Mult(c float64) ScalarExpression {
	for i, coeff := range e.L {
		e.L[i] = coeff * c
	}
	e.C *= c

	return e
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (e *ScalarLinearExpr) LessEq(other ScalarExpression) ScalarConstraint {
	return LessEq(e, other)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (e *ScalarLinearExpr) GreaterEq(other ScalarExpression) ScalarConstraint {
	return GreaterEq(e, other)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (e *ScalarLinearExpr) Eq(other ScalarExpression) ScalarConstraint {
	return ScalarConstraint{e, other, SenseEqual}
}

/*
RewriteInTermsOfIndices
Description:

	Rewrites the current linear expression in terms of the new variables.

Usage:

	rewrittenLE, err := orignalLE.RewriteInTermsOfIndices(newXIndices1)
*/
func (e *ScalarLinearExpr) RewriteInTermsOfIndices(newXIndices []uint64) (*ScalarLinearExpr, error) {
	// Create new Linear Express
	var newLE ScalarLinearExpr = ScalarLinearExpr{
		XIndices: newXIndices,
	}

	// Find length of X indices
	numIndices := len(newXIndices)

	// Create L matrix of appropriate dimension
	var newL []float64
	for rowIndex := 0; rowIndex < numIndices; rowIndex++ {
		newL = append(newL, 0.0)
	}

	// Populate L
	for oi1Index, oldIndex1 := range e.XIndices {
		// Identify what term is associated with the pair (oldIndex1, oldIndex2)
		oldLterm := e.L[oi1Index]

		// Get the new indices corresponding to oi1 and oi2
		ni1Index, err := FindInSlice(oldIndex1, newXIndices)
		if err != nil {
			return &newLE, fmt.Errorf("The index %v was found in the old X indices, but it does not exist in the new ones!", oldIndex1)
		}
		newIndex1 := newXIndices[ni1Index]

		// Plug the oldQterm into newQ
		newL[newIndex1] += oldLterm
	}
	newLE.L = newL

	// Populate C
	newLE.C = e.C

	return &newLE, nil

}
