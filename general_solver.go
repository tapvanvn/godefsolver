package godepsolver

import "fmt"

type GeneralSolver struct {
	orideps map[string][]string       //origional define
	deps    map[string][]string       //this stuff depencence on those stuffs
	rank    map[string]map[string]int //
}

func (solver *GeneralSolver) solve() {

	solver.deps = map[string][]string{}
	solver.rank = map[string]map[string]int{}

	for from, deps := range solver.orideps {

		solver.deps[from] = make([]string, 0)
		rank := map[string]int{}

		for _, dep := range deps {
			rank[dep] = 1
		}
		for _, dep := range deps {
			solver.solveNest(dep, rank)
		}
		fmt.Printf("%s->%v\n", from, rank)
		if _, loopDetect := rank[from]; loopDetect {
			fmt.Printf("\tLoop detected\n")
		}
		solvedDeps := SortRank(from, rank)
		fmt.Printf("\tsolved:%v\n", solvedDeps)
		solver.deps[from] = solvedDeps
		solver.rank[from] = rank
	}
}

func (solver *GeneralSolver) solveNest(dep string, rank map[string]int) {
	unsolve := []string{}
	if nestDeps, hasNest := solver.orideps[dep]; hasNest {

		for _, dep := range nestDeps {

			if rankLevel, has := rank[dep]; has {
				rank[dep] = rankLevel + 1
			} else {
				rank[dep] = 1
				unsolve = append(unsolve, dep)
			}
		}
		for _, dep := range unsolve {
			solver.solveNest(dep, rank)
		}
	}
}

func (solver *GeneralSolver) SetDependencies(from string, to []string) {

	solver.orideps[from] = to
	solver.solve()
}

func NewGeneralSolver(dependencies map[string][]string) *GeneralSolver {

	solver := &GeneralSolver{

		orideps: dependencies,
		deps:    map[string][]string{},
		rank:    map[string]map[string]int{},
	}
	solver.solve()
	return solver
}

func (solver *GeneralSolver) GetDependency(from string) []string {
	if deps, has := solver.deps[from]; has {
		return deps
	}
	return []string{}
}
