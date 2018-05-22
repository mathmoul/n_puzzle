package solver

import (
	"N_Puzzle/npuzzle"
	"fmt"
)

// Start function
func Start(p npuzzle.Puzzle, h uint) {
	a := NewAstar(p, h)
	fmt.Println(a)
	// TODO heuristic between answer and puzzle
	//d.Answer.PrintPuzzle()
}
