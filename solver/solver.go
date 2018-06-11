package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
	"fmt"
)

// Start function
func Start(p npuzzle.Puzzle, h uint, c uint) {
	a := NewAstar(p, h)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	fmt.Println("Searching solution...")
	if n, err := a.Run(SortSwitch(c)); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
		fmt.Println("Number of turns:", a.Turns)
		fmt.Println("Max state:", a.MaxState)
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
	for len(a.OpenList) > 0 {
		for i, n := range a.OpenList {
			a.OpenList = append(a.OpenList[:i], a.OpenList[i+1:]...)
			a.Turns += 1
			n.Execute(a)
			if uint(len(a.OpenList)) > a.MaxState {
				a.MaxState = uint(len(a.OpenList))
			}
			if n.H == 0 {
				return n, nil
			}
			a.ClosedList = append(a.ClosedList, n)
			FCost(a.OpenList)
		}
	}
	fmt.Println("turns", a.Turns)
	return
}

func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	if err != nil {
		return
	}
	a.OpenList = append(a.OpenList, *NewNode(
		actions.None,
		0,
		h,
		nil,
		a.Puzzle))
	return
}

func printNodeSlice(nodes []Node) {
	for _, n := range nodes {
		fmt.Println("h->", n.H, "g->", n.G, "somm", n.H+n.G, "|", n.Somm)
	}
}
