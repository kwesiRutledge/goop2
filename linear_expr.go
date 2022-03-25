package goop2

// LinearExpr represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant
type LinearExpr struct {
	variables    []uint64
	coefficients []float64
	constant     float64
}

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewLinearExpr(c float64) Expr {
	return &LinearExpr{constant: c}
}

// NumVars returns the number of variables in the expression
func (e *LinearExpr) NumVars() int {
	return len(e.variables)
}

// Vars returns a slice of the Var ids in the expression
func (e *LinearExpr) Vars() []uint64 {
	return e.variables
}

// Coeffs returns a slice of the coefficients in the expression
func (e *LinearExpr) Coeffs() []float64 {
	return e.coefficients
}

// Constant returns the constant additive value in the expression
func (e *LinearExpr) Constant() float64 {
	return e.constant
}

// Plus adds the current expression to another and returns the resulting
// expression
func (e *LinearExpr) Plus(other Expr) Expr {
	e.variables = append(e.variables, other.Vars()...)
	e.coefficients = append(e.coefficients, other.Coeffs()...)
	e.constant += other.Constant()
	return e
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (e *LinearExpr) Mult(c float64) Expr {
	for i, coeff := range e.coefficients {
		e.coefficients[i] = coeff * c
	}
	e.constant *= c

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
