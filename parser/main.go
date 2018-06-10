package parser

import (
	"N_Puzzle/npuzzle"
	"github.com/go-errors/errors"
	"os"
	"bufio"
	"container/list"
	"strings"
	"fmt"
	"strconv"
)

type Datas struct {
	Size  int
	Board []int
}

func File(av []string) (puzzle *npuzzle.Puzzle, err error) {
	l := list.New()
	if len(av) > 1 {
		return nil, errors.New("Too much arguments")
	}
	file, err := os.Open(av[0])
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newList, err := ScanToList(scanner.Text())
		if err == nil && newList != nil && newList.Len() != 0 {
			l.PushBack(newList)
		}
	}
	d, err := DataFromList(l)
	if err != nil {
		return nil, err
	}
	return npuzzle.PuzzleFromDatas(d.Size, d.Board)
}

func ScanToList(text string) (*list.List, error) {
	l := list.New()
	tab := strings.Fields(text)
	for i := range tab {
		if rune(tab[i][0]) == rune('#') {
			return l, nil
		} else {
			l.PushBack(tab[i])
		}
	}
	return l, nil
}

func (d *Datas) ListCheckSize(l *list.List) (error) {
	if l.Len() != 1 {
		return errors.New("Issue with puzzle size")
	} else if s, err := strconv.Atoi(l.Front().Value.(string)); s <= 2 && err == nil {
		return errors.New(fmt.Sprintln("Size too short or negative : ", s))
	} else if err != nil {
		return err
	} else {
		d.Size = s
		d.Board = make([]int, s*s)
		for i := range d.Board {
			d.Board[i] = -1
		}
	}
	return nil
}

func DataFromList(l *list.List) (d *Datas, err error) {
	d = new(Datas)
	i := -1
	for e := l.Front(); e != nil; e = e.Next() {
		if i == -1 {
			if err = d.ListCheckSize(e.Value.(*list.List)); err != nil {
				return
			}
		} else {
			if l.Len() -1 > d.Size {
				return nil, errors.New("Too much lanes for board")
			}
			if e.Value.(*list.List).Len() != d.Size {
				return nil, errors.New(fmt.Sprintln("Issue with size for lane ", i+1))
			}
			var v int
			for ee := e.Value.(*list.List).Front(); ee != nil; ee = ee.Next() {
				v, err = strconv.Atoi(ee.Value.(string))
				if err != nil {
					return
				}
				if err = CheckNumberIntoBoard(v, d.Size, d.Board); err != nil {
					return nil, err
				}
				d.Board[i] = v
				i++
			}
			i--
		}
		// Si le nombre est superieur ou egal a 0 inferieur a size * size, pas deja utilise ou
		i++
	}
	return
}
