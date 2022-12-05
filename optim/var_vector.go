package optim

/*
var_vector.go
Description:
	The VarVector type will represent a
*/

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type VarVector struct {
	Elements []Var
}

// =========
// Functions
// =========

/*
Length
Description:

	Returns the length of the vector of optimization variables.
*/
func (v VarVector) Length() int {
	return len(v.Elements)
}

/*
Len
Description:

	This function is created to mirror the GoNum Vector API. Does the same thing as Length.
*/
func (v VarVector) Len() int {
	return v.Length()
}

func (v VarVector) At(x int) Var {
	return v.Elements[x]
}
