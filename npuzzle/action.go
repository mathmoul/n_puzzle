package npuzzle

import (
	"N_Puzzle/actions"
)

func (p *Puzzle) Move(action actions.Action) (a *Puzzle) {
	a = p
	switch action.Value {
	case actions.Top:
		a.MoveTop()
		break
	case actions.Bot:
		a.MoveBot()
		break
	case actions.Left:
		a.MoveLeft()
		break
	case actions.Right:
		a.MoveRight()
		break
	}

	a.zeroIndex()
	a.TabTiles()
	return
}

func (p *Puzzle) MoveTop() {
	tmp := p.Board[p.Zero.I-3]
	p.Board[p.Zero.I-3] = 0
	p.Board[p.Zero.I] = tmp
}

func (p *Puzzle) MoveBot() {
	tmp := p.Board[p.Zero.I+3]
	p.Board[p.Zero.I+3] = 0
	p.Board[p.Zero.I] = tmp
}

func (p *Puzzle) MoveLeft() {
	tmp := p.Board[p.Zero.I-1]
	p.Board[p.Zero.I-1] = 0
	p.Board[p.Zero.I] = tmp
}

func (p *Puzzle) MoveRight() {
	tmp := p.Board[p.Zero.I+1]
	p.Board[p.Zero.I+1] = 0
	p.Board[p.Zero.I] = tmp
}
