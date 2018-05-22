package npuzzle

// TileIndex struct
type TileIndex struct {
	I int
}

// Puzzle struct
type Puzzle struct {
	Zero TileIndex
	Case []int
	Size int
}
