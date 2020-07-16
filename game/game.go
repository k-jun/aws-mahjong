package game

import (
	"aws-mahjong/board"
	"errors"
)

var (
	UsernameIsEmpty       = errors.New("username can't be empty")
	GameReachMaxMemberErr = errors.New("game already fulled")
)

type Game interface {
	AddUsername(username string) error
	Capacity() int
}

type GameImpl struct {
	capacity  int
	usernames []string
	board     *board.Board
}

func NewGame(capacity int, username string) Game {
	return &GameImpl{
		capacity:  capacity,
		usernames: []string{username},
		board:     nil,
	}
}

func (g *GameImpl) Capacity() int {
	return g.capacity
}

func (g *GameImpl) AddUsername(username string) error {
	if username == "" {
		return UsernameIsEmpty
	}
	if len(g.usernames) >= g.capacity {
		return GameReachMaxMemberErr
	}
	g.usernames = append(g.usernames, username)
	return nil
}
