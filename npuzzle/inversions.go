package npuzzle

func (p *Puzzle) Inversions() (inversions int) {
	for i := range p.Board {
		inversions += inversion(p.Board, i)
	}
	return
}

func inversion(b Board, i int) (inversions int) {
	if b[i] == 0 {
		return 0
	}
	slice := b[i:]
	n := b[i]
	for r := range slice {
		if slice[r] == 0 {
			continue
		}
		if n > slice[r] {
			inversions++
		}
	}
	return inversions
}

func (p *Puzzle) Mod(i int) int {
	return p.Size % i
}
