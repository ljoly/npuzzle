package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getStates(bestState *State, e *Env, chanState chan<- State) {
	indexToMove := getIndexToMove(bestState.board)
	for i := 0; i < 4; i++ {
		go getNewState(*e, i, indexToMove, *bestState, chanState)
	}
}

func play(e *Env) {
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
		if sameArrays(bestState.board, e.finalState) || bestState.heuristic == 0 {
			// e.moves = len(closedList)
			fmt.Println("Puzzle solved", bestState.board)
			return
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
}
