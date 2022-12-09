package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PrintBoard() {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	fmt.Println("|   || a | b | c | d | e | f | g | h |")
	pL()
	pL()
	fmt.Println("| 1 " + getLine(Board[0]))
	pL()
	fmt.Println("| 2 " + getLine(Board[1]))
	pL()
	fmt.Println("| 3 " + getLine(Board[2]))
	pL()
	fmt.Println("| 4 " + getLine(Board[3]))
	pL()
	fmt.Println("| 5 " + getLine(Board[4]))
	pL()
	fmt.Println("| 6 " + getLine(Board[5]))
	pL()
	fmt.Println("| 7 " + getLine(Board[6]))
	pL()
	fmt.Println("| 8 " + getLine(Board[7]))
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
func getLine(line []*Piece) string {
	ret := "|| "
	for i := range make([]int, 8) {
		ret += boardNotation(line[i]) + " | "
	}
	return ret
}
func boardNotation(item *Piece) string {
	if item == nil {
		return " "
	}
	switch item.Piece {
	case Pawn{}:
		return "p"
	case Rook{}:
		return "r"
	case Knight{}:
		return "n"
	case Bishop{}:
		return "b"
	case Queen{}:
		return "q"
	case King{}:
		return "k"
	}
	panic("none of the above")
}

func HandleInput(c *gin.Context, st string, end string, turn bool) (out [][]int, htmlReturned bool) {
	out = [][]int{{-1, -1}, {-1, -1}}
	htmlReturned = false
	out[0] = LocHash[st]
	out[1] = LocHash[end]
	options := getAllOptions(out[0][0], out[0][1])
	if !Contains(options, out[1]) {
		fmt.Println("ivalid")
		htmlReturned = true
		E = "Invalid Move"
		c.JSON(200, gin.H{"err": E})
		return
	}
	start, goal := Board[out[0][0]][out[0][1]], Board[out[1][0]][out[1][1]]
	if start == nil {
		fmt.Println("st empty")
		htmlReturned = true
		E = "Starting Location is Empty"
		c.JSON(200, gin.H{"err": E})
		return
	}
	if start.Color != turn {
		fmt.Println("move opp")
		htmlReturned = true
		E = "You Can't Move an Enemy Piece!"
		c.JSON(200, gin.H{"err": E})
		return
	}
	if goal != nil {
		if goal.Color == turn {
			fmt.Println("self")
			htmlReturned = true
			E = "You Can't Attack Your Own Pieces!"
			c.JSON(200, gin.H{"err": E})
			return
		}
	}
	E = ""

	return
}

func GetCheckedInput(c *gin.Context, st string, end string, turn bool, posMoves [][]int) (out [][]int, htmlReturned bool) {
	out = [][]int{{}, {}}
	htmlReturned = false

	fmt.Println("You Are In Check rn, at the end of this round you must not be in check anymore")
	out[0] = LocHash[st]
	out[1] = LocHash[end]
	options := getAllOptions(out[0][0], out[0][1])
	if !Contains(posMoves, out[1]) {
		fmt.Println("still check")
		htmlReturned = true
		E = "King Still in Check!"
		c.JSON(200, gin.H{"err": E})
		return
	}
	if !Contains(options, out[1]) {
		fmt.Println("invalid")
		htmlReturned = true
		E = "Invalid Move"
		c.JSON(200, gin.H{"err": E})
		return
	}
	start, goal := Board[out[0][0]][out[0][1]], Board[out[1][0]][out[1][1]]
	if start == nil {
		fmt.Println("st empty")
		htmlReturned = true
		E = "Starting Location is Empty"
		c.JSON(200, gin.H{"err": E})
		return
	}
	if start.Color != turn {
		fmt.Println("moved enemy")
		htmlReturned = true
		E = "You Can't Move an Enemy Piece!"
		c.JSON(200, gin.H{"err": E})
		return
	}
	if goal != nil {
		if goal.Color == turn {
			fmt.Println("self")
			htmlReturned = true
			E = "You Can't Attack Your Own Pieces!"
			c.JSON(200, gin.H{"err": E})
			return
		}
	}

	return
}
