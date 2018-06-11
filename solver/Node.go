package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
	"container/list"
	"fmt"
)

type Node struct {
	Action actions.Action
	G      int
	H      int
	Somm   int
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
		Somm:   g + h,
		Parent: parent,
		State:  state,
	}
}

func BoardsEqual(new npuzzle.Board, old npuzzle.Board) bool {
	if len(new) == len(old) {
		for i := 0; i < len(old); i++ {
			if int(new[i]) != int(old[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (n *Node) AlreadyClosed(closedList *[]Node) bool {
	for _, closedNode := range *closedList {
		if BoardsEqual(n.State.Board, closedNode.State.Board) {
			return true
		}
	}
	return false
}

func (n *Node) Execute(a *Astar) {
	for _, b := range actions.L {
		if n.State.Zero.ToTile(n.State.Size).TestAction(b.Value, n.State.Size) {
			var y = n.State.Copy()
			y.Move(b)
			h, err := a.HeuristicFunction(*y, a.Goal)
			if err != nil {
				log.Fatal(err)
			}
			newNode := NewNode(b, n.G+1, h, n, *y)
			if !newNode.AlreadyClosed(a.ClosedList) {
				OpenListLowerCost(a.OpenList, newNode)
			}
		}
	}
	return
}

func OpenListLowerCost(openList *[]Node, newNode *Node) {
	o := *openList
	for i, n := range *openList {
		if BoardsEqual(n.State.Board, newNode.State.Board) {
			if newNode.G < n.G {
				*openList = append(o[:i], o[i+1:]...)
			} else {
				return
			}
		}
	}
	*openList = append(o, *newNode)
}

func TestNodes(ol *list.List, cl *list.List) (cpt int) {
	for c := cl.Front(); c != nil; c = c.Next() {
		for o := ol.Front(); o != nil; o = o.Next() {
			if BoardsEqual(c.Value.(*Node).State.Board, o.Value.(*Node).State.Board) {
				cpt++
			}
		}
	}
	return
}

func (n *Node) PrintNode() {
	fmt.Println("Move :", n.Action.Name)
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
