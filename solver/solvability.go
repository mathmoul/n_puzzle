package solver

import "fmt"

func (a *Astar) CheckSolvability() bool {
	a.Puzzle.PrintPuzzle()
	i := a.Puzzle.Inversions()
	j := a.Goal.Inversions()
	fmt.Println(i, j)
	return false
}
