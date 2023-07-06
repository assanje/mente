package main

import (
	. "mente/core"
)

func main() {
	board := NewChessboard()
	board.Display()
	//board.DisplayDebug()
	FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board2 := FENtoBitboard(FEN)
	board2.Display()
}
