package generator

import (
	p "N_Puzzle/puzzle"
	"fmt"
)

// Gen function
func Gen(size int, solvable bool, iterations uint) (puzzle *p.Puzzle, err error) {
	if solvable {
		fmt.Println("This puzzle is sovlable")
	} else {
		fmt.Println("This puzzle is unsolvable")
	}
	if puzzle, err = p.MakePuzzle(size, solvable, iterations); err != nil {
		return
	}
	puzzle.PrintPuzzle()
	return
}