package solver

import (
	"N_Puzzle/npuzzle"
	"fmt"
)

type HeuristicFunction func(board npuzzle.Board, dt npuzzle.Puzzle) (ret int, err error)

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
	return HeuristicFunction(func(board npuzzle.Board, dt npuzzle.Puzzle) (ret int, err error) {
		for i, b := range board {
			fmt.Println(i, b)
		}
		return 0, nil
	})
}

func LinearHeuristic() HeuristicFunction {
	return HeuristicFunction(func(board npuzzle.Board, dt npuzzle.Puzzle) (ret int, err error) {
		for i, b := range board {
			fmt.Println(i, b)
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	return HeuristicFunction(func(board npuzzle.Board, dt npuzzle.Puzzle) (ret int, err error) {
		for i, b := range board {
			fmt.Println(i, b)
		}
		return
	})
}
