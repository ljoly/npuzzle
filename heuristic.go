package main

const (
	horizontal = iota
	vertical
)

func misplacedTiles(e Env, state *State) int {
	var mp int
	for i := range state.board {
		if state.board[i] != 0 && i != getIndexInFinalState(e.finalState, state.board[i]) {
			mp++
		}
	}
	return mp
}

func getIndexInFinalRow(e Env, dir, index int, val int) int {
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

func verticalConflict(e Env, currentState []int, index int) int {
	var conflict int
	finalIndexOfCurrent := getIndexInFinalRow(e, vertical, index, currentState[index])
	start := index
	for start >= e.boardSize {
		start -= e.boardSize
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < e.boardSize*e.boardSize; i += e.boardSize {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(e, vertical, i, currentState[i])
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	return conflict
}

func horizontalConflict(e Env, currentState []int, index int) int {
	var conflict int
	finalIndexOfCurrent := getIndexInFinalRow(e, horizontal, index, currentState[index])
	start := index
	for start%e.boardSize > 0 {
		start--
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < start+e.boardSize; i++ {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(e, horizontal, i, currentState[i])
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	return conflict
}

func getConflicts(e Env, currentState []int, i int, chanLC chan<- int) {
	var l int
	if currentState[i] != 0 {
		l += verticalConflict(e, currentState, i)
		l += horizontalConflict(e, currentState, i)
	}
	chanLC <- l
}

func linearConflict(e Env, state *State) int {
	var l int
	chanLC := make(chan int, e.boardSize)
	for i := range state.board {
		go getConflicts(e, state.board, i, chanLC)
	}
	for i := 0; i < len(state.board); i++ {
		l += <-chanLC
	}
	close(chanLC)
	return l
}

func getDistance(current, final []int, index int, e Env, chanM chan<- int) {
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

func manhattanDistance(e Env, board []int) int {
	var m int
	chanM := make(chan int, e.boardSize)
	for i := 0; i < len(board); i++ {
		go getDistance(board, e.finalState, i, e, chanM)
	}
	for i := 0; i < len(board); i++ {
		m += <-chanM
	}
	close(chanM)
	return m
}

func heuristic(e Env, state *State) int {
	var h int
	switch e.heuristic {
	case manhattan:
		h = manhattanDistance(e, state.board)
	case misplaced:
		h = misplacedTiles(e, state)
	case manhattanLC:
		h = manhattanDistance(e, state.board)
		h += linearConflict(e, state)
	}
	return h
}
