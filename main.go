package main

import (
	"N_Puzzle/flags"
	"N_Puzzle/npuzzle"
	"N_Puzzle/parser"
	"N_Puzzle/solver"
	"log"
)

func main() {
	var p npuzzle.Puzzle
	f, err := flags.Parse()
	if err != nil {
		log.Fatal(err)
	}
	if len(f.Args) == 0 {
		p, err = npuzzle.Generate()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		u, err := parser.File(f.Args)
		if err != nil {
			log.Fatal(err)
		}
		p = *u
	}
	solver.Start(p, f.Heuristic-1, f.Cost)
}
