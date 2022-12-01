package main

import (
	"reflect"
)

func getAllOptions(x int, y int) [][]int {
	piece := board[x][y]
	switch piece.piece {
	case pawn{}:
		return getPawnOptions(x, y)
	case rook{}:
		return getRookOptions(x, y)
	case bishop{}:
		return getBishopOptions(x, y)
	case knight{}:
		return getKnightOptions(x, y)
	case queen{}:
		return getQueenOptions(x, y)
	case king{}:
		return getKingOptions(x, y)
	}

	return [][]int{}
}
func getPawnOptions(x int, y int) [][]int {
	var out [][]int = [][]int{}
	var piece *piece = board[x][y]

	if piece.color {
		if x != 7 {
			if board[x+1][y] == nil {
				out = append(out, []int{x + 1, y})
			}
			if x == 1 {
				if board[3][y] == nil {
					out = append(out, []int{3, y})
				}
			}
			if y != 7 {
				if board[x+1][y+1] != nil {
					if !board[x+1][y+1].color {
						out = append(out, []int{x + 1, y + 1})
					}
				}
			}
			if y != 0 {
				if board[x+1][y-1] != nil {
					if !board[x+1][y-1].color {
						out = append(out, []int{x + 1, y - 1})
					}
				}
			}
			if x == 4 {
				if y < 7 {
					if board[x][y+1] != nil {
						if !board[x][y+1].color {
							if board[x][y+1].piece == defPawn {
								out = append(out, []int{x + 1, y + 1})
								pass = -1
							}
						}
					}
				}
				if y > 0 {
					if board[x][y-1] != nil {
						if !board[x][y-1].color && board[x][y-1].piece == defPawn {
							out = append(out, []int{x + 1, y - 1})
							pass = -1
						}
					}
				}
			}
		}
	} else if !piece.color {
		if x != 0 {
			if board[x-1][y] == nil {
				out = append(out, []int{x - 1, y})
			}
			if x == 6 {
				if board[4][y] == nil {
					out = append(out, []int{4, y})
				}
			}
			if y != 7 {
				if board[x-1][y+1] != nil {
					if board[x-1][y+1].color {
						out = append(out, []int{x - 1, y + 1})
					}
				}
			}
			if y != 0 {
				if board[x-1][y-1] != nil {
					if board[x-1][y-1].color {
						out = append(out, []int{x - 1, y - 1})
					}
				}
			}
			if x == 3 {
				if y < 7{
					if board[x][y+1] != nil {
						if board[x][y+1].color && board[x][y+1].piece == defPawn {
							out = append(out, []int{x - 1, y + 1})
							pass = 1
						}
					}
				}
				if x > 0{
					if board[x][y-1] != nil {
						if board[x][y-1].color && board[x][y-1].piece == defPawn {
							out = append(out, []int{x - 1, y - 1})
							pass = 1
						}
					}
				}
			}
		}
	}
	return out
}
func getRookOptions(xx int, yy int) [][]int {
	var out [][]int
	var x, y int = xx, yy
	if x != 0 {
		for x > 0 {
			x--
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
		x = xx
	}
	if x != 7 {
		for x < 7 {
			x++
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
		x = xx
	}
	if y != 0 {
		for y > 0 {
			y--
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
		y = yy
	}
	if y != 7 {
		for y < 7 {
			y++
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
	}
	return out
}
func getBishopOptions(xx int, yy int) [][]int {
	var out [][]int
	var x, y int = xx, yy
	if x != 0 && y != 0 {
		for x > 0 && y > 0 {
			x--
			y--
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
	}
	x, y = xx, yy
	if x != 0 && y != 7 {
		for x > 0 && y < 7 {
			x--
			y++
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
	}
	x, y = xx, yy
	if x != 7 && y != 0 {
		for x < 7 && y > 0 {
			x++
			y--
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
	}
	x, y = xx, yy
	if x != 7 && y != 7 {
		for x < 7 && y < 7 {
			x++
			y++
			out = append(out, []int{x, y})
			if board[x][y] != nil {
				break
			}
		}
	}
	return out
}

func getKnightOptions(x int, y int) [][]int {
	var out [][]int
	var tmp [][]int
	tmp = append(tmp, []int{x + 2, y + 1})
	tmp = append(tmp, []int{x + 2, y - 1})
	tmp = append(tmp, []int{x - 2, y + 1})
	tmp = append(tmp, []int{x - 2, y - 1})
	tmp = append(tmp, []int{x + 1, y + 2})
	tmp = append(tmp, []int{x - 1, y + 2})
	tmp = append(tmp, []int{x + 1, y - 2})
	tmp = append(tmp, []int{x - 1, y - 2})
	for _, ival := range tmp {
		if ival[0] >= 0 && ival[0] <= 7 {
			if ival[1] >= 0 && ival[1] <= 7 {
				out = append(out, ival)
			}
		}
	}
	return out
}

func getQueenOptions(x int, y int) [][]int {
	return append(getRookOptions(x, y), getBishopOptions(x, y)...)
}

func getKingOptions(x int, y int) [][]int {
	var out, tmp, tmp1, allTaken [][]int

	tmp = append(tmp, []int{x + 1, y + 1})
	tmp = append(tmp, []int{x + 1, y})
	tmp = append(tmp, []int{x + 1, y - 1})
	tmp = append(tmp, []int{x, y + 1})
	tmp = append(tmp, []int{x, y - 1})
	tmp = append(tmp, []int{x - 1, y + 1})
	tmp = append(tmp, []int{x - 1, y})
	tmp = append(tmp, []int{x - 1, y - 1})

	allTaken = getAllChecks(board[x][y].color)

	for i := range tmp {
		if tmp[i][0] >= 0 && tmp[i][0] <= 7 {
			if tmp[i][1] >= 0 && tmp[i][1] <= 7 {
				tmp1 = append(tmp1, tmp[i])
			}
		}
	}

	for i := range tmp1 {
		if contains(allTaken, tmp1[i]) {
			continue
		}
		loc := board[tmp1[i][0]][tmp1[i][1]]
		if loc == nil {
			out = append(out, tmp1[i])
			continue
		}
		if loc.color == board[x][y].color {
			continue
		}
		out = append(out, tmp1[i])
	}

	return out
}

func getAllChecks(team bool) [][]int {
	var allTaken [][]int
	var chnge = 0
	for i := range board {
		for j := range board[i] {
			loc := board[i][j]
			if loc == nil {
				continue
			}
			if loc.color == team {
				continue
			}

			if loc.piece == defKing {
				allTaken = append(allTaken, getFakeKing(i, j)...)
			} else if loc.piece == defPawn {
				if loc.color {
					chnge = 1
				} else {
					chnge = -1
				}
				pInfo := getPawnOptions(i, j)
				if len(pInfo) == 0{
					continue
				}
				for i, ival := range pInfo[:len(pInfo)-1] {
					if reflect.DeepEqual(ival, []int{i + chnge, j}) {
						pInfo = append(pInfo[:i], pInfo[i+1:]...)
					}
					if reflect.DeepEqual(ival, []int{i + (chnge * 2), j}) {
						pInfo = append(pInfo[:i], pInfo[i+1:]...)
					}
				}
				if reflect.DeepEqual(pInfo[:len(pInfo)-1], []int{i + chnge, j}) {
					pInfo = pInfo[:len(pInfo)-1]
				}
				if reflect.DeepEqual(pInfo[:len(pInfo)-1], []int{i + (chnge * 2), j}) {
					pInfo = pInfo[:len(pInfo)-1]
				}
				allTaken = append(allTaken, pInfo...)
			} else {
				allTaken = append(allTaken, getAllOptions(i, j)...)
			}
		}
	}
	return allTaken
}

func getFakeKing(x int, y int) [][]int {
	var tmp [][]int

	tmp = append(tmp, []int{x + 1, y + 1})
	tmp = append(tmp, []int{x + 1, y})
	tmp = append(tmp, []int{x + 1, y - 1})
	tmp = append(tmp, []int{x, y + 1})
	tmp = append(tmp, []int{x, y - 1})
	tmp = append(tmp, []int{x - 1, y + 1})
	tmp = append(tmp, []int{x - 1, y})
	tmp = append(tmp, []int{x - 1, y - 1})
	return tmp
}
