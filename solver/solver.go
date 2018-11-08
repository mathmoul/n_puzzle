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
	if n, err := a.Run( /*SortSwitch(c)*/ ); err != nil {
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
func (a *Astar) Run( /*FCost SortList */ ) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	for a.OpenList.Num() > 0 {
		node := a.OpenList.ExtractMinValue()

		uuid := node.(*Node).State.CreateUuid()

		if node.(*Node).H == 0 {
			return node.(*Node), nil
		}
		a.Turns++
		node.(*Node).Execute(a)
		num := a.OpenList.Num()
		if num > a.MaxState {
			a.MaxState = num
		}
		a.ClosedList[uuid] = node.(*Node)

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
	a.OpenList.InsertValue(NewNode(
		actions.None,
		0,
		uint(h),
		nil,
		a.Puzzle))
	return
}
