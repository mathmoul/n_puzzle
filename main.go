package main

import (
	"N_Puzzle/flags"
	"N_Puzzle/npuzzle"
	"log"
)

func main() {
	var p npuzzle.Puzzle
	flags, err := flags.Parse()
	if err != nil {
		log.Fatal(err)
	}
	if len(flags.Args) == 0 {
		p, err = npuzzle.Generate()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		p, err = npuzzle.ParseArgs(flags.Args)
		if err != nil {
			log.Fatal(err)
		}
	}
	p.PrintPuzzle()
}
