package npuzzle

// TileIndex struct
type TileIndex struct {
	I int
}

type Board []int

// Puzzle struct
type Puzzle struct {
	Zero TileIndex
	Board
	Size int
	Tiles
}

func (b Board) Copy(i int) (Board) {
	nb := make([]int, i*i)
	if len(b) == len(nb) {
		for i, y := range b {
			nb[i] = y
		}
		return nb
	}
	return Board{}
}

func (p *Puzzle) Copy() (*Puzzle) {
	return &Puzzle{
		Zero:  p.Zero,
		Board: p.Board.Copy(p.Size),
		Size:  p.Size,
		Tiles: p.Tiles.Copy(p.Size),
	}
}

func (t Tiles) Copy(i int) Tiles {
	np := make([]Tile, i*i)
	if len(t) == len(np) {
		for i, y := range t {
			np[i] = y
		}
		return np
	}
	return Tiles{}
}
