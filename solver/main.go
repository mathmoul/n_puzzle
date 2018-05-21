package solver

import (
	"N_Puzzle/puzzle"
	"log"
)

type Datas struct {
	*puzzle.Puzzle
	Answer *puzzle.Puzzle
}

type List struct {
}

type Taquin struct {
}

var OpenList []List

var ClosedList []List

var TaquinInitial Taquin

var NumeroCaseVide int

func NewDatas(p *puzzle.Puzzle) (d *Datas) {
	newP, err := puzzle.MakePuzzle(p.Size, true, 0)
	if err != nil {
		log.Fatal(err)
	}
	d = &Datas{
		Puzzle: p,
		Answer: newP,
	}
	return
}

func Start(p *puzzle.Puzzle) {
	d := NewDatas(p)
	log.Println(d)
	// TODO heuristic between answer and puzzle
	//d.Answer.PrintPuzzle()
}
