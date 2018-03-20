package main

import (
	"fmt"
	"os"
	"strconv"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func getIndexToMove(state []int) int {
	for i := 0; i < len(state); i++ {
		if state[i] == 0 {
			return i
		}
	}
	return -1
}

func getFinalState(e *Env) {
	var cursor = 1
	var x = 0
	var ix = 1
	var y = 0
	var iy = 0

	e.finalState = make([]int, e.boardSize*e.boardSize)
	for i := 0; i < len(e.finalState); i++ {
		e.finalState[i] = -1
	}
	for {
		e.finalState[y*e.boardSize+x] = cursor
		if cursor == 0 {
			break
		}
		cursor++
		if x+ix == e.boardSize || x+ix < 0 || (ix != 0 && e.finalState[y*e.boardSize+x+ix] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == e.boardSize || y+iy < 0 || (iy != 0 && e.finalState[(y+iy)*e.boardSize+x] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cursor == e.boardSize*e.boardSize {
			cursor = 0
		}
	}
}
