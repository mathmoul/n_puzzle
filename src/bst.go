package main

import "fmt"

//Bst binary search tree struct
type Bst struct {
	Item
	Left  *Bst
	Right *Bst
}

// Item interface for easiest comparaisons on BST
type Item interface {
	Compare(than Item) int
}

// BstString implements the Item interface
type BstString string

// Integer implements the Item interface
type Integer int

// Compare func compare BstString to Item of new BST
func (a BstString) Compare(b Item) int {
	s1 := a
	s2 := b.(BstString)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// Compare func compare Bst Integer to Item of new BST
func (a Integer) Compare(b Item) int {
	a1 := a
	a2 := b.(Integer)
	switch {
	case a1 > a2:
		return 1
	case a1 < a2:
		return -1
	default:
		return 0
	}
}

// NewBst returns *BST
func NewBst(item *Item) *Bst {
	return &Bst{Item: *item}
}

// Insert new `Item` on BST
func (n *Bst) Insert(data Item) error {
	if n == nil {
		return fmt.Errorf("Cannot insert Value into a Nil tree")
	}

	switch {
	case n.Compare(data) == 0:
		return nil
	case n.Compare(data) > 0:
		if n.Left == nil {
			n.Left = &Bst{Item: data}
			return nil
		}
		return n.Left.Insert(data)
	case n.Compare(data) < 0:
		if n.Right == nil {
			n.Right = &Bst{Item: data}
			return nil
		}
		return n.Right.Insert(data)
	}
	return nil
}

// Find `Item` on BST returns nil, false if can't find item
func (n *Bst) Find(data Item) (Item, bool) {
	if n == nil {
		return nil, false
	}
	switch {
	case n.Compare(data) == 0:
		return data, true
	case n.Compare(data) > 0:
		return n.Left.Find(data)
	default:
		return n.Right.Find(data)
	}
}
