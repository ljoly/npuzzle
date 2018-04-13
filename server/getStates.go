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

func getStates(bestState *State, chanState chan<- State) {
	indexToMove := getIndexToMove(bestState.board)
	for i := 0; i < 4; i++ {
		go getNewState(i, indexToMove, *bestState, chanState)
	}
}

func getFinalState() {
	var cursor = 1
	var x = 0
	var ix = 1
	var y = 0
	var iy = 0

	e.finalState = make([]int, e.boardSize*e.boardSize)
	for i := 0; i < len(e.finalState); i++ {
		e.finalState[i] = -1
	}
	for {
		e.finalState[y*e.boardSize+x] = cursor
		if cursor == 0 {
			break
		}
		cursor++
		if x+ix == e.boardSize || x+ix < 0 || (ix != 0 && e.finalState[y*e.boardSize+x+ix] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == e.boardSize || y+iy < 0 || (iy != 0 && e.finalState[(y+iy)*e.boardSize+x] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cursor == e.boardSize*e.boardSize {
			cursor = 0
		}
	}
}
