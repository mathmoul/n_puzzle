package tools

import (
	"container/list"
	"fmt"
	"reflect"
)

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func PrintAddr(i ...interface{}) {
	for _, a := range i {
		fmt.Printf("%p ", &a)
	}
	fmt.Println()
}

func CloneValues(source interface{}, dest interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(dest).Elem().Set(y.Elem())
	} else {
		reflect.ValueOf(dest).Elem().Set(x)
	}
}