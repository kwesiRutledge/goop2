package optim

import "gonum.org/v1/gonum/mat"

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type Var struct {
	ID    uint64
	Lower float64
	Upper float64
	Vtype VarType
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the variable expression v.
*/
func (v Var) Variables() []Var {
	return []Var{v}
}

// NumVars returns the number of variables in the expression. For a variable, it
// always returns one.
func (v Var) NumVars() int {
	return 1
}

// Vars returns a slice of the Var ids in the expression. For a variable, it
// always returns a singleton slice with the given variable ID.
func (v Var) IDs() []uint64 {
	return []uint64{v.ID}
}

// Coeffs returns a slice of the coefficients in the expression. For a variable,
// it always returns a singleton slice containing the value one.
func (v Var) Coeffs() []float64 {
	return []float64{1}
}

// Constant returns the constant additive value in the expression. For a
// variable, it always returns zero.
func (v Var) Constant() float64 {
	return 0
}

// Plus adds the current expression to another and returns the resulting
// expression.
func (v Var) Plus(e ScalarExpression, extras ...interface{}) (ScalarExpression, error) {
	// Input Processing??

	// Algorithm
	vv := VarVector{
		UniqueVars(append([]Var{v}, e.Variables()...)),
	}
	coeffs := append([]float64{1}, e.Coeffs()...)
	newExpr := &ScalarLinearExpr{
		X: vv,
		L: *mat.NewVecDense(vv.Len(), coeffs),
		C: e.Constant(),
	}
	return newExpr, nil
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (v Var) Mult(m float64) (ScalarExpression, error) {
	// Constants
	// switch m.(type) {
	// case float64:

	vars := []Var{v}
	coeffs := []float64{m * v.Coeffs()[0]}

	// Algorithm
	newExpr := &ScalarLinearExpr{
		X: VarVector{vars},
		L: *mat.NewVecDense(1, coeffs),
		C: 0,
	}
	return newExpr, nil
	// case *Var:
	// 	return nil
	// }
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (v Var) LessEq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (v Var) GreaterEq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (v Var) Eq(other ScalarExpression) (ScalarConstraint, error) {
	return v.Comparison(other, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := v.Comparison(expr1,SenseGreaterThanEqual)
*/
func (v Var) Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error) {
	// Constants

	// Algorithm
	return ScalarConstraint{v, rhs, sense}, nil
}

/*
// ID returns the ID of the variable
func (v *Var) ID() uint64 {
	return v.ID
}

// Lower returns the lower value limit of the variable
func (v *Var) Lower() float64 {
	return v.Lower
}

// Upper returns the upper value limit of the variable
func (v *Var) Upper() float64 {
	return v.Upper
}

// Type returns the type of variable (continuous, binary, integer, etc)
func (v *Var) Type() VarType {
	return v.Vtype
}
*/

// VarType represents the type of the variable (continuous, binary,
// integer, etc) and uses Gurobi's encoding.
type VarType byte

// Multiple common variable types have been included as constants that conform
// to Gurobi's encoding.
const (
	Continuous VarType = 'C'
	Binary             = 'B'
	Integer            = 'I'
)

/*
UniqueVars
Description:

	This function creates a slice of unique variables from the slice given in
	varsIn
*/
func UniqueVars(varsIn []Var) []Var {
	// Constants

	// Algorithm
	var varsOut []Var
	for _, v := range varsIn {
		if vIndex, _ := FindInSlice(v, varsOut); vIndex == -1 { // If v is not yet in varsOut, then add it
			varsOut = append(varsOut, v)
		}
	}

	return varsOut

}
