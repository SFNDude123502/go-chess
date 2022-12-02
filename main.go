package main

import (
	"github.com/gin-gonic/gin"
)

// TODO: Castling

func init() {
	makeBoard()
}

func main() {
	srv := gin.Default()

	srv.LoadHTMLGlob("./templates/index.go.html")
	srv.Static("/templates", "./templates")
	srv.StaticFile("/favicon.ico", "./favicon.ico")

	srv.GET("/", getChess)

	srv.Run(":80")
}

func getChess(c *gin.Context) {

	c.HTML(200, "board", gin.H{"board": htmlBoard()})
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

var hash1 map[interface{}]string = map[interface{}]string{
	pawn{}:   "P",
	rook{}:   "R",
	knight{}: "N",
	bishop{}: "B",
	king{}:   "K",
	queen{}:  "Q",
}
