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

//		1
//	0	+	2
//		3

func swapTiles(tile1, tile2 int, board []int, e Env) bool {
	if tile2 >= 0 && tile2 < e.boardSize*e.boardSize {
		tmp := board[tile1]
		board[tile1] = board[tile2]
		board[tile2] = tmp
	}
	return true
}

func getXYfromIndex(index int, e Env) (int, int) {
	x := index / e.boardSize
	y := index % e.boardSize
	return x, y
}

func getNewState(e Env, index, indexToMove int, currentState *State, chanState chan<- State) {
	new := &State{
		board:    nil,
		priority: -1,
		parent:   nil,
	}
	var board []int
	var passed = false
	board = make([]int, e.boardSize*e.boardSize)
	copy(board, currentState.board)
	x, y := getXYfromIndex(indexToMove, e)

	switch {
	case index == 0 && y-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-1, board, e)
	case index == 1 && x-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-e.boardSize, board, e)
	case index == 2 && y+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+1, board, e)
	case index == 3 && x+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+e.boardSize, board, e)
	}
	if passed == false {
		chanState <- *new
	} else {
		new.board = board
		new.parent = currentState
		new.priority = heuristic(e, new)
		chanState <- *new
	}
}

func printState(e Env, state State) {
	fmt.Println("State Index : ", state.index)
	fmt.Println("State Parent : ", state.parent)
	fmt.Println("State Priority : ", state.priority)
	if state.board != nil {
		fmt.Println("State Board : ")
		for i := 0; i < e.boardSize; i++ {
			for j := 0; j < e.boardSize; j++ {
				fmt.Printf("%d\t", state.board[i*e.boardSize+j])
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("State Board : ", state.board)
	}
}

func play(e Env) {
	getFinalState(&e)
	indexToMove := getIndexToMove(e.initState)
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
		go getNewState(e, i, indexToMove, pq[0], chanState)
	}

	for i := 0; i < 4; i++ {
		state := <-chanState
		printState(e, state)
		fmt.Print("\n")
	}
}
