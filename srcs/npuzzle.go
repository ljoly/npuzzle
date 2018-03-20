package main

import (
	"container/heap"
	"fmt"
)

type State struct {
	board    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index  int // The index of the item in the heap.
	parent *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use smaller than here.
	return pq[i].priority < pq[j].priority
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

func getNewState(index, indexToMove int, currentState *State, chanState chan<- State) {
	// calculer nouveaux states
}

func play(e Env) {
	getFinalState(&e)
	indexToMove := getIndexToMove(e.initState)
	fmt.Println(indexToMove)
	pq := make(PriorityQueue, 1)
	pq[0] = &State{
		board:    e.initState,
		priority: -1,
		parent:   nil,
	}
	heap.Init(&pq)
	// new := &State{
	// 	board:    e.finalState,
	// 	priority: 0,
	// 	parent:   nil,
	// }
	// heap.Push(&pq, new)
	// heuristic := heuristic(e, new)
	// fmt.Println(heuristic)

	chanState := make(chan State)
	for i := 0; i < 4; i++ {
		go getNewState(i, indexToMove, pq[0], chanState)
	}

	for i := 0; i < 4; i++ {
		state := <-chanState
		fmt.Println(state)
		// mettre dans openlist
	}
}
