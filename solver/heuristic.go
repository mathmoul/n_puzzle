package solver

import (
	"N_Puzzle/npuzzle"
	"N_Puzzle/tools"
	"fmt"

	"github.com/go-errors/errors"
)

type HeuristicFunction func(board npuzzle.Puzzle, dt npuzzle.Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

func FindHeuristic(h uint) HeuristicFunction {
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

func LinearHeuristic() HeuristicFunction {
	return HeuristicFunction(func(board npuzzle.Puzzle, dt npuzzle.Puzzle) (ret int, err error) {
		for i, b := range board.Tiles {
			fmt.Println(i, b)
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	return HeuristicFunction(func(board npuzzle.Puzzle, dt npuzzle.Puzzle) (ret int, err error) {
		for i, b := range board.Tiles {
			fmt.Println(i, b)
		}
		return
	})
}
