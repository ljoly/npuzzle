package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type State struct {
	board    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index     int // The index of the item in the heap.
	iteration int
	heuristic int
	parent    *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want the lowest priority so we use smaller than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
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

func getStates(bestState *State, e *Env, chanState chan<- State) {
	indexToMove := getIndexToMove(bestState.board)
	for i := 0; i < 4; i++ {
		go getNewState(*e, i, indexToMove, *bestState, chanState)
	}
}

func play(e *Env) *State {
	var (
		openList   PriorityQueue
		closedList PriorityQueue
	)
	getFinalState(e)
	if sameArrays(e.initState, e.finalState) {
		return initList(*e)[0]
	}
	// openList := initList(*e)
	// closedList := initList(*e)
	chanState := make(chan State)
	bestState := &State{board: e.initState, priority: -1, index: 0, iteration: 0, parent: nil, heuristic: -1}
	// bestState.heuristic = heuristic(*e, bestState)
	// bestState.priority = bestState.heuristic
	heap.Push(&openList, bestState)
	// for openList.Len() > 0 {
	for openList.Len() > 0 {
		//sort the open list
		sort.Sort(&openList)
		//remove the best state from the open list
		bestState = heap.Pop(&openList).(*State)
		if sameArrays(bestState.board, e.finalState) || bestState.heuristic == 0 {
			// e.moves = len(closedList)
			return bestState
		}

		go getStates(bestState, e, chanState)

		for i := 0; i < 4; i++ {
			ngbState := <-chanState
			//check if the state exists && if it is not in the closed list
			if ngbState.board != nil && findInList(&ngbState, closedList) == -1 {
				//check if the state is in the open list
				index := findInList(&ngbState, openList)
				if index != -1 {
					//modify priority if it is higher (== worse) in the open list
					if openList[index].priority > ngbState.priority {
						openList[index].priority = ngbState.priority
						openList[index].heuristic = ngbState.heuristic
						openList[index].iteration = ngbState.iteration
						openList[index].parent = &ngbState
					}
				} else {
					//push neighbour to open list
					heap.Push(&openList, &ngbState)
				}
			}
		}
		//push the best state in the closed list
		heap.Push(&closedList, bestState)
		//check if the puzzle is solved

	}
	// all states were reviewed
	fmt.Println("No Answer")
	return nil
}
