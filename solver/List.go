package solver

// type SortList func(Nodes)

// func (n OpenList) Len() int           { return len(n) }
// func (n OpenList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
// func (n OpenList) Less(i, j int) bool { return j < i }

// func SortH(n Nodes) {
// 	sort.SliceStable(n, func(i, j string) bool {
// 		return n[i].H < n[j].H
// 	})
// }

// func SortGreedy(n Nodes) {
// 	sort.Sort(Nodes(n))
// }

// func SortUniform(n Nodes) {
// 	sort.SliceStable(n, func(i, j int) bool {
// 		return n[i].G < n[j].G
// 	})
// }

// func SortSwitch(cost uint) SortList {
// 	switch cost {
// 	case 1:
// 		fmt.Println("Sort by Heuristic")
// 		return SortH
// 	case 2:
// 		fmt.Println("Sort Greedy")
// 		return SortGreedy
// 	case 3:
// 		fmt.Println("Sort uniform")
// 		return SortUniform
// 	}
// 	return SortGreedy
// }
