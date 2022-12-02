package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func printBoard() {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	fmt.Println("|   || a | b | c | d | e | f | g | h |")
	pL()
	pL()
	fmt.Println("| 1 " + getLine(board[0]))
	pL()
	fmt.Println("| 2 " + getLine(board[1]))
	pL()
	fmt.Println("| 3 " + getLine(board[2]))
	pL()
	fmt.Println("| 4 " + getLine(board[3]))
	pL()
	fmt.Println("| 5 " + getLine(board[4]))
	pL()
	fmt.Println("| 6 " + getLine(board[5]))
	pL()
	fmt.Println("| 7 " + getLine(board[6]))
	pL()
	fmt.Println("| 8 " + getLine(board[7]))
	pL()
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
func pL() {
	pr := "|---||"
	for range make([]int, 8) {
		pr += "---|"
	}
	fmt.Println(pr)
}
func getLine(line []*piece) string {
	ret := "|| "
	for i := range make([]int, 8) {
		ret += boardNotation(line[i]) + " | "
	}
	return ret
}
func boardNotation(item *piece) string {
	if item == nil {
		return " "
	}
	switch item.piece {
	case pawn{}:
		return "p"
	case rook{}:
		return "r"
	case knight{}:
		return "n"
	case bishop{}:
		return "b"
	case queen{}:
		return "q"
	case king{}:
		return "k"
	}
	panic("none of the above")
}

func handleInput(c *gin.Context, st string, end string, turn bool) (out [][]int, htmlReturned bool) {
	out = [][]int{{-1, -1}, {-1, -1}}
	htmlReturned = false
	out[0] = hash[st]
	out[1] = hash[end]
	options := getAllOptions(out[0][0], out[0][1])
	if !contains(options, out[1]) {
		fmt.Println("ivalid")
		htmlReturned = true
		e = "Invalid Move"
		c.Redirect(301, "/")
		return
	}
	start, goal := board[out[0][0]][out[0][1]], board[out[1][0]][out[1][1]]
	if start == nil {
		fmt.Println("st empty")
		htmlReturned = true
		e = "Starting Location is Empty"
		c.Redirect(301, "/")
		return
	}
	if start.color != turn {
		fmt.Println("move opp")
		htmlReturned = true
		e = "You Can't Move an Enemy Piece!"
		c.Redirect(301, "/")
		return
	}
	if goal != nil {
		if goal.color == turn {
			fmt.Println("self")
			htmlReturned = true
			e = "You Can't Attack Your Own Pieces!"
			c.Redirect(301, "/")
			return
		}
	}
	e = ""

	return
}

func getCheckedInput(c *gin.Context, st string, end string, turn bool, posMoves [][]int) (out [][]int, htmlReturned bool) {
	out = [][]int{{}, {}}
	htmlReturned = false

	fmt.Println("You Are In Check rn, at the end of this round you must not be in check anymore")
	out[0] = hash[st]
	out[1] = hash[end]
	options := getAllOptions(out[0][0], out[0][1])
	if !contains(posMoves, out[1]) {
		fmt.Println("still check")
		htmlReturned = true
		e = "King Still in Check!"
		c.Redirect(301, "/")
		return
	}
	if !contains(options, out[1]) {
		fmt.Println("invalid")
		htmlReturned = true
		e = "Invalid Move"
		c.Redirect(301, "/")
		return
	}
	start, goal := board[out[0][0]][out[0][1]], board[out[1][0]][out[1][1]]
	if start == nil {
		fmt.Println("st empty")
		htmlReturned = true
		e = "Starting Location is Empty"
		c.Redirect(301, "/")
		return
	}
	if start.color != turn {
		fmt.Println("moved enemy")
		htmlReturned = true
		e = "You Can't Move an Enemy Piece!"
		c.Redirect(301, "/")
		return
	}
	if goal != nil {
		if goal.color == turn {
			fmt.Println("self")
			htmlReturned = true
			e = "You Can't Attack Your Own Pieces!"
			c.Redirect(301, "/")
			return
		}
	}

	return
}
