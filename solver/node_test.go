package solver_test

import (
	"testing"
	"N_Puzzle/solver"
	"N_Puzzle/npuzzle"
)

func TestBoardsEqual(t *testing.T) {
	b := npuzzle.Board{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if !solver.BoardsEqual(b, b) {
		t.Error("Boards equals send false")
	}
	if solver.BoardsEqual(b, npuzzle.Board{8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Error("Boards differents send true")
	}
}
