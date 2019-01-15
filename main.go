/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mmoullec <mmoullec@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/01/15 17:01:00 by mmoullec          #+#    #+#             */
/*   Updated: 2019/01/15 17:01:05 by mmoullec         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"crypto/rand"
	"flag"
	"log"
	"math/big"
)

func main() {
	var p *Puzzle
	f, err := Parse()
	if err != nil {
		log.Fatal(err)
	}
	if len(f.Args) == 0 {
		p, err = Generate()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		p, err = File(f.Args)
		if err != nil {
			log.Fatal(err)
		}
	}
	Start(p, f.Heuristic-1, f.Cost)
}

type Flags struct {
	Size       int
	Solvable   bool
	Iterations uint
	Cost       uint
	Args       []string
	Heuristic  uint
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
	flag.IntVar(&f.Size, "size", 3, "Size of the puzzle's side. Must be > 3.")
	//flag.BoolVar(&f.Solvable, "s", true, "Forces generation of a solvable puzzle. Overrides -u.")
	flag.BoolVar(&unsolv, "u", false, "Forces generation of an unsolvable puzzle.\n(default: random solvable or unsolvable puzzle)")
	flag.UintVar(&f.Iterations, "iterations", 10000, "Number of iterations.")
	flag.UintVar(&f.Heuristic, "heu", 1,
		"Forces heuristic, must be between 1 to 3\n\t1 = mahnattan \n\t2 = linear \n\t3 = missplaced \n")
	flag.UintVar(&f.Cost, "c", 2, "Choose cost, must be between 1 to 3\n\t1 = Greedy Search (Only Heuristic) (faster)\n\t2 = Astar (average)\n\t3 = Uniform search (slower)\n")
	flag.Parse()
	f.Args = flag.Args()
	f.Solvable = !unsolv
	if f.Solvable {
		r, err = rand.Int(rand.Reader, big.NewInt(int64(len(t))))
		if err != nil {
			return
		}
		f.Solvable = t[r.Int64()]
	}
	if f.Heuristic < 1 || f.Heuristic > 3 {
		log.Fatal("Wrong heuristic")
	}
	if f.Cost < 1 || f.Cost > 3 {
		log.Fatal("Wrong cost")
	}
	if f.Size < 3 {
		log.Fatal("Size cant be lower than 3")
	}
	global = f
	return
}

// Get flags
func Get() Flags {
	return global
}
