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
}
