package game

import "aws-mahjong/board"

var _ Game = &GameMock{}

type GameMock struct {
	ExpectedCapacity int
	ExpectedError    error
	ExpectedBoard    board.Board
	ExpectedUsers    []*User
}

func (g *GameMock) Capacity() int {
	return g.ExpectedCapacity
}

func (g *GameMock) AddUser(user *User) error {
	g.ExpectedUsers = append(g.ExpectedUsers, user)
	return g.ExpectedError
}

func (g *GameMock) RemoveUser(user *User) error {
	for idx, u := range g.ExpectedUsers {
		if u.ID == user.ID {
			lastIdx := len(g.ExpectedUsers) - 1
			g.ExpectedUsers[idx], g.ExpectedUsers[lastIdx] = g.ExpectedUsers[lastIdx], g.ExpectedUsers[idx]
			g.ExpectedUsers = g.ExpectedUsers[:lastIdx]

		}

	}
	return g.ExpectedError
}

func (g *GameMock) Board() board.Board {
	return g.ExpectedBoard
}

func (g *GameMock) GameStart() error {
	return g.ExpectedError
}

func (g *GameMock) Users() []*User {
	return g.ExpectedUsers
}
