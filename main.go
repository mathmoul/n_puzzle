package main

import (
	"N_Puzzle/generator"
	"N_Puzzle/solver"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
)

// T var for simple boolean tab
var T = [2]bool{
	true,
	false,
}

// Flags function
func Flags(i *int, solv *bool, iter *uint) ([]string, error) {
	var unsolv bool
	flag.IntVar(i, "size", 3, "Size of the puzzle's side. Must be >3.")
	flag.BoolVar(solv, "s", false, "Forces generation of a solvable puzzle. Overrides -u.")
	flag.BoolVar(&unsolv, "u", false, "Forces generation of an unsolvable puzzle.")
	flag.UintVar(iter, "iterations", 10000, "Number of iterations.")
	flag.Parse()
	if *solv && unsolv {
		return nil, errors.New("can't be both solvable and unsolvable")
	}
	if !*solv && !unsolv {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(T))))
		if err != nil {
			return nil, err
		}
		if T[r.Int64()] {
			*solv = true
		}
	}
	if *i < 3 {
		return nil, errors.New("size cant be lower than 3")
	}
	return flag.Args(), nil
}

func main() {
	var size int
	var solv bool
	var iter uint
	args, err := Flags(&size, &solv, &iter)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("args =>", args)
	// file reader
	// if not => generate a puzzle
	p, err := generator.Gen(size, solv, iter)
	if err != nil {
		log.Fatal(err)
	}
	solver.Start(p)
	//fmt.Println(p)
}
