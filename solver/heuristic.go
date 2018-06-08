package solver

import (
	"N_Puzzle/npuzzle"
	"N_Puzzle/tools"
	"github.com/go-errors/errors"
	"fmt"
)

type HeuristicFunction func(board npuzzle.Puzzle, dt npuzzle.Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

func FindHeuristic(h uint) HeuristicFunction {
	fmt.Print("Chosen Heuristic function : ")
	switch h {
	case Manhattan:
		return ManhattanHeuristic()
		break
	case Linear:
		return LinearHeuristic()
		break
	case Misplaced:
		return MisplacedHeuristic()
		break
	default:
		return ManhattanHeuristic()
		break
	}
	return ManhattanHeuristic()
}

// Add on A the solv function depends on heuristic and fill Solution number

func ManhattanHeuristic() HeuristicFunction {
	fmt.Println("Manhattan")
	return HeuristicFunction(func(board npuzzle.Puzzle, final npuzzle.Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			ret += tools.Abs(current.X - final.X)
			ret += tools.Abs(current.Y - final.Y)
		}
		return
	})
}

func VerticalConflict(current, final npuzzle.Tile) (conflicts int) {
	if current.Y == final.Y {
		if current.X != final.X {
			conflicts += tools.Abs(current.X - final.X)
		}
	}
	return conflicts * 2
}

func HorizontalConflict(current, final npuzzle.Tile) (conflicts int) {
	if current.X == final.X {
		if current.Y != final.Y {
			conflicts += tools.Abs(current.Y - final.Y)
		}
	}
	return conflicts * 2
}

func LinearHeuristic() HeuristicFunction {
	fmt.Println("Manhattan with linear conflicts")
	return HeuristicFunction(func(board npuzzle.Puzzle, final npuzzle.Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			if current.X != final.X {
				ret += tools.Abs(current.X - final.X)
			} else {
				ret += HorizontalConflict(current, final)
			}
			if current.Y != final.Y {
				ret += tools.Abs(current.Y - final.Y)
			} else {
				ret += VerticalConflict(current, final)
			}
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	fmt.Println("Misplaced Tiles")
	return HeuristicFunction(func(board npuzzle.Puzzle, final npuzzle.Puzzle) (ret int, err error) {
		for i := range board.Tiles {
			if board.Board[i] != final.Board[i] {
				ret++
			}
		}
		return
	})
}
