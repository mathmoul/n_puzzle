package puzzle

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// Puzzle struct
type Puzzle struct {
	Case []int
	Size int
}

func (p *Puzzle) zeroIndex() int {
	for i := range p.Case {
		if p.Case[i] == 0 {
			return i
		}
	}
	return -1
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
	idx := p.zeroIndex()
	if idx < 0 {
		os.Exit(0)
	}
	poss := make([]int, 0)
	if (idx % p.Size) > 0 {
		poss = append(poss, idx-1)
	}
	if (idx % p.Size) < (p.Size - 1) {
		poss = append(poss, idx+1)
	}
	if (idx / p.Size) > 0 {
		poss = append(poss, idx-p.Size)
	}
	if (idx / p.Size) < (p.Size - 1) {
		poss = append(poss, idx+p.Size)
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(poss))))
	if err != nil {
		log.Fatal(err)
	}
	swi := poss[n.Int64()]
	p.Case[idx] = p.Case[swi]
	p.Case[swi] = 0
	return
}

// PrintPuzzle function
func (p *Puzzle) PrintPuzzle() {
	u := p.Case
	for y := 0; y < p.Size; y++ {
		for x := 0; x < p.Size; x++ {
			if u[y+x*p.Size] == 0 {
				color.New(color.FgRed).Printf("|%*d| ", len(strconv.Itoa(p.Size*p.Size))+1, u[x+y*p.Size])
			} else {
				fmt.Printf("|%*d| ", len(strconv.Itoa(p.Size*p.Size))+1, u[x+y*p.Size])
			}
		}
		fmt.Printf("\n")
	}
}
