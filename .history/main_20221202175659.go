package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO: Castling

func init() {
	makeBoard()
}

func main() {
	srv := gin.Default()

	srv.LoadHTMLFiles("./templates/index.go.html", "./templates/victory.go.html")
	srv.Static("/templates", "./templates")
	srv.StaticFile("/favicon.ico", "./favicon.ico")

	srv.GET("/", getChess)
	srv.GET("/move", getMove)

	srv.Run(":8080")
}

func getChess(c *gin.Context) {
	printBoard()
	c.HTML(200, "board", gin.H{"board": htmlBoard(), "err": e})
}

func getMove(c *gin.Context) {
	fmt.Println("moved")
	var winner string
	var htmlReturned bool
	st, end := c.Query("st"), c.Query("end")
	fmt.Println(st, end)
	var coords [][]int
	if kingInCheck(turn) {
		posMoves := getAllMoves(turn)
		useMoves := tryAllMoves(posMoves)
		coords, htmlReturned = getCheckedInput(c, st, end, turn, useMoves)
		if htmlReturned {
			return
		}
		if len(useMoves) == 0 {
			if !turn {
				winner = "White"
			} else {
				winner = "Black"
			}
			c.HTML(200, "victory", gin.H{"winner": winner})
			return
		}

	} else {
		coords, htmlReturned = handleInput(c, st, end, turn)
		if htmlReturned {
			return
		}
	}

	tmp := board[coords[0][0]][coords[0][1]]
	board[coords[0][0]][coords[0][1]] = nil

	if kingInCheck(turn) {
		fmt.Println("move into check")
		board[coords[0][0]][coords[0][1]] = tmp
		e = "This move will put your king into check!"
		c.Redirect(200, "/")
		return
	}
	board[coords[1][0]][coords[1][1]] = tmp

	if pass != 0 {
		board[coords[1][0]+pass][coords[1][1]] = nil
	}

	fmt.Println("success")

	pass = 0
	c.Redirect(301, "/")
	turn = !turn
}
