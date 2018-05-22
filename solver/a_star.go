package solver

import "N_Puzzle/npuzzle"

type List struct {
}

type Astar struct {
	*npuzzle.Puzzle
	Goal       npuzzle.Puzzle
	OpenList   []List
	ClosedList []List
}
