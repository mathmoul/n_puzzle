package npuzzle

func PuzzleFromDatas(size int, board []int) (p *Puzzle, err error) {
	p = initPuzzle(size)
	p.Board = board
	p.zeroIndex()
	p.TabTiles()
	return
}
