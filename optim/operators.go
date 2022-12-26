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
			SenseEqual,
		}, nil
	case VectorExpression:
		lhsAsVecExpr, _ := lhs.(VectorExpression)
		return lhsAsVecExpr.Eq(rhs)
	}

	return nil, fmt.Errorf("Not implemented!")
}

// LessEq returns a constraint representing lhs <= rhs
func LessEq(lhs, rhs ScalarExpression) ScalarConstraint {
	return ScalarConstraint{lhs, rhs, SenseLessThanEqual}
}

// GreaterEq returns a constraint representing lhs >= rhs
func GreaterEq(lhs, rhs ScalarExpression) ScalarConstraint {
	return ScalarConstraint{lhs, rhs, SenseGreaterThanEqual}
}
