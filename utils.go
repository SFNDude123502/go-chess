package main

import (
	"reflect"
)

func SetSide(color bool) {
	var end, pLine = 0, 1
	if !color {
		end, pLine = 7, 6
	}
	Board[end][0] = &Piece{Color: color, Piece: Rook{}}
	Board[end][7] = &Piece{Color: color, Piece: Rook{}}
	Board[end][1] = &Piece{Color: color, Piece: Knight{}}
	Board[end][6] = &Piece{Color: color, Piece: Knight{}}
	Board[end][2] = &Piece{Color: color, Piece: Bishop{}}
	Board[end][5] = &Piece{Color: color, Piece: Bishop{}}
	Board[end][3] = &Piece{Color: color, Piece: Queen{}}
	Board[end][4] = &Piece{Color: color, Piece: King{}}
	for i := range make([]int, 8) {
		Board[pLine][i] = &Piece{Color: color, Piece: Pawn{}}
	}
}

func Contains(list [][]int, target []int) bool {
	for _, ival := range list {
		if reflect.DeepEqual(ival, target) {
			return true
		}
	}
	return false
}

func FindKing(team bool) []int {
	for i := range Board {
		for j := range Board[i] {
			loc := Board[i][j]
			if loc == nil {
				continue
			}
			if loc.Piece == DefKing {
				if loc.Color == team {
					return []int{i, j}
				}
			}
		}
	}
	return []int{-1, -1}
}

func KingInCheck(team bool) bool {
	allChecks := GetAllChecks(team)
	kingLoc := FindKing(team)
	return Contains(allChecks, kingLoc)
}

func HtmlBoard() [][]string {
	var out [][]string
	var str string
	for i := range Board {
		out = append(out, make([]string, 8))
		for j := range Board[i] {
			loc := Board[i][j]
			if loc == nil {
				out[i][j] = ""
				continue
			}
			str = "/templates/pieces/"
			if Board[i][j].Color {
				str += "w"
			} else {
				str += "b"
			}
			str += PieceHash[Board[i][j].Piece]
			str += ".png"
			out[i][j] = str
		}
	}
	return out
}

func Eh(err error) {
	if err != nil {
		panic(err)
	}
}

func Make2dArr(h int, w int) [][]string {
	var arr [][]string
	for range make([]int, h) {
		arr = append(arr, make([]string, w))
	}
	return arr
}
