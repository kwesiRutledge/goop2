package optim

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type Var struct {
	ID    uint64
	Lower float64
	Upper float64
	Vtype VarType
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
func (v Var) Plus(e Expr) Expr {
	vars := append([]uint64{v.ID}, e.IDs()...)
	coeffs := append([]float64{1}, e.Coeffs()...)
	newExpr := &LinearExpr{
		XIndices: vars,
		L:        coeffs,
		C:        e.Constant(),
	}
	return newExpr
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (v Var) Mult(m float64) Expr {
	// Constants
	// switch m.(type) {
	// case float64:

	vars := []uint64{v.ID}
	coeffs := []float64{m * v.Coeffs()[0]}

	// Algorithm
	newExpr := &LinearExpr{
		XIndices: vars,
		L:        coeffs,
		C:        0,
	}
	return newExpr
	// case *Var:
	// 	return nil
	// }
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (v Var) LessEq(other Expr) *Constr {
	return LessEq(v, other)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (v Var) GreaterEq(other Expr) *Constr {
	return GreaterEq(v, other)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (v Var) Eq(other Expr) *Constr {
	return Eq(v, other)
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
