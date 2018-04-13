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

func play(e *Env) {
	tab, size := parseFile(string(e.file))
	e.initState = tab
	e.boardSize = size
	getFinalState(e)
	if sameArrays(e.initState, e.finalState) {
		fmt.Println("Puzzle already solved")
		return
	}
	checkSolvability(*e)
	aStarSolver(e)
	printResults(*e)
}

func main() {
	e := Env{}
	parseCommand(&e)
	if *flagServer {
		launchServer(e)
	} else {
		play(&e)
	}
}
