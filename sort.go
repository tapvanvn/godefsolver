package godepsolver

import (
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func SortRank(from string, rank map[string]int) []string {
	p := make(PairList, len(rank))

	i := 0
	for k, v := range rank {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	deps := []string{}
	for _, k := range p {
		if k.Key == from {
			continue
		}
		deps = append(deps, k.Key)
	}
	return deps
}
