package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var Messages []string

func Controller(wg *sync.WaitGroup) {
	defer wg.Done()
	srv := gin.Default()

	srv.LoadHTMLFiles("./templates/index.go.html", "./templates/victory.go.html")
	srv.Static("/templates", "./templates")
	srv.StaticFile("/favicon.ico", "./favicon.ico")

	srv.GET("/", getChess)
	srv.POST("/move", postMove)
	srv.POST("/message", postMessage)
	srv.GET("/reload", webSocket)

	err := srv.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func webSocket(c *gin.Context) {
	//if len(clients) == 2 {
	//	return
	//}
	WsConv(c.Writer, c.Request)
}

func postMessage(c *gin.Context) {
	Messages = append(Messages, c.PostForm("message"))
	c.JSON(200, nil)
}

func getChess(c *gin.Context) {
	PrintBoard()
	c.HTML(200, "board", gin.H{"board": HtmlBoard()})
}

func postMove(c *gin.Context) {
	fmt.Println("moved")
	var winner string
	var htmlReturned bool
	st, end := c.PostForm("st"), c.PostForm("end")
	fmt.Println(st, end)
	var coords [][]int

	if Checked {
		posMoves := GetAllMoves(Turn)
		useMoves := TryAllMoves(Turn, posMoves)
		coords, htmlReturned = GetCheckedInput(c, st, end, Turn, useMoves)
		if htmlReturned {
			return
		}

		if len(useMoves) == 0 {
			if !Turn {
				winner = "White"
			} else {
				winner = "Black"
			}
			c.JSON(200, gin.H{"winner": winner})
			return
		}
	} else {
		coords, htmlReturned = HandleInput(c, st, end, Turn)
		if htmlReturned {
			return
		}
	}

	tmp := Board[coords[0][0]][coords[0][1]]
	Board[coords[0][0]][coords[0][1]] = nil

	if KingInCheck(Turn) {
		fmt.Println("move into check")
		Board[coords[0][0]][coords[0][1]] = tmp
		E = "This move will put your king into check!"
		c.JSON(200, gin.H{"err": E})
		return
	}
	Board[coords[1][0]][coords[1][1]] = tmp
	PromotePawn(coords[1], Turn)
	if Pass != 0 {
		Board[coords[1][0]+Pass][coords[1][1]] = nil
	}

	if KingInCheck(!Turn) {
		Checked = true

		posMoves := GetAllMoves(!Turn)
		useMoves := TryAllMoves(!Turn, posMoves)

		if len(useMoves) == 0 {
			if Turn {
				winner = "White"
			} else {
				winner = "Black"
			}
			c.JSON(200, gin.H{"winner": winner})
			return
		}
		fmt.Println(useMoves)
	} else {
		Checked = false
	}
	fmt.Println("success")

	Pass = 0
	c.JSON(200, nil)
	Turn = !Turn
}
