package main

import (
	"sync"
)

// TODO: Castling
//       Drag and Drop / Click 2 locations

func init() {
	Wg = &sync.WaitGroup{}
	Wg.Add(1)
	MakeBoard()
}

func main() {

	go WsDealer()
	go Controller(Wg)

	Wg.Wait()
}
