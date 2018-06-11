package solver

import (
	"sort"
)

type H struct {
	Nodes
}

func (h H) Less(i, j int) bool {
	return h[i].H < h[j].H
}

type SortList func([]*Node)

func (n Nodes) Len() int               { return len(n) }
func (n Nodes) Swap(i, j int)          { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool     { return n[i].H < n[j].H }
func (n Nodes) LessG(i, j int) bool    { return n[i].G < n[j].G }
func (n Nodes) LessSomm(i, j int) bool { return n[i].Somm < n[j].Somm }

func SortH(n []*Node) {
	sort.Sort(H(Nodes(n)))
}

func SortGreedy(n []Node) {
	sort.SliceStable(n, func(i, j int) bool {
		return n[i].Somm < n[j].Somm
	})
}

func SortUniform(n []Node) {
	sort.SliceStable(n, func(i, j int) bool {
		return n[i].G < n[j].G
	})
}

func SortSwitch(cost uint) SortList {
	switch cost {
	case 1:
		return SortH
	case 2:
	default:
		return SortGreedy
	case 3:
		return SortUniform
	}
	return SortGreedy
}
