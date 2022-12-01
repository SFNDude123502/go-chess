package main

import "fmt"

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
