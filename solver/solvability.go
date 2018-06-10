package solver

import "fmt"

func (a *Astar) CheckSolvability() bool {
	i := a.Puzzle.Inversions()
	j := a.Goal.Inversions()
	fmt.Println(i, j)
	return false
}
