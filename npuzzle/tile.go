package npuzzle

import (
	"N_Puzzle/actions"
	"fmt"
)

type Tile struct {
	X int
	Y int
}
type Tiles []Tile

func (t *Tile) TestAction(action int, size int) bool {
	switch action {
	case actions.Top:
		return !(t.Y-1 < 0)
		break
	case actions.Bot:
		return t.Y+1 < size
		break
	case actions.Left:
		return !(t.X-1 < 0)
		break
	case actions.Right:
		return t.X+1 < size
		break
	}
	return false
}

func (t *Tile) Bot() bool {
	fmt.Println("bot")
	return false
}

func (t *Tile) Left() bool {
	fmt.Println("Left")
	return false
}

func (t *Tile) Right() bool {
	fmt.Println("right")
	return false
}
