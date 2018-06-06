package npuzzle

//func (p *Puzzle) Move(action actions.Action) (a *Puzzle) {
//	a = p
//	switch action.Value {
//	case actions.Top:
//		a.MoveTop()
//	case actions.Bot:
//		a.MoveBot()
//	case actions.Left:
//		a.MoveLeft()
//	case actions.Right:
//		a.MoveRight()
//	}
//	a.zeroIndex()
//	a.TabTiles()
//	return
//}
//
//func (p *Puzzle) MoveTop() {
//	tmp := *p.Board
//	tmp2 := tmp[p.Zero.I-3]
//	tmp[p.Zero.I-3] = 0
//	tmp[p.Zero.I] = tmp2
//	p.Board = &tmp
//}
//
//func (p *Puzzle) MoveBot() {
//	tmp := *p.Board
//	tmp2 := tmp[p.Zero.I+3]
//	tmp[p.Zero.I+3] = 0
//	tmp[p.Zero.I] = tmp2
//	p.Board = &tmp
//}
//
//func (p *Puzzle) MoveLeft() {
//	tmp := *p.Board
//	tmp2 := tmp[p.Zero.I-1]
//	tmp[p.Zero.I-1] = 0
//	tmp[p.Zero.I] = tmp2
//	p.Board = &tmp
//}
//
//func (p *Puzzle) MoveRight() {
//	tmp := *p.Board
//	tmp2 := tmp[p.Zero.I+1]
//	tmp[p.Zero.I+1] = 0
//	tmp[p.Zero.I] = tmp2
//	p.Board = &tmp
//}
