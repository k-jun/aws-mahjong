package game

import "aws-mahjong/board"

type Game struct {
	capacity int
	board    *board.Board
}
