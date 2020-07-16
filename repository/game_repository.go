package repository

import (
	"aws-mahjong/game"
	"errors"
)

var (
	BoardNotFoundErr = errors.New("board not found")
	BoardIsNil       = errors.New("baord is nil")
)

// TODO periodically sync with room repository to save memory

type GameRepository struct {
	games map[string]*game.Game
}

func NewGameRepository() *GameRepository {
	return &GameRepository{games: map[string]*game.Game{}}
}

func (r *GameRepository) Add(roomName string, board *game.Game) error {
	if board == nil {
		return BoardIsNil
	}
	r.games[roomName] = board
	return nil
}

func (r *GameRepository) Remove(roomName string) error {
	if r.games[roomName] == nil {
		return BoardNotFoundErr
	}
	delete(r.games, roomName)
	return nil
}

func (r *GameRepository) Find(roomName string) (*game.Game, error) {
	if r.games[roomName] == nil {
		return nil, BoardNotFoundErr
	}
	return r.games[roomName], nil
}
