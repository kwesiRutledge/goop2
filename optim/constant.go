package optim

import (
	"fmt"
)

// Integer constants represnting commonly used numbers. Makes for better
// readability
const (
	Zero = K(0)
	One  = K(1)
)

// K is a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
type K float64

/*
Variables
Description:

	Shares all variables included in the expression that is K.
	It is a constant, so there are none.
*/
func (c K) Variables() []Var {
	return []Var{}
}

// NumVars returns the number of variables in the expression. For constants,
// this is always 0
func (c K) NumVars() int {
	return 0
}

// Vars returns a slice of the Var ids in the expression. For constants,
// this is always nil
func (c K) IDs() []uint64 {
	return nil
}

// Coeffs returns a slice of the coefficients in the expression. For constants,
// this is always nil
func (c K) Coeffs() []float64 {
	return nil
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

// Plus adds the current expression to another and returns the resulting
// expression
func (c K) Plus(e ScalarExpression, extras ...interface{}) (ScalarExpression, error) {
	switch e.(type) {
	case K:
		return K(c.Constant() + e.Constant()), nil
	default:
		fmt.Sprintf("Unrecognized expression type %T!", e)
		return c, fmt.Errorf("Unexpected type in Plus() for constant %v: %T", e)
	}
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (c K) Mult(val float64) (ScalarExpression, error) {
	return K(float64(c) * val), nil
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(other ScalarExpression) (ScalarConstraint, error) {
	return c.Comparison(other, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(other ScalarExpression) (ScalarConstraint, error) {
	return c.Comparison(other, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(other ScalarExpression) (ScalarConstraint, error) {
	return c.Comparison(other, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.
*/
func (c K) Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error) {
	// Constants

	// Algorithm
	return ScalarConstraint{c, rhs, sense}, nil
}

/*
Multiply
Description:

	This method multiplies the input constant by another expression.
*/
func (c K) Multiply(term1 interface{}) (Expression, error) {
	// Constants

	// Algorithm
	switch term1.(type) {
	case float64:
		term1AsFloat, _ := term1.(float64)
		return c.Multiply(K(term1AsFloat))
	case K:
		term1AsK, _ := term1.(K)
		return c * term1AsK, nil
	//case ScalarLinearExpr:
	//	term1AsSLE, _ := term1.(ScalarLinearExpr)
	//	product := QuadraticExpr{
	//		Q: float64[][]{}, // TODO: Finish this. Maybe switch quadratic expression to use gonum.Matrix objects instead of the current system.
	//	}
	default:
		return K(0), fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1)

	}
}
