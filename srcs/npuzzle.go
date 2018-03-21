package main

import (
	"container/heap"
	"sort"
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

func play(e Env) *State {
	getFinalState(&e)
	indexToMove := getIndexToMove(e.initState)
	openList := initList(e)
	closedList := initList(e)
	chanState := make(chan State)
	for i := 0; i < 5; i++ {
		for i := 0; i < 4; i++ {
			go getNewState(e, i, indexToMove, openList[0], chanState)
		}
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
						openList[index].parent = &ngbState
					}
				} else {
					//push neighbour to open list
					heap.Push(&openList, &ngbState)
					// fmt.Println("ngbState: ", ngbState)
				}
			}
		}
		if len(openList) > 0 {
			//sort the open list
			// for i := 0; i < len(openList); i++ {
			// 	fmt.Println("BEFORE_SORT: ", openList[i].priority)
			// }
			sort.Sort(&openList)
			// for i := 0; i < len(openList); i++ {
			// 	fmt.Println("AFTER_SORT: ", openList[i].priority)
			// }
			bestState := openList[0]
			//push the best state in the closed list
			heap.Push(&closedList, bestState)
			//remove the best state from the open list
			heap.Remove(&openList, 0)
			// for i := 0; i < len(closedList); i++ {
			// 	fmt.Println(closedList[i])
			// }
			//check if the puzzle is solved
			printState(e, *bestState)
			if bestState.priority == 0 {
				return bestState
			}
		}
	}
	// all states were reviewed
	return nil
}
