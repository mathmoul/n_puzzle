package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
	"container/list"
	"fmt"
)

// Start function
func Start(p npuzzle.Puzzle, h uint, c uint) {
	a := NewAstar(p, h)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	if n, err := a.Run(SortSwitch(c)); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
	}
	//d.Answer.PrintPuzzle()
}

const (
	No = iota
)

func (a *Astar) Run(FCost SortList) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	//if a.OpenList.Back().Value.(*Node).H == 0 {
	//	return nil
	//}
	for a.OpenList.Len() > 0 {
		for e := a.OpenList.Front(); e != nil; e = e.Next() {
			a.OpenList.Remove(e)
			a.Turns += 1
			c := e.Value
			c.(*Node).Execute(a)
			if uint(a.OpenList.Len()) > a.MaxState {
				a.MaxState = uint(a.OpenList.Len())
			}
			if c.(*Node).H == 0 {
				fmt.Println("turns", a.Turns)
				fmt.Println(a.MaxState)
				return c.(*Node), nil
			}
			a.ClosedList.PushBack(c)
			FCost(&a.OpenList)
		}
	}
	fmt.Println("turns", a.Turns)
	return
}

func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	fmt.Println(h)
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

func PrintListH(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("G => ", e.Value.(*Node).H)
	}
}
