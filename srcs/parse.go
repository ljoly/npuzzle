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

func checkSolvability(e Env) {
	var inv_count = 0
	for i := 0; i < e.boardSize*e.boardSize; i++ {
		for j := i + 1; j < e.boardSize*e.boardSize; j++ {
			if e.initState[i] > e.initState[j] {
				inv_count++
			}
		}
	}
	if inv_count%2 == 0 && e.boardSize%2 != 0 || inv_count%2 != 0 && e.boardSize%2 == 0 {
		printError("Not solvable")
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

// /!\ careful with abnormal characters after board
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
