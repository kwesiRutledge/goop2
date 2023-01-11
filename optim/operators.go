package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"

	log "github.com/sirupsen/logrus"
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
	case float64:
		// Convert lhs to K
		lhsAsFloat64, _ := lhs.(float64)
		lhsAsK := K(lhsAsFloat64)

		// Create constraint
		return lhsAsK.Comparison(rhs, sense)
	case mat.VecDense:
		// Convert lhs to KVector.
		lhsAsVecDense, _ := lhs.(mat.VecDense)
		lhsAsKVector := KVector(lhsAsVecDense)

		// Create constraint
		return lhsAsKVector.Comparison(rhs, sense)
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

/*
Multiply
Description:

	Defines the multiplication between two objects.
*/
//func Multiply(term1, term2 interface{}) (Expression, error) {
//	// Constants
//
//	// Algorithm
//	switch term1.(type) {
//	case
//	}
//}

// Dot returns the dot product of a vector of variables and slice of floats.
func Dot(vs []Var, coeffs []float64) ScalarExpression {
	if len(vs) != len(coeffs) {
		log.WithFields(log.Fields{
			"num_vars":   len(vs),
			"num_coeffs": len(coeffs),
		}).Panic("Number of vars and coeffs mismatch")
	}

	newExpr := NewExpr(0)
	for i := range vs {
		newExpr.Plus(vs[i].Mult(coeffs[i]))
	}

	return newExpr
}
