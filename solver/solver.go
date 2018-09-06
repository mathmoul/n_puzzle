package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"fmt"
	"log"
)

// Start function init astar
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
	//No action
	No = iota
)

// Run function Runs the astar algorithm
func (a *Astar) Run(FCost SortList) (q *Node, err error) {
	fmt.Println("here")

	if err = a.RootNode(No); err != nil {
		return
	}
	for len(a.OpenList) > 0 {
		n := a.OpenList[0]
		a.OpenList = append(a.OpenList[:0], a.OpenList[1:]...)
		if n.H == 0 {
			return n, nil
		}
		a.Turns++
		n.Execute(a)
		if uint(len(a.OpenList)) > a.MaxState {
			a.MaxState = uint(len(a.OpenList))
		}
		a.ClosedList = append(a.ClosedList, n)
		//a.OpenList = FasterSort(a.OpenList, FCost)
		FCost(a.OpenList)
	}
	return
}

/*
RootNode func
*/
func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	if err != nil {
		return
	}
	a.OpenList = append(a.OpenList, NewNode(
		actions.None,
		0,
		uint(h),
		nil,
		a.Puzzle))
	return
}

func printNodeSlice(nodes Nodes) {
	for _, n := range nodes {
		fmt.Println("h->", n.H, "g->", n.G, "somm", n.H+n.G, "|", n.Somm)
	}
}
