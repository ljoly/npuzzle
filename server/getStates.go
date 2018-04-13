package main

func swapTiles(tile1, tile2 int, board []int) bool {
	if tile2 >= 0 && tile2 < e.boardSize*e.boardSize {
		tmp := board[tile1]
		board[tile1] = board[tile2]
		board[tile2] = tmp
	}
	return true
}

func getXYfromIndex(index int) (int, int) {
	x := index / e.boardSize
	y := index % e.boardSize
	return x, y
}

func getNewState(index, indexToMove int, currentState State, chanState chan<- State) {
	new := &State{
		board:     nil,
		priority:  -1,
		parent:    nil,
		iteration: 0,
		heuristic: 0,
	}
	var passed = false
	board := make([]int, e.boardSize*e.boardSize)
	copy(board, currentState.board)
	x, y := getXYfromIndex(indexToMove)
	switch {
	case index == 0 && y-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-1, board)
	case index == 1 && x-1 >= 0:
		passed = swapTiles(indexToMove, indexToMove-e.boardSize, board)
	case index == 2 && y+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+1, board)
	case index == 3 && x+1 < e.boardSize:
		passed = swapTiles(indexToMove, indexToMove+e.boardSize, board)
	}
	if passed == false {
		chanState <- *new
	} else {
		new.board = board
		new.parent = &currentState
		new.iteration = currentState.iteration + 1
		new.heuristic = getHeuristic(new)
		new.priority = new.heuristic
		if *flagGreed {
			new.priority += new.iteration
		}
		chanState <- *new
	}
}
