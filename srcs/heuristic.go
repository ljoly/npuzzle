package main

func abs(val int) int {
	if val < 0 {
		return (val * (-1))
	}
	return (val)
}

const (
	topDir = iota
	bottomDir
	leftDir
	rightDir
)

// func isInRow(final []int, direction int, index int, val int) bool {
// 	return false
// }

// func topConflict(currentState, final []int, index int) int {
// 	var conflict int
// 	for i := index; i >= 0; i-- {
// 		if isInRow(final, topDir, i, )
// 	}
// 	return conflict
// }

// func linearConflict(e Env, state *State) int {
// 	var l int
// 	for i := range state.board {
// 		// test with go routine
// 		if state.board[i] != 0 {
// 			if i-e.boardSize >= 0 {
// 				l += topConflict(state.board, e.finalState, i)
// 			}
// 			l += bottomConflict(state.board, i)
// 			l += rightConflict(state.board, i)
// 			l += leftConflict(state.board, i)
// 		}
// 	}
// 	return l
// }

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
	return m
}

func heuristic(e Env, state *State) int {
	var h int
	if e.heuristic == 1 {
		h = manhattanDistance(e, state)
	} else if e.heuristic == 2 {
		// h = manhattanDistance(e, state) + linearConflict(e, state)
	}
	return (h)
}
