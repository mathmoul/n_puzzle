package main

import (
	"fmt"
)

/* execute Worker */
func executeWorker(senderNodes <-chan *Node /*, receiverNodes chan<- *Node */, a *Astar, uuid <-chan BstString, state <-chan *Puzzle) {
	fmt.Println("coucou")
	id := make(chan int, len(L))
	nodes := make(chan *Node, len(L))
	for range L {
		go worker(id, (<-state).Copy(), a, <-senderNodes, nodes)
	}
	for _, v := range L {
		id <- v.Value
	}
	close(id)
	for range L {
		add(<-nodes, a, <-uuid)
	}
	close(nodes)
}
