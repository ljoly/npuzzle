package main

import (
	"fmt"
	"strings"
)

func getSize(line []string) (int, []int) {
	if len(line) > 1 && line[1][0] != '#' {
		printError("Error in file")
	}
	var size = atoi(line[0])
	var board = make([]int, size*size)
	return size, board
}

func countInversions(tab []int, e Env) int {
	var inversions int

	for i, val := range e.initState {
		for j := i + 1; j < len(e.initState); j++ {
			index := getIndexInFinalState(e.finalState, e.initState[j])
			if val != 0 && e.finalState[index] != 0 && val > e.finalState[index] {
				inversions++
			}
		}
	}
	return inversions
}

func checkSolvability(e Env) {
	var startInversions = countInversions(e.initState, e)
	var goalInversions = countInversions(e.finalState, e)

	if e.boardSize%2 == 0 { // In this case, the row of the '0' tile matters
		_, y := getXYfromIndex(getIndexToMove(e.initState), e)
		startInversions += y / e.boardSize
		goalInversions += y / e.boardSize
	}
	if startInversions%2 != goalInversions%2 {
		printError("Puzzle is not solvable")
	}
}

func checkBoard(board []int) {
	var zeros int
	for i := 0; i < len(board); i++ {
		if board[i] == 0 {
			zeros++
		}
	}
	if zeros > 1 {
		fmt.Println(zeros)
		printError("Letters instead of numbers")
	}
	var count, add int
	for i := 0; i < len(board); i++ {
		add += board[i]
		count += i
	}
	if add != count {
		printError("Wrong numbers")
	}
}

func parseFile(file string) ([]int, int) {
	var size, x int
	var board []int
	var lines = strings.Split(file, "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" && lines[i][0] != '#' {
			var line = strings.Fields(lines[i])
			if len(line) == 1 || size == 0 {
				size, board = getSize(line)
			} else if len(line) == size {
				for y := 0; y < size; y++ {
					board[x*size+y] = atoi(line[y])
				}
				x++
			} else if len(line) > size {
				if line[size][0] == '#' {
					for y := 0; y < size; y++ {
						board[x*size+y] = atoi(line[y])
					}
					x++
				} else {
					printError("Error in file")
				}
			} else {
				printError("Error in file")
			}
			if x >= size {
				break
			}
		}
	}
	checkBoard(board)
	return board, size
}
