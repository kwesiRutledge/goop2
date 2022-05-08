package goop2

// LinearExpr represents a linear general expression of the form
//	L' * x + C
// where L is a vector of coefficients that matches the dimension of x, the vector of variables
// variables and C is a constant
type LinearExpr struct {
	XIndices []uint64
	L        []float64 // Vector of coefficients. Should match the dimensions of XIndices
	C        float64
}

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewLinearExpr(c float64) Expr {
	return &LinearExpr{C: c}
}

// NumVars returns the number of variables in the expression
func (e *LinearExpr) NumVars() int {
	return len(e.XIndices)
}

// Vars returns a slice of the Var ids in the expression
func (e *LinearExpr) Vars() []uint64 {
	return e.XIndices
}

// Coeffs returns a slice of the coefficients in the expression
func (e *LinearExpr) Coeffs() []float64 {
	return e.L
}

// Constant returns the constant additive value in the expression
func (e *LinearExpr) Constant() float64 {
	return e.C
}

// Plus adds the current expression to another and returns the resulting
// expression
func (e *LinearExpr) Plus(other Expr) Expr {
	e.XIndices = append(e.XIndices, other.Vars()...)
	e.L = append(e.L, other.Coeffs()...)
	e.C += other.Constant()
	return e
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (e *LinearExpr) Mult(c float64) Expr {
	for i, coeff := range e.L {
		e.L[i] = coeff * c
	}
	e.C *= c

	return e
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (e *LinearExpr) LessEq(other Expr) *Constr {
	return LessEq(e, other)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (e *LinearExpr) GreaterEq(other Expr) *Constr {
	return GreaterEq(e, other)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (e *LinearExpr) Eq(other Expr) *Constr {
	return Eq(e, other)
}
