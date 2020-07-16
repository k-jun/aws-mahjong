package game

import "aws-mahjong/board"

type Game struct {
	capacity int
	board    *board.Board
}

func NewGame(capacity int) *Game {
	return &Game{
		capacity: capacity,
		board:    nil,
	}
}

func (g *Game) Capacity() int {
	return g.capacity

}
