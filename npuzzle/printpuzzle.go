package npuzzle

import (
  "fmt"
	"strconv"

	"github.com/fatih/color"
)

// PrintPuzzle function
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