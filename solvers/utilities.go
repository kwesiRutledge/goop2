package solvers

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"

	"github.com/kwesiRutledge/gurobi.go/gurobi"
)

func VarTypeToGRBVType(goopTypeIn optim.VarType) (rune, error) {
	// Double check

	switch goopTypeIn {
	case optim.Continuous:
		return gurobi.CONTINUOUS, nil
	case optim.Binary:
		return gurobi.BINARY, nil
	case optim.Integer:
		return gurobi.INTEGER, nil
	default:
		return -1, fmt.Errorf("The goop variable type \"%v\" is not currently supported by VarTypeToGRBVType.", goopTypeIn)

	}
}
