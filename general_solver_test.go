package godepsolver_test

import (
	"godepsolver"
	"testing"
)

//A->AB,AC,AD
//AB->BC,CD
//BC->CD,AC

func TestGeneralSolver(t *testing.T) {

	deps := map[string][]string{

		"A":  {"AB", "AC", "AD"},
		"AB": {"BC", "CD"},
		"BC": {"CD", "AC", "A"},
	}

	solver := godepsolver.NewGeneralSolver(deps)
	_ = solver
}
