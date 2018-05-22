package solver

import "N_Puzzle/npuzzle"

type List struct {
}

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

type Astar struct {
	npuzzle.Puzzle
	Goal       npuzzle.Puzzle
	OpenList   []List
	ClosedList []List
	Heuristic  func() int
}

func NewAstar(p npuzzle.Puzzle, h uint) *Astar {
	return &Astar{
		Puzzle:     p,
		Goal:       npuzzle.Goal(p.Size),
		OpenList:   []List{},
		ClosedList: []List{},
		Heuristic:  nil,
	}
}

type IAstar interface {
	ManhattanHeuristic() (ret int, err error)
	LinearHeuristic() (ret int, err error)
	MisplacedHeuristic() (ret int, err error)
}

func (a *Astar) FindHeuristic(h uint) {
	switch h {
	case Manhattan:
		a.ManhattanHeuristic()
		break
	case Linear:
		a.LinearHeuristic()
		break
	case Misplaced:
		a.MisplacedHeuristic()
		break
	}
}

func (a *Astar) ManhattanHeuristic() (ret int) {
	return
}

func (a *Astar) LinearHeuristic() (ret int) {
	return 0
}

func (a *Astar) MisplacedHeuristic() (ret int) {
	return
}
