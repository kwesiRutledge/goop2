package goop2

/*
quadratic_expr.go
Description:
	Defines some of the functions necessary to define polynomial expressions in terms of the variables
	of an optimization problem.
*/

// Type Definitions
// ================

/*
MonomialExpr
Description:
	A monomial of optimization variables (given by their indices).
	The monomial object defines a monomial written as follows:
		Coefficients[0] * Variables[0][0] * Variables[0][1] + Coefficients[1] * Variables[1][0] * Variables[1][1] + ... + Coefficients[n] * Variables[n][0] * Variables[n][1]
*/
type QuadraticExpr struct {
	Coefficients  []float64
	VariablePairs [][2]uint64
}

// Member Functions
// ================

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewQuadraticExpr(c float64, v1, v2 uint64) *QuadraticExpr {
	return &QuadraticExpr{
		Coefficients: []float64{c},
		VariablePairs: [][2]uint64{
			[2]uint64{v1, v2},
		},
	}
}

/*
NumVars
Description:
	Returns the number of variables in the expression.
	To make this actually meaningful, we only count the unique vars.
*/
func (e *QuadraticExpr) NumVars() int {

	// Counting Loop
	var uniqueVars []uint64
	for _, variablePairs := range e.VariablePairs {
		// Check to see if each vp

		// FINISH HERE !!!
		elt1 := variablePairs[0]
		if ind1, _ := FindInSlice(elt1, uniqueVars); ind1 == -1 {
			// If elt1 is not in uniqueVars, then add it.
			uniqueVars = append(uniqueVars, elt1)
		}

		elt2 := variablePairs[1]
		if ind2, _ := FindInSlice(elt2, uniqueVars); ind2 == -1 {
			// If elt1 is not in uniqueVars, then add it.
			uniqueVars = append(uniqueVars, elt2)
		}

	}

	return len(uniqueVars)
}

// // Vars returns a slice of the Var ids in the expression
// func (e *LinearExpr) Vars() []uint64 {
// 	return e.variables
// }

// // Coeffs returns a slice of the coefficients in the expression
// func (e *LinearExpr) Coeffs() []float64 {
// 	return e.coefficients
// }

// // Constant returns the constant additive value in the expression
// func (e *LinearExpr) Constant() float64 {
// 	return e.constant
// }

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
// func (e *LinearExpr) Mult(c float64) Expr {
// 	for i, coeff := range e.coefficients {
// 		e.coefficients[i] = coeff * c
// 	}
// 	e.constant *= c

// 	return e
// }

// // LessEq returns a less than or equal to (<=) constraint between the
// // current expression and another
// func (e *LinearExpr) LessEq(other Expr) *Constr {
// 	return LessEq(e, other)
// }

// // GreaterEq returns a greater than or equal to (>=) constraint between the
// // current expression and another
// func (e *LinearExpr) GreaterEq(other Expr) *Constr {
// 	return GreaterEq(e, other)
// }

// // Eq returns an equality (==) constraint between the current expression
// // and another
// func (e *LinearExpr) Eq(other Expr) *Constr {
// 	return Eq(e, other)
// }
