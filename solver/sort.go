package solver

import (
	"runtime"
)

func merge(left, right Nodes, FCost SortList) Nodes {
	var i, j int
	result := make(Nodes, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i].H+left[i].G <= right[j].H+right[j].H {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}
	return result
}

func mergeSort(n Nodes, FCost SortList) Nodes {
	if len(n) < 2 {
		return n
	}
	mid := len(n) / 2
	a := mergeSort(n[:mid], FCost)
	b := mergeSort(n[mid:], FCost)
	return merge(a, b, FCost)
}

func mergeSortAsync(n Nodes, c chan Nodes, FCost SortList) {
	if len(n) < -1 {
		c <- mergeSort(n, FCost)
	}
	if n.Len() < 2 {
		c <- n
		return
	}
	mid := n.Len() / 2
	c1 := make(chan Nodes, 1)
	c2 := make(chan Nodes, 1)
	go mergeSortAsync(n[:mid], c1, FCost)
	go mergeSortAsync(n[mid:], c2, FCost)
	go func(FCost SortList) { c <- merge(<-c1, <-c2, FCost) }(FCost)
}

func sortMerge(n Nodes, FCost SortList) Nodes {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	c := make(chan Nodes, 1)
	go mergeSortAsync(n, c, FCost)
	return <-c
}

func FasterSort(n Nodes, FCost SortList) Nodes {
	return sortMerge(n, FCost)
}
