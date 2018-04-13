package main

import (
	"fmt"
	"os"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func printBoard(e Env, state State) {
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

func printStates(e *Env, state *State) {
	if state != nil {
		printStates(e, state.parent)
		fmt.Println("State", e.moves)
		printBoard(*e, *state)
		e.moves++
	}
}

func printResults(e Env) {
	fmt.Println("Heuristic:", e.heuristic)
	fmt.Println("States selected in the openList:", e.timeComplexity)
	fmt.Println("Maximum number of states in memory:", e.sizeComplexity)
	fmt.Println("Number of moves:", e.moves-1)
	fmt.Println("greedySearch:", *flagGreed)
}
