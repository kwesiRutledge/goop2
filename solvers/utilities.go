package solvers

import (
	"fmt"

	"github.com/kwesiRutledge/goop2"
	"github.com/kwesiRutledge/gurobi.go/gurobi"
)

func VarTypeToGRBVType(goopTypeIn goop2.VarType) (rune, error) {
	// Double check

	switch goopTypeIn {
	case goop2.Continuous:
		return gurobi.CONTINUOUS, nil
	case goop2.Binary:
		return gurobi.BINARY, nil
	case goop2.Integer:
		return gurobi.INTEGER, nil
	default:
		return -1, fmt.Errorf("The goop variable type \"%v\" is not currently supported by VarTypeToGRBVType.", goopTypeIn)

	}
}
