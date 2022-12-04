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
	srv.POST("/move", postMove)
	srv.POST("/message", postMessage)
	srv.GET("/reload", webSocket)
	go wsDealer()

	srv.Run(":8080")
}

func webSocket(c *gin.Context) {
	wsConv(c.Writer, c.Request)
}

func postMessage(c *gin.Context) {
	messages = append(messages, c.PostForm("message"))
	c.Redirect(301, "/")
}

func getChess(c *gin.Context) {
	printBoard()
	c.HTML(200, "board", gin.H{"board": htmlBoard(), "err": e, "messages": messages})
}

func postMove(c *gin.Context) {
	fmt.Println("moved")
	var winner string
	var htmlReturned bool
	st, end := c.PostForm("st"), c.PostForm("end")
	fmt.Println(st, end)
	var coords [][]int

	if checked {
		posMoves := getAllMoves(turn)
		useMoves := tryAllMoves(turn, posMoves)
		coords, htmlReturned = getCheckedInput(c, st, end, turn, useMoves)
		if htmlReturned {
			fmt.Println("2")
			return
		}

		fmt.Println("3")
		if len(useMoves) == 0 {
			fmt.Println("4")
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
		c.Redirect(301, "/")
		return
	}
	board[coords[1][0]][coords[1][1]] = tmp

	if pass != 0 {
		board[coords[1][0]+pass][coords[1][1]] = nil
	}

	if kingInCheck(!turn) {
		fmt.Println("1")
		checked = true

		posMoves := getAllMoves(!turn)
		useMoves := tryAllMoves(!turn, posMoves)

		if len(useMoves) == 0 {
			fmt.Println("5")
			if turn {
				winner = "White"
			} else {
				winner = "Black"
			}
			c.HTML(200, "victory", gin.H{"winner": winner})
			return
		}
		fmt.Println(useMoves)
	} else {
		checked = false
	}
	fmt.Println("success")

	pass = 0
	c.Redirect(301, "/")
	turn = !turn
}
