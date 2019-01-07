package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"

	heap "github.com/theodesp/go-heaps"
	rank_paring "github.com/theodesp/go-heaps/rank_pairing"
)

// Start function init astar
func Start(p Puzzle, h uint, c uint) {
	a := NewAstar(p, h, c)
	costFunction = FindCostFunction(c)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	fmt.Println("Searching solution...")
	if n, err := a.Run(p.Size); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
		fmt.Println("Number of turns:", a.Turns)
		fmt.Println("Max state:", a.MaxState)
	}
}

const (
	//No action
	No = iota
)

// Run function Runs the astar algorithm
func (a *Astar) Run(size int /*FCost SortList */) (q *Node, err error) {
	return runN(a)
}

func run(a *Astar /* , FCost */) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	return
}

func runN(a *Astar /* , FCost */) (q *Node, err error) {
	if err = a.RootNode(No); err != nil {
		return
	}
	for a.OpenList.Size() > 0 {
		node := a.OpenList.DeleteMin()

		uuid := node.(*Node).State.CreateUuid()

		if node.(*Node).H == 0 {
			return node.(*Node), nil
		}
		a.Turns++
		node.(*Node).Execute(a)
		num := a.OpenList.Size()
		if num > int(a.MaxState) {
			a.MaxState = uint(num)
		}
		if a.ClosedList == nil {
			a.ClosedList = NewBst(BstString(uuid))
		} else {
			a.ClosedList.Insert(BstString(uuid))
		}
	}
	return
}

/*
RootNode func
*/
func (a *Astar) RootNode(action int) (err error) {
	var h int
	currentState := a.Puzzle
	h, err = a.HeuristicFunction(currentState, a.Goal)
	if err != nil {
		return
	}
	a.OpenList.Insert(NewNode(
		None,
		0,
		uint(h),
		nil,
		a.Puzzle))
	return
}

func (a *Astar) CheckSolvability() bool {
	a.Puzzle.PrintPuzzle()
	pI := a.Puzzle.Inversions()
	gI := a.Goal.Inversions()
	if a.Puzzle.Mod(2) == 0 {
		pI += a.Puzzle.Zero.I / a.Size
		gI += a.Goal.Zero.I / a.Size
	}
	return pI%2 == gI%2
}

/**
Node
*/
type Node struct {
	Action Action
	G      uint
	H      uint
	Parent *Node
	State  Puzzle
}

func (n *Node) Compare(than heap.Item) int {
	return costFunction(n, than.(*Node))
	// return int((n.G + n.H) - (than.(*Node).G + than.(*Node).H))
}

func (n *Node) Tag() interface{} {
	return n.State.CreateUuid()
}

func (n *Node) Key() float64 {
	return float64(n.G + n.H)
}

type INode interface {
	Execute() *Node
}

func NewNode(action Action, g uint, h uint, parent *Node, state Puzzle) *Node {
	return &Node{
		Action: action,
		G:      g,
		H:      h,
		Parent: parent,
		State:  state,
	}
}

func (n *Node) AlreadyClosed(closedList *Bst) bool {
	_, ok := closedList.Find(BstString(n.State.CreateUuid()))
	return ok
}

func move(action Action, state *Puzzle, astar *Astar, n *Node) chan *Node {
	tile := state.Zero.ToTile(state.Size)
	size := state.Size
	ch := make(chan *Node)
	go func() {
		if tile.TestAction(action.Value, size) {
			state.Move(action.Value)
			state.CreateUuid()
			h, err := astar.HeuristicFunction(*state, astar.Goal)
			if err != nil {
				log.Fatal(err)
			}
			ch <- NewNode(action, n.G+1, uint(h), n, *state)
		} else {
			ch <- nil
		}
		close(ch)
	}()
	return ch
}

func add(newNode *Node, a *Astar) {
	if newNode != nil {
		if !newNode.AlreadyClosed(a.ClosedList) {
			OpenListLowerCost(a.OpenList, newNode)
		}
	}
}

func (n *Node) Execute(a *Astar) {
	top, bot, left, right := <-move(L[0], n.State.Copy(), a, n), <-move(L[1], n.State.Copy(), a, n), <-move(L[2], n.State.Copy(), a, n), <-move(L[3], n.State.Copy(), a, n)
	add(top, a)
	add(bot, a)
	add(left, a)
	add(right, a)
}

func OpenListLowerCost(o *rank_paring.RPHeap, newNode *Node) {
	o.Insert(newNode)
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

/**
Heuristic
*/

type HeuristicFunction func(board Puzzle, dt Puzzle) (ret int, err error)

const (
	Manhattan = iota
	Linear
	Misplaced
	Pattern
)

func FindHeuristic(h uint) HeuristicFunction {
	fmt.Print("Chosen Heuristic function : ")
	switch h {
	case Manhattan:
		return ManhattanHeuristic()
		break
	case Linear:
		return LinearHeuristic()
		break
	case Misplaced:
		return MisplacedHeuristic()
		break
	default:
		return ManhattanHeuristic()
		break
	}
	return ManhattanHeuristic()
}

type CostFunction func(a, b *Node) int

const (
	greedy = iota
	aStar
	uniform
)

func FindCostFunction(c uint) CostFunction {
	fmt.Print("Chosen Cost function : ")
	switch c - 1 {
	case greedy:
		return greedyCost()
	case aStar:
		return astarCost()
	case uniform:
		return uniformCost()
	}
	return astarCost()
}

func greedyCost() CostFunction {
	fmt.Println("greedy cost")
	return (func(a, b *Node) int {
		return int(a.H - b.H)
	})
}

func astarCost() CostFunction {
	fmt.Println("astar cost")
	return (func(a, b *Node) int {
		return int(a.G+a.H) - int(b.G+b.H)
	})
}

func uniformCost() CostFunction {
	fmt.Println("Uniform cost")
	return (func(a, b *Node) int {
		return int(a.G) - int(b.G)
	})
}

// Add on A the solv function depends on heuristic and fill Solution number

func ManhattanHeuristic() HeuristicFunction {
	fmt.Println("Manhattan")
	return HeuristicFunction(func(board Puzzle, final Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			ret += Abs(current.X - final.X)
			ret += Abs(current.Y - final.Y)
		}
		return
	})
}

func VerticalConflict(current, final Tile) (conflicts int) {
	if current.Y == final.Y {
		if current.X != final.X {
			conflicts += Abs(current.X - final.X)
		}
	}
	return conflicts * 2
}

func HorizontalConflict(current, final Tile) (conflicts int) {
	if current.X == final.X {
		if current.Y != final.Y {
			conflicts += Abs(current.Y - final.Y)
		}
	}
	return conflicts * 2
}

func LinearHeuristic() HeuristicFunction {
	fmt.Println("Manhattan with linear conflicts")
	return HeuristicFunction(func(board Puzzle, final Puzzle) (ret int, err error) {
		ret = 0
		if len(board.Tiles) != len(final.Tiles) {
			return 0, errors.New("les tableaux de tiles ne sont pas de la meme taille")
		}
		for i := range board.Tiles {
			current := board.Tiles[i]
			final := final.Tiles[i]
			if current.X != final.X {
				ret += Abs(current.X - final.X)
			} else {
				ret += HorizontalConflict(current, final)
			}
			if current.Y != final.Y {
				ret += Abs(current.Y - final.Y)
			} else {
				ret += VerticalConflict(current, final)
			}
		}
		return
	})
}

func MisplacedHeuristic() HeuristicFunction {
	fmt.Println("Misplaced Tiles")
	return HeuristicFunction(func(board Puzzle, final Puzzle) (ret int, err error) {
		for i := range board.Tiles {
			if board.Board[i] != final.Board[i] {
				ret++
			}
		}
		return
	})
}

/**
Astar
*/

type List struct {
	Puzzle
	Next *list.List
}

var costFunction CostFunction

type Astar struct {
	Puzzle
	Goal       Puzzle
	OpenList   *rank_paring.RPHeap
	ClosedList *Bst
	Turns      uint
	MaxState   uint
	HeuristicFunction
}

type IAstar interface {
	ManhattanHeuristic() (ret int, err error)
	LinearHeuristic() (ret int, err error)
	MisplacedHeuristic() (ret int, err error)

	Run() (err error)

	RootNode(action int, parent *Node) (err error)

	PrintResult() (err error)

	S()

	Done() bool
}

func (a *Astar) Done() bool {
	return false
}

func NewAstar(p Puzzle, h, c uint) *Astar {
	return &Astar{
		Puzzle:            p,
		Goal:              Goal(p.Size),
		OpenList:          rank_paring.New().Init(),
		ClosedList:        nil,
		HeuristicFunction: FindHeuristic(h),
		Turns:             0,
		MaxState:          0,
	}
}

func (a *Astar) S() {
	fmt.Println("A* =>", a)
}
