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
		return Comparison(lhsAsK, rhs, sense)
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
func Multiply(term1, term2 interface{}) (Expression, error) {
	// Constants

	// Algorithm
	switch term1.(type) {
	case float64:
		// Convert lhs to K
		term1AsFloat64, _ := term1.(float64)
		term1AsK := K(term1AsFloat64)

		// Create constraint
		return Multiply(term1AsK, term2)
	case mat.VecDense:
		// Convert lhs to KVector.
		term1AsVecDense, _ := term1.(mat.VecDense)
		term1AsKVector := KVector(term1AsVecDense)

		// Create constraint
		return term1AsKVector.Multiply(term2)
	//case ScalarExpression:
	//	term1AsScalarExpression, _ := term1.(ScalarExpression)
	//
	//	// Create Constraint
	//	return term1AsScalarExpression.Multiply(term2)
	//
	//case VectorExpression:
	//	lhsAsVecExpr, _ := lhs.(VectorExpression)
	//	return lhsAsVecExpr.Comparison(rhs, sense)
	default:
		return nil, fmt.Errorf("Multiply of %T term with %T term is not yet defined!", term1, term2)
	}
}

// Dot returns the dot product of a vector of variables and slice of floats.
func Dot(vs []Variable, coeffs []float64) ScalarExpression {
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

// Sum returns the sum of the given expressions. It creates a new empty
// expression and adds to it the given expressions.
func Sum(exprs ...ScalarExpression) (ScalarExpression, error) {
	sum := NewExpr(0)
	var err error
	for _, e := range exprs {
		sum, err = sum.Plus(e)
		if err != nil {
			return sum, fmt.Errorf("Error computing Plus() on %v,%v: %v", sum, e, err)
		}
	}

	return sum, nil
}
