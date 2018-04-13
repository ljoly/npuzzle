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

var e Env

func play() {
	tab, size := parseFile(string(e.file))
	e.initState = tab
	e.boardSize = size
	getFinalState()
	if sameArrays(e.initState, e.finalState) {
		fmt.Println("Puzzle already solved")
		return
	}
	checkSolvability()
	aStarSolver()
	printResults()
}

func main() {
	parseCommand()
	if *flagServer {
		launchServer()
	} else {
		play()
	}
}
