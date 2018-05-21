package puzzle

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/fatih/color"
)

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

// ZeroIndex func
func (p *Puzzle) ZeroIndex() (err error) {
	for i := range p.Case {
		if p.Case[i] == 0 {
			p.Zero.I = i
			return
		}
	}
	return errors.New("No tile '0'")
}

// Init create new puzzle instance
func Init(size int) *Puzzle {
	var u = Puzzle{
		Size: size,
		Case: make([]int, size*size),
	}
	for i := range u.Case {
		u.Case[i] = -1
	}
	return &u
}

func (p *Puzzle) makeGoals() {
	u := p.Case
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
	p.Case = u
}

// MakePuzzle create a new puzzle
func MakePuzzle(size int, solvable bool, iterations uint) (p *Puzzle, err error) {
	p = Init(size)
	p.makeGoals()
	for i := 0; uint(i) < iterations; i++ {
		if err = p.SwapEmpty(); err != nil {
			return
		}
	}
	if !solvable {
		if p.Case[0] == 0 || p.Case[1] == 0 {
			p.Case[len(p.Case)-1], p.Case[len(p.Case)-2] = p.Case[len(p.Case)-2], p.Case[len(p.Case)-1]
		} else {
			p.Case[0], p.Case[1] = p.Case[1], p.Case[0]
		}
	}
	return
}

// SwapEmpty function swap empty
func (p *Puzzle) SwapEmpty() (err error) {
	if err = p.ZeroIndex(); err != nil {
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
	p.Case[p.Zero.I] = p.Case[swi]
	p.Case[swi] = 0
	return
}

// PrintPuzzle function
func (p *Puzzle) PrintPuzzle() {
	u := p.Case
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
