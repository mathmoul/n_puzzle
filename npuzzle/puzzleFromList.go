package npuzzle

import (
	"container/list"
	"github.com/go-errors/errors"
	"fmt"
	"strconv"
)

func PuzzleFromList(l *list.List) (p *Puzzle, err error) {
	p = new(Puzzle)
	i := -1
	for e := l.Front(); e != nil; e = e.Next() {
		if i == -1 {
			fmt.Println(strconv.Atoi(e.Value.(*list.List).Front().Value.(string)))
			if e.Value.(*list.List).Len() != 1 {
				return nil, errors.New("Issue with puzzle size")
			} else if s, err := strconv.Atoi(e.Value.(*list.List).Front().Value.(string)); s <= 2 && err == nil {
				return nil, errors.New(fmt.Sprintln("Size too short or negative : ", s))
			} else if err != nil {
				return nil, err
			} else {
				p.Size = s
			}
		}
		// Si le nombre est superieur ou egal a 0 inferieur a size * size, pas deja utilise ou
		i++
	}
	fmt.Println(p)
	return
}
