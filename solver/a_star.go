package solver

import (
	"N_Puzzle/npuzzle"
	"container/list"
	"fmt"

	"github.com/starwander/GoFibonacciHeap"
)

type List struct {
	npuzzle.Puzzle
	Next *list.List
}

type ClosedList map[string]*Node

type Astar struct {
	npuzzle.Puzzle
	Goal     npuzzle.Puzzle
	OpenList *fibHeap.FibHeap
	ClosedList
	Turns    uint
	MaxState uint
	HeuristicFunction
}

type IAstar interface {
	ManhattanHeuristic() (ret int, err error)
	LinearHeuristic() (ret int, err error)
	MisplacedHeuristic() (ret int, err error)

	Run() (err error)

	RootNode(action int, parent *Node) (err error)

	PrintResult() (err error)

	S()

	Done() bool
}

func (a *Astar) Done() bool {
	return false
}

func NewAstar(p npuzzle.Puzzle, h uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              npuzzle.Goal(p.Size),
		OpenList:          fibHeap.NewFibHeap(),
		ClosedList:        make(map[string]*Node),
		HeuristicFunction: FindHeuristic(h),
		Turns:             0,
		MaxState:          0,
	}
}

func (a *Astar) S() {
	fmt.Println("A* =>", a)
}
