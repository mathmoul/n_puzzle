package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
	"fmt"
)

type Node struct {
	Action actions.Action
	G      int
	H      int
	Parent *Node
	State  npuzzle.Puzzle
}

type INode interface {
	Execute() *Node
}

func NewNode(action actions.Action, g int, h int, parent *Node, state npuzzle.Puzzle) *Node {
	return &Node{
		Action: action,
		G:      g,
		H:      h,
		Parent: parent,
		State:  state,
	}
}

func (n *Node) Execute(a *Astar) (ret []*Node) {
	for _, b := range actions.L {
		if n.State.Zero.ToTile(n.State.Size).TestAction(b.Value, n.State.Size) {
			var y = n.State.Copy()
			y.Move(b)
			h, err := a.HeuristicFunction(*y, a.Goal)
			if err != nil {
				log.Fatal(err)
			}
			ret = append(ret, NewNode(b, n.G+1, h, n, *y))
			// TODO if newNode is in a.ClosedState dont add it dont do for greedy search
		}
	}
	return
}

func (n *Node) PrintNode() {
	fmt.Println(n.G, " ", n.H)
	n.State.PrintPuzzle()
}

func (n *Node) PrintResult() {
	for n != nil {
		//fmt.Println(y)
		defer n.PrintNode()
		n = n.Parent
	}
}
