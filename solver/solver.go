package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
)

// Start function
func Start(p npuzzle.Puzzle, h uint) {
	a := NewAstar(p, h)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
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

		e := a.OpenList.Back()
		c := e.Value
		_ = c.(*Node).Execute(a)

		//for _, n := range nodes {
		//	a.OpenList.PushBack(n)
		//}
		//a.ClosedList.PushBack(c)
		//a.OpenList.Remove(e)
		//tools.PrintList(a.OpenList)
		//fmt.Println()
		//tools.PrintList(a.ClosedList)
		//if a.Done() {
		//	return nil
		//}

		if a.Turns > 0 {
			return
		}
	}
	return
}

func (a *Astar) PrintResult() (err error) {
	return
}

func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	if err != nil {
		return
	}
	a.OpenList.PushBack(NewNode(
		actions.None,
		0,
		h,
		nil,
		a.Puzzle))
	return
}
