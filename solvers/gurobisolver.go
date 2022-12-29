package solvers

import (
	"fmt"
	"github.com/kwesiRutledge/goop2/optim"
	"io"
	"log"
	"os"

	gurobi "github.com/kwesiRutledge/gurobi.go/gurobi"
)

// Type Definition

type GurobiSolver struct {
	Env                    *gurobi.Env
	CurrentModel           *gurobi.Model
	ModelName              string
	GoopIDToGurobiIndexMap map[uint64]int32 // Maps each Goop ID (uint64) to the idx value used for each Gurobi variable.
}

// Function

/*
NewGurobiSolver
Description:

	Create a new gurobi solver object.
*/
func NewGurobiSolver() *GurobiSolver {
	// Constants
	modelName := "goopModel"

	// Algorithm

	newGS := GurobiSolver{}
	newGS.CreateModel(modelName)

	return &newGS

}

/*
ShowLog
Description:

	Decides whether or not to print logs to the terminal?
*/
func (gs *GurobiSolver) ShowLog(tf bool) error {
	// Constants
	logFileName := gs.ModelName + ".txt"

	// Check to see if logFile exists
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		//Do Nothing. The later lines will create the file.
	} else {
		//Delete the old file.
		err = os.Remove(logFileName)
		if err != nil {
			return fmt.Errorf("There was an issue deleting the old log file: %v", err)
		}
	}

	// Create Logging file
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// log.Fatal(err)
		return fmt.Errorf("There was an issue createing a log file: %v", err)
	}

	// Attach logger to terminal only if tf is true
	if tf {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.SetOutput(file)
	}

	// Create initial file
	log.Println("Log file created.")

	return nil

}

/*
SetTimeLimit
Description:

	Sets the time limit of the current model in gurobi solver gs.

Input:

	limitInS = Value of time limit in seconds (float)
*/
func (gs *GurobiSolver) SetTimeLimit(limitInS float64) error {

	err := gs.Env.SetDBLParam("TimeLimit", limitInS)
	if err != nil {
		return fmt.Errorf("There was an issue using SetDBLParam(): %v", err)
	}

	// If there was no error, return nil
	return nil
}

/*
GetTimeLimit
Description:

	Gets the time limit of the current model in gurobi solver gs.

Input:

	None

Output

	limitInS = Value of time limit in seconds (float)
*/
func (gs *GurobiSolver) GetTimeLimit() (float64, error) {

	limitOut, err := gs.Env.GetDBLParam("TimeLimit")
	if err != nil {
		return -1, fmt.Errorf("There was an error getting the double param TimeLimit: %v", err)
	}

	// If all things succeeded, return good data.
	return limitOut, err
}

/*
CreateModel
Description:
*/
func (gs *GurobiSolver) CreateModel(modelName string) {
	// Constants

	// Algorithm
	env, err := gurobi.NewEnv(modelName + ".log")
	if err != nil {
		panic(err.Error())
	}

	gs.Env = env

	// Create an empty model.
	model, err := gurobi.NewModel(modelName, env)
	if err != nil {
		panic(err.Error())
	}
	gs.CurrentModel = model

	// Create an empty map
	gs.GoopIDToGurobiIndexMap = make(map[uint64]int32)

}

/*
FreeEnv
Description:

	Frees the Env() method. Useful after the problem is solved.
*/
func (gs *GurobiSolver) FreeEnv() {
	gs.Env.Free()
}

/*
FreeModel
Description

	Frees the Model member. Useful after the problem is solved.
*/
func (gs *GurobiSolver) FreeModel() {
	gs.CurrentModel.Free()
}

/*
Free
Description:

	Frees the Env and Model elements of the system.
*/
func (gs *GurobiSolver) Free() {
	gs.FreeModel()
	gs.FreeEnv()
}

/*
AddVar
Description:

	Adds a variable to the Gurobi Model.
*/
func (gs *GurobiSolver) AddVar(varIn optim.Var) error {
	// Constants

	// Convert Variable Type
	vType, err := VarTypeToGRBVType(varIn.Vtype)
	if err != nil {
		return fmt.Errorf("There was an error defining gurobi type: %v", err)
	}

	// Add Variable to Current Model
	_, err = gs.CurrentModel.AddVar(int8(vType), 0.0, varIn.Lower, varIn.Upper, fmt.Sprintf("x%v", varIn.ID), []*gurobi.Constr{}, []float64{})

	fmt.Printf("%v: L=%v, U=%v, name=%v\n", int8(vType), varIn.Lower, varIn.Upper, fmt.Sprintf("x%v", varIn.ID))

	// Update Map from GoopID to Gurobi Idx
	gs.GoopIDToGurobiIndexMap[varIn.ID] = int32(len(gs.CurrentModel.Variables) - 1)

	return err
}

/*
AddVars
Description:

	Adds a set of variables to the Gurobi Model.
*/
func (gs *GurobiSolver) AddVars(varSliceIn []optim.Var) error {
	// Constants

	// Iterate through ALL variable address in varSliceIn
	for _, tempVar := range varSliceIn {
		err := gs.AddVar(tempVar)
		if err != nil {
			// Terminate early.
			return fmt.Errorf("Error in AddVar(): %v", err)
		}
	}

	// If we successfully made it through all Var objects, then return no errors.
	return nil
}

/*
AddConstraint
Description:

	Adds a single constraint to the gurobi model object inside of the current GurobiSolver object.
*/
func (gs *GurobiSolver) AddConstr(constrIn optim.ScalarConstraint) error {
	// Constants

	// Identify the variables in the left hand side of this constraint
	var tempVarSlice []*gurobi.Var
	for _, tempGoopID := range constrIn.LeftHandSide.IDs() {
		tempGurobiIdx := gs.GoopIDToGurobiIndexMap[tempGoopID]

		// Locate the gurobi variable in the current model that has matching ID
		for _, tempGurobiVar := range gs.CurrentModel.Variables {
			if tempGurobiIdx == tempGurobiVar.Index {
				tempVarSlice = append(tempVarSlice, &tempGurobiVar)
			}
		}
	}

	// Call Gurobi library's AddConstr() function
	_, err := gs.CurrentModel.AddConstr(
		tempVarSlice,
		constrIn.LeftHandSide.Coeffs(),
		int8(constrIn.Sense),
		constrIn.RightHandSide.Constant(),
		fmt.Sprintf("goop Constraint #%v", len(gs.CurrentModel.Constraints)),
	)
	if err != nil {
		return fmt.Errorf("There was an issue with adding the constraint to the gurobi model: %v", err)
	}

	// Create no errors if there were no errors!
	return nil
}

/*
SetObjective
Description:

	This algorithm should set the objective based on the value of the expression provided as input to this function.
*/
func (gs *GurobiSolver) SetObjective(objIn optim.Objective) error {

	objExpression := objIn.ScalarExpression

	// Handle this differently for different types of expression inputs
	switch objExpression.(type) {
	case *optim.ScalarLinearExpr:
		gurobiLE := &gurobi.LinExpr{}
		for varIndex, goopIndex := range objExpression.IDs() {
			gurobiIndex := gs.GoopIDToGurobiIndexMap[goopIndex]

			// Add each linear term to the expression.
			tempGurobiVar := gurobi.Var{
				Model: gs.CurrentModel,
				Index: gurobiIndex,
			}
			gurobiLE = gurobiLE.AddTerm(&tempGurobiVar, objExpression.Coeffs()[varIndex])
		}

		// Add a constant term to the expression
		gurobiLE = gurobiLE.AddConstant(objExpression.Constant())

		fmt.Println(gurobiLE)

		// Add linear expression to the objective.
		err := gs.CurrentModel.SetLinearObjective(gurobiLE, int32(objIn.Sense))
		if err != nil {
			return fmt.Errorf("There was an issue setting the linear objective with SetLinearObjective(): %v", err)
		}

		return nil

	case *optim.QuadraticExpr:
		objExpressionAsQE := objExpression.(*optim.QuadraticExpr)
		gurobiQE := &gurobi.QuadExpr{}

		// Create quadratic part of quadratic expression
		for varIndex1, goopIndex1 := range objExpression.IDs() {
			gurobiIndex1 := gs.GoopIDToGurobiIndexMap[goopIndex1]

			for varIndex2, goopIndex2 := range objExpression.IDs() {
				gurobiIndex2 := gs.GoopIDToGurobiIndexMap[goopIndex2]

				// Add each linear term to the expression.
				tempGurobiVar1 := gurobi.Var{
					Model: gs.CurrentModel,
					Index: gurobiIndex1,
				}
				tempGurobiVar2 := gurobi.Var{
					Model: gs.CurrentModel,
					Index: gurobiIndex2,
				}

				gurobiQE = gurobiQE.AddQTerm(&tempGurobiVar1, &tempGurobiVar2, objExpressionAsQE.Q[varIndex1][varIndex2])
			}
		}

		// Create linear part of quadratic expression
		for varIndex, goopIndex := range objExpression.IDs() {
			gurobiIndex := gs.GoopIDToGurobiIndexMap[goopIndex]

			// Add each linear term to the expression.
			tempGurobiVar := gurobi.Var{
				Model: gs.CurrentModel,
				Index: gurobiIndex,
			}
			gurobiQE = gurobiQE.AddTerm(&tempGurobiVar, objExpressionAsQE.L[varIndex])
		}

		// Create offset
		gurobiQE = gurobiQE.AddConstant(objExpressionAsQE.C)

		// Return
		fmt.Println(gurobiQE)

		err := gs.CurrentModel.SetQuadraticObjective(gurobiQE, int32(objIn.Sense))
		if err != nil {
			return fmt.Errorf("There was an issue setting the quadratic objective with SetQuadraticObjective(): %v", err)
		}

		return nil

	default:
		return fmt.Errorf("Unexpected objective type given to gurobisolver's SetObjective(): %T", objExpression)
	}
}

/*
Optimize
Description:
*/
func (gs *GurobiSolver) Optimize() (optim.Solution, error) {
	// Make sure that all changes are applied to the given model.
	err := gs.CurrentModel.Update()
	if err != nil {
		return optim.Solution{}, fmt.Errorf("There was an issue updating the current gurobi model: %v", err)
	}

	// Optimize
	err = gs.CurrentModel.Optimize()
	if err != nil {
		return optim.Solution{}, fmt.Errorf("There was an issue optimizing the current model: %v", err)
	}

	// Construct solution:
	// - Status
	tempSolution := optim.Solution{}
	tempStatus, err := gs.CurrentModel.GetIntAttr("Status")
	if err != nil {
		return tempSolution, fmt.Errorf("There was an issue collecting the model's status: %v", err)
	}
	tempSolution.Status = optim.OptimizationStatus(tempStatus)

	// - Values
	tempValues := make(map[uint64]float64)
	for _, tempGurobiVar := range gs.CurrentModel.Variables {
		val, err := tempGurobiVar.GetDouble("X")
		if err != nil {
			return tempSolution, fmt.Errorf("Error while retrieving the optimal values of the problem: %v", err)
		}
		// identify goop index that has this gurobi variables data
		for goopIndex, gurobiIndex := range gs.GoopIDToGurobiIndexMap {
			if gurobiIndex == tempGurobiVar.Index {
				tempValues[goopIndex] = val
				break // When you find it, save the value and return the value to the map.
			}
		}
	}
	tempSolution.Values = tempValues

	// - Objective
	tempObjective, err := gs.CurrentModel.GetDoubleAttr("ObjVal")
	if err != nil {
		return tempSolution, fmt.Errorf("There was an issue getting the objective value of the current model.")
	}
	tempSolution.Objective = tempObjective

	// All steps were successful, return solution!
	return tempSolution, nil
}

/*
DeleteSolver
Description:

	Attempts to delete all info about the current solver.
*/
func (gs *GurobiSolver) DeleteSolver() error {
	// Free model and environment
	gs.CurrentModel.Free()

	gs.Env.Free()

	return nil
}
