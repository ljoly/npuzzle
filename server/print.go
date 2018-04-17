package main

import (
	"fmt"
	"os"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func printBoard(state State) {
	if state.board != nil {
		for i := 0; i < e.boardSize; i++ {
			for j := 0; j < e.boardSize; j++ {
				fmt.Printf("%d\t", state.board[i*e.boardSize+j])
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("State Board : ", state.board)
	}
	fmt.Println("")
}

func printResults() {
	switch e.heuristic {
	case manhattan:
		fmt.Println("Heuristic:", "Manhattan Distance")
	case misplaced:
		fmt.Println("Heuristic:", "Misplaced Tiles")
	case manhattanLC:
		fmt.Println("Heuristic:", "Manhattan Distance + Linear Conflict")
	}
	if *flagGreed {
		fmt.Println("Search type: greedy")
	} else {
		fmt.Println("Search type: uniform-cost")
	}
	fmt.Println("States selected in the openList:", e.timeComplexity)
	fmt.Println("Maximum number of states in memory:", e.sizeComplexity)
	fmt.Println("Number of moves:", e.moves-1)
}

func printState(state *State) {
	fmt.Println("State", e.moves)
	printBoard(*state)
}

func printMoves(state *State) {
	if state != nil {
		printMoves(state.parent)
		printState(state)
		e.moves++
	}
}
