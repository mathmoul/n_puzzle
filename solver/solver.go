package solver

import (
	"N_Puzzle/npuzzle"
)

// Start function
func Start(p npuzzle.Puzzle, h uint) {
	a := NewAstar(p, h)
	a.Run()
	// TODO heuristic between answer and puzzle
	//d.Answer.PrintPuzzle()
}

const (
	No = iota
)

func (a *Astar) Run() (err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	for a.OpenList.Len() > 0 {
		a.Turns += 1

	}
	a.S()
	return
}

func (a *Astar) PrintResult() (err error) {
	return
}

func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentstate := a.Puzzle.Board
	h, err = a.HeuristicFunction(currentstate, a.Puzzle)
	if err != nil {
		return
	}
	a.OpenList.PushBack(NewNode(
		action,
		0,
		h,
		nil,
		a.Puzzle))
	return
}
