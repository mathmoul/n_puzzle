package npuzzle

import (
	"strconv"
	"strings"
)

// TileIndex struct
type TileIndex struct {
	I int
}

type Board []int

// type Board interface {
// 	FromBits(uint64) board
// 	ToBits(board) uint64
// }

// Puzzle struct
type Puzzle struct {
	Zero TileIndex
	Board
	Size int
	Tiles
}

func (p *Puzzle) CreateUuid() string {
	b := p.Board
	tab := make([]string, p.Size*p.Size)
	for k, v := range b {
		tab[k] = strconv.Itoa(v)
	}
	return strings.Join(tab, "|")
}

func (b Board) Copy(i int) Board {
	nb := make([]int, i*i)
	if len(b) == len(nb) {
		for i, y := range b {
			nb[i] = y
		}
		return nb
	}
	return Board{}
}

func (p *Puzzle) Copy() *Puzzle {
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
