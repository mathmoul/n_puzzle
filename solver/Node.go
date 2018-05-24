package solver

import "N_Puzzle/npuzzle"

type Node struct {
	Action int
	G      int
	H      int
	Parent *Node
	State  npuzzle.Puzzle
}

type INode interface {
	DoSomething()
}

func (n *Node) DoSomething() {

}

func NewNode(action int, g int, h int, parent *Node, state npuzzle.Puzzle) *Node {
	return &Node{
		Action: action,
		G:      g,
		H:      h,
		Parent: parent,
		State:  state,
	}
}

func (n *Node) Copy() (nn *Node) {
	nn = n
	return
}
