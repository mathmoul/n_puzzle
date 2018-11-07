package npuzzle

func (t *TileIndex) ToTile(s int) (y *Tile) {
	return &Tile{t.I % s, t.I / s}
}
