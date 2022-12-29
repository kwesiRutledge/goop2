package optim

import (
	"errors"
	"fmt"
	"github.com/kwesiRutledge/gurobi.go/gurobi"
	"time"

	"github.com/sirupsen/logrus"
)

// Model represents the overall constrained linear optimization model to be
// solved. Model contains all the variables associated with the optimization
// problem, constraints, objective, and parameters. New variables can only be
// created using an instantiated Model.
type Model struct {
	Variables []Var
	constrs   []ScalarConstraint
	obj       *Objective
	showLog   bool
	timeLimit time.Duration
}

// NewModel returns a new model with some default arguments such as not to show
// the log and no time limit.
func NewModel() *Model {
	return &Model{showLog: false}
}

// ShowLog instructs the solver to show the log or not.
func (m *Model) ShowLog(shouldShow bool) {
	m.showLog = shouldShow
}

// SetTimeLimit sets the solver time limit for the model.
func (m *Model) SetTimeLimit(dur time.Duration) {
	m.timeLimit = dur
}

/*
AddVar
Description:

	This method adds an "unbounded" continuous variable to the model.
*/
func (m *Model) AddVar() Var {
	return m.AddRealVar()
}

/*
AddRealVar
Description:

	Adds a Real variable to the model and returns said variable.
*/
func (m *Model) AddRealVar() Var {
	return m.AddVarClassic(-gurobi.INFINITY, gurobi.INFINITY, Continuous)
}

// AddVar adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
func (m *Model) AddVarClassic(lower, upper float64, vtype VarType) Var {
	id := uint64(len(m.Variables))
	newVar := Var{id, lower, upper, vtype}
	m.Variables = append(m.Variables, newVar)
	return newVar
}

// AddBinaryVar adds a binary variable to the model and returns said variable.
func (m *Model) AddBinaryVar() Var {
	return m.AddVarClassic(0, 1, Binary)
}

// AddVarVector adds a vector of variables of a given variable type to the
// model. It then returns the resulting slice.
/*
AddVarVector
Description:
	Creates a VarVector object using a constructor that assumes you want an "unbounded" vector of real optimization
	variables.
*/
func (m *Model) AddVarVector(dim int) VarVector {
	// Constants

	// Algorithm
	varSlice := make([]Var, dim)
	for eltIndex := 0; eltIndex < dim; eltIndex++ {
		varSlice[eltIndex] = m.AddVar()
	}
	return VarVector{varSlice}
}

/*
AddVarVectorClassic
Description:

	The classic version of AddVarVector defined in the original goop.
*/
func (m *Model) AddVarVectorClassic(
	num int, lower, upper float64, vtype VarType,
) VarVector {
	stID := uint64(len(m.Variables))
	vs := make([]Var, num)
	for i := range vs {
		vs[i] = Var{stID + uint64(i), lower, upper, vtype}
	}

	m.Variables = append(m.Variables, vs...)
	return VarVector{vs}
}

// AddBinaryVarVector adds a vector of binary variables to the model and
// returns the slice.
func (m *Model) AddBinaryVarVector(num int) VarVector {
	return m.AddVarVectorClassic(num, 0, 1, Binary)
}

// AddVarMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
func (m *Model) AddVarMatrix(
	rows, cols int, lower, upper float64, vtype VarType,
) [][]Var {
	vs := make([][]Var, rows)
	for i := range vs {
		tempVV := m.AddVarVectorClassic(cols, lower, upper, vtype)
		vs[i] = tempVV.Elements
	}

	return vs
}

// AddBinaryVarMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (m *Model) AddBinaryVarMatrix(rows, cols int) [][]Var {
	return m.AddVarMatrix(rows, cols, 0, 1, Binary)
}

// AddConstr adds a the given constraint to the model.
func (m *Model) AddConstr(constr ScalarConstraint, extras ...interface{}) {
	// Constants
	nExtraArguments := len(extras)

	// Input Processing
	if nExtraArguments > 1 {
		// Do nothing, but report an error.
		logrus.Error(
			fmt.Sprintf("The optimizer tried to add a constraint using a bad call to AddConstr! Skipping this constraint: %v , because of extra inputs %v", constr, extras),
		)
		return
	}

	optionalErrorArgument := extras[0]
	switch optionalErrorArgument.(type) {
	case error:
		// Cast argument
		err, _ := optionalErrorArgument.(error)
		if err != nil {
			logrus.Error(
				fmt.Sprintf("There was an error computing constraint %v: %v", constr, err),
			)
			return
		}
	default:
		logrus.Info(
			fmt.Sprintf("Unexpected input to AddConstr %v of type %T.", optionalErrorArgument, optionalErrorArgument),
		)
		return
	}

	// Algorithm
	m.constrs = append(m.constrs, constr)
}

// SetObjective sets the objective of the model given an expression and
// objective sense.
func (m *Model) SetObjective(e ScalarExpression, sense ObjSense) {
	m.obj = NewObjective(e, sense)
}

// Optimize optimizes the model using the given solver type and returns the
// solution or an error.
func (m *Model) Optimize(solver Solver) (*Solution, error) {
	// Variables
	var err error

	// Input Processing
	if len(m.Variables) == 0 {
		return nil, errors.New("no variables in model")
	}

	// lbs := make([]float64, len(m.Variables))
	// ubs := make([]float64, len(m.Variables))
	// types := new(bytes.Buffer)
	// for i, v := range m.Variables {
	// 	lbs[i] = v.Lower
	// 	ubs[i] = v.Upper
	// 	types.WriteByte(byte(v.Vtype))
	// }

	solver.ShowLog(m.showLog)

	if m.timeLimit > 0 {
		solver.SetTimeLimit(m.timeLimit.Seconds())
	}

	solver.AddVars(m.Variables)

	for _, constr := range m.constrs {
		solver.AddConstr(constr)
	}

	if m.obj != nil {
		logrus.WithField(
			"num_vars", m.obj.NumVars(),
		).Info("Number of variables in objective")
		err = solver.SetObjective(*m.obj)
		if err != nil {
			return nil, fmt.Errorf("There was an error setting the objective: %v", err)
		}
	}

	mipSol, err := solver.Optimize()
	defer solver.DeleteSolver()

	if mipSol.Status != OptimizationStatus_OPTIMAL {
		errorMessage, err := mipSol.Status.ToMessage()
		if err != nil {
			return nil, fmt.Errorf("There was an issue converting optimization status to a message: %v", err)
		}
		return nil, fmt.Errorf(
			"[Code = %d] %s",
			mipSol.Status,
			errorMessage,
		)
	}

	return &mipSol, nil
}
