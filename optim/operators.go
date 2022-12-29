package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
operators.go
Description:
	Defines the operators that transform variables and expressions into expressions or constraints.
*/

/*
Eq
Description:

	Returns a constraint representing lhs == rhs
*/
func Eq(lhs, rhs interface{}) (Constraint, error) {
	return Comparison(lhs, rhs, SenseEqual)
}

// LessEq returns a constraint representing lhs <= rhs
func LessEq(lhs, rhs interface{}) (Constraint, error) {
	return Comparison(lhs, rhs, SenseLessThanEqual)
}

// GreaterEq returns a constraint representing lhs >= rhs
func GreaterEq(lhs, rhs ScalarExpression) (Constraint, error) {
	return Comparison(lhs, rhs, SenseGreaterThanEqual)
}

/*
Comparison
Description:

	Compares the two inputs lhs (Left Hand Side) and rhs (Right Hand Side) in the sense provided in sense.

Usage:

	constr, err := Comparison(expr1, expr2, SenseGreaterThanEqual)
*/
func Comparison(lhs, rhs interface{}, sense ConstrSense) (Constraint, error) {
	// Constants

	// Algorithm
	switch lhs.(type) {
	case mat.VecDense:
		// Convert lhs to KVector.
		lhsAsVecDense, _ := lhs.(mat.VecDense)
		lhsAsKVector := KVector(lhsAsVecDense)

		// Create constraint
		return lhsAsKVector.Eq(rhs)
	case ScalarExpression:
		lhsAsScalarExpression, _ := lhs.(ScalarExpression)
		rhsAsScalarExpression, _ := rhs.(ScalarExpression)
		return ScalarConstraint{
			lhsAsScalarExpression,
			rhsAsScalarExpression,
			sense,
		}, nil
	case VectorExpression:
		lhsAsVecExpr, _ := lhs.(VectorExpression)
		return lhsAsVecExpr.Comparison(rhs, sense)
	default:
		return nil, fmt.Errorf("Comparison in sense '%v' is not defined for lhs type %T and rhs type %T!", sense, lhs, rhs)
	}

}
