package npuzzle

import (
	"N_Puzzle/flags"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
)

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
		Size: size,
		Board: make([]int, size*size),
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
	p.Board[p.Zero.I] = p.Board[swi]
	p.Board[swi] = 0
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
	flags := flags.Get()
	if flags.Solvable {
		fmt.Println("This puzzle is sovlable")
	} else {
		fmt.Println("This puzzle is unsolvable")
	}
	tmp := initPuzzle(flags.Size)
	if err = tmp.makePuzzle(flags.Solvable, flags.Iterations); err != nil {
		return
	}
	if err = tmp.zeroIndex(); err != nil {
		return
	}
	p = *tmp
	return
}

func Goal(size int) Puzzle {
	tmp := initPuzzle(size)
	tmp.makeGoals()
	tmp.zeroIndex()
	return *tmp
}
