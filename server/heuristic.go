package main

const (
	horizontal = iota
	vertical
)

func misplacedTiles(state *State) int {
	var mp int
	for i := range state.board {
		if state.board[i] != 0 && i != getIndexInFinalState(e.finalState, state.board[i]) {
			mp++
		}
	}
	return mp
}

func getIndexInFinalRow(dir, index int, val int) int {
	start := index
	if dir == vertical {
		for start >= e.boardSize {
			start -= e.boardSize
		}
		for i := start; i < e.boardSize*e.boardSize; i += e.boardSize {
			if e.finalState[i] == val {
				return i
			}
		}
	} else if dir == horizontal {
		for start%e.boardSize > 0 {
			start--
		}
		for i := start; i < start+e.boardSize; i++ {
			if e.finalState[i] == val {
				return i
			}
		}
	}
	return -1
}

func verticalConflict(currentState []int, index int) int {
	var conflict int
	finalIndexOfCurrent := getIndexInFinalRow(vertical, index, currentState[index])
	start := index
	for start >= e.boardSize {
		start -= e.boardSize
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < e.boardSize*e.boardSize; i += e.boardSize {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(vertical, i, currentState[i])
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	return conflict
}

func horizontalConflict(currentState []int, index int) int {
	var conflict int
	finalIndexOfCurrent := getIndexInFinalRow(horizontal, index, currentState[index])
	start := index
	for start%e.boardSize > 0 {
		start--
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < start+e.boardSize; i++ {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(horizontal, i, currentState[i])
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	return conflict
}

func getConflicts(currentState []int, i int, chanLC chan<- int) {
	var l int
	if currentState[i] != 0 {
		l += verticalConflict(currentState, i)
		l += horizontalConflict(currentState, i)
	}
	chanLC <- l
}

func linearConflict(state *State) int {
	var l int
	chanLC := make(chan int, e.boardSize)
	for i := range state.board {
		go getConflicts(state.board, i, chanLC)
	}
	for i := 0; i < len(state.board); i++ {
		l += <-chanLC
	}
	close(chanLC)
	return l
}

func getDistance(current, final []int, index int, chanM chan<- int) {
	var piece, xCurr, yCurr, xFinal, yFinal, distance int
	piece = current[index]
	xCurr = index / e.boardSize
	yCurr = index % e.boardSize
	for i := 0; i < e.boardSize*e.boardSize; i++ {
		if final[i] == piece {
			xFinal = i / e.boardSize
			yFinal = i % e.boardSize
			break
		}
	}
	distance = abs(xFinal-xCurr) + abs(yFinal-yCurr)
	chanM <- distance
}

func manhattanDistance(board []int) int {
	var m int
	chanM := make(chan int, e.boardSize)
	for i := 0; i < len(board); i++ {
		go getDistance(board, e.finalState, i, chanM)
	}
	for i := 0; i < len(board); i++ {
		m += <-chanM
	}
	close(chanM)
	return m
}

func getHeuristic(state *State) int {
	var h int
	switch e.heuristic {
	case manhattan:
		h = manhattanDistance(state.board)
	case misplaced:
		h = misplacedTiles(state)
	case manhattanLC:
		h = manhattanDistance(state.board)
		h += linearConflict(state)
	}
	return h
}
