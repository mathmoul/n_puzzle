package solver

import "container/list"

type SortList func(l **list.List)

func SortH(l **list.List) {
	ll := *l
	for e := ll.Front(); e != nil; e = e.Next() {
		if e.Next() != nil {
			if e.Value.(*Node).H > e.Next().Value.(*Node).H {
				ll.MoveAfter(e, e.Next())
				SortH(l)
			}
		}
	}
}

func SortGreedy(l **list.List) {
	ll := *l
	for e := ll.Front(); e != nil; e = e.Next() {
		if e.Next() != nil {
			if e.Value.(*Node).H+e.Value.(*Node).G >
				e.Next().Value.(*Node).H+e.Next().Value.(*Node).G {
				ll.MoveAfter(e, e.Next())
				SortGreedy(l)
			}
		}
	}
}

func SortUniform(l **list.List) {
	ll := *l
	for e := ll.Front(); e != nil; e = e.Next() {
		if e.Next() != nil {
			if e.Value.(*Node).G > e.Next().Value.(*Node).G {
				ll.MoveAfter(e, e.Next())
				SortUniform(l)
			}
		}
	}
}

func SortSwitch(cost uint) SortList {
	switch cost {
	case 1:
		return SortH
	case 2:
	default:
		return SortGreedy
	case 3:
		return SortUniform
	}
	return SortGreedy
}
