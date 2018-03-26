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

func manhattan(e Env, state *State) int {
	var heuristic int
	for i := 0; i < len(state.board); i++ {
		heuristic += distance(state.board, e.finalState, i, e)
	}
	return heuristic
}

func heuristic(e Env, state *State) int {
	if e.heuristic == 1 {
		return (manhattan(e, state))
	}
	if e.heuristic == 1 {
		return (-1)
	}
	return (-1)
}
