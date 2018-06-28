package npuzzle

import (
	"N_Puzzle/actions"
)

func (p *Puzzle) Move(action actions.Action) {
	switch action.Value {
	case actions.Top:
		p.Board.MoveTop(p.Zero.I, p.Size)
		break
	case actions.Bot:
		p.Board.MoveBot(p.Zero.I, p.Size)
		break
	case actions.Left:
		p.MoveLeft(p.Zero.I)
		break
	case actions.Right:
		p.MoveRight(p.Zero.I)
		break
	}

	p.zeroIndex()
	p.TabTiles()
}

func (b Board) MoveTop(i, size int) {
	tmp := b[i-size]
	b[i-size] = 0
	b[i] = tmp
}

func (b Board) MoveBot(i, size int) {
	tmp := b[i+size]
	b[i+size] = 0
	b[i] = tmp
}

func (b Board) MoveLeft(i int) {
	tmp := b[i-1]
	b[i-1] = 0
	b[i] = tmp
}

func (b Board) MoveRight(i int) {
	tmp := b[i+1]
	b[i+1] = 0
	b[i] = tmp
}
