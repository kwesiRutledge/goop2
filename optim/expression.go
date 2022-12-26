package optim

/*
expression.go
Description:
	This file holds all of the functions and methods related to the Expression
	interface.
*/

/*
Expression
Description:

	This interface should be implemented by and ScalarExpression and VectorExpression
*/
type Expression interface {
	// NumVars returns the number of variables in the expression
	NumVars() int

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e Expression) Expression

	// Mult multiplies the current expression to another and returns the
	// resulting expression
	Mult(c Expression) ScalarExpression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(e Expression) Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(e Expression) Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(e ScalarExpression) *ScalarConstraint
}
