package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getStates(bestState *State, chanState chan<- State) {
	indexToMove := getIndexToMove(bestState.board)
	for i := 0; i < 4; i++ {
		go getNewState(i, indexToMove, *bestState, chanState)
	}
}

func aStarSolver() {
	var (
		openList   PriorityQueue
		closedList PriorityQueue
	)
	chanState := make(chan State)
	bestState := &State{
		board:     e.initState,
		priority:  -1,
		index:     0,
		iteration: 0,
		parent:    nil,
		heuristic: -1,
	}
	heap.Push(&openList, bestState)
	for openList.Len() > 0 {
		//sort the open list
		sort.Sort(&openList)
		//select and remove the best state from the open list
		bestState = heap.Pop(&openList).(*State)
		//check if the puzzle is solved
		if sameArrays(bestState.board, e.finalState) /*|| bestState.heuristic == 0*/ {
			printStates(bestState)
			close(chanState)
			return
		}

		getStates(bestState, chanState)

		for i := 0; i < 4; i++ {
			childState := <-chanState
			//check if the state exists && if it is not in the closed list
			if childState.board != nil && findInList(&childState, closedList) == -1 {
				//check if the state is in the open list
				index := findInList(&childState, openList)
				if index != -1 {
					//modify priority if it is higher (== worse) in the open list
					if openList[index].priority > childState.priority {
						openList[index].priority = childState.priority
						openList[index].heuristic = childState.heuristic
						openList[index].iteration = childState.iteration
						openList[index].parent = &childState
					}
				} else {
					//push neighbour to open list
					heap.Push(&openList, &childState)
					e.timeComplexity++
					if e.sizeComplexity < len(openList) {
						e.sizeComplexity = len(openList)
					}
				}
			}
		}
		//push the best state in the closed list
		heap.Push(&closedList, bestState)
	}
	// all states were reviewed
	fmt.Println("No Answer")
}
