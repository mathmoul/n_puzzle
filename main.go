package main

import (
	"N_Puzzle/flags"
	"N_Puzzle/npuzzle"
	"log"
	"fmt"
	"N_Puzzle/solver"
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
		//p, err = npuzzle.ParseArgs(flags.Args)
		//if err != nil {
		//	log.Fatal(err)
		//}
		fmt.Println("TODO handle file reading")
		return
	}
	fmt.Printf("Puzzle =>")
	p.PrintPuzzle()
	fmt.Println()
	//fmt.Println(p)
	solver.Start(p, flags.Heuristic-1, flags.Cost)
}
