package main

func abs(val int) int {
	if val < 0 {
		return (val * (-1))
	}
	return (val)
}

func distance(actual, final []int, index int, e Env) int {
	var piece int
	var x_act, y_act, x_final, y_final, distance int
	piece = actual[index]
	x_act = index / e.boardSize
	y_act = index % e.boardSize
	for i := 0; i < e.boardSize*e.boardSize; i++ {
		if final[i] == piece {
			x_final = i / e.boardSize
			y_final = i % e.boardSize
			break
		}
	}
	distance = abs(x_final-x_act) + abs(y_final-y_act)
	return distance
}

func manhattan(e Env, state *State) int {
	var heristic int = 0
	for i := 0; i < len(state.board); i++ {
		heristic += distance(state.board, e.finalState, i, e)
	}
	return heristic
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
