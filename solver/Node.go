package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"log"
	"fmt"
	"container/list"
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

func (n *Node) AlreadyClosed(closedList *list.List) bool {
	for e := closedList.Front(); e != nil; e = e.Next() {
		if BoardsEqual(n.State.Board, e.Value.(*Node).State.Board) {
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

func OpenListLowerCost(l *list.List, newnode *Node) {
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Node)
		if BoardsEqual(v.State.Board, newnode.State.Board) {
			if newnode.G < v.G {
				l.Remove(e)
			} else {
				return
			}
		}
	}
	l.PushBack(newnode)
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
	fmt.Println(n.G, " ", n.H)
	n.State.PrintPuzzle()
}

func (n *Node) PrintResult() {
	if n != nil {
		n.Parent.PrintResult()
		n.PrintNode()
	}
}
