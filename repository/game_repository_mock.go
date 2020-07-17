package repository

import "aws-mahjong/game"

var _ GameRepository = &GameRepositoryMock{}

type GameRepositoryMock struct {
	ExpectedGame  game.Game
	ExpectedError error
}

func (r *GameRepositoryMock) Add(roomName string, inGame game.Game) error {
	return r.ExpectedError
}

func (r *GameRepositoryMock) Remove(roomName string) error {
	return r.ExpectedError
}

func (r *GameRepositoryMock) Find(roomName string) (game.Game, error) {
	return r.ExpectedGame, r.ExpectedError
}
