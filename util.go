package goop2

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Sum returns the sum of the given expressions. It creates a new empty
// expression and adds to it the given expressions.
func Sum(exprs ...Expr) Expr {
	newExpr := NewExpr(0)
	for _, e := range exprs {
		newExpr.Plus(e)
	}

	return newExpr
}

// SumVars returns the sum of the given variables. It creates a new empty
// expression and adds to it the given variables.
func SumVars(vs ...*Var) Expr {
	newExpr := NewExpr(0)
	for _, v := range vs {
		newExpr.Plus(v)
	}
	return newExpr
}

// SumRow returns the sum of all the variables in a single specified row of
// a variable matrix.
func SumRow(vs [][]*Var, row int) Expr {
	newExpr := NewExpr(0)
	for col := 0; col < len(vs[0]); col++ {
		newExpr.Plus(vs[row][col])
	}
	return newExpr
}

// SumCol returns the sum of all variables in a single specified column of
// a variable matrix.
func SumCol(vs [][]*Var, col int) Expr {
	newExpr := NewExpr(0)
	for row := 0; row < len(vs); row++ {
		newExpr.Plus(vs[row][col])
	}
	return newExpr
}

// Dot returns the dot product of a vector of variables and slice of floats.
func Dot(vs []*Var, coeffs []float64) Expr {
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

/*
FindInSlice
Description:
	Identifies if the  input xIn is in the slice sliceIn.
	If it is, then this function returns the index such that xIn = sliceIn[index] and no errors.
	If it is not, then this function returns the index -1 and the boolean value false.
*/
func FindInSlice(xIn interface{}, sliceIn interface{}) (int, error) {

	switch xIn.(type) {
	case string:
		x := xIn.(string)
		slice := sliceIn.([]string)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case int:
		x := xIn.(int)
		slice := sliceIn.([]int)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case uint64:
		x := xIn.(uint64)
		slice := sliceIn.([]uint64)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	default:
		return -1, fmt.Errorf("The FindInSlice() function was only defined for type string, not type %T:", xIn)
	}

}
