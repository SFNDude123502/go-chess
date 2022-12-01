package main

import (
	"fmt"
	"reflect"
)

// TODO: Castling

type piece struct {
	color bool
	piece interface{}
}

type pawn struct{}
type rook struct{}
type knight struct{}
type bishop struct{}
type queen struct{}
type king struct{}

var board [][]*piece
var pass int = 0
var turn bool = true
var defPawn = pawn{}
var defKing = king{}

func init() {
	board = make([][]*piece, 8)
	for i := range make([]int, 8) {
		board[i] = make([]*piece, 8)
	}

	setSide(true)
	setSide(false)
}
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

func main() {
	printBoard()
	for {
		coords := getInput(turn)

		board[coords[1][0]][coords[1][1]] = board[coords[0][0]][coords[0][1]]
		board[coords[0][0]][coords[0][1]] = nil
		if pass != 0 {
			board[coords[1][0]+pass][coords[1][1]] = nil
		}
		allChecks := getAllChecks(turn)
		if contains(allChecks, ) {
			
		}

		pass = 0
		printBoard()
		turn = !turn
	}
}

func getInput(turn bool) [2][2]int {
	var in1, in2 string
	var in [2]string
	var out [2][2]int = [2][2]int{{}, {}}
	for {
		fmt.Println("Enter current location of desired piece, followed by desired location formatted like this: \ng1 f3")
		fmt.Scan(&in1, &in2)
		in = [2]string{in1, in2}
		out[0] = hash[in[0]]
		out[1] = hash[in[1]]
		options := getAllOptions(out[0][0], out[0][1])
		if !contains(options, []int{out[1][0], out[1][1]}) {
			fmt.Println("Illegal Move")
			continue
		}
		start, goal := board[out[0][0]][out[0][1]], board[out[1][0]][out[1][1]]
		if start == nil {
			fmt.Println("Starting Location is Empty")
			continue
		}
		if start.color != turn {
			fmt.Println("You Can't Move an Enemy Piece!")
			continue
		}
		if goal != nil {
			if goal.color == turn {
				fmt.Println("You Can't Attack Your Own Pieces!", out)
				continue
			}
		}

		break
	}
	return out
}

func contains(list [][]int, target []int) bool {
	for _, ival := range list {
		if reflect.DeepEqual(ival, target) {
			return true
		}
	}
	return false
}

var hash map[string]([2]int) = map[string]([2]int){
	"a1": {0, 0}, "b1": {0, 1}, "c1": {0, 2}, "d1": {0, 3}, "e1": {0, 4}, "f1": {0, 5}, "g1": {0, 6}, "h1": {0, 7},
	"a2": {1, 0}, "b2": {1, 1}, "c2": {1, 2}, "d2": {1, 3}, "e2": {1, 4}, "f2": {1, 5}, "g2": {1, 6}, "h2": {1, 7},
	"a3": {2, 0}, "b3": {2, 1}, "c3": {2, 2}, "d3": {2, 3}, "e3": {2, 4}, "f3": {2, 5}, "g3": {2, 6}, "h3": {2, 7},
	"a4": {3, 0}, "b4": {3, 1}, "c4": {3, 2}, "d4": {3, 3}, "e4": {3, 4}, "f4": {3, 5}, "g4": {3, 6}, "h4": {3, 7},
	"a5": {4, 0}, "b5": {4, 1}, "c5": {4, 2}, "d5": {4, 3}, "e5": {4, 4}, "f5": {4, 5}, "g5": {4, 6}, "h5": {4, 7},
	"a6": {5, 0}, "b6": {5, 1}, "c6": {5, 2}, "d6": {5, 3}, "e6": {5, 4}, "f6": {5, 5}, "g6": {5, 6}, "h6": {5, 7},
	"a7": {6, 0}, "b7": {6, 1}, "c7": {6, 2}, "d7": {6, 3}, "e7": {6, 4}, "f7": {6, 5}, "g7": {6, 6}, "h7": {6, 7},
	"a8": {7, 0}, "b8": {7, 1}, "c8": {7, 2}, "d8": {7, 3}, "e8": {7, 4}, "f8": {7, 5}, "g8": {7, 6}, "h8": {7, 7},
}
