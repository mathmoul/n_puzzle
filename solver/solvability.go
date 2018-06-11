package solver

func (a *Astar) CheckSolvability() bool {
	a.Puzzle.PrintPuzzle()
	pI := a.Puzzle.Inversions()
	gI := a.Goal.Inversions()
	if a.Puzzle.Mod(2) == 0 {
		pI += a.Puzzle.Zero.I / a.Size
		gI += a.Goal.Zero.I / a.Size
	}
	return pI%2 == gI%2
}
