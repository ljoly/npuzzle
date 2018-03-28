package main

func abs(val int) int {
	if val < 0 {
		return (val * (-1))
	}
	return (val)
}

func distance(actual, final []int, index int, e Env) int {
	var piece, xAct, yAct, xFinal, yFinal, distance int
	piece = actual[index]
	xAct = index / e.boardSize
	yAct = index % e.boardSize
	for i := 0; i < e.boardSize*e.boardSize; i++ {
		if final[i] == piece {
			xFinal = i / e.boardSize
			yFinal = i % e.boardSize
			break
		}
	}
	distance = abs(xFinal-xAct) + abs(yFinal-yAct)
	return distance
}

func linearConflict(e Env, state *State) int {
	var l int
	return l
}

func manhattanDistance(e Env, state *State) int {
	var m int
	for i := 0; i < len(state.board); i++ {
		m += distance(state.board, e.finalState, i, e)
	}
	return m
}

func heuristic(e Env, state *State) int {
	var h int
	if e.heuristic == 1 {
		h = manhattanDistance(e, state)
	} else if e.heuristic == 2 {
		h = manhattanDistance(e, state) + linearConflict(e, state)
	}
	return (h)
}
