package goop2

import (
	"fmt"
	"os"
)

/*
quadratic_expr.go
Description:
	Defines some of the functions necessary to define polynomial expressions in terms of the variables
	of an optimization problem.
*/

// Type Definitions
// ================

/*
QuadraticExpr
Description:
	A quadratic expression of optimization variables (given by their indices).
	The quadratic expression object defines a quadratic written as follows:
		x' * Q * x + q * x + b
*/
type QuadraticExpr struct {
	Q        [][]float64
	q        []float64
	b        float64
	XIndices []uint64
}

// Member Functions
// ================

/*
NewQuadraticExpr_qb0
Description:
	NewQuadraticExpr_q0 returns a basic Quadratic expression with only the matrix Q being defined,
	all other values are assumed to be zero.
*/
func NewQuadraticExpr_qb0(QIn [][]float64, xIndicesIn []uint64) (*QuadraticExpr, error) {
	// Constants
	numXIndices := len(xIndicesIn)

	// Input Checking

	// Algorithm
	var qZero []float64
	for qInd := 0; qInd < numXIndices; qInd++ {
		qZero = append(qZero, 0.0)
	}

	return NewQuadraticExpr(QIn, qZero, 0.0, xIndicesIn)
}

/*
NewQuadraticExpr
Description:
	NewQuadraticExpr returns a basic Quadratic expression whuch is defined by QIn, qIn and bIn.
*/
func NewQuadraticExpr(QIn [][]float64, qIn []float64, bIn float64, xIndicesIn []uint64) (*QuadraticExpr, error) {
	// Constants
	numXIndices := len(xIndicesIn)

	// Input Checking
	if len(QIn) != numXIndices {
		return &QuadraticExpr{}, fmt.Errorf("The number of indices was %v which did not match the first dimension of QIn (%v).", numXIndices, len(QIn))
	}

	for rowIndex, QRow := range QIn {
		if len(QRow) != numXIndices {
			return &QuadraticExpr{}, fmt.Errorf("The number of indices was %v which did not match the length of QIn's %vth row (%v).", numXIndices, rowIndex, len(QRow))
		}
	}

	if len(qIn) != numXIndices {
		return &QuadraticExpr{}, fmt.Errorf("The number of indices was %v which did not match the length of qIn (%v).", numXIndices, len(qIn))
	}

	// Algorithm

	return &QuadraticExpr{
		Q:        QIn,
		q:        qIn,
		b:        bIn,
		XIndices: xIndicesIn,
	}, nil
}

/*
Check
Description:
	This function checks the dimensions of all of the members of the quadratic expression which are slices.
	They should have compatible dimensions.
*/
func (e *QuadraticExpr) Check() error {
	// Make the number of elements in q be the dimension of the x in the expression.
	numXIndices := len(e.q)

	// Check Number of Rows in Q
	if len(e.Q) != numXIndices {
		return fmt.Errorf("The nuber of indices was %v which did not match the first dimension of QIn (%v).", numXIndices, len(e.Q))
	}

	for rowIndex, QRow := range e.Q {
		if len(QRow) != numXIndices {
			return fmt.Errorf("The nuber of indices was %v which did not match the length of QIn's %vth row (%v).", numXIndices, rowIndex, len(QRow))
		}
	}

	// Otherwise, return no errors.
	return nil
}

/*
NumVars
Description:
	Returns the number of variables in the expression.
	To make this actually meaningful, we only count the unique vars.
*/
func (e *QuadraticExpr) NumVars() int {

	return len(e.Vars())
}

/*
Vars
Description:
	Returns the ids of all of the variables in the quadratic expression.
*/
func (e *QuadraticExpr) Vars() []uint64 {
	return e.XIndices
}

/*
Coeffs
Description:
	Returns the slice of all coefficient values for each pair of variable tuples.
	The coefficients of the quadratic expression are created in an ordering that comes from the following vector.

	Consider xI (the indices of the input expression e). The output coefficients will be c.
	The coefficients of the expression
		e = x' Q x + q' * x + b
	will be
		e = c' mx + b
	where
		mx = [ x[0]*x[0], x[0]*x[1], ... , x[0]*x[N-1], x[1]*x[1] , x[1]*x[2], ... , x[1]*x[N-1], x[2]*x[2], ... , x[N-1]*x[N-1], x[0], x[1], ... , x[N-1] ]
*/
func (e *QuadraticExpr) Coeffs() []float64 {
	// Create container for all coefficients
	var coefficientList []float64
	var numVars int = e.NumVars()

	// Consider all pairs of indices in x.
	var xPairs [][2]uint64
	for vIIndex, varIndex := range e.XIndices {
		for vIIndex2 := vIIndex; vIIndex2 < numVars; vIIndex2++ {
			varIndex2 := e.XIndices[vIIndex2]

			// Save pairs of indices and the associated coefficients
			xPairs = append(xPairs, [2]uint64{varIndex, varIndex2})

			if vIIndex == vIIndex2 {
				coefficientList = append(coefficientList, e.Q[vIIndex][vIIndex2])
			} else {
				coefficientList = append(coefficientList, e.Q[vIIndex][vIIndex2]+e.Q[vIIndex2][vIIndex])
			}

		}
	}

	return coefficientList
}

/*
Constant
Description:
	Returns the constant value associated with a quadratic expression.
*/
func (e *QuadraticExpr) Constant() float64 {
	return e.b
}

/*
Plus
Description:

*/
func (e *QuadraticExpr) Plus(eIn Expr) Expr {
	// Constants
	var newQExpr QuadraticExpr = *e

	// Algorithm depends
	switch eIn.(type) {
	case *QuadraticExpr:
		// Add matrices together
		quadraticEIn := eIn.(*QuadraticExpr)
		for rowInd, Qrow := range quadraticEIn.Q {
			for colInd, Qval := range Qrow {
				newQExpr.Q[rowInd][colInd] += Qval
			}
		}

		// Add vectors together
		for eltInd, qElt := range quadraticEIn.q {
			newQExpr.q[eltInd] += qElt
		}

		// Add constants together
		newQExpr.b += quadraticEIn.b

	default:
		fmt.Println("Unexpected type given to Plus().")
		os.Exit(1)
	}

	return &newQExpr

}

// // Plus adds the current expression to another and returns the resulting
// // expression
// func (e *LinearExpr) Plus(other Expr) Expr {
// 	e.variables = append(e.variables, other.Vars()...)
// 	e.coefficients = append(e.coefficients, other.Coeffs()...)
// 	e.constant += other.Constant()
// 	return e
// }

// // Mult multiplies the current expression to another and returns the
// // resulting expression
/*
Mult
Description:
	Mult multiplies the current expression to another and returns the
	resulting expression
*/
func (e *QuadraticExpr) Mult(c float64) Expr {
	// Iterate through all of the rows and columns of Q
	nV := e.NumVars()
	for i := 0; i < nV; i++ {
		for j := 0; j < nV; j++ {
			e.Q[i][j] = e.Q[i][j] * c
		}
	}

	// Iterate through the linear coefficients
	for i := 0; i < nV; i++ {
		e.q[i] = e.q[i] * c
	}

	// Update through the constant
	e.b *= c

	return e
}

/*
LessEq
Description:
	LessEq returns a less than or equal to (<=) constraint between the
	current expression and another
*/
func (e *QuadraticExpr) LessEq(other Expr) *Constr {
	return LessEq(e, other)
}

/*
GreaterEq
Description:
	GreaterEq returns a greater than or equal to (>=) constraint between the
	current expression and another
*/
func (e *QuadraticExpr) GreaterEq(other Expr) *Constr {
	return GreaterEq(e, other)
}

/*
Eq
Description:
	Form an equality constraint with this equality constraint and another
	Eq returns an equality (==) constraint between the current expression
	and another
*/
func (e *QuadraticExpr) Eq(other Expr) *Constr {
	return Eq(e, other)
}
