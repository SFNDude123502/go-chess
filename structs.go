package main

type Piece struct {
	Color bool
	Piece interface{}
}

type Pawn struct{}
type Rook struct{}
type Knight struct{}
type Bishop struct{}
type Queen struct{}
type King struct{}

type webReq struct {
	Board    [][]string
	Messages []string
}
