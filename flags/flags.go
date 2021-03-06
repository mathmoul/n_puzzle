package flags

import (
	"crypto/rand"
	"errors"
	"flag"
	"log"
	"math/big"
)

// Flags Struct
type Flags struct {
	Size       int
	Solvable   bool
	Iterations uint
	Args       []string
}

var t = [2]bool{
	true,
	false,
}

var global Flags

// Parse func
func Parse() (f Flags, err error) {
	var r *big.Int
	var unsolv bool
	flag.IntVar(&f.Size, "size", 3, "Size of the puzzle's side. Must be >3.")
	flag.BoolVar(&f.Solvable, "s", true, "Forces generation of a solvable puzzle. Overrides -u.")
	flag.BoolVar(&unsolv, "u", false, "Forces generation of an unsolvable puzzle.")
	flag.UintVar(&f.Iterations, "iterations", 10000, "Number of iterations.")
	flag.Parse()
	f.Args = flag.Args()
	if f.Solvable && unsolv {
		err = errors.New("")
		return
	}
	if !f.Solvable && unsolv {
		r, err = rand.Int(rand.Reader, big.NewInt(int64(len(t))))
		if err != nil {
			return
		}
		f.Solvable = t[r.Int64()]
	}
	if f.Size < 3 {
		log.Fatal("size cant be lower than 3")
	}
	global = f
	return
}

// Get flags
func Get() Flags {
	return global
}
