package npuzzle

const (
	No = iota
)

func (p *Puzzle) Action(action int) func() {
	switch action {
	default:
		return DoNoting
	}
}

func DoNoting() {

}
