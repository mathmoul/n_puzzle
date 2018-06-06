package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
)

// Start function
func Start(p npuzzle.Puzzle, h uint) {
	a := NewAstar(p, h)
	if n, err := a.Run(); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
	}
	// TODO heuristic between answer and puzzle
	//d.Answer.PrintPuzzle()
}

const (
	No = iota
)

func (a *Astar) Run() (q *Node, err error) {
	//fmt.Printf("%+v\n", a)
	if err = a.RootNode(No); err != nil {
		return
	}
	//if a.OpenList.Back().Value.(*Node).H == 0 {
	//	return nil
	//}
	for a.OpenList.Len() > 0 {
		for e := a.OpenList.Front(); e != nil; e = e.Next() {
			a.Turns += 1
			c := e.Value
			nodes := c.(*Node).Execute(a)
			if c.(*Node).H == 0 {
				//c.(*Node).State.PrintPuzzle()
				return c.(*Node), nil
			}
			for _, n := range nodes {
				a.OpenList.PushBack(n)
			}
			a.ClosedList.PushBack(c)
			a.OpenList.Remove(e)
			//tools.PrintList(a.OpenList)
			//fmt.Println()
			//tools.PrintList(a.ClosedList)
			//if a.Done() {
			//	return nil
			//}

			//if a.Turns > 0 {
			//	return
			//}
		}
	}
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
