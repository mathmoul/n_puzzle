package main

import (
	"fmt"
	"log"
)

var costFunction CostFunction

// Start function init astar
func Start(p *Puzzle, h uint, c uint) {
	a := NewAstar(p, h, c)
	costFunction = FindCostFunction(c)
	if !a.CheckSolvability() {
		log.Fatal("This puzzle is unsolvable")
	}
	if n, err := runN(a); err != nil {
		log.Fatal(err)
	} else {
		n.PrintResult()
		fmt.Println("Number of turns:", a.Turns)
		fmt.Println("Max state:", a.MaxState)
	}
}

func runN(a *Astar /* , FCost */) (q *Node, err error) {
	if err = a.RootNode(); err != nil {
		return
	}

	for a.OpenList.Size() > 0 {
		senderNodes := make(chan *Node, 1)
		// receiverNodes := make(chan *Node, 1)
		uuidChan := make(chan BstString, 1)
		stateChan := make(chan *Puzzle, 1)

		go executeWorker(senderNodes, a, uuidChan, stateChan)

		heapItem := a.OpenList.DeleteMin()
		stateChan <- decompute(heapItem.(*Node).State)
		uuidChan <- (<-stateChan).CreateUuid()

		if *heapItem.(*Node).H == 0 {
			return heapItem.(*Node), nil
		}

		(*a).Turns++

		senderNodes <- heapItem.(*Node)
		num := a.OpenList.Size()
		if num > int(a.MaxState) {
			a.MaxState = uint(num)
		}
		if a.ClosedList == nil {
			a.ClosedList = NewBst(<-uuidChan)
		} else {
			a.ClosedList.Insert(<-uuidChan)
		}
	}
	return
}

// func Execute(senderNodes <-chan *Node, receiverNodes chan<- *Node, a *Astar, uuid <-chan BstString, state <-chan *Puzzle) {
// 	id := make(chan int, len(L))
// 	nodes := make(chan *Node, len(L))
// 	for range L {
// 		go worker(id, state.Copy(), a, n, nodes)
// 	}
// 	for _, v := range L {
// 		id <- v.Value
// 	}
// 	close(id)
// 	for range L {
// 		add(<-nodes, a, uuid)
// 	}
// 	close(nodes)
// }

func move(action Action, state *Puzzle, astar *Astar, n *Node, results chan<- *Node) {
	tile := state.Zero.ToTile(state.Size)
	size := state.Size
	if tile.TestAction(action.Value, size) {
		state.Move(action.Value)
		h, err := astar.HeuristicFunction(state, astar.Goal)
		if err != nil {
			log.Fatal(err)
		}
		results <- NewNode(&action.Name, *n.G+1, uint(h), n, state)
	} else {
		results <- nil
	}
	tile = nil
}

func add(newNode *Node, a *Astar, uuid BstString) {
	if newNode != nil {
		if !newNode.AlreadyClosed(a.ClosedList, uuid) {
			a.OpenList.Insert(newNode)
		} else {
			newNode = nil
		}
	}
}

func worker(id <-chan int, puzzle *Puzzle, a *Astar, n *Node, results chan<- *Node) {
	move(L[<-id], puzzle, a, n, results)
	puzzle = nil
	a = nil
	n = nil
}
