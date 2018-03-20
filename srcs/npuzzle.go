package main

import (
	"container/heap"
	"fmt"
)

type State struct {
	board    []int
	heur     int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index  int // The index of the item in the heap.
	parent *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	fmt.Println(x)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.index = -1 // for safety
	*pq = old[0 : n-1]
	return state
}

func play(e Env) {
	getFinalState(&e)
	for i := 0; i < e.boardSize; i++ {
		fmt.Println(e.finalState[i])
	}
	x, y := getIndexToMove(e, e.initState)
	// possibilities := getPossibilities(e, e.initState, x, y)
	fmt.Println(x, y)

	pq := make(PriorityQueue, 1)
	pq[0] = &State{
		heur:     1,
		priority: 1,
		// index:    0,
		parent: nil,
	}
	heap.Init(&pq)
	new := &State{
		heur:     2,
		priority: 2,
		// index:    0,
		board:  e.initState,
		parent: nil,
	}
	// fmt.Println(pq[0])
	heap.Push(&pq, new)
	heuristic := heuristic(e, new)
	fmt.Println(heuristic)
}
