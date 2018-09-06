package solver

import (
	"N_Puzzle/actions"
	"N_Puzzle/npuzzle"
	"container/list"
	"fmt"
	"log"
	"sync"
)

type Node struct {
	Action actions.Action
	G      uint
	H      uint
	Somm   uint
	Parent *Node
	State  npuzzle.Puzzle
}

type INode interface {
	Execute() *Node
}

func NewNode(action actions.Action, g uint, h uint, parent *Node, state npuzzle.Puzzle) *Node {
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

func (n *Node) AlreadyClosed(closedList *Nodes) bool {
	for _, closedNode := range *closedList {
		if BoardsEqual(n.State.Board, closedNode.State.Board) {
			return true
		}
	}
	return false
}

func (n *Node) Execute(a *Astar) {
	ch := make(chan *Node, 4)
	var wg sync.WaitGroup
	for _, b := range actions.L {
		if n.State.Zero.ToTile(n.State.Size).TestAction(b.Value, n.State.Size) {
			wg.Add(1)
			go func(wg *sync.WaitGroup, state *npuzzle.Puzzle, ch chan *Node, b actions.Action, a *Astar) {
				wg.Done()
				state.Move(b)
				h, err := a.HeuristicFunction(*state, a.Goal)
				if err != nil {
					log.Fatal(err)
				}
				ch <- NewNode(b, n.G+1, uint(h), n, *state)
			}(&wg, n.State.Copy(), ch, b, a)
			newNode := <-ch
			if !newNode.AlreadyClosed(&a.ClosedList) {
				OpenListLowerCost(&a.OpenList, newNode)
			}
		}
	}
	wg.Wait()
	return
}

func OpenListLowerCost(openList *Nodes, newNode *Node) {
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
	*openList = append(o, newNode)
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
