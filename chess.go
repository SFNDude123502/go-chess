package main

import "fmt"

func makeBoard() {
	board = make([][]*piece, 8)
	for i := range make([]int, 8) {
		board[i] = make([]*piece, 8)
	}

	setSide(true)
	setSide(false)
}

func newGame() {
	makeBoard()
	printBoard()
	var coords [][]int
	var winner string
	for {
		if kingInCheck(turn) {
			posMoves := getAllMoves(turn)
			useMoves := tryAllMoves(turn, posMoves)
			//coords = getCheckedInput(turn, useMoves)
			if len(useMoves) == 0 {
				break
			}

		} else {
			//coords = handleInput(turn)
		}

		board[coords[1][0]][coords[1][1]] = board[coords[0][0]][coords[0][1]]
		board[coords[0][0]][coords[0][1]] = nil
		if pass != 0 {
			board[coords[1][0]+pass][coords[1][1]] = nil
		}

		pass = 0
		printBoard()
		turn = !turn
	}
	if !turn {
		winner = "White"
	} else {
		winner = "Black"
	}
	fmt.Println("Checkmate!", winner, "Wins!")
}
