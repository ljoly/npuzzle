package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	for i, val := range tab {
		for j := i + 1; j < len(tab); j++ {
			index := getIndexInFinalState(e.finalState, tab[j])
			if val != 0 && e.finalState[index] != 0 && tab[i] > tab[j] {
				inversions++
			}
		}
	}
	return inversions
}

func checkSolvability(e Env) {
	var startInversions = countInversions(e.initState, e)
	var goalInversions = countInversions(e.finalState, e)

	if e.boardSize%2 == 0 {
		startInversions += getIndexToMove(e.initState) / e.boardSize
		goalInversions += getIndexToMove(e.finalState) / e.boardSize
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

var (
	flagGreed *bool
)

func parseCommand(e *Env) {
	flagFile := flag.String("f", "", "File containing the puzzle to solve")
	flagH := flag.String("heuristic", "MLC", "Heuristics: Manhattan (\"M\"), Misplaced Tiles (\"MT\") or Manhattan + Linear Conflict (\"MLC\")")
	flagGreed = flag.Bool("greedy", false, "Use greedy search (if not true, uniform-cost search is used")
	flag.Parse()

	file, err := ioutil.ReadFile(*flagFile)
	if err != nil {
		msg := "Not a valid file or flag -f is missing"
		printError(msg)
	}
	e.file = string(file)

	// default heuristic: ManhattanDistance + LinearConflict
	e.heuristic = manhattanLC
	switch {
	case *flagH == "M":
		e.heuristic = manhattan
	case *flagH == "MT":
		e.heuristic = misplaced
	case *flagH == "MLC":
		e.heuristic = manhattanLC
	}
}
