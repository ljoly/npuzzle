package main

import "fmt"

func get_answer(size int) {

	var puzzle [][]int
	var cursor = 1
	var x = 0
	var ix = 1
	var y = 0
	var iy = 0

	puzzle = make([][]int, size)
	for i := 0; i < len(puzzle); i++ {
		puzzle[i] = make([]int, size)
		for j := 0; j < len(puzzle); j++ {
			puzzle[i][j] = -1
		}
	}

	for {
		puzzle[y][x] = cursor
		if cursor == 0 {
			break
		}
		cursor += 1
		if x+ix == size || x+ix < 0 || (ix != 0 && puzzle[y][x+ix] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && puzzle[y+iy][x] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cursor == size*size {
			cursor = 0
		}
	}
	fmt.Println(puzzle)
}
