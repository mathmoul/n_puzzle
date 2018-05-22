package solver

import "N_Puzzle/npuzzle"

func Start(p npuzzle.Puzzle /* + Heuristic */) {
	a := NewAstar(p)
	// TODO heuristic between answer and puzzle
	//d.Answer.PrintPuzzle()
}
