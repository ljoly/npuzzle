package main

func swapTiles(tile1, tile2 int, board []int, e Env) bool {
	if tile2 >= 0 && tile2 < e.boardSize*e.boardSize {
		tmp := board[tile1]
		board[tile1] = board[tile2]
		board[tile2] = tmp
	}
	return true
}

func getXYfromIndex(index int, e Env) (int, int) {
	x := index / e.boardSize
	y := index % e.boardSize
	return x, y
}

func getNewState(e Env, index, indexToMove int, currentState State, chanState chan<- State) {
	// fmt.Println("Start : ", currentState)
	new := &State{
		board:    nil,
		priority: -1,
		parent:   nil,
	}
	var passed = false
	board := make([]int, e.boardSize*e.boardSize)
	copy(board, currentState.board)
	x, y := getXYfromIndex(indexToMove, e)

	// fmt.Println(indexToMove, x, y)
	switch {
	case index == 0 && y-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-1, board, e)
	case index == 1 && x-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-e.boardSize, board, e)
	case index == 2 && y+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+1, board, e)
	case index == 3 && x+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+e.boardSize, board, e)
	}
	if passed == false {
		chanState <- *new
	} else {
		new.board = board
		new.parent = &currentState
		new.iteration = currentState.iteration + 1
		new.heuristic = heuristic(e, new)
		new.priority = new.heuristic + new.iteration
		chanState <- *new
	}
	// fmt.Println("End : ", currentState)
}
