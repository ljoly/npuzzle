package main

func abs(val int) int {
	if val < 0 {
		return (val * (-1))
	}
	return (val)
}

const (
	horizontal = iota
	vertical
)

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
	// fmt.Println("final index vertical", finalIndexOfCurrent)
	start := index
	for start >= e.boardSize {
		start -= e.boardSize
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < e.boardSize*e.boardSize; i += e.boardSize {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(e, vertical, i, currentState[i])
				// fmt.Println("final index comp", finalIndexComp, i)
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	// fmt.Println("vertical", conflict)
	return conflict
}

func horizontalConflict(e Env, currentState []int, index int) int {
	var conflict int
	finalIndexOfCurrent := getIndexInFinalRow(e, horizontal, index, currentState[index])
	// fmt.Println("final index horizontal", finalIndexOfCurrent)
	start := index
	for start%e.boardSize > 0 {
		start--
	}
	if finalIndexOfCurrent != -1 {
		for i := start; i < start+e.boardSize; i++ {
			if i != index && currentState[i] != 0 {
				finalIndexComp := getIndexInFinalRow(e, horizontal, i, currentState[i])
				// fmt.Println("final index comp", finalIndexComp, i)
				if finalIndexComp != -1 && ((index > i && finalIndexOfCurrent < finalIndexComp) || (index < i && finalIndexOfCurrent > finalIndexComp)) {
					conflict++
				}
			}
		}
	}
	// fmt.Println("horizontal", conflict)
	return conflict
}

func linearConflict(e Env, state *State) int {
	var l int
	for i := range state.board {
		// test with go routine
		if state.board[i] != 0 {
			// fmt.Println("state[", i, "] = ", state.board[i])
			l += verticalConflict(e, state.board, i)
			l += horizontalConflict(e, state.board, i)
		}
	}
	// fmt.Println("LC", l)
	return l
}

func distance(current, final []int, index int, e Env) int {
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
	return distance
}

func manhattanDistance(e Env, state *State) int {
	var m int
	for i := 0; i < len(state.board); i++ {
		// test with go routine
		m += distance(state.board, e.finalState, i, e)
	}
	// fmt.Println("MANHATTAN", m)
	return m
}

func heuristic(e Env, state *State) int {
	var h int
	if e.heuristic == 1 {
		h = manhattanDistance(e, state)
	} else if e.heuristic == 2 {
		h = manhattanDistance(e, state)
		h += linearConflict(e, state)
	}
	return (h)
}
