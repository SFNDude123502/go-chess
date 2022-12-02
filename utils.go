package main

import "reflect"

func setSide(color bool) {
	var end, pLine = 0, 1
	if !color {
		end, pLine = 7, 6
	}
	board[end][0] = &piece{color: color, piece: rook{}}
	board[end][7] = &piece{color: color, piece: rook{}}
	board[end][1] = &piece{color: color, piece: knight{}}
	board[end][6] = &piece{color: color, piece: knight{}}
	board[end][2] = &piece{color: color, piece: bishop{}}
	board[end][5] = &piece{color: color, piece: bishop{}}
	board[end][3] = &piece{color: color, piece: queen{}}
	board[end][4] = &piece{color: color, piece: king{}}
	for i := range make([]int, 8) {
		board[pLine][i] = &piece{color: color, piece: pawn{}}
	}
}

func contains(list [][]int, target []int) bool {
	for _, ival := range list {
		if reflect.DeepEqual(ival, target) {
			return true
		}
	}
	return false
}

func findKing(team bool) []int {
	for i := range board {
		for j := range board[i] {
			loc := board[i][j]
			if loc == nil {
				continue
			}
			if loc.piece == defKing {
				if loc.color == team {
					return []int{i, j}
				}
			}
		}
	}
	return []int{-1, -1}
}

func kingInCheck(team bool) bool {
	allChecks := getAllChecks(team)
	kingLoc := findKing(team)
	return contains(allChecks, kingLoc)
}

func htmlBoard() [][]string {
	var out [][]string
	var str string
	for i := range board {
		out = append(out, make([]string, 8))
		for j := range board[i] {
			loc := board[i][j]
			if loc == nil {
				out[i][j] = ""
				continue
			}
			str = "/templates/pieces/"
			if board[i][j].color {
				str += "w"
			} else {
				str += "b"
			}
			str += hash1[board[i][j].piece]
			str += ".png"
			out[i][j] = str
		}
	}
	return out
}

