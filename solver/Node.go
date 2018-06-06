package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"fmt"
	"N_Puzzle/tools"
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

func (n *Node) Execute(a *Astar) (new []*Node) {
	for _, b := range actions.L {
		fmt.Println(n.State.Zero.ToTile(n.State.Size).TestAction(b.Value, n.State.Size))
		fmt.Println(n.State)
		if n.State.Zero.ToTile(n.State.Size).TestAction(b.Value, n.State.Size) {
			y := n.State
			tools.PrintAddr(y, n.State)
			y.Move(b)
			//h, err := a.HeuristicFunction(*cs, a.Goal)
			//if err != nil {
				//log.Fatal(err)
			//}
			//new = append(new, NewNode(b, n.G+1, h, n, *cs))
		}
	}
	return
}
