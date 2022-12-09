package main

import (
	"fmt"
)

func MakeBoard() {
	Board = make([][]*Piece, 8)
	for i := range make([]int, 8) {
		Board[i] = make([]*Piece, 8)
	}

	SetSide(true)
	SetSide(false)
}

func newGame() {
	MakeBoard()
	PrintBoard()
	var coords [][]int
	var winner string
	for {
		if KingInCheck(Turn) {
			posMoves := GetAllMoves(Turn)
			useMoves := TryAllMoves(Turn, posMoves)
			//coords = getCheckedInput(turn, useMoves)
			if len(useMoves) == 0 {
				break
			}

		} else {
			//coords = handleInput(turn)
		}

		Board[coords[1][0]][coords[1][1]] = Board[coords[0][0]][coords[0][1]]
		Board[coords[0][0]][coords[0][1]] = nil
		if Pass != 0 {
			Board[coords[1][0]+Pass][coords[1][1]] = nil
		}

		Pass = 0
		PrintBoard()
		Turn = !Turn
	}
	if !Turn {
		winner = "White"
	} else {
		winner = "Black"
	}
	fmt.Println("Checkmate!", winner, "Wins!")
}

func PromotePawn(loc []int, turn bool) {
	if Board[loc[0]][loc[1]] == nil {
		return
	}
	if Board[loc[0]][loc[1]].Color != turn {
		return
	}
	if Board[loc[0]][loc[1]].Piece != DefPawn {
		return
	}

	if turn {
		if loc[0] == 7 {
			Board[loc[0]][loc[1]].Piece = Queen{}
		}
	} else {
		if loc[0] == 0 {
			Board[loc[0]][loc[1]].Piece = Queen{}
		}
	}
}
