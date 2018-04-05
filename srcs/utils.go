package main

import (
	"strconv"
)

func abs(val int) int {
	if val < 0 {
		return (val * (-1))
	}
	return (val)
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func sameArrays(tab1, tab2 []int) bool {
	var count int
	for i := range tab1 {
		if tab1[i] == tab2[i] {
			count++
		}
	}
	if count == len(tab2) {
		return true
	}
	return false
}

func findInList(state *State, queue PriorityQueue) int {
	for i := range queue {
		if sameArrays(queue[i].board, state.board) {
			return i
		}
	}
	return -1
}

func getIndexInFinalState(state []int, toFind int) int {
	for i, value := range state {
		if value == toFind {
			return i
		}
	}
	return -1
}

func getIndexToMove(state []int) int {
	for i := 0; i < len(state); i++ {
		if state[i] == 0 {
			return i
		}
	}
	return -1
}
