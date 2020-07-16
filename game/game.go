package game

import "aws-mahjong/board"

type Game struct {
	capacity int
	names    []string
	board    *board.Board
}

func NewGame(capacity int, username string) *Game {
	return &Game{
		capacity: capacity,
		names:    []string{username},
		board:    nil,
	}
}

func (g *Game) AddUser(username string) error {
	g.names = append(g.names, username)
	return nil
}

func (g *Game) Capacity() int {
	return g.capacity
}
