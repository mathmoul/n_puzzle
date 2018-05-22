package solver

import "N_Puzzle/npuzzle"

type List struct {
}

type Astar struct {
	npuzzle.Puzzle
	Goal       npuzzle.Puzzle
	OpenList   []List
	ClosedList []List
}


func NewAstar(p npuzzle.Puzzle) (*Astar) {
	return &Astar{
		Puzzle: p,
		Goal: npuzzle.Goal(p.Size),
		OpenList: []List{},
		ClosedList: []List{},
	}
}

type IAstar interface {

}