package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func (p *Puzzle) Move(action int) {
	switch action {
	case Top:
		p.Board.MoveTop(p.Zero.I, p.Size)
		break
	case Bot:
		p.Board.MoveBot(p.Zero.I, p.Size)
		break
	case Left:
		p.Board.MoveLeft(p.Zero.I)
		break
	case Right:
		p.Board.MoveRight(p.Zero.I)
		break
	}

	p.zeroIndex()
	p.TabTiles()
}

func (b Board) MoveTop(i, size int) {
	tmp := b[i-size]
	b[i-size] = 0
	b[i] = tmp
}

func (b Board) MoveBot(i, size int) {
	tmp := b[i+size]
	b[i+size] = 0
	b[i] = tmp
}

func (b Board) MoveLeft(i int) {
	tmp := b[i-1]
	b[i-1] = 0
	b[i] = tmp
}

func (b Board) MoveRight(i int) {
	tmp := b[i+1]
	b[i+1] = 0
	b[i] = tmp
}

func (p *Puzzle) makeGoals() {
	u := p.Board
	s := p.Size
	cur, ix := 1, 1
	x, y, iy := 0, 0, 0
	for true {
		u[x+y*s] = cur
		if cur == 0 {
			break
		}
		cur++
		if x+ix == s || x+ix < 0 || (ix != 0 && u[x+ix+y*s] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == s || y+iy < 0 || (iy != 0 && u[x+(y+iy)*s] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cur == s*s {
			cur = 0
		}
	}
	p.Board = u
}

func initPuzzle(size int) *Puzzle {
	var u = &Puzzle{
		Size:  size,
		Board: make([]int, size*size),
		Tiles: make([]Tile, size*size),
	}
	for i := range u.Board {
		u.Board[i] = -1
	}
	return u
}

func (p *Puzzle) swapEmpty() (err error) {
	if err = p.zeroIndex(); err != nil {
		log.Fatal(err)
	}
	poss := make([]int, 0)
	if (p.Zero.I % p.Size) > 0 {
		poss = append(poss, p.Zero.I-1)
	}
	if (p.Zero.I % p.Size) < (p.Size - 1) {
		poss = append(poss, p.Zero.I+1)
	}
	if (p.Zero.I / p.Size) > 0 {
		poss = append(poss, p.Zero.I-p.Size)
	}
	if (p.Zero.I / p.Size) < (p.Size - 1) {
		poss = append(poss, p.Zero.I+p.Size)
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(poss))))
	if err != nil {
		log.Fatal(err)
	}
	swi := poss[n.Int64()]
	p.Board[p.Zero.I], p.Board[swi] = p.Board[swi], 0
	return
}

func (p *Puzzle) makePuzzle(solvable bool, iterations uint) (err error) {
	p.makeGoals()
	for i := 0; uint(i) < iterations; i++ {
		if err = p.swapEmpty(); err != nil {
			return
		}
	}
	if !solvable {
		if p.Board[0] == 0 || p.Board[1] == 0 {
			p.Board[len(p.Board)-1], p.Board[len(p.Board)-2] = p.Board[len(p.Board)-2], p.Board[len(p.Board)-1]
		} else {
			p.Board[0], p.Board[1] = p.Board[1], p.Board[0]
		}
	}
	return
}

// ZeroIndex func
func (p *Puzzle) zeroIndex() (err error) {
	for i := range p.Board {
		if p.Board[i] == 0 {
			p.Zero.I = i
			return
		}
	}
	return errors.New("No tile '0'")
}

// Generate function
func Generate() (p Puzzle, err error) {
	f := Get()
	tmp := initPuzzle(f.Size)
	if err = tmp.makePuzzle(f.Solvable, f.Iterations); err != nil {
		return
	}
	if err = tmp.zeroIndex(); err != nil {
		return
	}
	tmp.TabTiles()
	p = *tmp
	return
}

func Tiling(size int, pos int) (t Tile) {
	t.X = pos % size
	t.Y = pos / size
	return
}

func (p *Puzzle) TabTiles() {
	for i := 0; i < p.Size*p.Size; i++ {
		p.Tiles[p.Board[i]] = Tiling(p.Size, i)
	}
}

func Goal(size int) Puzzle {
	tmp := initPuzzle(size)
	tmp.makeGoals()
	tmp.zeroIndex()
	tmp.TabTiles()
	return *tmp
}

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

func (p *Puzzle) CreateUuid() *string {
	b := p.Board
	tab := make([]string, p.Size*p.Size)
	for k, v := range b {
		tab[k] = strconv.Itoa(v)
	}
	var s *string
	*s = strings.Join(tab, "|")
	return s
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

func (p *Puzzle) PrintPuzzle() {
	u := p.Board
	for y := 0; y < p.Size; y++ {
		for x := 0; x < p.Size; x++ {
			if u[x+y*p.Size] == 0 {
				color.New(color.FgRed).Printf("|%*d| ", len(strconv.Itoa(p.Size*p.Size))+1, u[x+y*p.Size])
			} else {
				fmt.Printf("|%*d| ", len(strconv.Itoa(p.Size*p.Size))+1, u[x+y*p.Size])
			}
		}
		fmt.Printf("\n")
	}
}

func PuzzleFromDatas(size int, board []int) (p *Puzzle, err error) {
	p = initPuzzle(size)
	p.Board = board
	p.zeroIndex()
	p.TabTiles()
	return
}

type Tile struct {
	X int
	Y int
}

type Tiles []Tile

func (t *Tile) TestAction(action int, size int) bool {
	switch action {
	case Top:
		return !(t.Y-1 < 0)
	case Bot:
		return t.Y+1 < size
	case Left:
		return !(t.X-1 < 0)
	case Right:
		return t.X+1 < size
	}
	return false
}

func (t *Tile) Bot() bool {
	fmt.Println("bot")
	return false
}

func (t *Tile) Left() bool {
	fmt.Println("Left")
	return false
}

func (t *Tile) Right() bool {
	fmt.Println("right")
	return false
}

func (t *TileIndex) ToTile(s int) (y *Tile) {
	return &Tile{t.I % s, t.I / s}
}
