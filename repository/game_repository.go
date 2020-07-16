package repository

import (
	"aws-mahjong/game"
	"errors"
)

var (
	GameNotFoundErr = errors.New("board not found")
	GameIsNil       = errors.New("baord is nil")
	RoomNameIsEmpry = errors.New("roomName is empty")
)

type GameRepository interface {
	Add(roomName string, board game.Game) error
	Remove(roomName string) error
	Find(roomName string) (game.Game, error)
}

// TODO periodically sync with room repository to save memory

type GameRepositoryImpl struct {
	games map[string]game.Game
}

func NewGameRepository() GameRepository {
	return &GameRepositoryImpl{
		games: map[string]game.Game{},
	}
}

func (r *GameRepositoryImpl) Add(roomName string, board game.Game) error {
	if roomName == "" {
		return RoomNameIsEmpry
	}

	if board == nil {
		return GameIsNil
	}
	r.games[roomName] = board
	return nil
}

func (r *GameRepositoryImpl) Remove(roomName string) error {
	if r.games[roomName] == nil {
		return GameNotFoundErr
	}
	delete(r.games, roomName)
	return nil
}

func (r *GameRepositoryImpl) Find(roomName string) (game.Game, error) {
	if r.games[roomName] == nil {
		return nil, GameNotFoundErr
	}
	return r.games[roomName], nil
}
