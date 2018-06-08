package parser

import (
	"N_Puzzle/npuzzle"
	"github.com/go-errors/errors"
	"os"
	"bufio"
	"container/list"
	"strings"
	"fmt"
)

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
		fmt.Println(newList.Len())
		if err == nil && newList != nil && newList.Len() != 0 {
			l.PushBack(newList)
		}
	}
	//tools.PrintList(l)
	return npuzzle.PuzzleFromList(l)
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
