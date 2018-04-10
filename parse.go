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

// func countInversions(tab []int, e Env) int {
// 	var inversions int

// 	for i, val := range tab {
// 		for j := i + 1; j < len(tab); j++ {
// 			index := getIndexInFinalState(e.finalState, tab[j])
// 			if val != 0 && e.finalState[index] != 0 && arr[i] > arr[j] {
// 				inversions++
// 			}
// 		}
// 	}
// 	return inversions
// }

func countInversions(arr []int, e Env) int {
	var inversion int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				inversion++
			}
		}
	}
	return inversion
}

func checkSolvability(e Env) {
	var startInversions = countInversions(e.initState, e)
	var goalInversions = countInversions(e.finalState, e)

	fmt.Println(startInversions, goalInversions)
	if e.boardSize%2 == 0 {
		_, y := getXYfromIndex(getIndexToMove(e.initState), e)
		startInversions += y / e.boardSize
		goalInversions += y / e.boardSize
	}
	if startInversions%2 != goalInversions%2 {
		printError("Puzzle is not solvable")
	}
}

// func isSolvable(arr []int, e Env) {
//     inversion := getInversions(arr);

//     if (N & 1)
//         return !(inversion & 1);

//     else     // grid is even
//     {
//         int pos = findXPosition(puzzle);
//         if (pos & 1)
//             return !(inversion & 1);
//         else
//             return inversion & 1;
//     }
// }

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

func parseCommand(e *Env) {
	flagFile := flag.String("f", "", "")
	flagH := flag.String("heuristic", "", "")
	flagGreed := flag.Bool("greedy", false, "")
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

	if *flagGreed {
		e.greedySearch = true
	}
}
