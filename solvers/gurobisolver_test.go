package solvers

import (
	"fmt"
	"os"
	"testing"
)

/*
TestGurobiSolver_CreateModel1
Description:
	Tests to see if CreateModel() actually works.
*/
func TestGurobiSolver_CreateModel1(t *testing.T) {
	// Constants
	gs1 := GurobiSolver{}
	modelName1 := "Anniversary"

	// Algorithm
	gs1.CreateModel(modelName1)
	if gs1.CurrentModel == nil {
		t.Errorf("The model was not successfully created!")
	}

	defer gs1.Free()
}

/*
TestGurobiSolver_ShowLog1
Description:
	Verifies that something is plotted to the screen when running 'go test' on this test.
*/
func TestGurobiSolver_ShowLog1(t *testing.T) {
	// Constants
	gs1 := NewGurobiSolver()
	modelName1 := "Anniversary"

	// Algorithm
	gs1.CreateModel(modelName1)
	if gs1.CurrentModel == nil {
		t.Errorf("The model was not successfully created!")
	}
	defer gs1.Free()

	gs1.ShowLog(true)

	fmt.Println("Something should be plotted on screen for logging...")
}

/*
TestGurobiSolver_SetTimeLimit1
Description:
	Verifies that we can properly change the time limit when commanding.
*/
func TestGurobiSolver_SetTimeLimit1(t *testing.T) {
	// Constants
	gs1 := NewGurobiSolver()
	modelName1 := "Anniversary1"

	newTimeLimit := 1.4

	// Create Model
	gs1.CreateModel(modelName1)
	if gs1.CurrentModel == nil {
		t.Errorf("The model was not successfully created!")
	}
	defer gs1.Free()
	defer os.Remove(modelName1 + ".log")

	// Set new timelimit
	gs1.SetTimeLimit(newTimeLimit)

	// Check timelimit
	timeLimitOut, err := gs1.GetTimeLimit()
	if err != nil {
		t.Errorf("There was an issue getting the time limit variable: %v", err)
	}

	if timeLimitOut != newTimeLimit {
		t.Errorf("The time limit returned from gurobi is %v; expected %v", timeLimitOut, newTimeLimit)
	}
}

/*
TestGurobiSolver_AddVar1
Description:
	Verifies that we can properly change the time limit when commanding.
*/
func TestGurobiSolver_AddVar1(t *testing.T) {
	// Constants
	gs1 := NewGurobiSolver()
	modelName1 := "Anniversary1"

	newTimeLimit := 1.4

	// Create Model
	gs1.CreateModel(modelName1)
	if gs1.CurrentModel == nil {
		t.Errorf("The model was not successfully created!")
	}
	defer gs1.Free()
	defer os.Remove(modelName1 + ".log")

	// Set new timelimit
	gs1.SetTimeLimit(newTimeLimit)

	// Check timelimit
	timeLimitOut, err := gs1.GetTimeLimit()
	if err != nil {
		t.Errorf("There was an issue getting the time limit variable: %v", err)
	}

	if timeLimitOut != newTimeLimit {
		t.Errorf("The time limit returned from gurobi is %v; expected %v", timeLimitOut, newTimeLimit)
	}
}
