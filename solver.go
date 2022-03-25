package goop2

/*
solver.go
Description:
	Defines the new interface Solver which should define
*/

type Solver interface {
	ShowLog(tf bool) error
	SetTimeLimit(timeLimit float64) error
	AddVar(varIn *Var) error
	AddVars(varSlice []*Var) error
	AddConstr(constrIn *Constr) error
	SetObjective(objectiveIn Objective) error
	Optimize() (Solution, error)
	DeleteSolver() error
}
