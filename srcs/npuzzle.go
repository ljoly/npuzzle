package main

import "fmt"

var (
	finalState [][]int
)

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

func play(e Env) {
	getFinalState(&e)
	for i := 0; i < e.boardSize; i++ {
		fmt.Println(e.finalState[i])
	}
}
