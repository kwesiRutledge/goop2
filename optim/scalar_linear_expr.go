package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// ScalarLinearExpr represents a linear general expression of the form
//
//	L' * x + C
//
// where L is a vector of coefficients that matches the dimension of x, the vector of variables
// variables and C is a constant
type ScalarLinearExpr struct {
	X VarVector
	L mat.VecDense // Vector of coefficients. Should match the dimensions of XIndices
	C float64
}

// NewLinearExpr returns a new expression with a single additive constant
// value, c, and no variables.
func NewLinearExpr(c float64) ScalarExpression {
	return &ScalarLinearExpr{C: c}
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the linear expression le.
*/
func (le *ScalarLinearExpr) Variables() []Var {
	return UniqueVars(le.X.Elements)
}

// NumVars returns the number of variables in the expression
func (e *ScalarLinearExpr) NumVars() int {
	return e.X.Len()
}

// Vars returns a slice of the Var ids in the expression
func (e *ScalarLinearExpr) IDs() []uint64 {
	return e.X.IDs()
}

// Coeffs returns a slice of the coefficients in the expression
func (e *ScalarLinearExpr) Coeffs() []float64 {
	var coeffsOut []float64
	for i := 0; i < e.L.Len(); i++ {
		coeffsOut = append(coeffsOut, e.L.AtVec(i))
	}
	return coeffsOut
}

// Constant returns the constant additive value in the expression
func (e *ScalarLinearExpr) Constant() float64 {
	return e.C
}

// Plus adds the current expression to another and returns the resulting
// expression
func (e *ScalarLinearExpr) Plus(eIn ScalarExpression, extras ...interface{}) (ScalarExpression, error) {
	// Algorithm depends
	switch eIn.(type) {
	//case *QuadraticExpr:
	//
	//	var newQExpr QuadraticExpr = *qe // get copy of e
	//	quadraticEIn := eIn.(*QuadraticExpr)
	//
	//	// Get Combined set of Variables
	//	newX := UniqueVars(append(newQExpr.X.Elements, quadraticEIn.X.Elements...))
	//	newQExprAligned, _ := newQExpr.RewriteInTermsOf(VarVector{newX})
	//	quadraticEInAligned, _ := quadraticEIn.RewriteInTermsOf(VarVector{newX})
	//
	//	// Add matrices together
	//	var tempSum mat.Dense
	//	tempSum.Add(&newQExprAligned.Q, &quadraticEInAligned.Q)
	//	newQExprAligned.Q = tempSum
	//
	//	// Add vectors together
	//	//var tempVecSum mat.VecDense
	//	//tempVecSum.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)
	//	newQExprAligned.L.AddVec(&newQExprAligned.L, &quadraticEInAligned.L)
	//
	//	// Add constants together
	//	newQExprAligned.C += quadraticEInAligned.C
	//	return newQExprAligned, nil
	//
	//case *ScalarLinearExpr:
	//	// Collect Expressions
	//	var newQExpr QuadraticExpr = *qe // get copy of e
	//	linearEIn := eIn.(*ScalarLinearExpr)
	//
	//	// Get Combined set of Variables
	//	newX := UniqueVars(append(newQExpr.X.Elements, linearEIn.X.Elements...))
	//	newQExprAligned, _ := newQExpr.RewriteInTermsOf(VarVector{newX})
	//	linearEInAligned, _ := linearEIn.RewriteInTermsOf(VarVector{newX})
	//
	//	// Add linear vector together with the quadratic expression
	//	//var vectorSum mat.VecDense
	//	//vectorSum.AddVec(newQExprAligned.L, linearEInAligned.L)
	//	newQExprAligned.L.AddVec(&newQExprAligned.L, &linearEInAligned.L)
	//	//for eltInd, qElt := range linearEInAligned.L {
	//	//	newQExprAligned.L[eltInd] += qElt
	//	//}
	//
	//	// Add constants together
	//	newQExprAligned.C += linearEIn.C
	//	return newQExprAligned, nil
	default:
		fmt.Println("Unexpected type given to Plus().")

		return &QuadraticExpr{}, fmt.Errorf("Unexpected type given as first argument to Plus %T.", eIn)
	}
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (e *ScalarLinearExpr) Mult(c float64) (ScalarExpression, error) {
	e.L.ScaleVec(c, &e.L)
	e.C *= c

	return e, nil
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (e *ScalarLinearExpr) LessEq(other ScalarExpression) (ScalarConstraint, error) {
	return e.Comparison(other, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (e *ScalarLinearExpr) GreaterEq(other ScalarExpression) (ScalarConstraint, error) {
	return e.Comparison(other, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (e *ScalarLinearExpr) Eq(other ScalarExpression) (ScalarConstraint, error) {
	return e.Comparison(other, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := e.Comparison(expr1,SenseGreaterThanEqual)
*/
func (e *ScalarLinearExpr) Comparison(rhs ScalarExpression, sense ConstrSense) (ScalarConstraint, error) {
	return ScalarConstraint{e, rhs, sense}, nil
}

/*
RewriteInTermsOfIndices
Description:

	Rewrites the current linear expression in terms of the new variables.

Usage:

	rewrittenLE, err := orignalLE.RewriteInTermsOfIndices(newXIndices1)
*/
func (e *ScalarLinearExpr) RewriteInTermsOf(newX VarVector) (*ScalarLinearExpr, error) {
	// Create new Linear Express
	var newLE ScalarLinearExpr = ScalarLinearExpr{
		X: newX,
	}

	// Find length of X indices
	dimX := newX.Len()

	// Create L matrix of appropriate dimension
	var newLfloat []float64
	for rowIndex := 0; rowIndex < dimX; rowIndex++ {
		newLfloat = append(newLfloat, 0.0)
	}
	newL := mat.NewVecDense(dimX, newLfloat)

	// Populate L
	for oi1Index, oldElt1 := range e.X.Elements {
		// Identify what term is associated with the pair (oldIndex1, oldIndex2)
		oldLterm := e.L.AtVec(oi1Index)

		// Get the new indices corresponding to oi1 and oi2
		ni1Index, err := FindInSlice(oldElt1, newX.Elements)
		if err != nil {
			return &newLE, fmt.Errorf("The element %v was found in the old X indices, but it does not exist in the new ones!", oldElt1)
		}
		//newElt1 := newX.Elements[ni1Index]

		// Plug the oldLinearterm into newLinear expression
		offset := ZerosVector(e.X.Len())
		offset.SetVec(ni1Index, oldLterm)

		newL.AddVec(newL, &offset)
	}
	newLE.L = *newL

	// Populate C
	newLE.C = e.C

	return &newLE, nil

}
