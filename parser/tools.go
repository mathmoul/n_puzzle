package parser

import (
	"github.com/go-errors/errors"
	"fmt"
)

var NumberErr = [2]func(i int) error{
	func(i int) error {
		return errors.New(fmt.Sprintln("Number too low or too high :", i))
	},
	func(i int) error {
		return errors.New(fmt.Sprintln("Number already exists in Board (twice or more):", i))
	},
}

func CheckNumberIntoBoard(n, size int, board []int) error {
	if n < 0 || n >= size*size {
		return NumberErr[0](n)
	}
	cpt := 0
	for _, b := range board {
		if b == n {
			cpt++
		}
	}
	if cpt > 0 {
		return NumberErr[1](n)
	}
	return nil
}
