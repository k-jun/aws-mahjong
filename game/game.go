package game

import (
	"aws-mahjong/board"
	"errors"
)

var (
	UserIsEmptyErr        = errors.New("user can't be empty")
	GameReachMaxMemberErr = errors.New("game already fulled")
)

type User struct {
	ID   string
	Name string
}

type Game interface {
	AddUser(user *User) error
	Capacity() int
}

type GameImpl struct {
	capacity int
	users    []*User
	board    *board.Board
}

func NewGame(capacity int, user *User) Game {
	return &GameImpl{
		capacity: capacity,
		users:    []*User{user},
		board:    nil,
	}
}

func (g *GameImpl) Capacity() int {
	return g.capacity
}

func (g *GameImpl) AddUser(user *User) error {
	if user == nil {
		return UserIsEmptyErr
	}
	if len(g.users) >= g.capacity {
		return GameReachMaxMemberErr
	}
	g.users = append(g.users, user)
	return nil
}
