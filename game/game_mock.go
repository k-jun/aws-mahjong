package game

import "aws-mahjong/board"

var _ Game = &GameMock{}

type GameMock struct {
	ExpectedCapacity int
	ExpectedError    error
	ExpectedBoard    board.Board
}

func (g *GameMock) Capacity() int {
	return g.ExpectedCapacity
}

func (g *GameMock) AddUser(user *User) error {
	return g.ExpectedError
}

func (g *GameMock) RemoveUser(user *User) error {
	return g.ExpectedError

}

func (g *GameMock) Board() board.Board {
	return g.ExpectedBoard
}

func (g *GameMock) GameStart() error {
	return g.ExpectedError
}
