package main

import "fmt"

type Env struct {
	initState      []int
	finalState     []int
	file           string
	boardSize      int
	heuristic      int
	timeComplexity int
	sizeComplexity int
	moves          int
}

const (
	manhattan = iota
	misplaced
	manhattanLC
)

func main() {
	e := Env{}
	parseCommand(&e)
	if *flagServer {
		initServer()
	}
	tab, size := parseFile(string(e.file))
	e.initState = tab
	e.boardSize = size
	getFinalState(&e)
	if sameArrays(e.initState, e.finalState) {
		fmt.Println("Puzzle already solved")
		return
	}
	checkSolvability(e)
	play(&e)
	fmt.Println("Heuristic:", e.heuristic)
	fmt.Println("States selected in the openList:", e.timeComplexity)
	fmt.Println("Maximum number of states in memory:", e.sizeComplexity)
	fmt.Println("Number of moves:", e.moves-1)
	fmt.Println("greedySearch:", *flagGreed)
}
