package repository

import (
	"aws-mahjong/game"
	"errors"
)

var (
	GameNotFoundErr = errors.New("board not found")
	GameIsNil       = errors.New("baord is nil")
)

// TODO periodically sync with room repository to save memory

type GameRepository struct {
	games map[string]*game.GameImpl
}

func NewGameRepository() *GameRepository {
	return &GameRepository{games: map[string]*game.GameImpl{}}
}

func (r *GameRepository) Add(roomName string, board *game.GameImpl) error {
	if board == nil {
		return GameIsNil
	}
	r.games[roomName] = board
	return nil
}

func (r *GameRepository) Remove(roomName string) error {
	if r.games[roomName] == nil {
		return GameNotFoundErr
	}
	delete(r.games, roomName)
	return nil
}

func (r *GameRepository) Find(roomName string) (*game.GameImpl, error) {
	if r.games[roomName] == nil {
		return nil, GameNotFoundErr
	}
	return r.games[roomName], nil
}
