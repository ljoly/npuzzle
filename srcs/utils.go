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

func getPossibilities(e Env, state [][]int, x int, y int) int {
	possibilities := 2
	if (x == 0 || x == e.boardSize-1) && y < e.boardSize-1 && y > 0 {
		possibilities = 3
	} else if x > 0 && x < e.boardSize-1 && (y == 0 || y == e.boardSize-1) {
		possibilities = 3
	} else {
		possibilities = 4
	}
	return possibilities
}

func getIndexToMove(e Env, state [][]int) (int, int) {
	for y := 0; y < e.boardSize; y++ {
		for x := 0; x < e.boardSize; x++ {
			if state[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func getFinalState(e *Env) {
	var cursor = 1
	var x = 0
	var ix = 1
	var y = 0
	var iy = 0

	e.finalState = make([][]int, e.boardSize)
	for i := 0; i < len(e.finalState); i++ {
		e.finalState[i] = make([]int, e.boardSize)
		for j := 0; j < len(e.finalState); j++ {
			e.finalState[i][j] = -1
		}
	}
	for {
		e.finalState[y][x] = cursor
		if cursor == 0 {
			break
		}
		cursor++
		if x+ix == e.boardSize || x+ix < 0 || (ix != 0 && e.finalState[y][x+ix] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == e.boardSize || y+iy < 0 || (iy != 0 && e.finalState[y+iy][x] != -1) {
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
