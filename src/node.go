package main

import (
	"fmt"

	heap "github.com/theodesp/go-heaps"
)

/**
Node
*/

type Node struct {
	Action string
	G      uint
	H      uint
	Parent *Node
	State  *Puzzle
}

func NewNode(action string, g uint, h uint, parent *Node, state *Puzzle) *Node {
	return &Node{
		Action: action,
		G:      g,
		H:      h,
		Parent: parent,
		State:  state,
	}
}

func (n *Node) Compare(than heap.Item) int {
	return costFunction(n, than.(*Node))
}

// func (n *Node) Tag() *string {
// 	return n.State.CreateUuid()
// }

type INode interface {
	Execute() *Node
}

func (n *Node) AlreadyClosed(closedList *Bst, uuid string) bool {
	_, ok := closedList.Find(BstString(uuid))
	return ok
}

func (n Node) Execute(a *Astar, uuid string) {
	id := make(chan int, len(L))
	nodes := make(chan *Node, len(L))
	for range L {
		go worker(id, n.State.Copy(), a, &n, nodes)
	}
	for _, v := range L {
		id <- v.Value
	}
	close(id)
	for range L {
		add(<-nodes, a, uuid)
	}
	close(nodes)
	id = nil
	nodes = nil
}

func (n *Node) PrintNode() {
	fmt.Println("Move :", n.Action)
	n.State.PrintPuzzle()
	fmt.Println("Cost:", n.H, "| Depth:", n.G)
	fmt.Println()
}

func (n *Node) PrintResult() {
	if n != nil {
		n.Parent.PrintResult()
		n.PrintNode()
	}
}
