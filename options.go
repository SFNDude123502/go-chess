package main

import (
	"reflect"
)

func getAllOptions(x int, y int) [][]int {
	piece := Board[x][y]
	switch piece.Piece {
	case Pawn{}:
		return getPawnOptions(x, y)
	case Rook{}:
		return getRookOptions(x, y)
	case Bishop{}:
		return getBishopOptions(x, y)
	case Knight{}:
		return getKnightOptions(x, y)
	case Queen{}:
		return getQueenOptions(x, y)
	case King{}:
		return getKingOptions(x, y)
	}

	return [][]int{}
}

func getPawnOptions(x int, y int) [][]int {
	var out [][]int
	var piece = Board[x][y]

	if piece.Color {
		if x != 7 {
			if Board[x+1][y] == nil {
				out = append(out, []int{x + 1, y})
			}
			if x == 1 {
				if Board[3][y] == nil {
					out = append(out, []int{3, y})
				}
			}
			if y < 7 {
				if Board[x+1][y+1] != nil {
					if !Board[x+1][y+1].Color {
						out = append(out, []int{x + 1, y + 1})
					}
				}
			}
			if y > 0 {
				if Board[x+1][y-1] != nil {
					if !Board[x+1][y-1].Color {
						out = append(out, []int{x + 1, y - 1})
					}
				}
			}
			if x == 4 {
				if y < 7 {
					if Board[x][y+1] != nil {
						if !Board[x][y+1].Color {
							if Board[x][y+1].Piece == DefPawn {
								out = append(out, []int{x + 1, y + 1})
								Pass = -1
							}
						}
					}
				}
				if y > 0 {
					if Board[x][y-1] != nil {
						if !Board[x][y-1].Color && Board[x][y-1].Piece == DefPawn {
							out = append(out, []int{x + 1, y - 1})
							Pass = -1
						}
					}
				}
			}
		}
	} else if !piece.Color {
		if x != 0 {
			if Board[x-1][y] == nil {
				out = append(out, []int{x - 1, y})
			}
			if x == 6 {
				if Board[4][y] == nil {
					out = append(out, []int{4, y})
				}
			}
			if y < 7 {
				if Board[x-1][y+1] != nil {
					if Board[x-1][y+1].Color {
						out = append(out, []int{x - 1, y + 1})
					}
				}
			}
			if y > 0 {
				if Board[x-1][y-1] != nil {
					if Board[x-1][y-1].Color {
						out = append(out, []int{x - 1, y - 1})
					}
				}
			}
			if x == 3 {
				if y < 7 {
					if Board[x][y+1] != nil {
						if Board[x][y+1].Color && Board[x][y+1].Piece == DefPawn {
							out = append(out, []int{x - 1, y + 1})
							Pass = 1
						}
					}
				}
				if y > 0 {
					if Board[x][y-1] != nil {
						if Board[x][y-1].Color && Board[x][y-1].Piece == DefPawn {
							out = append(out, []int{x - 1, y - 1})
							Pass = 1
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
			if Board[x][y] != nil {
				break
			}
		}
		x = xx
	}
	if x != 7 {
		for x < 7 {
			x++
			out = append(out, []int{x, y})
			if Board[x][y] != nil {
				break
			}
		}
		x = xx
	}
	if y != 0 {
		for y > 0 {
			y--
			out = append(out, []int{x, y})
			if Board[x][y] != nil {
				break
			}
		}
		y = yy
	}
	if y != 7 {
		for y < 7 {
			y++
			out = append(out, []int{x, y})
			if Board[x][y] != nil {
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
			if Board[x][y] != nil {
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
			if Board[x][y] != nil {
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
			if Board[x][y] != nil {
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
			if Board[x][y] != nil {
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

	allTaken = GetAllChecks(Board[x][y].Color)

	for i := range tmp {
		if tmp[i][0] >= 0 && tmp[i][0] <= 7 {
			if tmp[i][1] >= 0 && tmp[i][1] <= 7 {
				tmp1 = append(tmp1, tmp[i])
			}
		}
	}

	for i := range tmp1 {
		if Contains(allTaken, tmp1[i]) {
			continue
		}
		loc := Board[tmp1[i][0]][tmp1[i][1]]
		if loc == nil {
			out = append(out, tmp1[i])
			continue
		}
		if loc.Color == Board[x][y].Color {
			continue
		}
		out = append(out, tmp1[i])
	}

	return out
}

func GetAllChecks(team bool) [][]int {
	var allTaken [][]int
	var chnge = 0
	for i := range Board {
		for j := range Board[i] {
			loc := Board[i][j]
			if loc == nil {
				continue
			}
			if loc.Color == team {
				continue
			}

			if loc.Piece == DefKing {
				allTaken = append(allTaken, getFakeKing(i, j)...)
			} else if loc.Piece == DefPawn {
				if loc.Color {
					chnge = 1
				} else {
					chnge = -1
				}
				pInfo := getPawnOptions(i, j)
				if len(pInfo) == 0 {
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

func GetAllMoves(team bool) (out [][4]int) {
	for i := range Board {
		for j := range Board[i] {
			loc := Board[i][j]
			if loc == nil {
				continue
			}
			if loc.Color == !team {
				continue
			}
			moves := getAllOptions(i, j)
			for k := range moves {
				out = append(out, [4]int{i, j, moves[k][0], moves[k][1]})
			}

		}
	}
	return
}

func TryAllMoves(team bool, moves [][4]int) [][]int {
	var tmp *Piece
	var kingLoc = FindKing(team)
	var out [][]int
	for _, ival := range moves {
		tmp = Board[ival[2]][ival[3]]
		Board[ival[2]][ival[3]] = Board[ival[0]][ival[1]]
		if !KingInCheck(Board[moves[0][0]][moves[0][1]].Color) {
			out = append(out, []int{ival[2], ival[3]})
		}
		Board[ival[2]][ival[3]] = tmp
	}
	for i := 0; i < len(out); i++ {
		if reflect.DeepEqual(out[i], kingLoc) {
			if len(out) == 1 {
				return [][]int{}
			}
			out = append(out[:i], out[:i+1]...)
			i--
		}
	}
	return out
}
